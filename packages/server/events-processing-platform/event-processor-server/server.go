package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/commands"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore/store"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstroredb"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/subscriptions"
	email_validation_subscription "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/subscriptions/email_validation"
	graph_subscription "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/subscriptions/graph"
	location_validation_subscription "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/subscriptions/location_validation"
	phone_number_validation_subscription "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/subscriptions/phone_number_validation"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	waitShotDownDuration = 3 * time.Second
)

type server struct {
	cfg          *config.Config
	log          logger.Logger
	repositories *repository.Repositories
	commands     *domain.Commands
	echo         *echo.Echo
	doneCh       chan struct{}
	//	metrics            *metrics.ESMicroserviceMetrics
}

func NewServer(cfg *config.Config, log logger.Logger) *server {
	return &server{cfg: cfg, log: log, echo: echo.New(), doneCh: make(chan struct{})}
}

func (server *server) SetRepository(repository *repository.Repositories) {
	server.repositories = repository
}

func (server *server) SetCommands(commands *domain.Commands) {
	server.commands = commands
}

func (server *server) Run(parentCtx context.Context) error {
	ctx, cancel := signal.NotifyContext(parentCtx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := validator.GetValidator().Struct(server.cfg); err != nil {
		return errors.Wrap(err, "cfg validate")
	}

	// Setting up tracing
	if server.cfg.Jaeger.Enabled {
		tracer, closer, err := tracing.NewJaegerTracer(&server.cfg.Jaeger, server.log)
		if err != nil {
			server.log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
		}
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer)
	}

	//server.metrics = metrics.NewESMicroserviceMetrics(server.cfg)
	//server.interceptorManager = interceptors.NewInterceptorManager(server.log, server.getGrpcMetricsCb())
	//server.mw = middlewares.NewMiddlewareManager(server.log, server.cfg, server.getHttpMetricsCb())

	db, err := eventstroredb.NewEventStoreDB(server.cfg.EventStoreConfig, server.log)
	if err != nil {
		return err
	}
	defer db.Close() // nolint: errcheck

	// Setting up eventstore subscriptions
	err = subscriptions.NewSubscriptions(server.log, db, server.cfg).RefreshSubscriptions(ctx)
	if err != nil {
		server.log.Errorf("(graphConsumer.Connect) err: {%v}", err)
		cancel()
	}

	// Setting up Neo4j
	neo4jDriver, err := config.NewDriver(server.cfg)
	if err != nil {
		logrus.Fatalf("Could not establish connection with neo4j at: %v, error: %v", server.cfg.Neo4j.Target, err.Error())
	}
	defer neo4jDriver.Close(ctx)
	server.repositories = repository.InitRepos(&neo4jDriver)

	aggregateStore := store.NewAggregateStore(server.log, db)
	server.commands = commands.CreateCommands(server.log, server.cfg, aggregateStore)

	if server.cfg.Subscriptions.GraphSubscription.Enabled {
		graphSubscriber := graph_subscription.NewGraphSubscriber(server.log, db, server.repositories, server.cfg)
		go func() {
			err := graphSubscriber.Connect(ctx, graphSubscriber.ProcessEvents)
			if err != nil {
				server.log.Errorf("(graphSubscriber.Connect) err: {%v}", err)
				cancel()
			}
		}()
	}

	if server.cfg.Subscriptions.EmailValidationSubscription.Enabled {
		emailValidationSubscriber := email_validation_subscription.NewEmailValidationSubscriber(server.log, db, server.cfg, server.commands.EmailCommands)
		go func() {
			err := emailValidationSubscriber.Connect(ctx, emailValidationSubscriber.ProcessEvents)
			if err != nil {
				server.log.Errorf("(emailValidationSubscriber.Connect) err: {%v}", err)
				cancel()
			}
		}()
	}

	if server.cfg.Subscriptions.PhoneNumberValidationSubscription.Enabled {
		phoneNumberValidationSubscriber := phone_number_validation_subscription.NewPhoneNumberValidationSubscriber(server.log, db, server.cfg, server.commands.PhoneNumberCommands, server.repositories)
		go func() {
			err := phoneNumberValidationSubscriber.Connect(ctx, phoneNumberValidationSubscriber.ProcessEvents)
			if err != nil {
				server.log.Errorf("(phoneNumberValidationSubscriber.Connect) err: {%v}", err)
				cancel()
			}
		}()
	}

	if server.cfg.Subscriptions.LocationValidationSubscription.Enabled {
		locationValidationSubscriber := location_validation_subscription.NewLocationValidationSubscriber(server.log, db, server.cfg, server.commands.LocationCommands, server.repositories)
		go func() {
			err := locationValidationSubscriber.Connect(ctx, locationValidationSubscriber.ProcessEvents)
			if err != nil {
				server.log.Errorf("(locationValidationSubscriber.Connect) err: {%v}", err)
				cancel()
			}
		}()
	}

	//server.runMetrics(cancel)
	//server.runHealthCheck(ctx)

	closeGrpcServer, grpcServer, err := server.newEventProcessorGrpcServer()
	if err != nil {
		cancel()
		return err
	}
	defer closeGrpcServer() // nolint: errcheck

	<-ctx.Done()
	server.waitShootDown(waitShotDownDuration)

	grpcServer.GracefulStop()

	if err := server.echo.Shutdown(ctx); err != nil {
		server.log.Warnf("(Shutdown) err: {%validate}", err)
	}

	<-server.doneCh
	server.log.Infof("%server server exited properly", GetMicroserviceName(server.cfg))
	return nil
}

func (server *server) waitShootDown(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		server.doneCh <- struct{}{}
	}()
}

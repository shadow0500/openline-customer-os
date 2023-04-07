package projection

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	contact_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/event_handlers"
	contact_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/events"
	email_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/email/event_handlers"
	email_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/email/events"
	phone_number_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/phone_number/event_handlers"
	phone_number_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/phone_number/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type GraphProjection struct {
	log                     logger.Logger
	db                      *esdb.Client
	cfg                     *config.Config
	repositories            *repository.Repositories
	phoneNumberEventHandler *phone_number_event_handlers.GraphPhoneNumberEventHandler
	contactEventHandler     *contact_event_handlers.GraphContactEventHandler
	emailEventHandler       *email_event_handlers.GraphEmailEventHandler
}

func NewGraphProjection(log logger.Logger, db *esdb.Client, repositories *repository.Repositories, cfg *config.Config) *GraphProjection {
	return &GraphProjection{
		log:                     log,
		db:                      db,
		repositories:            repositories,
		cfg:                     cfg,
		phoneNumberEventHandler: &phone_number_event_handlers.GraphPhoneNumberEventHandler{Repositories: repositories},
		contactEventHandler:     &contact_event_handlers.GraphContactEventHandler{Repositories: repositories},
		emailEventHandler:       &email_event_handlers.GraphEmailEventHandler{Repositories: repositories},
	}
}

type Worker func(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error

func (gp *GraphProjection) Subscribe(ctx context.Context, prefixes []string, poolSize int, worker Worker) error {
	gp.log.Infof("(starting graph subscription) prefixes: {%+v}", prefixes)

	err := gp.db.CreatePersistentSubscriptionAll(ctx, gp.cfg.Subscriptions.GraphProjectionGroupName, esdb.PersistentAllSubscriptionOptions{
		Filter: &esdb.SubscriptionFilter{Type: esdb.StreamFilterType, Prefixes: prefixes},
	})
	if err != nil {
		if subscriptionError, ok := err.(*esdb.PersistentSubscriptionError); !ok || ok && (subscriptionError.Code != 6) {
			gp.log.Errorf("(CreatePersistentSubscriptionAll) err: {%v}", subscriptionError.Error())
		} else if ok && (subscriptionError.Code == 6) {
			// FIXME alexb refactor: call update only if current and new prefixes are different
			settings := esdb.SubscriptionSettingsDefault()
			err = gp.db.UpdatePersistentSubscriptionAll(ctx, gp.cfg.Subscriptions.GraphProjectionGroupName, esdb.PersistentAllSubscriptionOptions{
				Settings: &settings,
				Filter:   &esdb.SubscriptionFilter{Type: esdb.StreamFilterType, Prefixes: prefixes},
			})
			if err != nil {
				if subscriptionError, ok = err.(*esdb.PersistentSubscriptionError); !ok || ok && (subscriptionError.Code != 6) {
					gp.log.Errorf("(UpdatePersistentSubscriptionAll) err: {%v}", subscriptionError.Error())
				}
			}
		}
	}

	stream, err := gp.db.ConnectToPersistentSubscription(
		ctx,
		constants.EsAll,
		gp.cfg.Subscriptions.GraphProjectionGroupName,
		esdb.ConnectToPersistentSubscriptionOptions{},
	)
	if err != nil {
		return err
	}
	defer stream.Close()

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= poolSize; i++ {
		g.Go(gp.runWorker(ctx, worker, stream, i))
	}
	return g.Wait()
}

func (gp *GraphProjection) runWorker(ctx context.Context, worker Worker, stream *esdb.PersistentSubscription, i int) func() error {
	return func() error {
		return worker(ctx, stream, i)
	}
}

func (gp *GraphProjection) ProcessEvents(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error {

	for {
		event := stream.Recv()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if event.SubscriptionDropped != nil {
			gp.log.Errorf("(SubscriptionDropped) err: {%v}", event.SubscriptionDropped.Error)
			return errors.Wrap(event.SubscriptionDropped.Error, "Subscription Dropped")
		}

		if event.EventAppeared != nil {
			gp.log.ProjectionEvent(constants.GraphProjection, gp.cfg.Subscriptions.GraphProjectionGroupName, event.EventAppeared, workerID)

			err := gp.When(ctx, eventstore.NewEventFromRecorded(event.EventAppeared.Event))
			if err != nil {
				gp.log.Errorf("(GraphProjection.when) err: {%v}", err)

				if err := stream.Nack(err.Error(), esdb.Nack_Retry, event.EventAppeared); err != nil {
					gp.log.Errorf("(stream.Nack) err: {%v}", err)
					return errors.Wrap(err, "stream.Nack")
				}
			}

			err = stream.Ack(event.EventAppeared)
			if err != nil {
				gp.log.Errorf("(stream.Ack) err: {%v}", err)
				return errors.Wrap(err, "stream.Ack")
			}
			gp.log.Infof("(ACK) event commit: {%v}", *event.EventAppeared.Commit)
		}
	}
}

func (gp *GraphProjection) When(ctx context.Context, evt eventstore.Event) error {
	ctx, span := tracing.StartProjectionTracerSpan(ctx, "GraphProjection.When", evt)
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()), log.String("EventType", evt.GetEventType()))

	switch evt.GetEventType() {

	case phone_number_events.PhoneNumberCreated:
		return gp.phoneNumberEventHandler.OnPhoneNumberCreate(ctx, evt)
	case phone_number_events.PhoneNumberUpdated:
		return gp.phoneNumberEventHandler.OnPhoneNumberUpdate(ctx, evt)

	case email_events.EmailCreated:
		return gp.emailEventHandler.OnEmailCreate(ctx, evt)
	case email_events.EmailUpdated:
		return gp.emailEventHandler.OnEmailUpdate(ctx, evt)

	case contact_events.ContactCreated:
		return gp.contactEventHandler.OnContactCreate(ctx, evt)
	case contact_events.ContactUpdated:
		return gp.contactEventHandler.OnContactUpdate(ctx, evt)
	case contact_events.ContactPhoneNumberLinked:
		return gp.contactEventHandler.OnPhoneNumberLinkedToContact(ctx, evt)

	default:
		gp.log.Warnf("(GraphProjection) [When unknown EventType] eventType: {%s}", evt.EventType)
		return eventstore.ErrInvalidEventType
	}
}

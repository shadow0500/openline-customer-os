package commands

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/email/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

type EmailValidatedCommandHandler interface {
	Handle(ctx context.Context, command *EmailValidatedCommand) error
}

type emailValidatedCommandHandler struct {
	log logger.Logger
	cfg *config.Config
	es  eventstore.AggregateStore
}

func NewEmailValidatedCommandHandler(log logger.Logger, cfg *config.Config, es eventstore.AggregateStore) *emailValidatedCommandHandler {
	return &emailValidatedCommandHandler{log: log, cfg: cfg, es: es}
}

func (c *emailValidatedCommandHandler) Handle(ctx context.Context, command *EmailValidatedCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailValidatedCommandHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("Tenant", command.Tenant), log.String("ObjectID", command.ObjectID))

	emailAggregate := aggregate.NewEmailAggregateWithTenantAndID(command.Tenant, command.ObjectID)
	err := c.es.Exists(ctx, emailAggregate.GetID())
	if err != nil && !errors.Is(err, eventstore.ErrAggregateNotFound) {
		return err
	}

	emailAggregate, _ = aggregate.LoadEmailAggregate(ctx, c.es, command.Tenant, command.ObjectID)
	if err = emailAggregate.EmailValidated(ctx, command.Tenant, command.RawEmail, command.IsReachable, command.ValidationError, command.Domain, command.Username, command.EmailAddress,
		command.AcceptsMail, command.CanConnectSmtp, command.HasFullInbox, command.IsCatchAll, command.IsDeliverable, command.IsDisabled, command.IsValidSyntax); err != nil {
		return err
	}
	return c.es.Save(ctx, emailAggregate)
}

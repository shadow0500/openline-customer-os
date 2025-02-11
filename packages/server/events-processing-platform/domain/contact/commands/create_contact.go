package commands

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type CreateContactCommandHandler interface {
	Handle(ctx context.Context, command *CreateContactCommand) error
}

type createContactHandler struct {
	log        logger.Logger
	cfg        *config.Config
	eventStore eventstore.AggregateStore
}

func NewCreateContactCommandHandler(log logger.Logger, cfg *config.Config, es eventstore.AggregateStore) *createContactHandler {
	return &createContactHandler{log: log, cfg: cfg, eventStore: es}
}

func (c *createContactHandler) Handle(ctx context.Context, command *CreateContactCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createContactCommandHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("Tenant", command.Tenant), log.String("ObjectID", command.ObjectID))
	//span, ctx := opentracing.StartSpanFromContext(ctx, "createContactHandler.Handle")
	//defer span.Finish()
	//span.LogFields(log.String("ObjectID", command.GetAggregateID()))

	contactAggregate := aggregate.NewContactAggregateWithTenantAndID("openline", command.ObjectID)
	err := contactAggregate.CreateContact(ctx, &models.ContactDto{
		ID:          command.ObjectID,
		Tenant:      command.Tenant,
		FirstName:   command.FirstName,
		LastName:    command.LastName,
		Prefix:      command.Prefix,
		Description: command.Description,
		Source:      command.Source,
		CreatedAt:   command.CreatedAt,
		UpdatedAt:   command.CreatedAt,
	})
	if err != nil {
		return fmt.Errorf("CreateContactCommandHandler.Handle: failed to create contact: %w", err)
	}

	return c.eventStore.Save(ctx, contactAggregate)

}

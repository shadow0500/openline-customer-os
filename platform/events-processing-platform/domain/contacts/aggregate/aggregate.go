package aggregate

import (
	"github.com/openline-ai/openline-customer-os/platform/events-processing-platform/domain/contacts/events"
	"github.com/openline-ai/openline-customer-os/platform/events-processing-platform/domain/contacts/models"
	"github.com/openline-ai/openline-customer-os/platform/events-processing-platform/eventstore"
	"github.com/pkg/errors"
)

const (
	ContactAggregateType eventstore.AggregateType = "CONTACT"
)

type ContactAggregate struct {
	*eventstore.AggregateBase
	Contact *models.Contact
}

func NewContactAggregateWithID(id string) *ContactAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewContactAggregate()
	aggregate.SetID(id)
	aggregate.Contact.ID = id
	return aggregate
}

func NewContactAggregate() *ContactAggregate {
	contactAggregate := &ContactAggregate{Contact: models.NewContact()}
	base := eventstore.NewAggregateBase(contactAggregate.When)
	base.SetType(ContactAggregateType)
	contactAggregate.AggregateBase = base
	return contactAggregate
}

func (contactAggregate *ContactAggregate) When(event eventstore.Event) error {

	switch event.GetEventType() {

	case events.ContactCreated:
		return contactAggregate.onContactCreated(event)
	case events.ContactDeleted:
		return contactAggregate.onContactDeleted(event)
	case events.ContactUpdated:
		return contactAggregate.onShoppingCartUpdated(event)
	default:
		return eventstore.ErrInvalidEventType
	}
}

func (contactAggregate *ContactAggregate) onContactCreated(event eventstore.Event) error {
	var eventData events.ContactCreatedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	contactAggregate.Contact.Uuid = eventData.Uuid
	contactAggregate.Contact.FirstName = eventData.FirstName
	contactAggregate.Contact.LastName = eventData.LastName
	return nil
}

func (contactAggregate *ContactAggregate) onContactDeleted(event eventstore.Event) error {
	var eventData events.ContactDeletedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	contactAggregate.Contact.Uuid = eventData.Uuid
	return nil
}

func (contactAggregate *ContactAggregate) onShoppingCartUpdated(event eventstore.Event) error {
	var eventData events.ContactUpdatedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	contactAggregate.Contact.Uuid = eventData.Uuid
	contactAggregate.Contact.FirstName = eventData.FirstName
	contactAggregate.Contact.LastName = eventData.LastName

	return nil
}

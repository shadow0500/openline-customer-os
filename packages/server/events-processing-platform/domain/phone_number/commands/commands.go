package commands

import (
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"time"
)

type CreatePhoneNumberCommand struct {
	eventstore.BaseCommand
	Tenant      string
	PhoneNumber string
	Source      models.Source
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type UpsertPhoneNumberCommand struct {
	eventstore.BaseCommand
	Tenant         string
	RawPhoneNumber string
	Source         models.Source
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type FailedPhoneNumberValidationCommand struct {
	eventstore.BaseCommand
	Tenant          string
	ValidationError string
}

type SkippedPhoneNumberValidationCommand struct {
	eventstore.BaseCommand
	Tenant               string
	ValidationSkipReason string
}

type PhoneNumberValidatedCommand struct {
	eventstore.BaseCommand
	Tenant      string
	PhoneNumber string
	E164        string
}

func NewCreatePhoneNumberCommand(baseAggregateId, tenant, rawPhoneNumber, source, sourceOfTruth, appSource string, createdAt, updatedAt *time.Time) *CreatePhoneNumberCommand {
	return &CreatePhoneNumberCommand{
		BaseCommand: eventstore.NewBaseCommand(baseAggregateId),
		Tenant:      tenant,
		PhoneNumber: rawPhoneNumber,
		Source: models.Source{
			Source:        source,
			SourceOfTruth: sourceOfTruth,
			AppSource:     appSource,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewUpsertPhoneNumberCommand(baseAggregateId, tenant, rawPhoneNumber, source, sourceOfTruth, appSource string, createdAt, updatedAt *time.Time) *UpsertPhoneNumberCommand {
	return &UpsertPhoneNumberCommand{
		BaseCommand:    eventstore.NewBaseCommand(baseAggregateId),
		Tenant:         tenant,
		RawPhoneNumber: rawPhoneNumber,
		Source: models.Source{
			Source:        source,
			SourceOfTruth: sourceOfTruth,
			AppSource:     appSource,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewFailedPhoneNumberValidationCommand(baseAggregateId, tenant, validationError string) *FailedPhoneNumberValidationCommand {
	return &FailedPhoneNumberValidationCommand{
		BaseCommand:     eventstore.NewBaseCommand(baseAggregateId),
		Tenant:          tenant,
		ValidationError: validationError,
	}
}

func NewSkippedPhoneNumberValidationCommand(baseAggregateId, tenant, validationSkipReason string) *SkippedPhoneNumberValidationCommand {
	return &SkippedPhoneNumberValidationCommand{
		BaseCommand:          eventstore.NewBaseCommand(baseAggregateId),
		Tenant:               tenant,
		ValidationSkipReason: validationSkipReason,
	}
}

func NewPhoneNumberValidatedCommand(baseAggregateId, tenant, e164 string) *PhoneNumberValidatedCommand {
	return &PhoneNumberValidatedCommand{
		BaseCommand: eventstore.NewBaseCommand(baseAggregateId),
		Tenant:      tenant,
		E164:        e164,
	}
}

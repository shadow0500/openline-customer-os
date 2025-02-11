package commands

import (
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
)

type PhoneNumberCommands struct {
	UpsertPhoneNumber           UpsertPhoneNumberCommandHandler
	CreatePhoneNumber           CreatePhoneNumberCommandHandler
	FailedPhoneNumberValidation FailedPhoneNumberValidationCommandHandler
	SkipPhoneNumberValidation   SkippedPhoneNumberValidationCommandHandler
	PhoneNumberValidated        PhoneNumberValidatedCommandHandler
}

func NewPhoneNumberCommands(log logger.Logger, cfg *config.Config, es eventstore.AggregateStore) *PhoneNumberCommands {
	return &PhoneNumberCommands{
		CreatePhoneNumber:           NewCreatePhoneNumberCommandHandler(log, cfg, es),
		UpsertPhoneNumber:           NewUpsertPhoneNumberHandler(log, cfg, es),
		FailedPhoneNumberValidation: NewFailedPhoneNumberValidationCommandHandler(log, cfg, es),
		SkipPhoneNumberValidation:   NewSkippedPhoneNumberValidationCommandHandler(log, cfg, es),
		PhoneNumberValidated:        NewPhoneNumberValidatedCommandHandler(log, cfg, es),
	}
}

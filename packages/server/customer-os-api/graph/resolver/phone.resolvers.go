package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// PhoneNumberMergeToContact is the resolver for the phoneNumberMergeToContact field.
func (r *mutationResolver) PhoneNumberMergeToContact(ctx context.Context, contactID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	result, err := r.Services.PhoneNumberService.MergePhoneNumberToContact(ctx, contactID, mapper.MapPhoneNumberInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add phone number %s to contact %s", input.E164, contactID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberUpdateInContact is the resolver for the phoneNumberUpdateInContact field.
func (r *mutationResolver) PhoneNumberUpdateInContact(ctx context.Context, contactID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	result, err := r.Services.PhoneNumberService.UpdatePhoneNumberInContact(ctx, contactID, mapper.MapPhoneNumberUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update email %s in contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToPhoneNumber(result), nil
}

// PhoneNumberDeleteFromContact is the resolver for the phoneNumberDeleteFromContact field.
func (r *mutationResolver) PhoneNumberDeleteFromContact(ctx context.Context, contactID string, e164 string) (*model.Result, error) {
	result, err := r.Services.PhoneNumberService.Delete(ctx, contactID, e164)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove phone number %s from contact %s", e164, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// PhoneNumberDeleteFromContactByID is the resolver for the phoneNumberDeleteFromContactById field.
func (r *mutationResolver) PhoneNumberDeleteFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	result, err := r.Services.PhoneNumberService.DeleteById(ctx, contactID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove phone number %s from contact %s", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/customer-os-api/mapper"
)

// EmailMergeToContact is the resolver for the emailMergeToContact field.
func (r *mutationResolver) EmailMergeToContact(ctx context.Context, contactID string, input model.EmailInput) (*model.Email, error) {
	result, err := r.ServiceContainer.EmailService.MergeEmailToContact(ctx, contactID, mapper.MapEmailInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add email %s to contact %s", input.Email, contactID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailUpdateInContact is the resolver for the emailUpdateInContact field.
func (r *mutationResolver) EmailUpdateInContact(ctx context.Context, contactID string, input model.EmailUpdateInput) (*model.Email, error) {
	result, err := r.ServiceContainer.EmailService.UpdateEmailInContact(ctx, contactID, mapper.MapEmailUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update email %s in contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailRemoveFromContact is the resolver for the EmailRemoveFromContact field.
func (r *mutationResolver) EmailRemoveFromContact(ctx context.Context, contactID string, email string) (*model.Result, error) {
	result, err := r.ServiceContainer.EmailService.Delete(ctx, contactID, email)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove email %s from contact %s", email, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailRemoveFromContactByID is the resolver for the emailRemoveFromContactById field.
func (r *mutationResolver) EmailRemoveFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	result, err := r.ServiceContainer.EmailService.DeleteById(ctx, contactID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove email %s from contact %s", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
)

// EmailMergeToContact is the resolver for the emailMergeToContact field.
func (r *mutationResolver) EmailMergeToContact(ctx context.Context, contactID string, input model.EmailInput) (*model.Email, error) {
	result, err := r.Services.EmailService.MergeEmailTo(ctx, repository.CONTACT, contactID, mapper.MapEmailInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add email %s to contact %s", input.Email, contactID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailUpdateInContact is the resolver for the emailUpdateInContact field.
func (r *mutationResolver) EmailUpdateInContact(ctx context.Context, contactID string, input model.EmailUpdateInput) (*model.Email, error) {
	result, err := r.Services.EmailService.UpdateEmailFor(ctx, repository.CONTACT, contactID, mapper.MapEmailUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update email %s in contact %s", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailRemoveFromContact is the resolver for the EmailRemoveFromContact field.
func (r *mutationResolver) EmailRemoveFromContact(ctx context.Context, contactID string, email string) (*model.Result, error) {
	result, err := r.Services.EmailService.Delete(ctx, contactID, email)
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
	result, err := r.Services.EmailService.DeleteById(ctx, contactID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove email %s from contact %s", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// EmailMergeToUser is the resolver for the emailMergeToUser field.
func (r *mutationResolver) EmailMergeToUser(ctx context.Context, userID string, input model.EmailInput) (*model.Email, error) {
	result, err := r.Services.EmailService.MergeEmailTo(ctx, repository.USER, userID, mapper.MapEmailInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add email %s to user %s", input.Email, userID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailUpdateInUser is the resolver for the emailUpdateInUser field.
func (r *mutationResolver) EmailUpdateInUser(ctx context.Context, userID string, input model.EmailUpdateInput) (*model.Email, error) {
	result, err := r.Services.EmailService.UpdateEmailFor(ctx, repository.USER, userID, mapper.MapEmailUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update email %s in user %s", input.ID, userID)
		return nil, err
	}
	return mapper.MapEntityToEmail(result), nil
}

// EmailRemoveFromUserByID is the resolver for the emailRemoveFromUserById field.
func (r *mutationResolver) EmailRemoveFromUserByID(ctx context.Context, userID string, id string) (*model.Result, error) {
	result, err := r.Services.EmailService.DeleteById(ctx, userID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove email %s from user %s", id, userID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

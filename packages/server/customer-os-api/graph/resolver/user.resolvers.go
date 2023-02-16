package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserInput) (*model.User, error) {
	createdUserEntity, err := r.Services.UserService.Create(ctx, &service.UserCreateData{
		UserEntity:  mapper.MapUserInputToEntity(input),
		EmailEntity: mapper.MapEmailInputToEntity(input.Email),
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create user %s %s", input.FirstName, input.LastName)
		return nil, err
	}
	return mapper.MapEntityToUser(createdUserEntity), nil
}

// UserUpdate is the resolver for the user_Update field.
func (r *mutationResolver) UserUpdate(ctx context.Context, input model.UserUpdateInput) (*model.User, error) {
	updatedUserEntity, err := r.Services.UserService.Update(ctx, mapper.MapUserUpdateInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update user %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToUser(updatedUserEntity), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.UserPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.UserService.FindAll(ctx, pagination.Page, pagination.Limit, where, sort)
	return &model.UserPage{
		Content:       mapper.MapEntitiesToUsers(paginatedResult.Rows.(*entity.UserEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	userEntity, err := r.Services.UserService.FindUserById(ctx, id)
	if err != nil || userEntity == nil {
		graphql.AddErrorf(ctx, "User with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToUser(userEntity), nil
}

// UserByEmail is the resolver for the user_ByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	userEntity, err := r.Services.UserService.FindUserByEmail(ctx, email)
	if err != nil || userEntity == nil {
		graphql.AddErrorf(ctx, "User with email %s not identified", email)
		return nil, err
	}
	return mapper.MapEntityToUser(userEntity), nil
}

// Emails is the resolver for the emails field.
func (r *userResolver) Emails(ctx context.Context, obj *model.User) ([]*model.Email, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	emailEntities, err := r.Services.EmailService.FindAllFor(ctx, entity.USER, obj.ID)
	return mapper.MapEntitiesToEmails(emailEntities), err
}

// Conversations is the resolver for the conversations field.
func (r *userResolver) Conversations(ctx context.Context, obj *model.User, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ConversationService.GetConversationsForUser(ctx, obj.ID, pagination.Page, pagination.Limit, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get user %s conversations", obj.ID)
		return nil, err
	}
	return &model.ConversationPage{
		Content:       mapper.MapEntitiesToConversations(paginatedResult.Rows.(*entity.ConversationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

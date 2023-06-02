package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// Contacts is the resolver for the contacts field.
func (r *conversationResolver) Contacts(ctx context.Context, obj *model.Conversation) ([]*model.Contact, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "ConversationResolver.Contacts", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.conversationID", obj.ID))

	contactEntities, err := r.Services.ContactService.GetAllForConversation(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to fetch contacts for conversation %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToContacts(contactEntities), nil
}

// Users is the resolver for the users field.
func (r *conversationResolver) Users(ctx context.Context, obj *model.Conversation) ([]*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "ConversationResolver.Users", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.conversationID", obj.ID))

	userEntities, err := r.Services.UserService.GetAllForConversation(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to fetch users for conversation %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToUsers(userEntities), nil
}

// ConversationCreate is the resolver for the conversationCreate field.
func (r *mutationResolver) ConversationCreate(ctx context.Context, input model.ConversationInput) (*model.Conversation, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ConversationCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	conversationEntity, err := r.Services.ConversationService.CreateNewConversation(ctx, input.UserIds, input.ContactIds, mapper.MapConversationInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to create conversation between users: %v and contacts: %v", input.UserIds, input.ContactIds)
		return nil, err
	}
	return mapper.MapEntityToConversation(conversationEntity), nil
}

// ConversationUpdate is the resolver for the conversation_Update field.
func (r *mutationResolver) ConversationUpdate(ctx context.Context, input model.ConversationUpdateInput) (*model.Conversation, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ConversationUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.conversationID", input.ID))

	conversationEntity, err := r.Services.ConversationService.UpdateConversation(
		ctx, input.UserIds, input.ContactIds, mapper.MapConversationUpdateInputToEntity(input), input.SkipMessageCountIncrement)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to update conversation %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToConversation(conversationEntity), nil
}

// ConversationClose is the resolver for the conversation_Close field.
func (r *mutationResolver) ConversationClose(ctx context.Context, conversationID string) (*model.Conversation, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ConversationClose", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.conversationID", conversationID))

	conversationEntity, err := r.Services.ConversationService.CloseConversation(ctx, conversationID, entity.DataSourceOpenline)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to close conversation %s", conversationID)
		return nil, err
	}
	return mapper.MapEntityToConversation(conversationEntity), nil
}

// Conversation returns generated.ConversationResolver implementation.
func (r *Resolver) Conversation() generated.ConversationResolver { return &conversationResolver{r} }

type conversationResolver struct{ *Resolver }

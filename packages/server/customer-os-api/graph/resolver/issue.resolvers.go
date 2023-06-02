package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// Tags is the resolver for the tags field.
func (r *issueResolver) Tags(ctx context.Context, obj *model.Issue) ([]*model.Tag, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "IssueResolver.Tags", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.issueID", obj.ID))

	tagEntities, err := dataloader.For(ctx).GetTagsForIssue(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get tags for issue %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToTags(tagEntities), nil
}

// MentionedByNotes is the resolver for the mentionedByNotes field.
func (r *issueResolver) MentionedByNotes(ctx context.Context, obj *model.Issue) ([]*model.Note, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "IssueResolver.MentionedByNotes", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.issueID", obj.ID))

	noteEntities, err := dataloader.For(ctx).GetMentionedByNotesForIssue(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get notes for issue %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToNotes(noteEntities), nil
}

// InteractionEvents is the resolver for the interactionEvents field.
func (r *issueResolver) InteractionEvents(ctx context.Context, obj *model.Issue) ([]*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "IssueResolver.InteractionEvents", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.issueID", obj.ID))

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForIssue(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get interaction events for issue %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEvents(interactionEventEntities), nil
}

// Issue is the resolver for the issue field.
func (r *queryResolver) Issue(ctx context.Context, id string) (*model.Issue, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Issue", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.issueID", id))

	issueEntity, err := r.Services.IssueService.GetById(ctx, id)
	if err != nil || issueEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Issue with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToIssue(issueEntity), nil
}

// Issue returns generated.IssueResolver implementation.
func (r *Resolver) Issue() generated.IssueResolver { return &issueResolver{r} }

type issueResolver struct{ *Resolver }

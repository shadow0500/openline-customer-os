package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// Tags is the resolver for the tags field.
func (r *issueResolver) Tags(ctx context.Context, obj *model.Issue) ([]*model.Tag, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	tagEntities, err := dataloader.For(ctx).GetTagsForIssue(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get tags for issue %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToTags(tagEntities), nil
}

// MentionedByNotes is the resolver for the mentionedByNotes field.
func (r *issueResolver) MentionedByNotes(ctx context.Context, obj *model.Issue) ([]*model.Note, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	noteEntities, err := dataloader.For(ctx).GetMentionedByNotesForIssue(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get notes for issue %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToNotes(noteEntities), nil
}

// InteractionEvents is the resolver for the interactionEvents field.
func (r *issueResolver) InteractionEvents(ctx context.Context, obj *model.Issue) ([]*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEvents - interactionEvents"))
}

// Issue is the resolver for the issue field.
func (r *queryResolver) Issue(ctx context.Context, id string) (*model.Issue, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	issueEntity, err := r.Services.IssueService.GetById(ctx, id)
	if err != nil || issueEntity == nil {
		graphql.AddErrorf(ctx, "Issue with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToIssue(issueEntity), nil
}

// Issue returns generated.IssueResolver implementation.
func (r *Resolver) Issue() generated.IssueResolver { return &issueResolver{r} }

type issueResolver struct{ *Resolver }

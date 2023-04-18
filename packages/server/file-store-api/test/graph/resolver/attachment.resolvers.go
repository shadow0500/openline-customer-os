package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// AttachmentCreate is the resolver for the attachment_Create field.
func (r *mutationResolver) AttachmentCreate(ctx context.Context, input model.AttachmentInput) (*model.Attachment, error) {
	if r.Resolver.AttachmentCreate != nil {
		return r.Resolver.AttachmentCreate(ctx, input)
	}
	panic(fmt.Errorf("not implemented: AttachmentCreate - attachment_Create"))
}

// Attachment is the resolver for the attachment field.
func (r *queryResolver) Attachment(ctx context.Context, id string) (*model.Attachment, error) {
	if r.Resolver.Attachment != nil {
		return r.Resolver.Attachment(ctx, id)
	}
	panic(fmt.Errorf("not implemented: Attachment - attachment"))
}

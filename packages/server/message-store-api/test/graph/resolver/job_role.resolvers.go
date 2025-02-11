package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// Organization is the resolver for the organization field.
func (r *jobRoleResolver) Organization(ctx context.Context, obj *model.JobRole) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Contact is the resolver for the contact field.
func (r *jobRoleResolver) Contact(ctx context.Context, obj *model.JobRole) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: Contact - contact"))
}

// JobRoleDelete is the resolver for the jobRole_Delete field.
func (r *mutationResolver) JobRoleDelete(ctx context.Context, contactID string, roleID string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: JobRoleDelete - jobRole_Delete"))
}

// JobRoleCreate is the resolver for the jobRole_Create field.
func (r *mutationResolver) JobRoleCreate(ctx context.Context, contactID string, input model.JobRoleInput) (*model.JobRole, error) {
	panic(fmt.Errorf("not implemented: JobRoleCreate - jobRole_Create"))
}

// JobRoleUpdate is the resolver for the jobRole_Update field.
func (r *mutationResolver) JobRoleUpdate(ctx context.Context, contactID string, input model.JobRoleUpdateInput) (*model.JobRole, error) {
	panic(fmt.Errorf("not implemented: JobRoleUpdate - jobRole_Update"))
}

// JobRole returns generated.JobRoleResolver implementation.
func (r *Resolver) JobRole() generated.JobRoleResolver { return &jobRoleResolver{r} }

type jobRoleResolver struct{ *Resolver }

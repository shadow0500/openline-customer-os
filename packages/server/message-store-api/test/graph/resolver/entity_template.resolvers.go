package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// FieldSetTemplates is the resolver for the fieldSetTemplates field.
func (r *entityTemplateResolver) FieldSetTemplates(ctx context.Context, obj *model.EntityTemplate) ([]*model.FieldSetTemplate, error) {
	panic(fmt.Errorf("not implemented: FieldSetTemplates - fieldSetTemplates"))
}

// CustomFieldTemplates is the resolver for the customFieldTemplates field.
func (r *entityTemplateResolver) CustomFieldTemplates(ctx context.Context, obj *model.EntityTemplate) ([]*model.CustomFieldTemplate, error) {
	panic(fmt.Errorf("not implemented: CustomFieldTemplates - customFieldTemplates"))
}

// CustomFieldTemplates is the resolver for the customFieldTemplates field.
func (r *fieldSetTemplateResolver) CustomFieldTemplates(ctx context.Context, obj *model.FieldSetTemplate) ([]*model.CustomFieldTemplate, error) {
	panic(fmt.Errorf("not implemented: CustomFieldTemplates - customFieldTemplates"))
}

// EntityTemplateCreate is the resolver for the entityTemplateCreate field.
func (r *mutationResolver) EntityTemplateCreate(ctx context.Context, input model.EntityTemplateInput) (*model.EntityTemplate, error) {
	panic(fmt.Errorf("not implemented: EntityTemplateCreate - entityTemplateCreate"))
}

// EntityTemplate returns generated.EntityTemplateResolver implementation.
func (r *Resolver) EntityTemplate() generated.EntityTemplateResolver {
	return &entityTemplateResolver{r}
}

// FieldSetTemplate returns generated.FieldSetTemplateResolver implementation.
func (r *Resolver) FieldSetTemplate() generated.FieldSetTemplateResolver {
	return &fieldSetTemplateResolver{r}
}

type entityTemplateResolver struct{ *Resolver }
type fieldSetTemplateResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *entityTemplateResolver) FieldSets(ctx context.Context, obj *model.EntityTemplate) ([]*model.FieldSetTemplate, error) {
	panic(fmt.Errorf("not implemented: FieldSets - fieldSets"))
}
func (r *entityTemplateResolver) CustomFields(ctx context.Context, obj *model.EntityTemplate) ([]*model.CustomFieldTemplate, error) {
	panic(fmt.Errorf("not implemented: CustomFields - customFields"))
}
func (r *fieldSetTemplateResolver) CustomFields(ctx context.Context, obj *model.FieldSetTemplate) ([]*model.CustomFieldTemplate, error) {
	panic(fmt.Errorf("not implemented: CustomFields - customFields"))
}

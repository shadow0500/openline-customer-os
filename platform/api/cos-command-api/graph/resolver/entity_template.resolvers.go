package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// FieldSets is the resolver for the fieldSets field.
func (r *entityTemplateResolver) FieldSets(ctx context.Context, obj *model.EntityTemplate) ([]*model.FieldSetTemplate, error) {
	result, err := r.Services.FieldSetTemplateService.FindAll(obj.ID)
	return mapper.MapEntitiesToFieldSetTemplates(result), err
}

// CustomFields is the resolver for the customFields field.
func (r *entityTemplateResolver) CustomFields(ctx context.Context, obj *model.EntityTemplate) ([]*model.CustomFieldTemplate, error) {
	result, err := r.Services.CustomFieldTemplateService.FindAllForEntityTemplate(obj.ID)
	return mapper.MapEntitiesToCustomFieldTemplates(result), err
}

// CustomFields is the resolver for the customFields field.
func (r *fieldSetTemplateResolver) CustomFields(ctx context.Context, obj *model.FieldSetTemplate) ([]*model.CustomFieldTemplate, error) {
	result, err := r.Services.CustomFieldTemplateService.FindAllForFieldSetTemplate(obj.ID)
	return mapper.MapEntitiesToCustomFieldTemplates(result), err
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

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// FieldSetTemplates is the resolver for the fieldSetTemplates field.
func (r *entityTemplateResolver) FieldSetTemplates(ctx context.Context, obj *model.EntityTemplate) ([]*model.FieldSetTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, err := r.Services.FieldSetTemplateService.FindAll(ctx, obj.ID)
	return mapper.MapEntitiesToFieldSetTemplates(result), err
}

// CustomFieldTemplates is the resolver for the customFieldTemplates field.
func (r *entityTemplateResolver) CustomFieldTemplates(ctx context.Context, obj *model.EntityTemplate) ([]*model.CustomFieldTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, err := r.Services.CustomFieldTemplateService.FindAllForEntityTemplate(ctx, obj.ID)
	return mapper.MapEntitiesToCustomFieldTemplates(result), err
}

// CustomFieldTemplates is the resolver for the customFieldTemplates field.
func (r *fieldSetTemplateResolver) CustomFieldTemplates(ctx context.Context, obj *model.FieldSetTemplate) ([]*model.CustomFieldTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, err := r.Services.CustomFieldTemplateService.FindAllForFieldSetTemplate(ctx, obj.ID)
	return mapper.MapEntitiesToCustomFieldTemplates(result), err
}

// EntityTemplateCreate is the resolver for the entityTemplateCreate field.
func (r *mutationResolver) EntityTemplateCreate(ctx context.Context, input model.EntityTemplateInput) (*model.EntityTemplate, error) {
	entityTemplateEntity, err := r.Services.EntityTemplateService.Create(ctx, mapper.MapEntityTemplateInputToEntity(input))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create entity template: %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToEntityTemplate(entityTemplateEntity), nil
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

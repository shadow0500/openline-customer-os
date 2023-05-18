package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// EntityTemplates is the resolver for the entityTemplates field.
func (r *queryResolver) EntityTemplates(ctx context.Context, extends *model.EntityTemplateExtension) ([]*model.EntityTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	result, err := r.Services.EntityTemplateService.FindAll(ctx, utils.StringPtr(extends.String()))
	return mapper.MapEntitiesToEntityTemplates(result), err
}

// DashboardViewContacts is the resolver for the dashboardView_Contacts field.
func (r *queryResolver) DashboardViewContacts(ctx context.Context, pagination model.Pagination, where *model.Filter) (*model.ContactsPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	paginatedResult, err := r.Services.QueryService.GetDashboardViewContactsData(ctx, pagination.Page, pagination.Limit, where)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organizations and contacts data")
		return nil, err
	}
	return &model.ContactsPage{
		Content:       mapper.MapEntitiesToContacts(paginatedResult.Rows.(*entity.ContactEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// DashboardViewOrganizations is the resolver for the dashboardView_Organizations field.
func (r *queryResolver) DashboardViewOrganizations(ctx context.Context, pagination model.Pagination, where *model.Filter) (*model.OrganizationPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	paginatedResult, err := r.Services.QueryService.GetDashboardViewOrganizationsData(ctx, pagination.Page, pagination.Limit, where)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organizations and contacts data")
		return nil, err
	}
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

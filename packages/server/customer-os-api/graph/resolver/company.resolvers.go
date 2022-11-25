package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// CompaniesByNameLike is the resolver for the companies_ByNameLike field.
func (r *queryResolver) CompaniesByNameLike(ctx context.Context, pagination *model.Pagination, companyName string) (*model.CompanyPage, error) {
	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.ServiceContainer.CompanyService.FindCompaniesByNameLike(ctx, pagination.Page, pagination.Limit, companyName)
	return &model.CompanyPage{
		Content:       mapper.MapEntitiesToCompanies(paginatedResult.Rows.(*entity.CompanyEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

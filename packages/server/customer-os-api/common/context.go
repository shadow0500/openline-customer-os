package common

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"net/http"
)

type CustomContext struct {
	Tenant                   string
	UserId                   string
	IdentityId               string
	Roles                    []model.Role
	GraphqlRootOperationName string
}

var customContextKey = "CUSTOM_CONTEXT"

func WithContext(customContext *CustomContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), customContextKey, customContext))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func WithCustomContext(ctx context.Context, customContext *CustomContext) context.Context {
	return context.WithValue(ctx, customContextKey, customContext)
}

func GetContext(ctx context.Context) *CustomContext {
	customContext, ok := ctx.Value(customContextKey).(*CustomContext)
	if !ok {
		return new(CustomContext)
	}
	return customContext
}

func GetTenantFromContext(ctx context.Context) string {
	return GetContext(ctx).Tenant
}

func GetRolesFromContext(ctx context.Context) []model.Role {
	return GetContext(ctx).Roles
}

func GetUserIdFromContext(ctx context.Context) string {
	return GetContext(ctx).UserId
}

func GetIdentityIdFromContext(ctx context.Context) string {
	return GetContext(ctx).IdentityId
}

func GetGraphqlRootOperationNameFromContext(ctx context.Context) string {
	return GetContext(ctx).GraphqlRootOperationName
}

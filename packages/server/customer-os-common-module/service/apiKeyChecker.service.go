package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	repository "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/postgres"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/postgres/entity"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type App string

const (
	CUSTOMER_OS_API   App = "customer-os-api"
	FILE_STORE_API    App = "file-store-api"
	SETTINGS_API      App = "settings-api"
	MESSAGE_STORE_API App = "message-store-api"
	OASIS_API         App = "oasis-api"
	VALIDATION_API    App = "validation-api"
)

const ApiKeyHeader = "X-Openline-API-KEY"

func ApiKeyCheckerHTTP(appKeyRepo repository.AppKeyRepository, app App) func(c *gin.Context) {
	return func(c *gin.Context) {
		kh := c.GetHeader(ApiKeyHeader)
		if kh != "" {

			keyResult := appKeyRepo.FindByKey(c, string(app), kh)

			if keyResult.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"errors": []gin.H{{"message": fmt.Sprintf("Error while checking api key: %s", keyResult.Error.Error())}},
				})
				c.Abort()
				return
			}

			appKey := keyResult.Result.(*entity.AppKey)

			if appKey == nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"errors": []gin.H{{"message": "Invalid api key"}},
				})
				c.Abort()
				return
			}

			c.Next()
			// illegal request, terminate the current process
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errors": []gin.H{{"message": "Api key is required"}},
			})
			c.Abort()
			return
		}

	}
}

func ApiKeyCheckerGRPC(ctx context.Context, appKeyRepo repository.AppKeyRepository, app App) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	kh := md.Get(ApiKeyHeader)
	if kh != nil && len(kh) == 1 {
		keyResult := appKeyRepo.FindByKey(ctx, string(app), kh[0])
		if keyResult.Error != nil {
			return false
		}
		appKey := keyResult.Result.(*entity.AppKey)
		if appKey == nil {
			return false
		}
		return true
	}
	return false
}

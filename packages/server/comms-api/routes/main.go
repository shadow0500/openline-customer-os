package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-customer-os/packages/server/comms-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/routes/ContactHub"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/service"
	"log"
	"strings"
)

// Run will start the server
func Run(config *c.Config, hub *ContactHub.ContactHub, services *service.Services) {
	router := getRouter(config, hub, services)
	if err := router.Run(config.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(config *c.Config, hub *ContactHub.ContactHub, services *service.Services) *gin.Engine {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = strings.Split(config.Service.CorsUrl, " ")
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "POST", "GET")
	corsConfig.AddAllowHeaders("WebChatApiKey")

	router.Use(cors.New(corsConfig))
	route := router.Group("/")

	addMailRoutes(config, route, services.MailService, hub)
	addVconRoutes(config, route, services.CustomerOsService, hub)

	addWebSocketRoutes(route, config.WebChat.PingInterval, hub)
	addCallCredentialRoutes(route, config)
	route2 := router.Group("/")

	addHealthRoutes(route2)
	return router
}

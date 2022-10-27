package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	model2 "github.com/openline-ai/openline-customer-os/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/customer-os-api/graph/resolver"
	"github.com/openline-ai/openline-customer-os/customer-os-api/service"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/openline-ai/openline-customer-os/customer-os-api/graph/generated"
)

const customerOSApiPort = "1010"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/health", healthCheckHandler)
	r.GET("/readiness", healthCheckHandler)
	r.GET("/testDB", testDb)

	port := os.Getenv("PORT")
	if port == "" {
		port = customerOSApiPort
	}

	r.Run(":" + port)
}

func testDb(c *gin.Context) {
	contact := model2.ContactNode{
		FirstName:   "Test",
		LastName:    "mata",
		Label:       "asdasdasd",
		ContactType: "WTF",
	}
	aNewSavedContact, err := service.NewContactService().Create(contact)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"wtf_message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"wtf_message": aNewSavedContact,
	})
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
}

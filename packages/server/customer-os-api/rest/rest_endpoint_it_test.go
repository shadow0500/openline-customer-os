package rest

import (
	"context"
	"database/sql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/postgres"
	commonService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/testcontainers/testcontainers-go"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var (
	neo4jContainer testcontainers.Container
	driver         *neo4j.DriverWithContext

	postgresContainer testcontainers.Container
	postgresGormDB    *gorm.DB
	postgresSqlDB     *sql.DB
	serviceContainer  *service.Services
)

const tenantName = "openline"
const testUserId = "test-user-id"
const testContactId = "test-contact-id"
const testPlayerId = "test-player-id"

func TestMain(m *testing.M) {
	neo4jContainer, driver = neo4jt.InitTestNeo4jDB()
	defer func(dbContainer testcontainers.Container, driver neo4j.DriverWithContext, ctx context.Context) {
		neo4jt.CloseDriver(driver)
		neo4jt.Terminate(dbContainer, ctx)
	}(neo4jContainer, *driver, context.Background())

	postgresContainer, postgresGormDB, postgresSqlDB = postgres.InitTestDB()
	defer func(postgresContainer testcontainers.Container, ctx context.Context) {
		err := postgresContainer.Terminate(ctx)
		if err != nil {
			log.Fatal("Error during container termination")
		}
	}(postgresContainer, context.Background())

	prepareClient()

	os.Exit(m.Run())
}

func tearDownTestCase(ctx context.Context) func(tb testing.TB) {
	return func(tb testing.TB) {
		tb.Logf("Teardown test %v, cleaning neo4j DB", tb.Name())
		neo4jt.CleanupAllData(ctx, driver)
	}
}

func prepareClient() {
	log.Printf("******************PREPARING CLIENT******************")
	appLogger := logger.NewAppLogger(&logger.Config{
		DevMode: true,
	})
	appLogger.InitLogger()
	commonServices := commonService.InitServices(postgresGormDB, driver)
	serviceContainer = service.InitServices(appLogger, driver, commonServices, nil)
	dataloader.NewDataLoader(serviceContainer)
	log.Printf("%v", serviceContainer)
}

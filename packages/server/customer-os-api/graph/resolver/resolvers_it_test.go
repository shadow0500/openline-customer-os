package resolver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	cosHandler "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/handler"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/grpc/event_store"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/postgres"
	commonService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"gorm.io/gorm"
	"log"
	"os"
	"reflect"
	"sort"
	"testing"
)

var (
	neo4jContainer testcontainers.Container
	driver         *neo4j.DriverWithContext

	postgresContainer        testcontainers.Container
	postgresGormDB           *gorm.DB
	postgresSqlDB            *sql.DB
	c                        *client.Client
	cOwner                   *client.Client
	cCustomerOsPlatformOwner *client.Client
	cAdmin                   *client.Client
	cAdminWithTenant         *client.Client
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
	appLogger := logger.NewAppLogger(&logger.Config{
		DevMode: true,
	})
	appLogger.InitLogger()

	commonServices := commonService.InitServices(postgresGormDB, driver)
	testDialFactory := event_store.NewTestDialFactory()
	gRPCconn, _ := testDialFactory.GetEventsProcessingPlatformConn()
	serviceContainer := service.InitServices(appLogger, driver, commonServices, grpc_client.InitClients(gRPCconn))
	graphResolver := NewResolver(appLogger, serviceContainer, grpc_client.InitClients(gRPCconn))
	loader := dataloader.NewDataLoader(serviceContainer)
	customCtx := &common.CustomContext{
		Tenant:     tenantName,
		UserId:     testUserId,
		IdentityId: testPlayerId,
		Roles:      []model.Role{model.RoleUser},
	}

	customOwnerCtx := &common.CustomContext{
		Tenant:     tenantName,
		UserId:     testUserId,
		IdentityId: testPlayerId,
		Roles:      []model.Role{model.RoleUser, model.RoleOwner},
	}
	customCustomerOsPlatformOwnerCtx := &common.CustomContext{
		Tenant: tenantName,
		UserId: testUserId,
		Roles:  []model.Role{model.RoleUser, model.RoleCustomerOsPlatformOwner},
	}
	customAdminCtx := &common.CustomContext{
		Roles: []model.Role{model.RoleAdmin},
	}

	customAdminWTenantCtx := &common.CustomContext{
		Tenant: tenantName,
		Roles:  []model.Role{model.RoleAdmin},
	}
	schemaConfig := generated.Config{Resolvers: graphResolver}
	schemaConfig.Directives.HasRole = cosHandler.GetRoleChecker()
	schemaConfig.Directives.HasTenant = cosHandler.GetTenantChecker()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(schemaConfig))
	dataloaderServer := dataloader.Middleware(loader, server)
	handler := common.WithContext(customCtx, dataloaderServer)
	c = client.New(handler)
	cOwner = client.New(common.WithContext(customOwnerCtx, dataloaderServer))
	cCustomerOsPlatformOwner = client.New(common.WithContext(customCustomerOsPlatformOwnerCtx, dataloaderServer))
	cAdmin = client.New(common.WithContext(customAdminCtx, dataloaderServer))
	cAdminWithTenant = client.New(common.WithContext(customAdminWTenantCtx, dataloaderServer))
}

func getQuery(fileName string) string {
	b, err := os.ReadFile(fmt.Sprintf("test_queries/%s.txt", fileName))
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func assertRawResponseSuccess(t *testing.T, response *client.Response, err error) {
	require.Nil(t, err)
	require.NotNil(t, response)
	if response.Errors != nil {
		log.Println(fmt.Sprintf("Error in response: %v", string(response.Errors)))
	}
	require.NotNil(t, response.Data)
	require.Nil(t, response.Errors)
}

func assertNeo4jLabels(ctx context.Context, t *testing.T, driver *neo4j.DriverWithContext, expectedLabels []string) {
	actualLabels := neo4jt.GetAllLabels(ctx, driver)
	sort.Strings(expectedLabels)
	sort.Strings(actualLabels)
	if !reflect.DeepEqual(actualLabels, expectedLabels) {
		t.Errorf("Expected labels: %v, \nActual labels: %v", expectedLabels, actualLabels)
	}
}

func callGraphQL(t *testing.T, queryLocation string, vars map[string]interface{}) (rawResponse *client.Response) {
	// Transform map into var args of options
	options := make([]client.Option, 0, len(vars))
	for key, value := range vars {
		options = append(options, client.Var(key, value))
	}

	// Call RawPost with options
	rawResponse, err := c.RawPost(getQuery(queryLocation), options...)
	require.Nil(t, err)
	assertRawResponseSuccess(t, rawResponse, err)
	return
}

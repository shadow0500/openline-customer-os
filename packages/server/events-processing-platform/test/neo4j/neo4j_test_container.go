package neo4j

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

const username = "neo4j"
const password = "new-s3cr3t"

func startContainer(ctx context.Context, username, password string) (testcontainers.Container, error) {
	request := testcontainers.ContainerRequest{
		Image:        "neo4j:5-community",
		ExposedPorts: []string{"7687/tcp"},
		Env:          map[string]string{"NEO4J_AUTH": fmt.Sprintf("%s/%s", username, password)},
		WaitingFor:   wait.ForLog("Bolt enabled").WithStartupTimeout(300 * time.Second),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}

func InitTestNeo4jDB() (testcontainers.Container, *neo4j.DriverWithContext) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	neo4jContainer, err := startContainer(ctxWithTimeout, username, password)
	if err != nil {
		log.Panic(err)
	}
	port, err := neo4jContainer.MappedPort(ctxWithTimeout, "7687")
	if err != nil {
		log.Panic(err)
	}
	address := fmt.Sprintf("bolt://localhost:%d", port.Int())
	driver, err := neo4j.NewDriverWithContext(address, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Panic(err)
	}
	return neo4jContainer, &driver
}

func CloseDriver(driver neo4j.DriverWithContext) {
	err := driver.Close(context.Background())
	if err != nil {
		log.Panic("Neo4j driver should close")
	}
}

func Terminate(container testcontainers.Container, ctx context.Context) {
	err := container.Terminate(ctx)
	if err != nil {
		log.Fatal("Container should stop")
	}
}

type TestDatabase struct {
	Neo4jContainer testcontainers.Container
	Driver         *neo4j.DriverWithContext
	Repositories   *repository.Repositories
}

func SetupTestDatabase() (TestDatabase, func()) {
	database := TestDatabase{}
	database.Neo4jContainer, database.Driver = InitTestNeo4jDB()

	appLogger := logger.NewAppLogger(&logger.Config{
		DevMode: true,
	})
	appLogger.InitLogger()
	database.Repositories = repository.InitRepos(database.Driver)

	shutdown := func() {
		CloseDriver(*database.Driver)
		Terminate(database.Neo4jContainer, context.Background())
	}
	return database, shutdown
}

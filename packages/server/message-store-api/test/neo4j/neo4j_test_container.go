package neo4j

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"log"
)

const username = "neo4j"
const password = "s3cr3t"

func startContainer(ctx context.Context, username, password string) (testcontainers.Container, error) {
	request := testcontainers.ContainerRequest{
		Image:        "neo4j:5.2.0-community",
		ExposedPorts: []string{"7687/tcp"},
		Env:          map[string]string{"NEO4J_AUTH": fmt.Sprintf("%s/%s", username, password)},
		WaitingFor:   wait.ForLog("Bolt enabled"),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}

func InitTestNeo4jDB() (testcontainers.Container, *neo4j.Driver) {
	var ctx = context.Background()
	var err error
	neo4jContainer, err := startContainer(ctx, username, password)
	if err != nil {
		log.Panic("Container should start")
	}
	port, err := neo4jContainer.MappedPort(ctx, "7687")
	if err != nil {
		log.Panic("Port should be resolved")
	}
	address := fmt.Sprintf("bolt://localhost:%d", port.Int())
	driver, err := neo4j.NewDriver(address, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Panic("Driver should be created")
	}
	return neo4jContainer, &driver
}

func Close(closer io.Closer, resourceName string) {
	err := closer.Close()
	if err != nil {
		log.Panicf("%s should close", resourceName)
	}
}

func Terminate(container testcontainers.Container, ctx context.Context) {
	err := container.Terminate(ctx)
	if err != nil {
		log.Fatal("Container should stop")
	}
}

package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type userRepository struct {
	driver *neo4j.DriverWithContext
}

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (string, string, []string, error)
}

func NewUserRepository(driver *neo4j.DriverWithContext) UserRepository {
	return &userRepository{
		driver: driver,
	}
}
func (u *userRepository) toStringList(values []interface{}) []string {
	var result []string
	for _, value := range values {
		result = append(result, value.(string))
	}
	return result
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (string, string, []string, error) {
	session := (*u.driver).NewSession(
		ctx,
		neo4j.SessionConfig{
			AccessMode: neo4j.AccessModeRead,
			BoltLogger: neo4j.ConsoleBoltLogger(),
		},
	)
	defer session.Close(ctx)

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, `
			MATCH (e:Email {email:$email})<-[:HAS]-(u:User)-[:USER_BELONGS_TO_TENANT]->(t:Tenant)
			RETURN t.name, u.id, u.roles`,
			map[string]interface{}{
				"email": email,
			})
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)
	})
	if err != nil {
		return "", "", []string{}, err
	}
	if len(records.([]*neo4j.Record)) > 0 {
		tenant := records.([]*neo4j.Record)[0].Values[0].(string)
		userId := records.([]*neo4j.Record)[0].Values[1].(string)
		roleList, ok := records.([]*neo4j.Record)[0].Values[2].([]interface{})
		var roles []string
		if !ok {
			roles = []string{}
		} else {
			roles = u.toStringList(roleList)
		}
		return userId, tenant, roles, nil
	} else {
		return "", "", []string{}, nil
	}
}

package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/customer-os-api/utils"
)

type FieldSetDefinitionRepository interface {
	createFieldSetDefinitionInTx(entityDefId string, entity *entity.FieldSetDefinitionEntity, tx neo4j.Transaction) error
}

type fieldSetDefinitionRepository struct {
	driver *neo4j.Driver
	repos  *RepositoryContainer
}

func NewFieldSetDefinitionRepository(driver *neo4j.Driver, repos *RepositoryContainer) FieldSetDefinitionRepository {
	return &fieldSetDefinitionRepository{
		driver: driver,
		repos:  repos,
	}
}

func (r *fieldSetDefinitionRepository) createFieldSetDefinitionInTx(entityDefId string, entity *entity.FieldSetDefinitionEntity, tx neo4j.Transaction) error {
	queryResult, err := tx.Run(`
			MATCH (e:EntityDefinition {id:$entityDefinitionId})
			MERGE (e)-[:CONTAINS]->(f:FieldSetDefinition {
				id: randomUUID(),
				name: $name
			}) ON CREATE SET f.order=$order
			RETURN f`,
		map[string]any{
			"entityDefinitionId": entityDefId,
			"name":               entity.Name,
			"order":              entity.Order,
		})

	record, err := queryResult.Single()
	if err != nil {
		return err
	}
	fieldSetDefinitionId := utils.GetPropsFromNode(record.Values[0].(dbtype.Node))["id"].(string)
	for _, v := range entity.CustomFields {
		err := r.repos.CustomFieldDefinitionRepository.createCustomFieldDefinitionForFieldSetInTx(fieldSetDefinitionId, v, tx)
		if err != nil {
			return err
		}
	}
	return nil
}

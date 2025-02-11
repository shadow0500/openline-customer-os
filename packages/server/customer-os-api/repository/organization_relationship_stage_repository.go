package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type OrganizationRelationshipRepository interface {
	GetOrganizationRelationshipsForOrganizations(ctx context.Context, tenant string, organizationIds []string) ([]*utils.DbNodeAndId, error)
	GetOrganizationRelationshipsWithStagesForOrganizations(ctx context.Context, tenant string, organizationIds []string) ([]*utils.DbNodePairAndId, error)
	CreateDefaultStagesForNewTenant(ctx context.Context, tenant string) error
}

type organizationRelationshipRepository struct {
	driver *neo4j.DriverWithContext
}

func NewOrganizationRelationshipRepository(driver *neo4j.DriverWithContext) OrganizationRelationshipRepository {
	return &organizationRelationshipRepository{
		driver: driver,
	}
}

func (r *organizationRelationshipRepository) GetOrganizationRelationshipsForOrganizations(ctx context.Context, tenant string, organizationIds []string) ([]*utils.DbNodeAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationRelationshipRepository.GetOrganizationRelationshipsForOrganizations")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization)-[:IS]->(or:OrganizationRelationship)
			WHERE org.id IN $organizationIds
			RETURN or, org.id order by or.name`

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, query,
			map[string]any{
				"tenant":          tenant,
				"organizationIds": organizationIds,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeAndId(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeAndId), err
}

func (r *organizationRelationshipRepository) GetOrganizationRelationshipsWithStagesForOrganizations(ctx context.Context, tenant string, organizationIds []string) ([]*utils.DbNodePairAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationRelationshipRepository.GetOrganizationRelationshipsForOrganizations")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization)-[:IS]->(or:OrganizationRelationship)
			WHERE org.id IN $organizationIds
			OPTIONAL MATCH (or)-[:HAS_STAGE]->(ors:OrganizationRelationshipStage)<-[:HAS_STAGE]-(org)
			RETURN or, ors, org.id order by or.name`

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, query,
			map[string]any{
				"tenant":          tenant,
				"organizationIds": organizationIds,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodePairAndId(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodePairAndId), err
}

func (r *organizationRelationshipRepository) CreateDefaultStagesForNewTenant(ctx context.Context, tenant string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationRelationshipRepository.CreateDefaultStagesForNewTenant")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := fmt.Sprintf(`WITH [{name:"Target",order:10},
					{name:"Lead",order:20},
					{name:"Prospect",order:30},
					{name:"Trial",order:40},
					{name:"Lost",order:50},
					{name:"Live",order:60},
					{name:"Former",order:70}] AS stages
				UNWIND stages AS stage
				MATCH (t:Tenant {name:$tenant}), (or:OrganizationRelationship)
				MERGE (t)<-[:STAGE_BELONGS_TO_TENANT]-(s:OrganizationRelationshipStage {name:stage.name})<-[:HAS_STAGE]-(or)
				ON CREATE SET 	s.id=randomUUID(), 
								s.order=stage.order,
								s.createdAt=$now, 
								s:OrganizationRelationshipStage_%s`, tenant)

	span.LogFields(log.String("query", query))

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query,
			map[string]any{
				"tenant": tenant,
				"now":    utils.Now(),
			})
		return nil, err
	})
	return err
}

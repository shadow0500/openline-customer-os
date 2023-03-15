package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/utils"
	"golang.org/x/net/context"
	"time"
)

type InteractionEventRepository interface {
	MergeInteractionSession(ctx context.Context, tenant string, date time.Time, message entity.EmailMessageData) (string, error)
	MergeInteractionEvent(ctx context.Context, tenant, externalSystemId string, date time.Time, message entity.EmailMessageData) (string, error)
	MergeInteractionEventToSession(ctx context.Context, tenant, interactionEventId, interactionSessionId string) error

	InteractionEventSentByEmail(ctx context.Context, tenant, interactionEventId, emailId string) error

	InteractionEventSentToEmails(ctx context.Context, tenant, interactionEventId, sentType string, emails []string) error
}

type interactionEventRepository struct {
	driver *neo4j.DriverWithContext
}

func NewInteractionEventRepository(driver *neo4j.DriverWithContext) InteractionEventRepository {
	return &interactionEventRepository{
		driver: driver,
	}
}

func (r *interactionEventRepository) MergeInteractionSession(ctx context.Context, tenant string, syncDate time.Time, message entity.EmailMessageData) (string, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MERGE (is:InteractionSession_%s {identifier:$identifier, source:$source, channel:$channel}) " +
		" ON CREATE SET " +
		"  is:InteractionSession, " +
		"  is.id=randomUUID(), " +
		"  is.syncDate=$syncDate, " +
		"  is.createdAt=$createdAt, " +
		"  is.name=$name, " +
		"  is.status=$status," +
		"  is.type=$type," +
		"  is.sourceOfTruth=$sourceOfTruth, " +
		"  is.appSource=$appSource " +
		" WITH is " +
		" RETURN is.id"

	dbRecord, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":        tenant,
				"source":        message.ExternalSystem,
				"sourceOfTruth": message.ExternalSystem,
				"appSource":     message.ExternalSystem,
				"identifier":    message.EmailThreadId,
				"name":          message.Subject,
				"syncDate":      syncDate,
				"createdAt":     message.CreatedAt,
				"status":        "ACTIVE",
				"type":          "THREAD",
				"channel":       "EMAIL",
			})
		if err != nil {
			return nil, err
		}
		record, err := queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record, nil
	})
	if err != nil {
		return "", err
	}
	return dbRecord.(*db.Record).Values[0].(string), nil
}

func (r *interactionEventRepository) MergeInteractionEvent(ctx context.Context, tenant, externalSystemId string, syncDate time.Time, message entity.EmailMessageData) (string, error) {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystemId}) " +
		" MERGE (ie:InteractionEvent_%s {externalId:$externalId, source:$source, channel:$channel}) " +
		" MERGE (ie)-[:IS_LINKED_WITH {externalId:$externalId, syncDate: $syncDate}]->(e)" +
		" ON CREATE SET " +
		"  ie:InteractionEvent, " +
		"  ie:Action, " +
		"  ie:Action_%s, " +
		//"  ie:TimelineEvent, " +
		//"  ie:TimelineEvent_%s, " +
		"  ie.createdAt=$createdAt, " +
		"  ie.id=randomUUID(), " +
		"  ie.identifier=$identifier, " +
		"  ie.content=$content, " +
		"  ie.contentType=$contentType, " +
		"  ie.sourceOfTruth=$sourceOfTruth, " +
		"  ie.appSource=$appSource " +
		" WITH ie " +
		" RETURN ie.id"

	dbRecord, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		params := map[string]interface{}{
			"tenant":           tenant,
			"externalSystemId": externalSystemId,
			"identifier":       message.EmailMessageId,
			"source":           message.ExternalSystem,
			"sourceOfTruth":    message.ExternalSystem,
			"appSource":        message.ExternalSystem,
			"externalId":       message.ExternalId,
			"syncDate":         syncDate,
			"createdAt":        message.CreatedAt,
			"channel":          "EMAIL",
		}

		if message.Html != "" {
			params["content"] = message.Html
			params["contentType"] = "text/html"
		} else {
			params["content"] = message.Text
			params["contentType"] = "text/plain"
		}

		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			params)
		if err != nil {
			return nil, err
		}
		record, err := queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record, nil
	})
	if err != nil {
		return "", err
	}

	return dbRecord.(*db.Record).Values[0].(string), nil
}

func (r *interactionEventRepository) MergeInteractionEventToSession(ctx context.Context, tenant, interactionEventId, interactionSessionId string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (is:InteractionSession_%s {id:$interactionSessionId}) " +
		" MATCH (ie:InteractionEvent {id:$interactionEventId})" +
		" MERGE (ie)-[:PART_OF]->(is) "
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":               tenant,
				"interactionSessionId": interactionSessionId,
				"interactionEventId":   interactionEventId,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) InteractionEventSentByEmail(ctx context.Context, tenant, interactionEventId, emailId string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (is:InteractionEvent_%s {id:$interactionEventId}) " +
		" MATCH (e:Email_%s {id: $emailId}) " +
		" MERGE (is)-[:SENT_BY]->(e) "
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": interactionEventId,
				"emailId":            emailId,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) InteractionEventSentToEmails(ctx context.Context, tenant, interactionEventId, sentType string, emails []string) error {
	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) " +
		" MATCH (e:Email_%s) WHERE e.rawEmail in $emails " +
		" MERGE (ie)-[:SENT_TO {type: $sentType}]->(e) "
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": interactionEventId,
				"sentType":           sentType,
				"emails":             emails,
			})
		return nil, err
	})

	return err
}

package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"time"
)

type SendDirection string

const (
	SENT_TO SendDirection = "SENT_TO"
	SENT_BY SendDirection = "SENT_BY"
)

type PartOfType string

const (
	PART_OF_INTERACTION_SESSION PartOfType = "InteractionSession"
	PART_OF_MEETING             PartOfType = "Meeting"
)

type InteractionEventRepository interface {
	GetAllForInteractionSessions(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error)
	GetAllForMeetings(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error)
	GetAllForIssues(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error)
	GetSentByParticipantsForInteractionEvents(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeWithRelationAndId, error)
	GetSentToParticipantsForInteractionEvents(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeWithRelationAndId, error)
	GetReplyToInteractionEventsForInteractionEvents(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error)

	LinkWithPartOfXXInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, interactionEventId string, partOfId string, partOfType PartOfType) error
	LinkWithRepliesToInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, interactionEventId, repliesToEventId string) error

	LinkWithSentXXParticipantInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, entityType entity.EntityType, interactionEventId, participantId string, sentType *string, direction SendDirection) error
	LinkWithSentXXEmailInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, interactionEventId, email string, sentType *string, direction SendDirection) error
	LinkWithSentXXPhoneNumberInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, interactionEventId, e164 string, sentType *string, direction SendDirection) error

	Create(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, newInteractionEvent entity.InteractionEventEntity, source, sourceOfTruth entity.DataSource) (*dbtype.Node, error)
}

type interactionEventRepository struct {
	driver *neo4j.DriverWithContext
}

func NewInteractionEventRepository(driver *neo4j.DriverWithContext) InteractionEventRepository {
	return &interactionEventRepository{
		driver: driver,
	}
}

func (r *interactionEventRepository) LinkWithSentXXEmailInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, interactionEventId, email string, sentType *string, direction SendDirection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkWithSentXXEmailInTx")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (e:Email_%s) `, tenant)
	query += fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$eventId}) `, tenant)
	query += `WHERE e.email = $email OR e.rawEmail = $email `

	if direction == SENT_TO {
		if sentType != nil {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_TO {type:$sentType}]->(e) RETURN r`)
		} else {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_TO]->(e) RETURN r`)
		}
	} else {
		if sentType != nil {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_BY {type:$sentType}]->(e) RETURN r`)
		} else {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_BY]->(e) RETURN r`)
		}
	}
	queryResult, err := tx.Run(ctx, query,
		map[string]any{
			"email":    email,
			"eventId":  interactionEventId,
			"sentType": sentType,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}

func (r *interactionEventRepository) LinkWithSentXXPhoneNumberInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, interactionEventId, e164 string, sentType *string, direction SendDirection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkWithSentXXPhoneNumberInTx")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (p:PhoneNumber_%s) `, tenant)
	query += fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$eventId}) `, tenant)
	query += `WHERE p.e164 = $e164 OR p.rawPhoneNumber = $e164 `

	if direction == SENT_TO {
		if sentType != nil {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_TO {type:$sentType}]->(p) RETURN r`)
		} else {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_TO]->(p) RETURN r`)
		}
	} else {
		if sentType != nil {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_BY {type:$sentType}]->(p) RETURN r`)
		} else {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_BY]->(p) RETURN r`)
		}
	}
	queryResult, err := tx.Run(ctx, query,
		map[string]any{
			"e164":    e164,
			"eventId": interactionEventId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}

func (r *interactionEventRepository) LinkWithSentXXParticipantInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, entityType entity.EntityType, interactionEventId, participantId string, sentType *string, direction SendDirection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkWithSentXXParticipantInTx")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := ""
	switch entityType {
	case entity.CONTACT:
		query = fmt.Sprintf(`MATCH (p:Contact_%s {id:$participantId}) `, tenant)
	case entity.USER:
		query = fmt.Sprintf(`MATCH (p:User_%s {id:$participantId}) `, tenant)
	}
	query += fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$eventId}) `, tenant)

	if direction == SENT_TO {
		if sentType != nil {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_TO {type:$sentType}]->(p) RETURN r`)
		} else {
			query += fmt.Sprintf(`MERGE (ie)-[r:SENT_TO]->(p) RETURN r`)
		}
	} else {
		if sentType != nil {
			query += fmt.Sprintf(`MERGE (ie)<-[r:SENT_BY {type:$sentType}]-(p) RETURN r`)
		} else {
			query += fmt.Sprintf(`MERGE (ie)<-[r:SENT_BY]-(p) RETURN r`)
		}
	}
	queryResult, err := tx.Run(ctx, query,
		map[string]any{
			"participantId": participantId,
			"eventId":       interactionEventId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}

func (r *interactionEventRepository) LinkWithRepliesToInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant, interactionEventId, repliesToEventId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkWithRepliesToInTx")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	queryResult, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (ie:InteractionEvent_%s {id:$eventId})
			MATCH (rie:InteractionEvent_%s {id:$repliesToEventId})
			MERGE (ie)-[r:REPLIES_TO]->(rie)
			RETURN r`, tenant, tenant),
		map[string]any{
			"eventId":          interactionEventId,
			"repliesToEventId": repliesToEventId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}

func (r *interactionEventRepository) LinkWithPartOfXXInTx(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, interactionEventId string, partOfId string, partOfType PartOfType) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkWithPartOfXXInTx")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	queryResult, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (ie:InteractionEvent_%s {id:$eventId})
			MATCH (is:%s_%s {id:$interactionSessionId})
			MERGE (ie)-[r:PART_OF]->(is)
			RETURN r`, tenant, partOfType, tenant),
		map[string]any{
			"eventId":              interactionEventId,
			"interactionSessionId": partOfId,
		})
	if err != nil {
		return err
	}
	_, err = queryResult.Single(ctx)
	return err
}

func (r *interactionEventRepository) GetAllForInteractionSessions(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetAllForInteractionSessions")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (s:InteractionSession_%s)<-[:PART_OF]-(e:InteractionEvent) " +
		" WHERE s.id IN $ids " +
		" RETURN e, s.id ORDER BY e.createdAt ASC"

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]any{
				"tenant": tenant,
				"ids":    ids,
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

func (r *interactionEventRepository) GetAllForMeetings(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetAllForMeetings")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (m:Meeting_%s)<-[:PART_OF]-(e:InteractionEvent) " +
		" WHERE m.id IN $ids " +
		" RETURN e, m.id ORDER BY e.createdAt ASC"

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]any{
				"tenant": tenant,
				"ids":    ids,
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

func (r *interactionEventRepository) GetAllForIssues(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetAllForIssues")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (i:Issue_%s)<-[:PART_OF]-(e:InteractionEvent) " +
		" WHERE i.id IN $ids " +
		" RETURN e, i.id ORDER BY e.createdAt ASC"

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]any{
				"tenant": tenant,
				"ids":    ids,
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

func (r *interactionEventRepository) GetSentByParticipantsForInteractionEvents(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeWithRelationAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetSentByParticipantsForInteractionEvents")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (ie:InteractionEvent_%s)-[rel:SENT_BY]->(p) " +
		" WHERE ie.id IN $ids " +
		" RETURN p, rel, ie.id"

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]any{
				"tenant": tenant,
				"ids":    ids,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeWithRelationAndId(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeWithRelationAndId), err
}

func (r *interactionEventRepository) Create(ctx context.Context, tx neo4j.ManagedTransaction, tenant string, newInteractionEvent entity.InteractionEventEntity, source, sourceOfTruth entity.DataSource) (*dbtype.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.Create")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	var createdAt time.Time
	createdAt = utils.Now()
	if newInteractionEvent.CreatedAt != nil {
		createdAt = *newInteractionEvent.CreatedAt
	}

	query := "MERGE (ie:InteractionEvent_%s {id:randomUUID()}) ON CREATE SET " +
		"  ie:InteractionEvent, " +
		"  ie:TimelineEvent, " +
		"  ie:TimelineEvent_%s, " +
		" ie.source=$source, " +
		" ie.channel=$channel, " +
		" ie.channelData=$channelData, " +
		" ie.createdAt=$createdAt, " +
		" ie.identifier=$identifier, " +
		" ie.content=$content, " +
		" ie.contentType=$contentType, " +
		" ie.eventType=$eventType, " +
		" ie.sourceOfTruth=$sourceOfTruth, " +
		" ie.appSource=$appSource " +
		" RETURN ie"

	if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
		map[string]interface{}{
			"tenant":        tenant,
			"source":        source,
			"channel":       newInteractionEvent.Channel,
			"channelData":   newInteractionEvent.ChannelData,
			"createdAt":     createdAt,
			"identifier":    newInteractionEvent.EventIdentifier,
			"content":       newInteractionEvent.Content,
			"contentType":   newInteractionEvent.ContentType,
			"eventType":     newInteractionEvent.EventType,
			"sourceOfTruth": sourceOfTruth,
			"appSource":     newInteractionEvent.AppSource,
		}); err != nil {
		return nil, err
	} else {
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}
}

func (r *interactionEventRepository) GetReplyToInteractionEventsForInteractionEvents(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetReplyToInteractionEventsForInteractionEvents")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (ie:InteractionEvent_%s)-[rel:REPLIES_TO]->(rie:InteractionEvent_%s) " +
		" WHERE ie.id IN $ids " +
		" RETURN rie, ie.id"

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			map[string]any{
				"tenant": tenant,
				"ids":    ids,
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

func (r *interactionEventRepository) GetSentToParticipantsForInteractionEvents(ctx context.Context, tenant string, ids []string) ([]*utils.DbNodeWithRelationAndId, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetSentToParticipantsForInteractionEvents")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (ie:InteractionEvent_%s)-[rel:SENT_TO]->(p) " +
		" WHERE ie.id IN $ids " +
		" RETURN p, rel, ie.id"

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]any{
				"tenant": tenant,
				"ids":    ids,
			}); err != nil {
			return nil, err
		} else {
			return utils.ExtractAllRecordsAsDbNodeWithRelationAndId(ctx, queryResult, err)
		}
	})
	if err != nil {
		return nil, err
	}
	return result.([]*utils.DbNodeWithRelationAndId), err
}

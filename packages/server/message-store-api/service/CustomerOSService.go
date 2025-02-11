package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	commonModuleService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/config"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/repository/entity"
	"log"
	"time"
)

type CustomerOSService struct {
	driver               *neo4j.DriverWithContext
	postgresRepositories *repository.PostgresRepositories
	commonStoreService   *commonStoreService
	graphqlClient        *graphql.Client
	conf                 *config.Config
}

type EmailContent struct {
	MessageId string   `json:"messageId"`
	Html      string   `json:"html"`
	Subject   string   `json:"subject"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Cc        []string `json:"cc"`
	Bcc       []string `json:"bcc"`
	InReplyTo []string `json:"InReplyTo"`
	Reference []string `json:"Reference"`
}

type CustomerOSServiceInterface interface {
	ContactByPhoneExists(e164 string) (bool, error)

	GetUserByEmail(email string) (*User, error)
	GetContactById(id string) (*Contact, error)
	GetContactByEmail(email string) (*Contact, error)
	GetContactByPhone(phoneNumber string) (*Contact, error)

	CreateContactWithEmail(tenant string, email string) (*Contact, error)
	CreateContactWithPhone(tenant string, phoneNumber string) (*Contact, error)

	ConversationByIdExists(tenant string, conversationId string) (bool, error)

	GetConversations(tenant string) ([]Conversation, error)
	GetConversationById(tenant string, conversationId string) (Conversation, error)

	CreateConversation(tenant string, initiatorId string, initiatorFirstName string, initiatorLastName string, initiatorUsername string, initiatorType entity.SenderType, channel entity.EventType) (*Conversation, error)
	UpdateConversation(tenant string, conversationId string, participantId string, participantType entity.SenderType, lastSenderFirstName string, lastSenderLastName string, lastContentPreview string) (string, error)
}

type Conversation struct {
	Id        string
	StartedAt time.Time
	UpdatedAt time.Time
	Channel   string
	Status    string

	InitiatorFirstName  string
	InitiatorLastName   string
	InitiatorUsername   string
	InitiatorType       string
	LastSenderId        string
	LastSenderType      string
	LastSenderFirstName string
	LastSenderLastName  string
	LastContentPreview  string
}

type emailObject struct {
	Email   string  `json:"email"`
	Primary bool    `json:"primary"`
	Label   *string `json:"label"`
}

type phoneNumberObject struct {
	E164    string  `json:"e164"`
	Primary bool    `json:"primary"`
	Label   *string `json:"label"`
}

type Contact struct {
	FirstName    string              `json:"firstName"`
	LastName     string              `json:"lastName"`
	Id           string              `json:"id"`
	Emails       []emailObject       `json:"emails"`
	PhoneNumbers []phoneNumberObject `json:"phoneNumbers"`
}

type User struct {
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Id        string        `json:"id"`
	Emails    []emailObject `json:"emails"`
}

const contactFieldSelection = `firstName,lastName,id,
    				emails {
  					  email
                      primary
					  label
            		},
					phoneNumbers {
						  e164,
						  primary,
						  label
            		}`

const userFieldSelection = `firstName,lastName,id,
    				emails {
  					  email
                      primary
					  label
            		}`

func (s *CustomerOSService) addHeadersToGraphRequest(req *graphql.Request, ctx context.Context, tenant string) error {
	req.Header.Add("X-Openline-API-KEY", s.conf.Service.CustomerOsAPIKey)
	user, err := commonModuleService.GetUsernameMetadataForGRPC(ctx)
	if err != nil && user != nil {
		req.Header.Add("X-Openline-USERNAME", *user)
	}

	req.Header.Add("X-Openline-TENANT", tenant)
	return nil
}

func (s *CustomerOSService) GetUserByEmail(ctx context.Context, email string, tenant string) (*User, error) {
	graphqlRequest := graphql.NewRequest(`
  				query ($email: String!) {
  					user_ByEmail(email: $email){` + userFieldSelection + `}
  				}
    `)
	graphqlRequest.Var("email", email)

	err := s.addHeadersToGraphRequest(graphqlRequest, ctx, tenant)

	if err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}

	var graphqlResponse struct {
		UserByEmail *User `json:"user_ByEmail"`
	}

	if err = s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}
	bytes, _ := json.Marshal(graphqlResponse)
	log.Printf("GetUserByEmail: email=%s graphqlResponse = %s", email, bytes)
	return graphqlResponse.UserByEmail, nil
}

func (s *CustomerOSService) GetContactById(ctx context.Context, id string, tenant string) (*Contact, error) {
	graphqlRequest := graphql.NewRequest(`
  				query ($id: ID!) {
  					contact(id: $id){` + contactFieldSelection + `}
  				}
    `)
	graphqlRequest.Var("id", id)

	err := s.addHeadersToGraphRequest(graphqlRequest, ctx, tenant)
	if err != nil {
		return nil, fmt.Errorf("GetContactById: %w", err)
	}

	var graphqlResponse struct {
		Contact *Contact `json:"contact"`
	}

	if err = s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetContactById: %w", err)
	}

	bytes, _ := json.Marshal(graphqlResponse)
	log.Printf("GetContactById: id=%s graphqlResponse = %s", id, bytes)
	return graphqlResponse.Contact, nil
}

func (s *CustomerOSService) GetContactByEmail(ctx context.Context, email string, tenant string) (*Contact, error) {
	graphqlRequest := graphql.NewRequest(`
  				query ($email: String!) {
  					contact_ByEmail(email: $email){` + contactFieldSelection + `}
  				}
    `)
	graphqlRequest.Var("email", email)

	err := s.addHeadersToGraphRequest(graphqlRequest, ctx, tenant)
	if err != nil {
		return nil, fmt.Errorf("GetContactByEmail: %w", err)
	}

	var graphqlResponse struct {
		ContactByEmail Contact `json:"contact_ByEmail"`
	}

	if err = s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetContactByEmail: %w", err)
	}
	bytes, _ := json.Marshal(graphqlResponse)
	log.Printf("GetContactByEmail: email=%s graphqlResponse = %s", email, bytes)
	return &graphqlResponse.ContactByEmail, nil
}

func (s *CustomerOSService) GetContactByPhone(ctx context.Context, phoneNumber string, tenant string) (*Contact, error) {
	graphqlRequest := graphql.NewRequest(`
  				query ($phoneNumber: String!) {
  					contact_ByPhone(e164: $phoneNumber){` + contactFieldSelection + `}
  				}
    `)
	graphqlRequest.Var("phoneNumber", phoneNumber)

	err := s.addHeadersToGraphRequest(graphqlRequest, ctx, tenant)
	if err != nil {
		return nil, fmt.Errorf("GetContactByPhone: %w", err)
	}

	var graphqlResponse struct {
		ContactByPhone Contact `json:"contact_ByPhone"`
	}

	if err = s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetContactByPhone: %w", err)
	}
	bytes, _ := json.Marshal(graphqlResponse)
	log.Printf("GetContactByPhone: phoneNumber=%s graphqlResponse = %s", phoneNumber, bytes)
	return &graphqlResponse.ContactByPhone, nil
}

func (s *CustomerOSService) CreateContactWithEmail(ctx context.Context, tenant string, email string) (*Contact, error) {
	graphqlRequest := graphql.NewRequest(`
		mutation CreateContact ($email: String!) {
		  contact_Create(input: {
		  email:{email:  $email, label: WORK}}) {
			id
          }
		}
    `)

	graphqlRequest.Var("email", email)
	err := s.addHeadersToGraphRequest(graphqlRequest, ctx, tenant)

	if err != nil {
		return nil, fmt.Errorf("CreateContactWithEmail: %w", err)
	}

	var graphqlResponse map[string]map[string]string
	if err := s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return nil, err
	}
	id := graphqlResponse["contact_Create"]["id"]
	bytes, _ := json.Marshal(graphqlResponse)
	log.Printf("CreateContactWithEmail: email=%s graphqlResponse = %s", email, bytes)
	return s.GetContactById(ctx, id, tenant)

}

func (s *CustomerOSService) CreateContactWithPhone(ctx context.Context, tenant string, phoneNumber string) (*Contact, error) {
	graphqlRequest := graphql.NewRequest(`
		mutation CreateContact ($phoneNumber: String!) {
		  contact_Create(input: {
		  phoneNumber:{phoneNumber:  $phoneNumber, label: MAIN}}) {
			id
          }
		}
    `)

	graphqlRequest.Var("phoneNumber", phoneNumber)
	err := s.addHeadersToGraphRequest(graphqlRequest, ctx, tenant)

	if err != nil {
		return nil, fmt.Errorf("CreateContactWithPhone: %w", err)
	}

	var graphqlResponse map[string]map[string]string
	if err := s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateContactWithPhone: %w", err)
	}
	id := graphqlResponse["contact_Create"]["id"]
	bytes, _ := json.Marshal(graphqlResponse)
	log.Printf("CreateContactWithPhone: phoneNumber=%s graphqlResponse = %s", phoneNumber, bytes)
	return s.GetContactById(ctx, id, tenant)

}

func (s *CustomerOSService) ConversationByIdExists(ctx context.Context, tenant string, conversationId string) (bool, error) {
	session := utils.NewNeo4jWriteSession(ctx, *s.driver)
	defer session.Close(ctx)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, "MATCH (c:Conversation_"+tenant+"{id:$conversationId}) RETURN c",
			map[string]any{
				"conversationId": conversationId,
			}); err != nil {
			return nil, err
		} else {
			return queryResult.Collect(ctx)
		}
	})

	if err != nil {
		return false, fmt.Errorf("ConversationByIdExists: %w", err)
	}

	if len(dbRecords.([]*db.Record)) == 0 {
		return false, nil
	}

	return true, nil
}

func (s *CustomerOSService) GetConversations(ctx context.Context, tenant string, onlyContacts bool) ([]Conversation, error) {
	session := utils.NewNeo4jReadSession(ctx, *s.driver)
	defer session.Close(ctx)

	dbRecords, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		//todo move order by as param
		cypher := ""
		if onlyContacts {
			//cypher = "match (t:Tenant{name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User)-[:HAS]->(e:Email {email:$user}), (u)-[:PARTICIPATES]->(o:Conversation)<-[:PARTICIPATES]-(c:Contact) return distinct o"
			cypher = "match (t:Tenant{name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User), (u)-[:PARTICIPATES]->(o:Conversation)<-[:PARTICIPATES]-(c:Contact) return distinct o order by o.updatedAt desc"
		} else {
			cypher = "MATCH (c:Conversation_" + tenant + ") RETURN c order by c.updatedAt desc"
		}

		if queryResult, err := tx.Run(ctx, cypher, map[string]any{
			"tenant": tenant,
		}); err != nil {
			return nil, fmt.Errorf("GetConversations: %w", err)
		} else {
			return queryResult.Collect(ctx)
		}
	})

	if err != nil {
		return nil, fmt.Errorf("GetConversations: %w", err)
	}

	var conversations []Conversation
	for _, v := range dbRecords.([]*neo4j.Record) {
		node := v.Values[0].(neo4j.Node)
		conversations = append(conversations, *mapNodeToConversation(&node))
	}

	return conversations, nil
}

func (s *CustomerOSService) GetConversationById(ctx context.Context, tenant string, conversationId string) (*Conversation, error) {
	session := utils.NewNeo4jReadSession(ctx, *s.driver)
	defer session.Close(ctx)

	conversationNode, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		if queryResult, err := tx.Run(ctx, "MATCH (c:Conversation_"+tenant+"{id: $conversationId}) RETURN c", map[string]any{
			"conversationId": conversationId,
		}); err != nil {
			return nil, err
		} else {
			record, err := queryResult.Single(ctx)
			if err != nil {
				return nil, err
			}
			if record == nil {
				return nil, errors.New("conversation not found")
			}

			return utils.NodePtr(record.Values[0].(neo4j.Node)), nil
		}
	})
	if err != nil {
		return nil, fmt.Errorf("GetConversationById: %w", err)
	}

	return mapNodeToConversation(conversationNode.(*dbtype.Node)), nil
}

func (s *CustomerOSService) GetConversationParticipants(ctx context.Context, tenant string, conversationId string) ([]string, error) {
	session := utils.NewNeo4jReadSession(ctx, *s.driver)
	defer session.Close(ctx)

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, "MATCH (c:Conversation_"+tenant+"{id:$conversationId})<-[PARTICIPATES]-(p)-[HAS]->(e:Email) WITH COALESCE(e.email, e.rawEmail) AS email WHERE email IS NOT NULL RETURN DISTINCT email",
			map[string]interface{}{
				"conversationId": conversationId,
			})
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)
	})
	if err != nil {
		return []string{}, fmt.Errorf("GetConversationParticipants: %w", err)
	}
	emails := make([]string, 0)
	if len(records.([]*neo4j.Record)) > 0 {
		for _, record := range records.([]*neo4j.Record) {
			if record != nil {
				if len(record.Values) > 0 {
					val, ok := record.Values[0].(string)
					if ok {
						emails = append(emails, val)
					}
				}
			}
		}
		return emails, nil
	} else {
		return []string{}, nil
	}
}

func (s *CustomerOSService) GetConversationParticipantsIds(ctx context.Context, tenant string, conversationId string) ([]Participant, error) {
	session := utils.NewNeo4jReadSession(ctx, *s.driver)
	defer session.Close(ctx)

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, "MATCH (c:Conversation_"+tenant+"{id:$conversationId})<-[PARTICIPATES]-(p) RETURN DISTINCT labels(p), p.id",
			map[string]interface{}{
				"conversationId": conversationId,
			})
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)
	})
	if err != nil {
		return []Participant{}, fmt.Errorf("GetConversationParticipantsIds: %w", err)
	}
	participants := make([]Participant, 0)
	if len(records.([]*neo4j.Record)) > 0 {
		for _, record := range records.([]*neo4j.Record) {
			if record != nil {
				if len(record.Values) > 0 {
					p := Participant{}
					val, ok := record.Values[0].([]any)
					if ok {
						detected := false
						for _, v := range val {
							strv, ok := v.(string)
							if ok {
								if strv == "Contact" {
									p.Type = entity.CONTACT
									detected = true
								} else if strv == "User" {
									p.Type = entity.USER
									detected = true
								}
							}
						}
						if !detected {
							log.Printf("GetConversationParticipantsIds: Error trying to get record type, skipping")
							continue
						}
					} else {
						log.Printf("GetConversationParticipantsIds: Error trying to get record type, skipping")
						continue
					}
					valId, ok := record.Values[1].(string)
					if ok {
						p.Id = valId
					} else {
						log.Printf("GetConversationParticipantsIds: Error trying to get record id, skipping")
						continue
					}
					participants = append(participants, p)

				}
			}
		}
		return participants, nil
	} else {
		return []Participant{}, nil
	}
}

func (s *CustomerOSService) CreateConversation(ctx context.Context, tenant string, initiator Participant, initiatorUsername string, channel entity.EventType, threadId string) (*Conversation, error) {
	session := utils.NewNeo4jWriteSession(ctx, *s.driver)
	defer session.Close(ctx)

	contactIds := []string{}
	userIds := []string{}
	initiatorTypeStr := ""

	if initiator.Type == entity.CONTACT {
		contactIds = append(contactIds, initiator.Id)
		initiatorTypeStr = "CONTACT"
	} else if initiator.Type == entity.USER {
		userIds = append(userIds, initiator.Id)
		initiatorTypeStr = "USER"
	}

	if result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "MATCH (t:Tenant {name:$tenant}) " +
			" MERGE (o:Conversation {id:randomUUID()}) " +
			" ON CREATE SET o.startedAt=$startedAt, " +
			"				o.updatedAt=$updatedAt, " +
			"				o.threadId=$threadId, " +
			" 				o.messageCount=0, " +
			"				o.channel=$channel, " +
			"				o.status=$status, " +
			"				o.source=$source, " +
			"				o.sourceOfTruth=$sourceOfTruth, " +
			" 				o.initiatorFirstName=$initiatorFirstName, " +
			"				o.initiatorLastName=$initiatorLastName, " +
			"				o.initiatorUsername=$initiatorUsername, " +
			"				o.initiatorType=$initiatorType, " +
			" 				o.source=$source, " +
			"				o.sourceOfTruth=$sourceOfTruth, " +
			" 				o.lastSenderId=$lastSenderId, " +
			"				o.lastSenderType=$lastSenderType, " +
			"				o.lastSenderFirstName=$lastSenderFirstName, " +
			"				o.lastSenderLastName=$lastSenderLastName, " +
			"				o.lastContentPreview=$lastContentPreview, " +
			" 				o.appSource=$appSource, " +
			"				o:Conversation_%s," +
			"				o:TimelineEvent," +
			"				o:TimelineEvent_%s " +
			" %s %s " +
			" RETURN DISTINCT o"
		queryLinkWithContacts := ""
		if len(contactIds) > 0 {
			queryLinkWithContacts = " WITH DISTINCT t, o " +
				" OPTIONAL MATCH (c:Contact)-[:CONTACT_BELONGS_TO_TENANT]->(t) WHERE c.id in $contactIds " +
				" MERGE (c)-[:PARTICIPATES]->(o) " +
				" WITH DISTINCT t, o " +
				" OPTIONAL MATCH (c:Contact)-[:CONTACT_BELONGS_TO_TENANT]->(t) WHERE c.id in $contactIds " +
				" MERGE (c)-[:INITIATED]->(o) "
		}
		queryLinkWithUsers := ""
		if len(userIds) > 0 {
			queryLinkWithUsers = " WITH DISTINCT t, o " +
				" OPTIONAL MATCH (u:User)-[:USER_BELONGS_TO_TENANT]->(t) WHERE u.id in $userIds " +
				" MERGE (u)-[:PARTICIPATES]->(o) " +
				" WITH DISTINCT t, o " +
				" OPTIONAL MATCH (u:User)-[:USER_BELONGS_TO_TENANT]->(t) WHERE u.id in $userIds " +
				" MERGE (u)-[:INITIATED]->(o) "
		}
		utc := time.Now().UTC()
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant, queryLinkWithContacts, queryLinkWithUsers),
			map[string]interface{}{
				"tenant":              tenant,
				"source":              "openline",
				"sourceOfTruth":       "openline",
				"appSource":           "manual",
				"status":              "ACTIVE",
				"initiatorFirstName":  "",
				"initiatorLastName":   "",
				"initiatorUsername":   initiatorUsername,
				"initiatorType":       initiatorTypeStr,
				"startedAt":           utc,
				"updatedAt":           utc,
				"threadId":            threadId,
				"channel":             channel,
				"contactIds":          contactIds,
				"userIds":             userIds,
				"lastSenderId":        "",
				"lastSenderType":      "",
				"lastSenderFirstName": "",
				"lastSenderLastName":  "",
				"lastContentPreview":  "",
			})

		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}); err != nil {
		return nil, fmt.Errorf("CreateConversation: %w", err)
	} else {
		dbNode := result.(*dbtype.Node)
		return mapNodeToConversation(dbNode), nil
	}
}

func (s *CustomerOSService) UpdateConversation(ctx context.Context, tenant string, conversationId string, lastSenderId string, lastSenderType string, contactIds []string, userIds []string, lastContentPreview string) (string, error) {
	session := utils.NewNeo4jWriteSession(ctx, *s.driver)
	defer session.Close(ctx)

	if result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "MATCH (t:Tenant {name:$tenant}) " +
			" MATCH (o:Conversation_%s {id:$conversationId}) " +
			" SET " +
			" 	o.messageCount=o.messageCount+1, " +
			"	o.updatedAt=$updatedAt, " +
			" 	o.lastSenderId=$lastSenderId, " +
			"	o.lastSenderType=$lastSenderType, " +
			"	o.lastSenderFirstName=$lastSenderFirstName, " +
			"	o.lastSenderLastName=$lastSenderLastName, " +
			"	o.lastContentPreview=$lastContentPreview " +
			" %s %s " +
			" RETURN DISTINCT o"
		queryLinkWithContacts := ""
		if len(contactIds) > 0 {
			queryLinkWithContacts = " WITH DISTINCT t, o " +
				" OPTIONAL MATCH (c:Contact)-[:CONTACT_BELONGS_TO_TENANT]->(t) WHERE c.id in $contactIds " +
				" MERGE (c)-[:PARTICIPATES]->(o) "
		}
		queryLinkWithUsers := ""
		if len(userIds) > 0 {
			queryLinkWithUsers = " WITH DISTINCT t, o " +
				" OPTIONAL MATCH (u:User)-[:USER_BELONGS_TO_TENANT]->(t) WHERE u.id in $userIds " +
				" MERGE (u)-[:PARTICIPATES]->(o) "
		}
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, queryLinkWithContacts, queryLinkWithUsers),
			map[string]interface{}{
				"tenant":              tenant,
				"conversationId":      conversationId,
				"updatedAt":           time.Now().UTC(),
				"contactIds":          contactIds,
				"userIds":             userIds,
				"lastSenderId":        lastSenderId,
				"lastSenderType":      lastSenderType,
				"lastSenderFirstName": "",
				"lastSenderLastName":  "",
				"lastContentPreview":  lastContentPreview,
			})
		return utils.ExtractSingleRecordFirstValueAsNode(ctx, queryResult, err)
	}); err != nil {
		return "", fmt.Errorf("UpdateConversation: %w", err)
	} else {
		dbNode := result.(*dbtype.Node)
		return utils.GetPropsFromNode(*dbNode)["id"].(string), err
	}
}

func mapNodeToConversation(node *dbtype.Node) *Conversation {
	if node == nil {
		return nil
	}

	props := utils.GetPropsFromNode(*node)

	conversation := new(Conversation)
	if id, err := utils.GetPropsFromNode(*node)["id"].(string); err {
		conversation.Id = id
	}
	if status, err := utils.GetPropsFromNode(*node)["status"].(string); err {
		conversation.Status = status
	}
	if channel, err := utils.GetPropsFromNode(*node)["channel"].(string); err {
		conversation.Channel = channel
	}
	if channel, err := utils.GetPropsFromNode(*node)["startedAt"].(time.Time); err {
		conversation.StartedAt = channel
	}
	if updatedAt, err := utils.GetPropsFromNode(*node)["updatedAt"].(time.Time); err {
		conversation.UpdatedAt = updatedAt
	}
	if initatorFirstName, err := utils.GetPropsFromNode(*node)["initiatorFirstName"].(string); err {
		conversation.InitiatorFirstName = initatorFirstName
	}
	if initatorLastName, err := utils.GetPropsFromNode(*node)["initiatorLastName"].(string); err {
		conversation.InitiatorLastName = initatorLastName
	}
	if initatorUsername, err := utils.GetPropsFromNode(*node)["initiatorUsername"].(string); err {
		conversation.InitiatorUsername = initatorUsername
	}
	if initatorType, err := utils.GetPropsFromNode(*node)["initiatorType"].(string); err {
		conversation.InitiatorType = initatorType
	}

	conversation.LastSenderId = utils.GetStringPropOrEmpty(props, "lastSenderId")
	conversation.LastSenderType = utils.GetStringPropOrEmpty(props, "lastSenderType")
	conversation.LastSenderFirstName = utils.GetStringPropOrEmpty(props, "lastSenderFirstName")
	conversation.LastSenderLastName = utils.GetStringPropOrEmpty(props, "lastSenderLastName")
	conversation.LastContentPreview = utils.GetStringPropOrEmpty(props, "lastContentPreview")

	return conversation
}

func (s *CustomerOSService) GetContactWithEmailOrCreate(ctx context.Context, tenant string, email string) (Contact, error) {
	contact, err := s.GetContactByEmail(ctx, email, tenant)
	if err != nil {
		contact, err = s.CreateContactWithEmail(ctx, tenant, email)
		if err != nil {
			return Contact{}, fmt.Errorf("GetContactWithEmailOrCreate: %w", err)
		}
		if contact == nil {
			return Contact{}, errors.New("GetContactWithEmailOrCreate: contact not found and could not be created")
		}
		return *contact, nil
	} else {
		return *contact, nil
	}
}

func (s *CustomerOSService) GetContactWithPhoneOrCreate(ctx context.Context, tenant string, phoneNumber string) (Contact, error) {
	contact, err := s.GetContactByPhone(ctx, phoneNumber, tenant)
	if err != nil {
		contact, err = s.CreateContactWithPhone(ctx, tenant, phoneNumber)
		if err != nil {
			return Contact{}, fmt.Errorf("GetContactWithPhoneOrCreate: %w", err)
		}
		if contact == nil {
			return Contact{}, errors.New("GetContactWithPhoneOrCreate: contact not found and could not be created")
		}
		return *contact, nil
	} else {
		return *contact, nil
	}
}

func (s *CustomerOSService) GetActiveConversationOrCreate(
	ctx context.Context,
	tenant string,
	initiator Participant,
	initiatorUsername *msProto.ParticipantId,
	eventType entity.EventType,
	threadId string,
) (*Conversation, error) {
	var conversation *Conversation
	var err error

	// for webchat, thread id is empty string
	if threadId == "" {
		if initiator.Type == entity.CONTACT {
			conversation, err = s.GetConversationWithContactInitiator(ctx, tenant, initiator.Id, eventType)
		} else if initiator.Type == entity.USER {
			conversation, err = s.GetConversationWithUserInitiator(ctx, tenant, initiator.Id, eventType)
		}
	} else {
		conversation, err = s.GetConversationWithThreadId(ctx, tenant, eventType, threadId)
	}

	if err != nil {
		return nil, fmt.Errorf("GetActiveConversationOrCreate: %w", err)
	}

	if conversation == nil {
		conversation, err = s.CreateConversation(ctx, tenant, initiator, s.commonStoreService.ConvertMSParticipantIdToUsername(initiatorUsername), eventType, threadId)
	}
	if err != nil {
		return nil, fmt.Errorf("GetActiveConversationOrCreate: %w", err)
	}

	return conversation, nil
}
func (s *CustomerOSService) GetConversationWithThreadId(ctx context.Context, tenant string, eventType entity.EventType, threadId string) (*Conversation, error) {
	session := utils.NewNeo4jWriteSession(ctx, *s.driver)
	defer session.Close(ctx)

	conversationNode, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `match(o:Conversation_`+tenant+`{status:"ACTIVE", channel: $eventType, threadId: $threadId}) return o`,
			map[string]any{
				"tenant":    tenant,
				"threadId":  threadId,
				"eventType": eventType,
			}); err != nil {
			return nil, err
		} else {
			record, err := queryResult.Single(ctx)
			if err != nil && err.Error() != "Result contains no more records" {
				return nil, err
			}
			if record != nil {
				return record.Values[0].(dbtype.Node), nil
			} else {
				return nil, nil
			}
		}
	})

	if err != nil {
		return nil, fmt.Errorf("GetConversationWithContactInitiator: %w", err)
	}

	if conversationNode != nil {
		node := conversationNode.(dbtype.Node)
		return mapNodeToConversation(&node), nil
	} else {
		return nil, nil
	}
}

func (s *CustomerOSService) GetConversationWithContactInitiator(ctx context.Context, tenant string, contactId string, eventType entity.EventType) (*Conversation, error) {
	session := utils.NewNeo4jWriteSession(ctx, *s.driver)
	defer session.Close(ctx)

	conversationNode, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `match(o:Conversation{status:"ACTIVE", channel: $eventType})<-[:INITIATED]-(c:Contact{id:$contactId})-[:CONTACT_BELONGS_TO_TENANT]->(t:Tenant{name:$tenant}) return o`,
			map[string]any{
				"tenant":    tenant,
				"contactId": contactId,
				"eventType": eventType,
			}); err != nil {
			return nil, err
		} else {
			record, err := queryResult.Single(ctx)
			if err != nil && err.Error() != "Result contains no more records" {
				return nil, err
			}
			if record != nil {
				return record.Values[0].(dbtype.Node), nil
			} else {
				return nil, nil
			}
		}
	})

	if err != nil {
		return nil, fmt.Errorf("GetConversationWithContactInitiator: %w", err)
	}

	if conversationNode != nil {
		node := conversationNode.(dbtype.Node)
		return mapNodeToConversation(&node), nil
	} else {
		return nil, nil
	}
}

func (s *CustomerOSService) GetConversationWithUserInitiator(ctx context.Context, tenant string, userId string, eventType entity.EventType) (*Conversation, error) {
	session := utils.NewNeo4jReadSession(ctx, *s.driver)
	defer session.Close(ctx)

	conversationNode, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if queryResult, err := tx.Run(ctx, `match(o:Conversation{status:"ACTIVE", channel: $eventType})<-[:INITIATED]-(u:User{id:$userId})-[:USER_BELONGS_TO_TENANT]->(t:Tenant{name:$tenant}) return o`,
			map[string]any{
				"tenant":    tenant,
				"userId":    userId,
				"eventType": eventType,
			}); err != nil {
			return nil, err
		} else {
			record, err := queryResult.Single(ctx)
			if err != nil && err.Error() != "Result contains no more records" {
				return nil, err
			}
			if record != nil {
				return record.Values[0].(dbtype.Node), nil
			} else {
				return nil, nil
			}
		}
	})

	if err != nil {
		return nil, fmt.Errorf("GetConversationWithUserInitiator: %w", err)
	}

	if conversationNode != nil {
		node := conversationNode.(dbtype.Node)
		return mapNodeToConversation(&node), nil
	} else {
		return nil, nil
	}
}

func NewCustomerOSService(driver *neo4j.DriverWithContext, graphqlClient *graphql.Client, postgresRepositories *repository.PostgresRepositories, commonStoreService *commonStoreService, config *config.Config) *CustomerOSService {
	customerOsService := new(CustomerOSService)
	customerOsService.driver = driver
	customerOsService.postgresRepositories = postgresRepositories
	customerOsService.commonStoreService = commonStoreService
	customerOsService.graphqlClient = graphqlClient
	customerOsService.conf = config
	return customerOsService
}

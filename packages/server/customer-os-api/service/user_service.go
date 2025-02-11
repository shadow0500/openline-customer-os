package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	user_grpc_service "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/user"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"reflect"
)

type UserService interface {
	Create(ctx context.Context, userCreateData *UserCreateData) (*entity.UserEntity, error)
	Update(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
	GetAll(ctx context.Context, page, limit int, filter *model.Filter, sortBy []*model.SortBy) (*utils.Pagination, error)
	FindUserById(ctx context.Context, userId string) (*entity.UserEntity, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	GetContactOwner(ctx context.Context, contactId string) (*entity.UserEntity, error)
	GetNoteCreator(ctx context.Context, noteId string) (*entity.UserEntity, error)
	GetAllForConversation(ctx context.Context, conversationId string) (*entity.UserEntities, error)
	GetUsersForEmails(ctx context.Context, emailIds []string) (*entity.UserEntities, error)
	GetUsersForPhoneNumbers(ctx context.Context, phoneNumberIds []string) (*entity.UserEntities, error)
	GetUsersForPlayers(ctx context.Context, playerIds []string) (*entity.UserEntities, error)
	GetUserOwnersForOrganizations(ctx context.Context, organizationIDs []string) (*entity.UserEntities, error)
	GetDistinctOrganizationOwners(ctx context.Context) (*entity.UserEntities, error)

	UpsertInEventStore(ctx context.Context, size int) (int, int, error)
	UpsertPhoneNumberRelationInEventStore(ctx context.Context, size int) (int, int, error)
	UpsertEmailRelationInEventStore(ctx context.Context, size int) (int, int, error)

	AddRole(ctx context.Context, userId string, role model.Role) (*entity.UserEntity, error)
	AddRoleInTenant(ctx context.Context, userId string, tenant string, role model.Role) (*entity.UserEntity, error)
	DeleteRole(ctx context.Context, userId string, role model.Role) (*entity.UserEntity, error)
	DeleteRoleInTenant(ctx context.Context, userId string, tenant string, role model.Role) (*entity.UserEntity, error)

	mapDbNodeToUserEntity(dbNode dbtype.Node) *entity.UserEntity
	addPlayerDbRelationshipToUser(relationship dbtype.Relationship, userEntity *entity.UserEntity)
}

type UserCreateData struct {
	UserEntity   *entity.UserEntity
	EmailEntity  *entity.EmailEntity
	PlayerEntity *entity.PlayerEntity
}

type userService struct {
	log          logger.Logger
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
}

func NewUserService(log logger.Logger, repositories *repository.Repositories, grpcClients *grpc_client.Clients) UserService {
	return &userService{
		log:          log,
		repositories: repositories,
		grpcClients:  grpcClients,
	}
}

func (s *userService) getNeo4jDriver() neo4j.DriverWithContext {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *userService) ContainsRole(parentCtx context.Context, allowedRoles []model.Role) bool {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.ContainsRole")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	myRoles := common.GetRolesFromContext(ctx)
	for _, allowedRole := range allowedRoles {
		for _, myRole := range myRoles {
			if myRole == allowedRole {
				return true
			}
		}
	}
	return false
}

func (s *userService) CanAddRemoveRole(parentCtx context.Context, role model.Role) bool {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.CanAddRemoveRole")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	switch role {
	case model.RoleAdmin:
		return false // this role is a special endpoint and can not be given to a user
	case model.RoleOwner:
		return s.ContainsRole(ctx, []model.Role{model.RoleAdmin, model.RoleCustomerOsPlatformOwner, model.RoleOwner})
	case model.RoleCustomerOsPlatformOwner:
		return s.ContainsRole(ctx, []model.Role{model.RoleAdmin, model.RoleCustomerOsPlatformOwner})
	case model.RoleUser:
		return s.ContainsRole(ctx, []model.Role{model.RoleAdmin, model.RoleCustomerOsPlatformOwner, model.RoleOwner})
	default:
		s.log.Errorf("unknown role: %s", role)
		return false
	}
}

func (s *userService) AddRole(parentCtx context.Context, userId string, role model.Role) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.AddRole")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	if !s.CanAddRemoveRole(ctx, role) {
		return nil, fmt.Errorf("user can not add role: %s", role)
	}

	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	updateDbUser, err := s.repositories.UserRepository.AddRole(ctx, session, common.GetContext(ctx).Tenant, userId, mapper.MapRoleToEntity(role))
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*updateDbUser), nil
}

func (s *userService) AddRoleInTenant(parentCtx context.Context, userId string, tenant string, role model.Role) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.AddRoleInTenant")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	if !s.CanAddRemoveRole(ctx, role) {
		return nil, fmt.Errorf("user can not add role: %s", role)
	}

	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	updateDbUser, err := s.repositories.UserRepository.AddRole(ctx, session, tenant, userId, mapper.MapRoleToEntity(role))
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*updateDbUser), nil
}

func (s *userService) DeleteRole(parentCtx context.Context, userId string, role model.Role) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.DeleteRole")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	if !s.CanAddRemoveRole(ctx, role) {
		return nil, fmt.Errorf("user can not delete role: %s", role)
	}

	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	updateDbUser, err := s.repositories.UserRepository.DeleteRole(ctx, session, common.GetContext(ctx).Tenant, userId, mapper.MapRoleToEntity(role))
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*updateDbUser), nil
}

func (s *userService) DeleteRoleInTenant(parentCtx context.Context, userId string, tenant string, role model.Role) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.DeleteRoleInTenant")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	if !s.CanAddRemoveRole(ctx, role) {
		return nil, fmt.Errorf("user can not delete role: %s", role)
	}

	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	updateDbUser, err := s.repositories.UserRepository.DeleteRole(ctx, session, tenant, userId, mapper.MapRoleToEntity(role))
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*updateDbUser), nil
}

func (s *userService) Create(parentCtx context.Context, userCreateData *UserCreateData) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.Create")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	userDbNode, err := session.ExecuteWrite(ctx, s.createUserInDBTxWork(ctx, userCreateData))
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*userDbNode.(*dbtype.Node)), nil
}

func (s *userService) Update(parentCtx context.Context, entity *entity.UserEntity) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.Update")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	if entity.Id != common.GetContext(ctx).UserId {
		if !s.ContainsRole(ctx, []model.Role{model.RoleAdmin, model.RoleCustomerOsPlatformOwner, model.RoleOwner}) {
			return nil, fmt.Errorf("user can not update other user")
		}
	}
	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	userDbNode, err := s.repositories.UserRepository.Update(ctx, session, common.GetContext(ctx).Tenant, *entity)
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*userDbNode), nil
}

func (s *userService) GetAll(parentCtx context.Context, page, limit int, filter *model.Filter, sortBy []*model.SortBy) (*utils.Pagination, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetAll")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	var paginatedResult = utils.Pagination{
		Limit: limit,
		Page:  page,
	}
	cypherSort, err := buildSort(sortBy, reflect.TypeOf(entity.UserEntity{}))
	if err != nil {
		return nil, err
	}
	cypherFilter, err := buildFilter(filter, reflect.TypeOf(entity.UserEntity{}))
	if err != nil {
		return nil, err
	}

	dbNodesWithTotalCount, err := s.repositories.UserRepository.GetPaginatedUsers(
		ctx,
		session,
		common.GetContext(ctx).Tenant,
		paginatedResult.GetSkip(),
		paginatedResult.GetLimit(),
		cypherFilter,
		cypherSort)
	if err != nil {
		return nil, err
	}
	paginatedResult.SetTotalRows(dbNodesWithTotalCount.Count)

	users := make(entity.UserEntities, 0, len(dbNodesWithTotalCount.Nodes))
	for _, v := range dbNodesWithTotalCount.Nodes {
		users = append(users, *s.mapDbNodeToUserEntity(*v))
	}
	paginatedResult.SetRows(&users)
	return &paginatedResult, nil
}

func (s *userService) GetContactOwner(parentCtx context.Context, contactId string) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetContactOwner")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	ownerDbNode, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return s.repositories.UserRepository.GetOwnerForContact(ctx, tx, common.GetContext(ctx).Tenant, contactId)
	})
	if err != nil {
		return nil, err
	} else if ownerDbNode.(*dbtype.Node) == nil {
		return nil, nil
	} else {
		return s.mapDbNodeToUserEntity(*ownerDbNode.(*dbtype.Node)), nil
	}
}

func (s *userService) GetNoteCreator(parentCtx context.Context, noteId string) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetNoteCreator")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	userDbNode, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return s.repositories.UserRepository.GetCreatorForNote(ctx, tx, common.GetContext(ctx).Tenant, noteId)
	})
	if err != nil {
		return nil, err
	} else if userDbNode.(*dbtype.Node) == nil {
		return nil, nil
	} else {
		return s.mapDbNodeToUserEntity(*userDbNode.(*dbtype.Node)), nil
	}
}

func (s *userService) FindUserById(parentCtx context.Context, userId string) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.FindUserById")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	if userDbNode, err := s.repositories.UserRepository.GetById(ctx, session, common.GetContext(ctx).Tenant, userId); err != nil {
		return nil, err
	} else {
		return s.mapDbNodeToUserEntity(*userDbNode), nil
	}
}

func (s *userService) FindUserByEmail(parentCtx context.Context, email string) (*entity.UserEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.FindUserByEmail")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	userDbNode, err := s.repositories.UserRepository.FindUserByEmail(ctx, session, common.GetContext(ctx).Tenant, email)
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToUserEntity(*userDbNode), nil
}

func (s *userService) GetAllForConversation(parentCtx context.Context, conversationId string) (*entity.UserEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetAllForConversation")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	dbNodes, err := s.repositories.UserRepository.GetAllForConversation(ctx, session, common.GetContext(ctx).Tenant, conversationId)
	if err != nil {
		return nil, err
	}

	userEntities := make(entity.UserEntities, 0, len(dbNodes))
	for _, dbNode := range dbNodes {
		userEntities = append(userEntities, *s.mapDbNodeToUserEntity(*dbNode))
	}
	return &userEntities, nil
}

func (s *userService) createUserInDBTxWork(parentCtx context.Context, newUser *UserCreateData) func(tx neo4j.ManagedTransaction) (any, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.createUserInDBTxWork")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	return func(tx neo4j.ManagedTransaction) (any, error) {
		tenant := common.GetContext(ctx).Tenant
		userDbNode, err := s.repositories.UserRepository.Create(ctx, tx, tenant, *newUser.UserEntity)
		if err != nil {
			return nil, err
		}
		var userId = utils.GetPropsFromNode(*userDbNode)["id"].(string)

		playerDbNode, err := s.repositories.PlayerRepository.Merge(ctx, tx, newUser.PlayerEntity)
		if err != nil {
			return nil, err
		}
		var playerId = utils.GetPropsFromNode(*playerDbNode)["id"].(string)

		err = s.repositories.PlayerRepository.LinkWithUserInTx(ctx, tx, playerId, userId, tenant, entity.IDENTIFIES)
		if err != nil {
			return nil, err
		}

		if newUser.EmailEntity != nil {
			_, _, err := s.repositories.EmailRepository.MergeEmailToInTx(ctx, tx, tenant, entity.USER, userId, *newUser.EmailEntity)
			if err != nil {
				return nil, err
			}
		}
		return userDbNode, nil
	}
}

func (s *userService) GetUsersForEmails(parentCtx context.Context, emailIds []string) (*entity.UserEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetUsersForEmails")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	users, err := s.repositories.UserRepository.GetAllForEmails(ctx, common.GetTenantFromContext(ctx), emailIds)
	if err != nil {
		return nil, err
	}
	userEntities := make(entity.UserEntities, 0, len(users))
	for _, v := range users {
		userEntity := s.mapDbNodeToUserEntity(*v.Node)
		userEntity.DataloaderKey = v.LinkedNodeId
		userEntities = append(userEntities, *userEntity)
	}
	return &userEntities, nil
}

func (s *userService) GetUsersForPhoneNumbers(parentCtx context.Context, phoneNumberIds []string) (*entity.UserEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetUsersForPhoneNumbers")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	users, err := s.repositories.UserRepository.GetAllForPhoneNumbers(ctx, common.GetTenantFromContext(ctx), phoneNumberIds)
	if err != nil {
		return nil, err
	}
	userEntities := make(entity.UserEntities, 0, len(users))
	for _, v := range users {
		userEntity := s.mapDbNodeToUserEntity(*v.Node)
		userEntity.DataloaderKey = v.LinkedNodeId
		userEntities = append(userEntities, *userEntity)
	}
	return &userEntities, nil
}

func (s *userService) GetUsersForPlayers(parentCtx context.Context, playerIds []string) (*entity.UserEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetUsersForPlayers")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	users, err := s.repositories.PlayerRepository.GetUsersForPlayer(ctx, playerIds)
	if err != nil {
		return nil, err
	}
	userEntities := make(entity.UserEntities, 0, len(users))
	for _, v := range users {
		userEntity := s.mapDbNodeToUserEntity(*v.Node)
		userEntity.DataloaderKey = v.LinkedNodeId
		s.addPlayerDbRelationshipToUser(*v.Relationship, userEntity)
		userEntity.Tenant = v.Tenant
		userEntities = append(userEntities, *userEntity)
	}
	return &userEntities, nil
}

func (s *userService) GetUserOwnersForOrganizations(parentCtx context.Context, organizationIDs []string) (*entity.UserEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetUserOwnersForOrganizations")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("organizationIDs", organizationIDs))

	users, err := s.repositories.UserRepository.GetAllOwnersForOrganizations(ctx, common.GetTenantFromContext(ctx), organizationIDs)
	if err != nil {
		return nil, err
	}
	userEntities := make(entity.UserEntities, 0, len(users))
	for _, v := range users {
		userEntity := s.mapDbNodeToUserEntity(*v.Node)
		userEntity.DataloaderKey = v.LinkedNodeId
		userEntities = append(userEntities, *userEntity)
	}
	return &userEntities, nil
}

func (s *userService) GetDistinctOrganizationOwners(parentCtx context.Context) (*entity.UserEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(parentCtx, "UserService.GetDistinctOrganizationOwners")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	dbNodes, err := s.repositories.UserRepository.GetDistinctOrganizationOwners(ctx, common.GetTenantFromContext(ctx))
	if err != nil {
		return nil, err
	}

	userEntities := make(entity.UserEntities, 0, len(dbNodes))
	for _, dbNode := range dbNodes {
		userEntities = append(userEntities, *s.mapDbNodeToUserEntity(*dbNode))
	}
	return &userEntities, nil
}

func (s *userService) UpsertInEventStore(ctx context.Context, size int) (int, int, error) {
	processedRecords := 0
	failedRecords := 0
	outputErr := error(nil)
	for size > 0 {
		batchSize := constants.Neo4jBatchSize
		if size < constants.Neo4jBatchSize {
			batchSize = size
		}
		records, err := s.repositories.UserRepository.GetAllCrossTenants(ctx, batchSize)
		if err != nil {
			return 0, 0, err
		}
		for _, v := range records {
			_, err := s.grpcClients.UserClient.UpsertUser(context.Background(), &user_grpc_service.UpsertUserGrpcRequest{
				Id:            utils.GetStringPropOrEmpty(v.Node.Props, "id"),
				Tenant:        v.LinkedNodeId,
				Name:          utils.GetStringPropOrEmpty(v.Node.Props, "name"),
				FirstName:     utils.GetStringPropOrEmpty(v.Node.Props, "firstName"),
				LastName:      utils.GetStringPropOrEmpty(v.Node.Props, "lastName"),
				Source:        utils.GetStringPropOrEmpty(v.Node.Props, "source"),
				SourceOfTruth: utils.GetStringPropOrEmpty(v.Node.Props, "sourceOfTruth"),
				AppSource:     utils.GetStringPropOrEmpty(v.Node.Props, "appSource"),
				CreatedAt:     utils.ConvertTimeToTimestampPtr(utils.GetTimePropOrNil(v.Node.Props, "createdAt")),
				UpdatedAt:     utils.ConvertTimeToTimestampPtr(utils.GetTimePropOrNil(v.Node.Props, "updatedAt")),
			})
			if err != nil {
				failedRecords++
				if outputErr != nil {
					outputErr = err
				}
				s.log.Errorf("(%s) Failed to call method: %v", utils.GetFunctionName(), err.Error())
			} else {
				processedRecords++
			}
		}

		size -= batchSize
	}

	return processedRecords, failedRecords, outputErr
}

func (s *userService) UpsertPhoneNumberRelationInEventStore(ctx context.Context, size int) (int, int, error) {
	processedRecords := 0
	failedRecords := 0
	outputErr := error(nil)
	for size > 0 {
		batchSize := constants.Neo4jBatchSize
		if size < constants.Neo4jBatchSize {
			batchSize = size
		}
		records, err := s.repositories.UserRepository.GetAllUserPhoneNumberRelationships(ctx, batchSize)
		if err != nil {
			return 0, 0, err
		}
		for _, v := range records {
			_, err := s.grpcClients.UserClient.LinkPhoneNumberToUser(context.Background(), &user_grpc_service.LinkPhoneNumberToUserGrpcRequest{
				Primary:       utils.GetBoolPropOrFalse(v.Values[0].(neo4j.Relationship).Props, "primary"),
				Label:         utils.GetStringPropOrEmpty(v.Values[0].(neo4j.Relationship).Props, "label"),
				UserId:        v.Values[1].(string),
				PhoneNumberId: v.Values[2].(string),
				Tenant:        v.Values[3].(string),
			})
			if err != nil {
				failedRecords++
				if outputErr != nil {
					outputErr = err
				}
				s.log.Errorf("(%s) Failed to call method: {%v}", utils.GetFunctionName(), err.Error())
			} else {
				processedRecords++
			}
		}

		size -= batchSize
	}

	return processedRecords, failedRecords, outputErr
}

func (s *userService) UpsertEmailRelationInEventStore(ctx context.Context, size int) (int, int, error) {
	processedRecords := 0
	failedRecords := 0
	outputErr := error(nil)
	for size > 0 {
		batchSize := constants.Neo4jBatchSize
		if size < constants.Neo4jBatchSize {
			batchSize = size
		}
		records, err := s.repositories.UserRepository.GetAllUserEmailRelationships(ctx, batchSize)
		if err != nil {
			return 0, 0, err
		}
		for _, v := range records {
			_, err := s.grpcClients.UserClient.LinkEmailToUser(context.Background(), &user_grpc_service.LinkEmailToUserGrpcRequest{
				Primary: utils.GetBoolPropOrFalse(v.Values[0].(neo4j.Relationship).Props, "primary"),
				Label:   utils.GetStringPropOrEmpty(v.Values[0].(neo4j.Relationship).Props, "label"),
				UserId:  v.Values[1].(string),
				EmailId: v.Values[2].(string),
				Tenant:  v.Values[3].(string),
			})
			if err != nil {
				failedRecords++
				if outputErr != nil {
					outputErr = err
				}
				s.log.Errorf("(%s) Failed to call method: {%v}", utils.GetFunctionName(), err.Error())
			} else {
				processedRecords++
			}
		}

		size -= batchSize
	}

	return processedRecords, failedRecords, outputErr
}

func (s *userService) mapDbNodeToUserEntity(dbNode dbtype.Node) *entity.UserEntity {
	props := utils.GetPropsFromNode(dbNode)
	return &entity.UserEntity{
		Id:            utils.GetStringPropOrEmpty(props, "id"),
		FirstName:     utils.GetStringPropOrEmpty(props, "firstName"),
		LastName:      utils.GetStringPropOrEmpty(props, "lastName"),
		CreatedAt:     utils.GetTimePropOrEpochStart(props, "createdAt"),
		UpdatedAt:     utils.GetTimePropOrEpochStart(props, "updatedAt"),
		Source:        entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth: entity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
		AppSource:     utils.GetStringPropOrEmpty(props, "appSource"),
		Roles:         utils.GetListStringPropOrEmpty(props, "roles"),
	}
}

func (s *userService) addPlayerDbRelationshipToUser(relationship dbtype.Relationship, userEntity *entity.UserEntity) {
	props := utils.GetPropsFromRelationship(relationship)
	userEntity.DefaultForPlayer = utils.GetBoolPropOrFalse(props, "default")
}

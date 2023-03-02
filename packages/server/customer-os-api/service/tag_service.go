package service

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

type TagService interface {
	Merge(ctx context.Context, tag *entity.TagEntity) (*entity.TagEntity, error)
	Update(ctx context.Context, tag *entity.TagEntity) (*entity.TagEntity, error)
	UnlinkAndDelete(ctx context.Context, id string) (bool, error)
	GetAll(ctx context.Context) (*entity.TagEntities, error)
	GetTagsForContact(ctx context.Context, contactId string) (*entity.TagEntities, error)
	GetTagsForContacts(ctx context.Context, contactIds []string) (*entity.TagEntities, error)
	GetTagsForTickets(ctx context.Context, ticketIds []string) (*entity.TagEntities, error)
	GetTagsForOrganizations(ctx context.Context, organizationIds []string) (*entity.TagEntities, error)
}

type tagService struct {
	repositories *repository.Repositories
}

func NewTagService(repository *repository.Repositories) TagService {
	return &tagService{
		repositories: repository,
	}
}

func (s *tagService) Merge(ctx context.Context, tag *entity.TagEntity) (*entity.TagEntity, error) {
	tagNodePtr, err := s.repositories.TagRepository.Merge(ctx, common.GetTenantFromContext(ctx), *tag)
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToTagEntity(*tagNodePtr), nil
}

func (s *tagService) Update(ctx context.Context, tag *entity.TagEntity) (*entity.TagEntity, error) {
	tagNodePtr, err := s.repositories.TagRepository.Update(ctx, common.GetTenantFromContext(ctx), *tag)
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToTagEntity(*tagNodePtr), nil
}

func (s *tagService) UnlinkAndDelete(ctx context.Context, id string) (bool, error) {
	err := s.repositories.TagRepository.UnlinkAndDelete(ctx, common.GetTenantFromContext(ctx), id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *tagService) GetAll(ctx context.Context) (*entity.TagEntities, error) {
	tagDbNodes, err := s.repositories.TagRepository.GetAll(ctx, common.GetTenantFromContext(ctx))
	if err != nil {
		return nil, err
	}
	tagEntities := entity.TagEntities{}
	for _, dbNodePtr := range tagDbNodes {
		tagEntities = append(tagEntities, *s.mapDbNodeToTagEntity(*dbNodePtr))
	}
	return &tagEntities, nil
}

func (s *tagService) GetTagsForContact(ctx context.Context, contactId string) (*entity.TagEntities, error) {
	tagDbNodes, err := s.repositories.TagRepository.GetForContact(ctx, common.GetTenantFromContext(ctx), contactId)
	if err != nil {
		return nil, err
	}
	tagEntities := entity.TagEntities{}
	for _, dbNodePtr := range tagDbNodes {
		tagEntities = append(tagEntities, *s.mapDbNodeToTagEntity(*dbNodePtr))
	}
	return &tagEntities, nil
}

func (s *tagService) GetTagsForContacts(ctx context.Context, contactIds []string) (*entity.TagEntities, error) {
	tags, err := s.repositories.TagRepository.GetForContacts(ctx, common.GetTenantFromContext(ctx), contactIds)
	if err != nil {
		return nil, err
	}
	tagEntities := entity.TagEntities{}
	for _, v := range tags {
		tagEntity := s.mapDbNodeToTagEntity(*v.Node)
		s.addDbRelationshipToTagEntity(*v.Relationship, tagEntity)
		tagEntity.DataloaderKey = v.LinkedNodeId
		tagEntities = append(tagEntities, *tagEntity)
	}
	return &tagEntities, nil
}

func (s *tagService) GetTagsForTickets(ctx context.Context, ticketIds []string) (*entity.TagEntities, error) {
	tags, err := s.repositories.TagRepository.GetForTickets(ctx, common.GetTenantFromContext(ctx), ticketIds)
	if err != nil {
		return nil, err
	}
	tagEntities := entity.TagEntities{}
	for _, v := range tags {
		tagEntity := s.mapDbNodeToTagEntity(*v.Node)
		s.addDbRelationshipToTagEntity(*v.Relationship, tagEntity)
		tagEntity.DataloaderKey = v.LinkedNodeId
		tagEntities = append(tagEntities, *tagEntity)
	}
	return &tagEntities, nil
}

func (s *tagService) GetTagsForOrganizations(ctx context.Context, organizationIds []string) (*entity.TagEntities, error) {
	tags, err := s.repositories.TagRepository.GetForOrganizations(ctx, common.GetTenantFromContext(ctx), organizationIds)
	if err != nil {
		return nil, err
	}
	tagEntities := entity.TagEntities{}
	for _, v := range tags {
		tagEntity := s.mapDbNodeToTagEntity(*v.Node)
		s.addDbRelationshipToTagEntity(*v.Relationship, tagEntity)
		tagEntity.DataloaderKey = v.LinkedNodeId
		tagEntities = append(tagEntities, *tagEntity)
	}
	return &tagEntities, nil
}

func (s *tagService) mapDbNodeToTagEntity(dbNode dbtype.Node) *entity.TagEntity {
	props := utils.GetPropsFromNode(dbNode)
	tag := entity.TagEntity{
		Id:        utils.GetStringPropOrEmpty(props, "id"),
		Name:      utils.GetStringPropOrEmpty(props, "name"),
		CreatedAt: utils.GetTimePropOrEpochStart(props, "createdAt"),
		UpdatedAt: utils.GetTimePropOrEpochStart(props, "updatedAt"),
		AppSource: utils.GetStringPropOrEmpty(props, "appSource"),
		Source:    entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
	}
	return &tag
}

func (s *tagService) addDbRelationshipToTagEntity(relationship dbtype.Relationship, tagEntity *entity.TagEntity) {
	props := utils.GetPropsFromRelationship(relationship)
	tagEntity.TaggedAt = utils.GetTimePropOrEpochStart(props, "taggedAt")
}

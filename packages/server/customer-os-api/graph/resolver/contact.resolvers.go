package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
)

// ContactType is the resolver for the contactType field.
func (r *contactResolver) ContactType(ctx context.Context, obj *model.Contact) (*model.ContactType, error) {
	entity, err := r.Services.ContactTypeService.FindContactTypeForContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact type for contact %s", obj.ID)
		return nil, err
	}
	if entity == nil {
		return nil, nil
	}
	return mapper.MapEntityToContactType(entity), nil
}

// CompanyPositions is the resolver for the companyPositions field.
func (r *contactResolver) CompanyPositions(ctx context.Context, obj *model.Contact) ([]*model.CompanyPosition, error) {
	companyPositionEntities, err := r.Services.CompanyService.GetCompanyPositionsForContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get company positions %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToCompanyPositiones(companyPositionEntities), err
}

// Groups is the resolver for the groups field.
func (r *contactResolver) Groups(ctx context.Context, obj *model.Contact) ([]*model.ContactGroup, error) {
	contactGroupEntities, err := r.Services.ContactGroupService.FindAllForContact(ctx, obj)
	return mapper.MapEntitiesToContactGroups(contactGroupEntities), err
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *contactResolver) PhoneNumbers(ctx context.Context, obj *model.Contact) ([]*model.PhoneNumber, error) {
	phoneNumberEntities, err := r.Services.PhoneNumberService.FindAllForContact(ctx, obj)
	return mapper.MapEntitiesToPhoneNumbers(phoneNumberEntities), err
}

// Emails is the resolver for the emails field.
func (r *contactResolver) Emails(ctx context.Context, obj *model.Contact) ([]*model.Email, error) {
	emailEntities, err := r.Services.EmailService.FindAllForContact(ctx, obj)
	return mapper.MapEntitiesToEmails(emailEntities), err
}

// CustomFields is the resolver for the customFields field.
func (r *contactResolver) CustomFields(ctx context.Context, obj *model.Contact) ([]*model.CustomField, error) {
	var customFields []*model.CustomField
	customFieldEntities, err := r.Services.CustomFieldService.FindAllForContact(ctx, obj)
	for _, v := range mapper.MapEntitiesToCustomFields(customFieldEntities) {
		customFields = append(customFields, v)
	}
	return customFields, err
}

// FieldSets is the resolver for the fieldSets field.
func (r *contactResolver) FieldSets(ctx context.Context, obj *model.Contact) ([]*model.FieldSet, error) {
	fieldSetEntities, err := r.Services.FieldSetService.FindAllForContact(ctx, obj)
	return mapper.MapEntitiesToFieldSets(fieldSetEntities), err
}

// Definition is the resolver for the definition field.
func (r *contactResolver) Definition(ctx context.Context, obj *model.Contact) (*model.EntityDefinition, error) {
	entity, err := r.Services.EntityDefinitionService.FindLinkedWithContact(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact definition for contact %s", obj.ID)
		return nil, err
	}
	if entity == nil {
		return nil, nil
	}
	return mapper.MapEntityToEntityDefinition(entity), err
}

// Owner is the resolver for the owner field.
func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	owner, err := r.Services.UserService.FindContactOwner(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get owner for contact %s", obj.ID)
		return nil, err
	}
	if owner == nil {
		return nil, nil
	}
	return mapper.MapEntityToUser(owner), err
}

// Conversations is the resolver for the conversations field.
func (r *contactResolver) Conversations(ctx context.Context, obj *model.Contact, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ConversationService.GetConversationsForContact(ctx, obj.ID, pagination.Page, pagination.Limit, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact %s conversations", obj.ID)
		return nil, err
	}
	return &model.ConversationPage{
		Content:       mapper.MapEntitiesToConversations(paginatedResult.Rows.(*entity.ConversationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Actions is the resolver for the actions field.
func (r *contactResolver) Actions(ctx context.Context, obj *model.Contact, from time.Time, to time.Time, actionTypes []model.ActionType) ([]model.Action, error) {
	actions, err := r.Services.ActionsService.GetContactActions(ctx, obj.ID, from, to, actionTypes)
	if err != nil {
		graphql.AddErrorf(ctx, "failed to get actions for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToActions(actions), nil
}

// ContactCreate is the resolver for the contact_Create field.
func (r *mutationResolver) ContactCreate(ctx context.Context, input model.ContactInput) (*model.Contact, error) {
	contactNodeCreated, err := r.Services.ContactService.Create(ctx, &service.ContactCreateData{
		ContactEntity:     mapper.MapContactInputToEntity(input),
		CustomFields:      mapper.MapCustomFieldInputsToEntities(input.CustomFields),
		FieldSets:         mapper.MapFieldSetInputsToEntities(input.FieldSets),
		PhoneNumberEntity: mapper.MapPhoneNumberInputToEntity(input.PhoneNumber),
		EmailEntity:       mapper.MapEmailInputToEntity(input.Email),
		DefinitionId:      input.DefinitionID,
		ContactTypeId:     input.ContactTypeID,
		OwnerUserId:       input.OwnerID,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create contact %s %s", *input.FirstName, *input.LastName)
		return nil, err
	}
	return mapper.MapEntityToContact(contactNodeCreated), nil
}

// ContactUpdate is the resolver for the contact_Update field.
func (r *mutationResolver) ContactUpdate(ctx context.Context, input model.ContactUpdateInput) (*model.Contact, error) {
	updatedContact, err := r.Services.ContactService.Update(ctx, &service.ContactUpdateData{
		ContactEntity: mapper.MapContactUpdateInputToEntity(input),
		ContactTypeId: input.ContactTypeID,
		OwnerUserId:   input.OwnerID,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update contact %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToContact(updatedContact), nil
}

// ContactHardDelete is the resolver for the contact_HardDelete field.
func (r *mutationResolver) ContactHardDelete(ctx context.Context, contactID string) (*model.Result, error) {
	result, err := r.Services.ContactService.HardDelete(ctx, contactID)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not hard delete contact %s", contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// ContactSoftDelete is the resolver for the contact_SoftDelete field.
func (r *mutationResolver) ContactSoftDelete(ctx context.Context, contactID string) (*model.Result, error) {
	result, err := r.Services.ContactService.SoftDelete(ctx, contactID)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not soft delete contact %s", contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// ContactMergeCompanyPosition is the resolver for the contact_MergeCompanyPosition field.
func (r *mutationResolver) ContactMergeCompanyPosition(ctx context.Context, contactID string, input model.CompanyPositionInput) (*model.CompanyPosition, error) {
	result, err := r.Services.CompanyService.MergeCompanyToContact(ctx, contactID, mapper.MapCompanyPositionInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add company position to contact %s", contactID)
		return nil, err
	}
	return mapper.MapEntityToCompanyPosition(result), nil
}

// ContactUpdateCompanyPosition is the resolver for the contact_UpdateCompanyPosition field.
func (r *mutationResolver) ContactUpdateCompanyPosition(ctx context.Context, contactID string, companyPositionID string, input model.CompanyPositionInput) (*model.CompanyPosition, error) {
	result, err := r.Services.CompanyService.UpdateCompanyPosition(ctx, contactID, companyPositionID, mapper.MapCompanyPositionInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update company position%s", companyPositionID)
		return nil, err
	}
	return mapper.MapEntityToCompanyPosition(result), nil
}

// ContactDeleteCompanyPosition is the resolver for the contact_DeleteCompanyPosition field.
func (r *mutationResolver) ContactDeleteCompanyPosition(ctx context.Context, contactID string, companyPositionID string) (*model.Result, error) {
	result, err := r.Services.CompanyService.DeleteCompanyPositionFromContact(ctx, contactID, companyPositionID)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove company position %s from contact %s", companyPositionID, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// Contact is the resolver for the contact field.
func (r *queryResolver) Contact(ctx context.Context, id string) (*model.Contact, error) {
	contactEntity, err := r.Services.ContactService.FindContactById(ctx, id)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ContactService.FindAll(ctx, pagination.Page, pagination.Limit, where, sort)
	return &model.ContactsPage{
		Content:       mapper.MapEntitiesToContacts(paginatedResult.Rows.(*entity.ContactEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// ContactByEmail is the resolver for the contactByEmail field.
func (r *queryResolver) ContactByEmail(ctx context.Context, email string) (*model.Contact, error) {
	contactEntity, err := r.Services.ContactService.FindContactByEmail(ctx, email)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with email %s not identified", email)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// ContactByPhone is the resolver for the contactByPhone field.
func (r *queryResolver) ContactByPhone(ctx context.Context, e164 string) (*model.Contact, error) {
	contactEntity, err := r.Services.ContactService.FindContactByPhoneNumber(ctx, e164)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with phone number %s not identified", e164)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

type contactResolver struct{ *Resolver }

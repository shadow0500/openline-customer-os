package mapper

import (
	"github.com/openline-ai/openline-customer-os/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/customer-os-api/graph/model"
)

func MapContactInputToEntity(input model.ContactInput) *entity.ContactEntity {
	contactEntity := entity.ContactEntity{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	if input.Label != nil {
		contactEntity.Label = *input.Label
	}
	if input.Company != nil {
		contactEntity.Company = *input.Company
	}
	if input.Title != nil {
		contactEntity.Title = *input.Title
	}
	if input.CompanyTitle != nil {
		contactEntity.CompanyTitle = *input.CompanyTitle
	}
	if input.Notes != nil {
		contactEntity.Notes = *input.Notes
	}
	if input.ContactType != nil {
		contactEntity.ContactType = *input.ContactType
	}
	return &contactEntity
}

func MapEntityToContact(contact *entity.ContactEntity) *model.Contact {
	var title = contact.Title
	var label = contact.Label
	var company = contact.Company
	var companyTitle = contact.CompanyTitle
	var notes = contact.Notes
	var contactType = contact.ContactType
	return &model.Contact{
		ID:           contact.Id,
		Title:        &title,
		FirstName:    contact.FirstName,
		LastName:     contact.LastName,
		Label:        &label,
		Company:      &company,
		CompanyTitle: &companyTitle,
		Notes:        &notes,
		ContactType:  &contactType,
		CreatedAt:    contact.CreatedAt,
	}
}

func MapEntitiesToContacts(contactEntities *entity.ContactNodes) []*model.Contact {
	var contacts []*model.Contact
	for _, contactEntity := range *contactEntities {
		contacts = append(contacts, MapEntityToContact(&contactEntity))
	}
	return contacts
}

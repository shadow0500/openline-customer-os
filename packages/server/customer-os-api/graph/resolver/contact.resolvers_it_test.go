package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueryResolver_ContactByEmail(t *testing.T) {
	defer tearDownTestCase()(t)
	otherTenant := "other"
	neo4jt.CreateTenant(driver, tenantName)
	neo4jt.CreateTenant(driver, otherTenant)
	contactId1 := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, otherTenant)
	neo4jt.AddEmailTo(driver, entity.CONTACT, tenantName, contactId1, "test@test.com", true, "MAIN")
	neo4jt.AddEmailTo(driver, entity.CONTACT, otherTenant, contactId2, "test@test.com", true, "MAIN")

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_by_email"), client.Var("email", "test@test.com"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_ByEmail model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.Equal(t, contactId1, contact.Contact_ByEmail.ID)
}

func TestQueryResolver_ContactByPhone(t *testing.T) {
	defer tearDownTestCase()(t)
	otherTenant := "other"
	neo4jt.CreateTenant(driver, tenantName)
	neo4jt.CreateTenant(driver, otherTenant)
	contactId1 := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, otherTenant)
	neo4jt.AddPhoneNumberToContact(driver, contactId1, "+1234567890", false, "OTHER")
	neo4jt.AddPhoneNumberToContact(driver, contactId2, "+1234567890", true, "MAIN")

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_by_phone"), client.Var("e164", "+1234567890"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_ByPhone model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.Equal(t, contactId1, contact.Contact_ByPhone.ID)
}

func TestMutationResolver_ContactCreate_Min(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	rawResponse, err := c.RawPost(getQuery("contact/create_contact_min"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_Create model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.Equal(t, "", contact.Contact_Create.Title.String())
	require.Equal(t, "", *contact.Contact_Create.FirstName)
	require.Equal(t, "", *contact.Contact_Create.LastName)
	require.Equal(t, "", *contact.Contact_Create.Label)
	require.Equal(t, model.DataSourceOpenline, contact.Contact_Create.Source)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact_"+tenantName))
	require.Equal(t, 2, neo4jt.GetTotalCountOfNodes(driver))

	assertNeo4jLabels(t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName})
}

func TestMutationResolver_ContactCreate(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	neo4jt.CreateTenant(driver, "otherTenant")

	rawResponse, err := c.RawPost(getQuery("contact/create_contact"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_Create model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.Equal(t, "MR", contact.Contact_Create.Title.String())
	require.Equal(t, "first", *contact.Contact_Create.FirstName)
	require.Equal(t, "last", *contact.Contact_Create.LastName)
	require.Equal(t, "Some label", *contact.Contact_Create.Label)
	require.Equal(t, model.DataSourceOpenline, contact.Contact_Create.Source)

	require.Equal(t, 5, len(contact.Contact_Create.CustomFields))

	boolField := contact.Contact_Create.CustomFields[0]
	require.NotNil(t, boolField.GetID())
	require.Equal(t, "boolField", boolField.Name)
	require.Equal(t, model.CustomFieldDataTypeBool, boolField.Datatype)
	require.Equal(t, true, boolField.Value.RealValue())

	decimalField := contact.Contact_Create.CustomFields[1]
	require.NotNil(t, decimalField.GetID())
	require.Equal(t, "decimalField", decimalField.Name)
	require.Equal(t, model.CustomFieldDataTypeDecimal, decimalField.Datatype)
	require.Equal(t, 0.001, decimalField.Value.RealValue())

	integerField := contact.Contact_Create.CustomFields[2]
	require.NotNil(t, integerField.GetID())
	require.Equal(t, "integerField", integerField.Name)
	require.Equal(t, model.CustomFieldDataTypeInteger, integerField.Datatype)
	// issue in decoding, int converted to float 64
	require.Equal(t, float64(123), integerField.Value.RealValue())

	textField := contact.Contact_Create.CustomFields[3]
	require.NotNil(t, textField.GetID())
	require.Equal(t, "textField", textField.Name)
	require.Equal(t, model.CustomFieldDataTypeText, textField.Datatype)
	require.Equal(t, "value1", textField.Value.RealValue())

	timeField := contact.Contact_Create.CustomFields[4]
	require.NotNil(t, timeField.GetID())
	require.Equal(t, "timeField", timeField.Name)
	require.Equal(t, model.CustomFieldDataTypeDatetime, timeField.Datatype)
	require.Equal(t, "2022-11-13T20:21:56.732Z", timeField.Value.RealValue())

	require.Equal(t, 1, len(contact.Contact_Create.Emails))
	require.NotNil(t, contact.Contact_Create.Emails[0].ID)
	require.Equal(t, "contact@abc.com", contact.Contact_Create.Emails[0].Email)
	require.Equal(t, "WORK", contact.Contact_Create.Emails[0].Label.String())
	require.Equal(t, false, contact.Contact_Create.Emails[0].Primary)
	require.Equal(t, model.DataSourceOpenline, contact.Contact_Create.Emails[0].Source)

	require.Equal(t, 1, len(contact.Contact_Create.PhoneNumbers))
	require.NotNil(t, contact.Contact_Create.PhoneNumbers[0].ID)
	require.Equal(t, "+1234567890", contact.Contact_Create.PhoneNumbers[0].E164)
	require.Equal(t, "MOBILE", contact.Contact_Create.PhoneNumbers[0].Label.String())
	require.Equal(t, true, contact.Contact_Create.PhoneNumbers[0].Primary)
	require.Equal(t, model.DataSourceOpenline, contact.Contact_Create.PhoneNumbers[0].Source)

	require.Equal(t, 0, len(contact.Contact_Create.Groups))

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact_"+tenantName))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "ContactGroup"))
	require.Equal(t, 5, neo4jt.GetCountOfNodes(driver, "CustomField"))
	require.Equal(t, 5, neo4jt.GetCountOfNodes(driver, "CustomField_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "TextField"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "IntField"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "FloatField"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "BoolField"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "TimeField"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Email_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "PhoneNumber"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "PhoneNumber_"+tenantName))
	require.Equal(t, 10, neo4jt.GetTotalCountOfNodes(driver))

	assertNeo4jLabels(t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName,
		"Email", "Email_" + tenantName, "PhoneNumber", "PhoneNumber_" + tenantName,
		"CustomField", "BoolField", "TextField", "FloatField", "TimeField", "IntField", "CustomField_" + tenantName})
}

func TestMutationResolver_ContactCreate_WithCustomFields(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	entityTemplateId := neo4jt.CreateEntityTemplate(driver, tenantName, model.EntityTemplateExtensionContact.String())
	fieldTemplateId := neo4jt.AddFieldTemplateToEntity(driver, entityTemplateId)
	setTemplateId := neo4jt.AddSetTemplateToEntity(driver, entityTemplateId)
	fieldInSetTemplateId := neo4jt.AddFieldTemplateToSet(driver, setTemplateId)

	rawResponse, err := c.RawPost(getQuery("contact/create_contact_with_custom_fields"),
		client.Var("entityTemplateId", entityTemplateId),
		client.Var("fieldTemplateId", fieldTemplateId),
		client.Var("setTemplateId", setTemplateId),
		client.Var("fieldInSetTemplateId", fieldInSetTemplateId))
	assertRawResponseSuccess(t, rawResponse, err)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact_"+tenantName))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "ContactGroup"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "CustomField"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "CustomField_"+tenantName))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "TextField"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "Email"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(driver, "PhoneNumber"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "EntityTemplate"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "CustomFieldTemplate"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "FieldSetTemplate"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "FieldSet"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "FieldSet_"+tenantName))
	require.Equal(t, 12, neo4jt.GetTotalCountOfNodes(driver))

	var contact struct {
		Contact_Create model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)

	createdContact := contact.Contact_Create
	require.Equal(t, model.DataSourceOpenline, createdContact.Source)
	require.Equal(t, entityTemplateId, createdContact.Template.ID)
	require.Equal(t, 2, len(createdContact.CustomFields))
	require.Equal(t, "field1", createdContact.CustomFields[0].Name)
	require.Equal(t, "TEXT", createdContact.CustomFields[0].Datatype.String())
	require.Equal(t, "value1", createdContact.CustomFields[0].Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, createdContact.CustomFields[0].Source)
	require.Equal(t, fieldTemplateId, createdContact.CustomFields[0].Template.ID)
	require.NotNil(t, createdContact.CustomFields[0].ID)
	require.NotNil(t, createdContact.CustomFields[0].CreatedAt)
	require.Equal(t, "field2", createdContact.CustomFields[1].Name)
	require.Equal(t, "TEXT", createdContact.CustomFields[1].Datatype.String())
	require.Equal(t, "value2", createdContact.CustomFields[1].Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, createdContact.CustomFields[1].Source)
	require.NotNil(t, createdContact.CustomFields[1].ID)
	require.NotNil(t, createdContact.CustomFields[1].CreatedAt)
	require.Equal(t, 2, len(createdContact.FieldSets))
	var set1, set2 *model.FieldSet
	if createdContact.FieldSets[0].Name == "set1" {
		set1 = createdContact.FieldSets[0]
		set2 = createdContact.FieldSets[1]
	} else {
		set1 = createdContact.FieldSets[1]
		set2 = createdContact.FieldSets[0]
	}
	require.NotNil(t, set1.ID)
	require.NotNil(t, set1.CreatedAt)
	require.Equal(t, "set1", set1.Name)
	require.Equal(t, 2, len(set1.CustomFields))
	require.NotNil(t, set1.CustomFields[0].CreatedAt)
	require.Equal(t, "field3InSet", set1.CustomFields[0].Name)
	require.Equal(t, "value3", set1.CustomFields[0].Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, set1.CustomFields[0].Source)
	require.Equal(t, "TEXT", set1.CustomFields[0].Datatype.String())
	require.Equal(t, fieldInSetTemplateId, set1.CustomFields[0].Template.ID)
	require.NotNil(t, set1.CustomFields[1].CreatedAt)
	require.Equal(t, "field4InSet", set1.CustomFields[1].Name)
	require.Equal(t, "value4", set1.CustomFields[1].Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, set1.CustomFields[1].Source)
	require.Equal(t, "TEXT", set1.CustomFields[1].Datatype.String())
	require.Nil(t, set1.CustomFields[1].Template)
	require.Equal(t, model.DataSourceOpenline, set1.Source)
	require.NotNil(t, set2.ID)
	require.NotNil(t, set2.CreatedAt)
	require.Equal(t, "set2", set2.Name)
	require.Equal(t, model.DataSourceOpenline, set2.Source)

	assertNeo4jLabels(t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName,
		"CustomFieldTemplate", "EntityTemplate", "FieldSet", "FieldSet_" + tenantName, "FieldSetTemplate",
		"CustomField", "TextField", "CustomField_" + tenantName})
}

func TestMutationResolver_ContactCreate_WithOwner(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	userId := neo4jt.CreateUser(driver, tenantName, entity.UserEntity{
		FirstName: "Agent",
		LastName:  "Smith",
	})

	rawResponse, err := c.RawPost(getQuery("contact/create_contact_with_owner"),
		client.Var("ownerId", userId))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_Create model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	createdContact := contact.Contact_Create
	require.Equal(t, "", createdContact.Title.String())
	require.Equal(t, "first", *createdContact.FirstName)
	require.Equal(t, "last", *createdContact.LastName)
	require.Equal(t, userId, createdContact.Owner.ID)
	require.Equal(t, "Agent", createdContact.Owner.FirstName)
	require.Equal(t, "Smith", createdContact.Owner.LastName)
	require.Equal(t, model.DataSourceOpenline, createdContact.Source)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Tenant"))
	require.Equal(t, 3, neo4jt.GetTotalCountOfNodes(driver))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "OWNS"))

	assertNeo4jLabels(t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "User"})
}

func TestMutationResolver_ContactCreate_WithExternalReference(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	neo4jt.CreateHubspotExternalSystem(driver, tenantName)

	rawResponse, err := c.RawPost(getQuery("contact/create_contact_with_external_reference"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_Create model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.NotNil(t, contact.Contact_Create.ID)
	require.Equal(t, model.DataSourceOpenline, contact.Contact_Create.Source)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "ExternalSystem"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "ExternalSystem_"+tenantName))
	require.Equal(t, 3, neo4jt.GetTotalCountOfNodes(driver))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "IS_LINKED_WITH"))

	assertNeo4jLabels(t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "ExternalSystem", "ExternalSystem_" + tenantName})
}

func TestMutationResolver_UpdateContact(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	origOwnerId := neo4jt.CreateDefaultUser(driver, tenantName)
	newOwnerId := neo4jt.CreateDefaultUser(driver, tenantName)
	contactId := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     model.PersonTitleMr.String(),
		FirstName: "first",
		LastName:  "last",
		Label:     "label",
	})

	neo4jt.UserOwnsContact(driver, origOwnerId, contactId)

	rawResponse, err := c.RawPost(getQuery("contact/update_contact"),
		client.Var("contactId", contactId),
		client.Var("ownerId", newOwnerId))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_Update model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.Equal(t, "DR", contact.Contact_Update.Title.String())
	require.Equal(t, "updated first", *contact.Contact_Update.FirstName)
	require.Equal(t, "updated last", *contact.Contact_Update.LastName)
	require.Equal(t, "updated label", *contact.Contact_Update.Label)
	require.Equal(t, newOwnerId, contact.Contact_Update.Owner.ID)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact_"+tenantName))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "OWNS"))

	assertNeo4jLabels(t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "User"})
}

func TestMutationResolver_UpdateContact_ClearTitle(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     model.PersonTitleMr.String(),
		FirstName: "first",
		LastName:  "last",
	})

	rawResponse, err := c.RawPost(getQuery("contact/update_contact_clear_title"),
		client.Var("contactId", contactId))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact_Update model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	updatedContact := contact.Contact_Update
	require.Equal(t, "", updatedContact.Title.String())
	require.Equal(t, "updated first", *updatedContact.FirstName)
	require.Equal(t, "updated last", *updatedContact.LastName)
}

func TestQueryResolver_Contact_WithJobRoles_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	organizationId1 := neo4jt.CreateFullOrganization(driver, tenantName, entity.OrganizationEntity{
		Name:        "name1",
		Description: "description1",
		Domain:      "domain1",
		Website:     "website1",
		Industry:    "industry1",
		IsPublic:    true,
	})
	organizationId2 := neo4jt.CreateFullOrganization(driver, tenantName, entity.OrganizationEntity{
		Name:        "name2",
		Description: "description2",
		Domain:      "domain2",
		Website:     "website2",
		Industry:    "industry2",
		IsPublic:    false,
	})
	role1 := neo4jt.ContactWorksForOrganization(driver, contactId, organizationId1, "CTO", false)
	role2 := neo4jt.ContactWorksForOrganization(driver, contactId, organizationId2, "CEO", true)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "JobRole"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "ROLE_IN"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "WORKS_AS"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_job_roles_by_id"),
		client.Var("contactId", contactId))
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedContact struct {
		Contact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedContact)
	require.Nil(t, err)
	require.Equal(t, contactId, searchedContact.Contact.ID)

	roles := searchedContact.Contact.JobRoles
	require.Equal(t, 2, len(roles))
	var cto, ceo *model.JobRole
	ceo = roles[0]
	cto = roles[1]
	require.Equal(t, role1, cto.ID)
	require.Equal(t, "CTO", *cto.JobTitle)
	require.Equal(t, false, cto.Primary)
	require.Equal(t, organizationId1, cto.Organization.ID)
	require.Equal(t, "name1", cto.Organization.Name)
	require.Equal(t, "description1", *cto.Organization.Description)
	require.Equal(t, "domain1", *cto.Organization.Domain)
	require.Equal(t, "website1", *cto.Organization.Website)
	require.Equal(t, "industry1", *cto.Organization.Industry)
	require.Equal(t, true, *cto.Organization.IsPublic)
	require.NotNil(t, cto.Organization.CreatedAt)

	require.Equal(t, role2, ceo.ID)
	require.Equal(t, "CEO", *ceo.JobTitle)
	require.Equal(t, true, ceo.Primary)
	require.Equal(t, organizationId2, ceo.Organization.ID)
	require.Equal(t, "name2", ceo.Organization.Name)
	require.Equal(t, "description2", *ceo.Organization.Description)
	require.Equal(t, "domain2", *ceo.Organization.Domain)
	require.Equal(t, "website2", *ceo.Organization.Website)
	require.Equal(t, "industry2", *ceo.Organization.Industry)
	require.Equal(t, false, *ceo.Organization.IsPublic)
	require.NotNil(t, ceo.Organization.CreatedAt)
}

func TestQueryResolver_Contact_WithNotes_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	userId := neo4jt.CreateDefaultUser(driver, tenantName)
	noteId1 := neo4jt.CreateNoteForContact(driver, tenantName, contactId, "note1")
	noteId2 := neo4jt.CreateNoteForContact(driver, tenantName, contactId, "note2")
	neo4jt.NoteCreatedByUser(driver, noteId1, userId)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "User"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Note"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "NOTED"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "CREATED"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_notes_by_id"),
		client.Var("contactId", contactId))
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedContact struct {
		Contact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedContact)
	require.Nil(t, err)
	require.Equal(t, contactId, searchedContact.Contact.ID)

	notes := searchedContact.Contact.Notes.Content
	require.Equal(t, 2, len(notes))
	var noteWithUser, noteWithoutUser *model.Note
	if noteId1 == notes[0].ID {
		noteWithUser = notes[0]
		noteWithoutUser = notes[1]
	} else {
		noteWithUser = notes[1]
		noteWithoutUser = notes[0]
	}
	require.Equal(t, noteId1, noteWithUser.ID)
	require.Equal(t, "note1", noteWithUser.HTML)
	require.NotNil(t, noteWithUser.CreatedAt)
	require.NotNil(t, noteWithUser.CreatedBy)
	require.Equal(t, userId, noteWithUser.CreatedBy.ID)
	require.Equal(t, "first", noteWithUser.CreatedBy.FirstName)
	require.Equal(t, "last", noteWithUser.CreatedBy.LastName)

	require.Equal(t, noteId2, noteWithoutUser.ID)
	require.Equal(t, "note2", noteWithoutUser.HTML)
	require.NotNil(t, noteWithoutUser.CreatedAt)
	require.Nil(t, noteWithoutUser.CreatedBy)
}

func TestQueryResolver_Contact_WithTags_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, tenantName)
	tagId1 := neo4jt.CreateTag(driver, tenantName, "tag1")
	tagId2 := neo4jt.CreateTag(driver, tenantName, "tag2")
	tagId3 := neo4jt.CreateTag(driver, tenantName, "tag3")
	neo4jt.TagContact(driver, contactId, tagId1)
	neo4jt.TagContact(driver, contactId, tagId2)
	neo4jt.TagContact(driver, contactId2, tagId3)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(driver, "Tag"))
	require.Equal(t, 3, neo4jt.GetCountOfRelationships(driver, "TAGGED"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_tags_by_id"),
		client.Var("contactId", contactId))
	assertRawResponseSuccess(t, rawResponse, err)

	var contactStruct struct {
		Contact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contactStruct)
	contact := contactStruct.Contact

	require.Nil(t, err)
	require.Equal(t, contactId, contact.ID)

	tags := contact.Tags
	require.Equal(t, 2, len(tags))
	require.Equal(t, tagId1, tags[0].ID)
	require.Equal(t, "tag1", tags[0].Name)
	require.Equal(t, tagId2, tags[1].ID)
	require.Equal(t, "tag2", tags[1].Name)
}

func TestQueryResolver_Contact_WithLocationsAndPlaces_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	neo4jt.CreateDefaultContact(driver, tenantName)
	locationId1 := neo4jt.CreateLocation(driver, tenantName, entity.LocationEntity{
		Name:      "WORK",
		Source:    entity.DataSourceOpenline,
		AppSource: "test",
	})
	locationId2 := neo4jt.CreateLocation(driver, tenantName, entity.LocationEntity{
		Name:      "UNKNOWN",
		Source:    entity.DataSourceOpenline,
		AppSource: "test",
	})
	placeInput := entity.PlaceEntity{
		Source:    entity.DataSourceOpenline,
		AppSource: "test",
		Country:   "testCountry",
		State:     "testState",
		City:      "testCity",
		Address:   "testAddress",
		Address2:  "testAddress2",
		Zip:       "testZip",
		Phone:     "testPhone",
		Fax:       "testFax",
	}
	placeId := neo4jt.CreatePlaceForLocation(driver, placeInput, locationId1)
	neo4jt.ContactAssociatedWithLocation(driver, contactId, locationId1)
	neo4jt.ContactAssociatedWithLocation(driver, contactId, locationId2)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Location"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Place"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "ASSOCIATED_WITH"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "LOCATED_AT"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_locations_and_places_by_id"),
		client.Var("contactId", contactId),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var contactStruct struct {
		Contact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contactStruct)
	require.Nil(t, err)

	contact := contactStruct.Contact
	require.NotNil(t, contact)
	require.Equal(t, 2, len(contact.Locations))

	var locationWithPlace, locationWithoutPlace *model.Location
	if contact.Locations[0].ID == locationId1 {
		locationWithPlace = contact.Locations[0]
		locationWithoutPlace = contact.Locations[1]
	} else {
		locationWithPlace = contact.Locations[1]
		locationWithoutPlace = contact.Locations[0]
	}

	require.Equal(t, locationId1, locationWithPlace.ID)
	require.Equal(t, "WORK", locationWithPlace.Name)
	require.NotNil(t, locationWithPlace.CreatedAt)
	require.NotNil(t, locationWithPlace.UpdatedAt)
	require.Equal(t, "test", *locationWithPlace.AppSource)
	require.Equal(t, model.DataSourceOpenline, *locationWithPlace.Source)
	require.NotNil(t, locationWithPlace.Place)

	place := locationWithPlace.Place
	require.Equal(t, placeId, place.ID)
	require.Equal(t, model.DataSourceOpenline, *place.Source)
	require.Equal(t, placeInput.Country, *place.Country)
	require.Equal(t, placeInput.City, *place.City)
	require.Equal(t, placeInput.State, *place.State)
	require.Equal(t, placeInput.Address, *place.Address)
	require.Equal(t, placeInput.Address2, *place.Address2)
	require.Equal(t, placeInput.Fax, *place.Fax)
	require.Equal(t, placeInput.Phone, *place.Phone)
	require.Equal(t, placeInput.Zip, *place.Zip)

	require.Equal(t, locationId2, locationWithoutPlace.ID)
	require.Equal(t, "UNKNOWN", locationWithoutPlace.Name)
	require.NotNil(t, locationWithoutPlace.CreatedAt)
	require.NotNil(t, locationWithoutPlace.UpdatedAt)
	require.Equal(t, "test", *locationWithoutPlace.AppSource)
	require.Equal(t, model.DataSourceOpenline, *locationWithoutPlace.Source)
	require.Nil(t, locationWithoutPlace.Place)
}

func TestQueryResolver_Contacts_SortByTitleAscFirstNameAscLastNameDesc(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	contact1 := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "MR",
		FirstName: "contact",
		LastName:  "1",
	})
	contact2 := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "DR",
		FirstName: "contact",
		LastName:  "9",
	})
	contact3 := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "",
		FirstName: "contact",
		LastName:  "222",
	})
	contact4 := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "MR",
		FirstName: "other contact",
		LastName:  "A",
	})

	rawResponse, err := c.RawPost(getQuery("contact/get_contacts_with_sorting"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contacts struct {
		Contacts model.ContactGroupPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contacts)
	require.Nil(t, err)
	require.NotNil(t, contacts.Contacts)
	require.Equal(t, 4, len(contacts.Contacts.Content))
	require.Equal(t, contact3, contacts.Contacts.Content[0].ID)
	require.Equal(t, contact2, contacts.Contacts.Content[1].ID)
	require.Equal(t, contact1, contacts.Contacts.Content[2].ID)
	require.Equal(t, contact4, contacts.Contacts.Content[3].ID)

	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "Contact"))
}

func TestQueryResolver_Contact_BasicFilters_FindContactWithLetterAInName(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	contactFoundByFirstName := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "MR",
		FirstName: "aa",
		LastName:  "bb",
	})
	contactFoundByLastName := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "MR",
		FirstName: "bb",
		LastName:  "AA",
	})
	contactFilteredOut := neo4jt.CreateContact(driver, tenantName, entity.ContactEntity{
		Title:     "MR",
		FirstName: "bb",
		LastName:  "BB",
	})

	require.Equal(t, 3, neo4jt.GetCountOfNodes(driver, "Contact"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contacts_basic_filters"))
	assertRawResponseSuccess(t, rawResponse, err)

	var contacts struct {
		Contacts model.ContactsPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contacts)
	require.Nil(t, err)
	require.NotNil(t, contacts.Contacts)
	require.Equal(t, 2, len(contacts.Contacts.Content))
	require.Equal(t, contactFoundByFirstName, contacts.Contacts.Content[0].ID)
	require.Equal(t, contactFoundByLastName, contacts.Contacts.Content[1].ID)
	require.Equal(t, 1, contacts.Contacts.TotalPages)
	require.Equal(t, int64(2), contacts.Contacts.TotalElements)
	// suppress unused warnings
	require.NotNil(t, contactFilteredOut)
}

func TestQueryResolver_Contact_WithConversations(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	user1 := neo4jt.CreateDefaultUser(driver, tenantName)
	user2 := neo4jt.CreateDefaultUser(driver, tenantName)
	contact1 := neo4jt.CreateDefaultContact(driver, tenantName)
	contact2 := neo4jt.CreateDefaultContact(driver, tenantName)
	contact3 := neo4jt.CreateDefaultContact(driver, tenantName)

	conv1_1 := neo4jt.CreateConversation(driver, user1, contact1)
	conv1_2 := neo4jt.CreateConversation(driver, user1, contact2)
	conv2_1 := neo4jt.CreateConversation(driver, user2, contact1)
	conv2_3 := neo4jt.CreateConversation(driver, user2, contact3)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "User"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "Conversation"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_conversations"),
		client.Var("contactId", contact1))
	assertRawResponseSuccess(t, rawResponse, err)

	var contact struct {
		Contact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)
	require.Equal(t, contact1, contact.Contact.ID)
	require.Equal(t, 1, contact.Contact.Conversations.TotalPages)
	require.Equal(t, int64(2), contact.Contact.Conversations.TotalElements)
	require.Equal(t, 2, len(contact.Contact.Conversations.Content))
	conversations := contact.Contact.Conversations.Content
	require.ElementsMatch(t, []string{conv1_1, conv2_1}, []string{conversations[0].ID, conversations[1].ID})
	require.ElementsMatch(t, []string{user1, user2}, []string{conversations[0].Users[0].ID, conversations[1].Users[0].ID})
	require.Equal(t, contact1, conversations[0].Contacts[0].ID)
	require.Nil(t, conversations[0].Subject)
	require.Equal(t, contact1, conversations[1].Contacts[0].ID)
	require.Nil(t, conversations[1].Subject)

	require.NotNil(t, conv1_2)
	require.NotNil(t, conv2_3)
}

func TestQueryResolver_Contact_WithActions(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, tenantName)
	userId := neo4jt.CreateDefaultUser(driver, tenantName)
	// Use below conversation when conversation is converted to Action
	neo4jt.CreateConversation(driver, userId, contactId)

	now := time.Now().UTC()
	secAgo1 := now.Add(time.Duration(-1) * time.Second)
	secAgo30 := now.Add(time.Duration(-30) * time.Second)
	from := now.Add(time.Duration(-10) * time.Minute)

	pageViewId1 := neo4jt.CreatePageView(driver, contactId, entity.PageViewEntity{
		StartedAt:      secAgo1,
		EndedAt:        now,
		TrackerName:    "tracker1",
		SessionId:      "session1",
		Application:    "application1",
		PageTitle:      "page1",
		PageUrl:        "http://app-1.ai",
		OrderInSession: 1,
		EngagedTime:    10,
	})

	pageViewId2 := neo4jt.CreatePageView(driver, contactId, entity.PageViewEntity{
		StartedAt:      secAgo30,
		EndedAt:        now,
		TrackerName:    "tracker2",
		SessionId:      "session2",
		Application:    "application2",
		PageTitle:      "page2",
		PageUrl:        "http://app-2.ai",
		OrderInSession: 2,
		EngagedTime:    20,
	})

	neo4jt.CreatePageView(driver, contactId2, entity.PageViewEntity{})

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "User"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(driver, "Action"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(driver, "PageView"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Conversation"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_actions"),
		client.Var("contactId", contactId),
		client.Var("from", from),
		client.Var("to", now))
	assertRawResponseSuccess(t, rawResponse, err)

	contact := rawResponse.Data.(map[string]interface{})["contact"]
	require.Equal(t, contactId, contact.(map[string]interface{})["id"])

	actions := contact.(map[string]interface{})["actions"].([]interface{})
	require.Equal(t, 2, len(actions))
	action1 := actions[0].(map[string]interface{})
	require.Equal(t, "PageViewAction", action1["__typename"].(string))
	require.Equal(t, pageViewId1, action1["id"].(string))
	require.NotNil(t, action1["startedAt"].(string))
	require.NotNil(t, action1["endedAt"].(string))
	require.Equal(t, "session1", action1["sessionId"].(string))
	require.Equal(t, "application1", action1["application"].(string))
	require.Equal(t, "page1", action1["pageTitle"].(string))
	require.Equal(t, "http://app-1.ai", action1["pageUrl"].(string))
	require.Equal(t, float64(1), action1["orderInSession"].(float64))
	require.Equal(t, float64(10), action1["engagedTime"].(float64))

	action2 := actions[1].(map[string]interface{})
	require.Equal(t, "PageViewAction", action2["__typename"].(string))
	require.Equal(t, pageViewId2, action2["id"].(string))
	require.NotNil(t, action2["startedAt"].(string))
	require.NotNil(t, action2["endedAt"].(string))
	require.Equal(t, "session2", action2["sessionId"].(string))
	require.Equal(t, "application2", action2["application"].(string))
	require.Equal(t, "page2", action2["pageTitle"].(string))
	require.Equal(t, "http://app-2.ai", action2["pageUrl"].(string))
	require.Equal(t, float64(2), action2["orderInSession"].(float64))
	require.Equal(t, float64(20), action2["engagedTime"].(float64))
}

func TestQueryResolver_Contact_WithActions_FilterByActionType(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	contactId := neo4jt.CreateDefaultContact(driver, tenantName)

	now := time.Now().UTC()
	secAgo1 := now.Add(time.Duration(-1) * time.Second)
	from := now.Add(time.Duration(-10) * time.Minute)

	actionId1 := neo4jt.CreatePageView(driver, contactId, entity.PageViewEntity{
		StartedAt:      secAgo1,
		EndedAt:        now,
		TrackerName:    "tracker1",
		SessionId:      "session1",
		Application:    "application1",
		PageTitle:      "page1",
		PageUrl:        "http://app-1.ai",
		OrderInSession: 1,
		EngagedTime:    10,
	})

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Action"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "PageView"))

	types := []model.ActionType{}
	types = append(types, model.ActionTypePageView)

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_actions_filter_by_action_type"),
		client.Var("contactId", contactId),
		client.Var("from", from),
		client.Var("to", now),
		client.Var("types", types))
	assertRawResponseSuccess(t, rawResponse, err)

	contact := rawResponse.Data.(map[string]interface{})["contact"]
	require.Equal(t, contactId, contact.(map[string]interface{})["id"])

	actions := contact.(map[string]interface{})["actions"].([]interface{})
	require.Equal(t, 1, len(actions))
	action1 := actions[0].(map[string]interface{})
	require.Equal(t, "PageViewAction", action1["__typename"].(string))
	require.Equal(t, actionId1, action1["id"].(string))
}

func TestQueryResolver_Contact_WithOrganizations_ById(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(driver, tenantName)
	organizationId1 := neo4jt.CreateOrganization(driver, tenantName, "organization1")
	organizationId2 := neo4jt.CreateOrganization(driver, tenantName, "organization2")
	organizationId3 := neo4jt.CreateOrganization(driver, tenantName, "organization3")
	organizationId0 := neo4jt.CreateOrganization(driver, tenantName, "organization0")
	neo4jt.LinkContactWithOrganization(driver, contactId, organizationId1)
	neo4jt.LinkContactWithOrganization(driver, contactId, organizationId2)
	neo4jt.LinkContactWithOrganization(driver, contactId, organizationId3)
	neo4jt.LinkContactWithOrganization(driver, contactId2, organizationId0)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(driver, "Organization"))
	require.Equal(t, 4, neo4jt.GetCountOfRelationships(driver, "CONTACT_OF"))

	rawResponse, err := c.RawPost(getQuery("contact/get_contact_with_organizations_by_id"),
		client.Var("contactId", contactId),
		client.Var("limit", 2),
		client.Var("page", 1),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedContact struct {
		Contact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedContact)
	require.Nil(t, err)
	require.Equal(t, contactId, searchedContact.Contact.ID)
	require.Equal(t, 2, searchedContact.Contact.Organizations.TotalPages)
	require.Equal(t, int64(3), searchedContact.Contact.Organizations.TotalElements)

	organizations := searchedContact.Contact.Organizations.Content
	require.Equal(t, 2, len(organizations))
	require.Equal(t, organizationId1, organizations[0].ID)
	require.Equal(t, organizationId2, organizations[1].ID)
}

func TestMutationResolver_ContactAddTagByID(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	tagId1 := neo4jt.CreateTag(driver, tenantName, "tag1")
	tagId2 := neo4jt.CreateTag(driver, tenantName, "tag2")
	neo4jt.TagContact(driver, contactId, tagId1)

	rawResponse, err := c.RawPost(getQuery("contact/add_tag_to_contact"),
		client.Var("contactId", contactId),
		client.Var("tagId", tagId2),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var contactStruct struct {
		Contact_AddTagById model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contactStruct)
	require.Nil(t, err)
	require.NotNil(t, contactStruct)
	tags := contactStruct.Contact_AddTagById.Tags
	require.Equal(t, contactId, contactStruct.Contact_AddTagById.ID)
	require.Equal(t, 2, len(tags))
	require.Equal(t, tagId1, tags[0].ID)
	require.Equal(t, "tag1", tags[0].Name)
	require.Equal(t, tagId2, tags[1].ID)
	require.Equal(t, "tag2", tags[1].Name)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Tag"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "TAGGED"))
}

func TestMutationResolver_ContactRemoveTagByID(t *testing.T) {
	defer tearDownTestCase()(t)
	neo4jt.CreateTenant(driver, tenantName)

	contactId := neo4jt.CreateDefaultContact(driver, tenantName)
	tagId1 := neo4jt.CreateTag(driver, tenantName, "tag1")
	tagId2 := neo4jt.CreateTag(driver, tenantName, "tag2")
	neo4jt.TagContact(driver, contactId, tagId1)
	neo4jt.TagContact(driver, contactId, tagId2)

	require.Equal(t, 2, neo4jt.GetCountOfRelationships(driver, "TAGGED"))

	rawResponse, err := c.RawPost(getQuery("contact/remove_tag_from_contact"),
		client.Var("contactId", contactId),
		client.Var("tagId", tagId2),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var contactStruct struct {
		Contact_RemoveTagById model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contactStruct)
	require.Nil(t, err)
	require.NotNil(t, contactStruct)
	tags := contactStruct.Contact_RemoveTagById.Tags
	require.Equal(t, contactId, contactStruct.Contact_RemoveTagById.ID)
	require.Equal(t, 1, len(tags))
	require.Equal(t, tagId1, tags[0].ID)
	require.Equal(t, "tag1", tags[0].Name)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(driver, "Contact"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(driver, "Tag"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(driver, "TAGGED"))
}

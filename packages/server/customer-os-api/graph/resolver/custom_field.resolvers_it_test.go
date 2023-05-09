package resolver

import (
	"context"
	"github.com/99designs/gqlgen/client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMutationResolver_CustomFieldsMergeAndUpdateInContact(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	contactId := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	entityTemplateId := neo4jt.CreateEntityTemplate(ctx, driver, tenantName, model.EntityTemplateExtensionContact.String())
	fieldTemplateId := neo4jt.AddFieldTemplateToEntity(ctx, driver, entityTemplateId)
	setTemplateId := neo4jt.AddSetTemplateToEntity(ctx, driver, entityTemplateId)
	fieldInSetTemplateId := neo4jt.AddFieldTemplateToSet(ctx, driver, setTemplateId)
	neo4jt.LinkEntityTemplateToContact(ctx, driver, entityTemplateId, contactId)
	fieldInContactId := neo4jt.CreateDefaultCustomFieldInContact(ctx, driver, contactId)
	fieldSetId := neo4jt.CreateDefaultFieldSet(ctx, driver, contactId)
	fieldInSetId := neo4jt.CreateDefaultCustomFieldInSet(ctx, driver, fieldSetId)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact_"+tenantName))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "CustomField"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "FieldSet"))

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "EntityTemplate"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "CustomFieldTemplate"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "FieldSetTemplate"))

	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "IS_DEFINED_BY"))

	rawResponse, err := c.RawPost(getQuery("update_custom_fields_and_filed_sets_in_contact"),
		client.Var("contactId", contactId),
		client.Var("customFieldId", fieldInContactId),
		client.Var("fieldSetId", fieldSetId),
		client.Var("customFieldInSetId", fieldInSetId),
		client.Var("fieldTemplateId", fieldTemplateId),
		client.Var("setTemplateId", setTemplateId),
		client.Var("fieldInSetTemplateId", fieldInSetTemplateId),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Contact_"+tenantName))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "CustomField"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "CustomField_"+tenantName))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "FieldSet"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "FieldSet_"+tenantName))

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "EntityTemplate"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "CustomFieldTemplate"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "FieldSetTemplate"))

	require.Equal(t, 4, neo4jt.GetCountOfRelationships(ctx, driver, "IS_DEFINED_BY"))

	var contact struct {
		CustomFieldsMergeAndUpdateInContact model.Contact
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &contact)
	require.Nil(t, err)
	require.NotNil(t, contact)

	updatedContact := contact.CustomFieldsMergeAndUpdateInContact
	require.Equal(t, entityTemplateId, updatedContact.Template.ID)
	require.Equal(t, 2, len(updatedContact.CustomFields))
	if updatedContact.CustomFields[0].ID == fieldInContactId {
		checkCustomField(t, *updatedContact.CustomFields[0], "field1", "value1", nil)
		checkCustomField(t, *updatedContact.CustomFields[1], "field2", "value2", &fieldTemplateId)
	} else {
		checkCustomField(t, *updatedContact.CustomFields[1], "field1", "value1", nil)
		checkCustomField(t, *updatedContact.CustomFields[0], "field2", "value2", &fieldTemplateId)
	}

	require.Equal(t, 2, len(updatedContact.FieldSets))
	require.Equal(t, model.DataSourceOpenline, updatedContact.FieldSets[0].Source)
	require.Equal(t, model.DataSourceOpenline, updatedContact.FieldSets[1].Source)
	require.ElementsMatch(t, []string{"set1", "set2"}, []string{updatedContact.FieldSets[0].Name, updatedContact.FieldSets[1].Name})

	if updatedContact.FieldSets[0].Template != nil {
		require.Equal(t, fieldInSetId, updatedContact.FieldSets[1].CustomFields[0].ID)
		checkCustomField(t, *updatedContact.FieldSets[1].CustomFields[0], "field3", "value3", nil)
		checkCustomField(t, *updatedContact.FieldSets[0].CustomFields[0], "field4", "value4", &fieldInSetTemplateId)
	} else {
		require.Equal(t, fieldInSetId, updatedContact.FieldSets[0].CustomFields[0].ID)
		checkCustomField(t, *updatedContact.FieldSets[0].CustomFields[0], "field3", "value3", nil)
		checkCustomField(t, *updatedContact.FieldSets[1].CustomFields[0], "field4", "value4", &fieldInSetTemplateId)
	}

	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Contact", "Contact_" + tenantName, "EntityTemplate", "CustomFieldTemplate",
		"FieldSetTemplate", "TextField", "CustomField", "CustomField_" + tenantName, "FieldSet", "FieldSet_" + tenantName})
}

func checkCustomField(t *testing.T, customField model.CustomField, name, value string, fieldTemplateId *string) {
	require.Equal(t, name, customField.Name)
	require.Equal(t, value, customField.Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, customField.Source)
	if fieldTemplateId != nil {
		require.Equal(t, *fieldTemplateId, customField.Template.ID)
	}
}

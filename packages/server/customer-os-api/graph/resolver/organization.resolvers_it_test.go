package resolver

import (
	"context"
	"github.com/99designs/gqlgen/client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/test/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils/decode"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueryResolver_Organizations_FilterByNameLike(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	neo4jt.CreateOrganization(ctx, driver, tenantName, "A closed organization")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "OPENLINE")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "the openline")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "some other open organization")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "OpEnLiNe")

	require.Equal(t, 5, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organizations"),
		client.Var("page", 1),
		client.Var("limit", 3),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizations struct {
		Organizations model.OrganizationPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizations)
	require.Nil(t, err)
	require.NotNil(t, organizations)
	pagedOrganizations := organizations.Organizations
	require.Equal(t, 2, pagedOrganizations.TotalPages)
	require.Equal(t, int64(4), pagedOrganizations.TotalElements)
	require.Equal(t, "OPENLINE", pagedOrganizations.Content[0].Name)
	require.Equal(t, "OpEnLiNe", pagedOrganizations.Content[1].Name)
	require.Equal(t, "some other open organization", pagedOrganizations.Content[2].Name)
}

func TestQueryResolver_Organization(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	organizationInput := entity.OrganizationEntity{
		Name:        "Organization name",
		Description: "Organization description",
		Website:     "Organization_website.com",
		Industry:    "tech",
		IsPublic:    true,
	}
	organizationId1 := neo4jt.CreateFullOrganization(ctx, driver, tenantName, organizationInput)
	neo4jt.AddDomainToOrg(ctx, driver, organizationId1, "domain1.com")
	neo4jt.AddDomainToOrg(ctx, driver, organizationId1, "domain2.com")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "otherOrganization")

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_by_id"),
		client.Var("organizationId", organizationId1),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organization struct {
		Organization model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organization)
	require.Nil(t, err)
	require.NotNil(t, organization)
	require.Equal(t, organizationId1, organization.Organization.ID)
	require.Equal(t, organizationInput.Name, organization.Organization.Name)
	require.Equal(t, organizationInput.Description, *organization.Organization.Description)
	require.Equal(t, []string{"domain1.com", "domain2.com"}, organization.Organization.Domains)
	require.Equal(t, organizationInput.Website, *organization.Organization.Website)
	require.Equal(t, organizationInput.IsPublic, *organization.Organization.IsPublic)
	require.Equal(t, organizationInput.Industry, *organization.Organization.Industry)
	require.NotNil(t, organization.Organization.CreatedAt)
}

func TestQueryResolver_Organizations_WithLocations(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	organizationId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "OPENLINE")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "some other organization")
	locationId1 := neo4jt.CreateLocation(ctx, driver, tenantName, entity.LocationEntity{
		Name:      "WORK",
		Source:    entity.DataSourceOpenline,
		AppSource: "test",
		Country:   "testCountry",
		Region:    "testRegion",
		Locality:  "testLocality",
		Address:   "testAddress",
		Address2:  "testAddress2",
		Zip:       "testZip",
	})
	locationId2 := neo4jt.CreateLocation(ctx, driver, tenantName, entity.LocationEntity{
		Name:      "UNKNOWN",
		Source:    entity.DataSourceOpenline,
		AppSource: "test",
	})
	neo4jt.OrganizationAssociatedWithLocation(ctx, driver, organizationId1, locationId1)
	neo4jt.OrganizationAssociatedWithLocation(ctx, driver, organizationId1, locationId2)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Location"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "ASSOCIATED_WITH"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organizations_with_locations"),
		client.Var("page", 1),
		client.Var("limit", 3),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationsStruct struct {
		Organizations model.OrganizationPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationsStruct)
	require.Nil(t, err)

	organizations := organizationsStruct.Organizations
	require.NotNil(t, organizations)
	require.Equal(t, int64(1), organizations.TotalElements)
	require.Equal(t, 2, len(organizations.Content[0].Locations))

	var locationWithAddressDtls, locationWithoutAddressDtls *model.Location
	if organizations.Content[0].Locations[0].ID == locationId1 {
		locationWithAddressDtls = organizations.Content[0].Locations[0]
		locationWithoutAddressDtls = organizations.Content[0].Locations[1]
	} else {
		locationWithAddressDtls = organizations.Content[0].Locations[1]
		locationWithoutAddressDtls = organizations.Content[0].Locations[0]
	}

	require.Equal(t, locationId1, locationWithAddressDtls.ID)
	require.Equal(t, "WORK", *locationWithAddressDtls.Name)
	require.NotNil(t, locationWithAddressDtls.CreatedAt)
	require.NotNil(t, locationWithAddressDtls.UpdatedAt)
	require.Equal(t, "test", locationWithAddressDtls.AppSource)
	require.Equal(t, model.DataSourceOpenline, locationWithAddressDtls.Source)
	require.Equal(t, "testCountry", *locationWithAddressDtls.Country)
	require.Equal(t, "testLocality", *locationWithAddressDtls.Locality)
	require.Equal(t, "testRegion", *locationWithAddressDtls.Region)
	require.Equal(t, "testAddress", *locationWithAddressDtls.Address)
	require.Equal(t, "testAddress2", *locationWithAddressDtls.Address2)
	require.Equal(t, "testZip", *locationWithAddressDtls.Zip)

	require.Equal(t, locationId2, locationWithoutAddressDtls.ID)
	require.Equal(t, "UNKNOWN", *locationWithoutAddressDtls.Name)
	require.NotNil(t, locationWithoutAddressDtls.CreatedAt)
	require.NotNil(t, locationWithoutAddressDtls.UpdatedAt)
	require.Equal(t, "test", locationWithoutAddressDtls.AppSource)
	require.Equal(t, model.DataSourceOpenline, locationWithoutAddressDtls.Source)
	require.Equal(t, "", *locationWithoutAddressDtls.Country)
	require.Equal(t, "", *locationWithoutAddressDtls.Region)
	require.Equal(t, "", *locationWithoutAddressDtls.Locality)
	require.Equal(t, "", *locationWithoutAddressDtls.Address)
	require.Equal(t, "", *locationWithoutAddressDtls.Address2)
	require.Equal(t, "", *locationWithoutAddressDtls.Zip)
}

func TestQueryResolver_Organizations_WithTags(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	organizationId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "Org 1 with 2 tags")
	organizationId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "Org 2 with 1 tag")
	neo4jt.CreateOrganization(ctx, driver, tenantName, "Org 3 with 0 tags")
	tag1 := neo4jt.CreateTag(ctx, driver, tenantName, "tag1")
	tag2 := neo4jt.CreateTag(ctx, driver, tenantName, "tag2")

	neo4jt.TagOrganization(ctx, driver, organizationId1, tag1)
	neo4jt.TagOrganization(ctx, driver, organizationId1, tag2)
	neo4jt.TagOrganization(ctx, driver, organizationId2, tag1)

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Tag"))
	require.Equal(t, 3, neo4jt.GetCountOfRelationships(ctx, driver, "TAGGED"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organizations_with_tags"),
		client.Var("page", 1),
		client.Var("limit", 10),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationsStruct struct {
		Organizations model.OrganizationPage
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationsStruct)
	require.Nil(t, err)

	organizations := organizationsStruct.Organizations
	require.NotNil(t, organizations)
	require.Equal(t, int64(3), organizations.TotalElements)
	require.Equal(t, 2, len(organizations.Content[0].Tags))
	require.ElementsMatch(t, []string{tag1, tag2},
		[]string{organizations.Content[0].Tags[0].ID, organizations.Content[0].Tags[1].ID})
	require.Equal(t, 1, len(organizations.Content[1].Tags))
	require.Equal(t, tag1, organizations.Content[1].Tags[0].ID)
	require.Equal(t, 0, len(organizations.Content[2].Tags))
}

func TestQueryResolver_Organization_WithNotes_ById(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "test org")
	userId := neo4jt.CreateDefaultUserWithId(ctx, driver, tenantName, testUserId)
	noteId1 := neo4jt.CreateNoteForOrganization(ctx, driver, tenantName, organizationId, "note1", utils.Now())
	noteId2 := neo4jt.CreateNoteForOrganization(ctx, driver, tenantName, organizationId, "note2", utils.Now())
	neo4jt.NoteCreatedByUser(ctx, driver, noteId1, userId)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Note"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "NOTED"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "CREATED"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_notes_by_id"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedOrganization struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedOrganization)
	require.Nil(t, err)
	require.Equal(t, organizationId, searchedOrganization.Organization.ID)

	notes := searchedOrganization.Organization.Notes.Content
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

func TestMutationResolver_OrganizationCreate(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	// Ensure that the tenant and organization type nodes were created in the database.
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetTotalCountOfNodes(ctx, driver))

	// Call the "create_organization" mutation.
	rawResponse := callGraphQL(t, "organization/create_organization", nil)

	// Unmarshal the response data into the "organization" struct.
	var organization struct {
		Organization_Create model.Organization
	}
	err := decode.Decode(rawResponse.Data.(map[string]any), &organization)
	require.Nil(t, err)
	require.NotNil(t, organization)

	// Assign the organization to a shorter variable for easier reference.
	createdOrganization := organization.Organization_Create

	// Ensure that the organization was created correctly.
	require.NotNil(t, createdOrganization.ID)
	require.NotNil(t, createdOrganization.CreatedAt)
	require.NotEqual(t, utils.GetEpochStart(), createdOrganization.CreatedAt)
	require.NotNil(t, createdOrganization.UpdatedAt)
	require.NotEqual(t, utils.GetEpochStart(), createdOrganization.UpdatedAt)
	require.Equal(t, "organization name", createdOrganization.Name)
	require.Equal(t, "organization description", *createdOrganization.Description)
	require.Equal(t, []string{"domain1", "domain2"}, createdOrganization.Domains)
	require.Equal(t, "organization website", *createdOrganization.Website)
	require.Equal(t, "organization industry", *createdOrganization.Industry)
	require.Equal(t, true, *createdOrganization.IsPublic)
	require.Equal(t, int64(10), *createdOrganization.Employees)
	require.Equal(t, model.MarketB2c, *createdOrganization.Market)
	require.Equal(t, model.DataSourceOpenline, createdOrganization.Source)
	require.Equal(t, model.DataSourceOpenline, createdOrganization.SourceOfTruth)
	require.Equal(t, "test", createdOrganization.AppSource)

	// Check the number of nodes and relationships in the Neo4j database
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization_"+tenantName))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Domain"))

	// Check the labels on the nodes in the Neo4j database
	assertNeo4jLabels(ctx, t, driver, []string{"Domain", "Tenant", "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationUpdate(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "some organization")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("organization/update_organization"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var organization struct {
		Organization_Update model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organization)
	require.Nil(t, err)
	require.NotNil(t, organization)
	updatedOrganization := organization.Organization_Update
	require.Equal(t, organizationId, updatedOrganization.ID)
	require.NotNil(t, updatedOrganization.UpdatedAt)
	require.NotEqual(t, utils.GetEpochStart(), updatedOrganization.UpdatedAt)
	require.Equal(t, "updated name", updatedOrganization.Name)
	require.Equal(t, "updated description", *updatedOrganization.Description)
	require.Equal(t, []string{"updated domain"}, updatedOrganization.Domains)
	require.Equal(t, "updated website", *updatedOrganization.Website)
	require.Equal(t, "updated industry", *updatedOrganization.Industry)
	require.Equal(t, true, *updatedOrganization.IsPublic)
	require.Equal(t, int64(100), *updatedOrganization.Employees)
	require.Equal(t, model.MarketB2b, *updatedOrganization.Market)
	require.Equal(t, model.DataSourceOpenline, updatedOrganization.SourceOfTruth)

	// Check still single organization node exists after update, no new node created
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
}

func TestMutationResolver_OrganizationDelete(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "LLC LLC")
	locationId := neo4jt.CreateLocation(ctx, driver, tenantName, entity.LocationEntity{
		Source: "manual",
	})
	neo4jt.OrganizationAssociatedWithLocation(ctx, driver, organizationId, locationId)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Location"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "ASSOCIATED_WITH"))

	rawResponse, err := c.RawPost(getQuery("organization/delete_organization"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var result struct {
		Organization_Delete model.Result
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &result)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, true, result.Organization_Delete.Result)

	require.Equal(t, 0, neo4jt.GetCountOfNodes(ctx, driver, "Location"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "ASSOCIATED_WITH"))

	assertNeo4jLabels(ctx, t, driver, []string{"Tenant"})
}

func TestQueryResolver_Organization_WithRoles_ById(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	contactId1 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "some organization")
	role1 := neo4jt.ContactWorksForOrganization(ctx, driver, contactId1, organizationId, "CTO", false)
	role2 := neo4jt.ContactWorksForOrganization(ctx, driver, contactId2, organizationId, "CEO", true)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "JobRole"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "ROLE_IN"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "WORKS_AS"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_job_roles_by_id"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedOrganization struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedOrganization)
	require.Nil(t, err)
	require.Equal(t, organizationId, searchedOrganization.Organization.ID)

	roles := searchedOrganization.Organization.JobRoles
	require.Equal(t, 2, len(roles))
	var cto, ceo *model.JobRole
	ceo = roles[0]
	cto = roles[1]
	require.Equal(t, role1, cto.ID)
	require.Equal(t, "CTO", *cto.JobTitle)
	require.Equal(t, false, cto.Primary)
	require.Equal(t, contactId1, cto.Contact.ID)

	require.Equal(t, role2, ceo.ID)
	require.Equal(t, "CEO", *ceo.JobTitle)
	require.Equal(t, true, ceo.Primary)
	require.Equal(t, contactId2, ceo.Contact.ID)
}

func TestQueryResolver_Organization_WithContacts_ById(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "organization1")
	organizationId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "organization2")
	contactId1 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	contactId3 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	contactId4 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId1, organizationId)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId2, organizationId)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId3, organizationId)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId4, organizationId2)

	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "JobRole"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 4, neo4jt.GetCountOfRelationships(ctx, driver, "WORKS_AS"))
	require.Equal(t, 4, neo4jt.GetCountOfRelationships(ctx, driver, "ROLE_IN"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_contacts_by_id"),
		client.Var("organizationId", organizationId),
		client.Var("limit", 1),
		client.Var("page", 1),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var searchedOrganization struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &searchedOrganization)
	require.Nil(t, err)
	require.Equal(t, organizationId, searchedOrganization.Organization.ID)
	require.Equal(t, 3, searchedOrganization.Organization.Contacts.TotalPages)
	require.Equal(t, int64(3), searchedOrganization.Organization.Contacts.TotalElements)

	contacts := searchedOrganization.Organization.Contacts.Content
	require.Equal(t, 1, len(contacts))
	require.Equal(t, contactId1, contacts[0].ID)
}

func TestQueryResolver_Organization_WithTimelineEvents_DirectAndFromMultipleContacts(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org1")
	contactId1 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId1, organizationId)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId2, organizationId)

	now := time.Now().UTC()
	secInFuture10 := now.Add(time.Duration(10) * time.Second)
	secAgo10 := now.Add(time.Duration(-10) * time.Second)
	secAgo20 := now.Add(time.Duration(-20) * time.Second)
	secAgo30 := now.Add(time.Duration(-30) * time.Second)
	secAgo50 := now.Add(time.Duration(-50) * time.Second)
	secAgo60 := now.Add(time.Duration(-60) * time.Second)
	secAgo70 := now.Add(time.Duration(-70) * time.Second)
	secAgo80 := now.Add(time.Duration(-80) * time.Second)
	secAgo90 := now.Add(time.Duration(-90) * time.Second)
	secAgo100 := now.Add(time.Duration(-100) * time.Second)
	secAgo1000 := now.Add(time.Duration(-1000) * time.Second)

	// prepare contact and org notes
	contactNoteId1 := neo4jt.CreateNoteForContact(ctx, driver, tenantName, contactId1, "contact note 1", secAgo10)
	contactNoteId2 := neo4jt.CreateNoteForContact(ctx, driver, tenantName, contactId2, "contact note 2", secAgo20)
	orgNoteId3 := neo4jt.CreateNoteForOrganization(ctx, driver, tenantName, organizationId, "org note 1", secAgo30)
	neo4jt.CreateNoteForOrganization(ctx, driver, tenantName, organizationId, "org note 2", secAgo1000)
	neo4jt.CreateNoteForOrganization(ctx, driver, tenantName, organizationId, "org note 3", secInFuture10)

	voiceSession := neo4jt.CreateInteractionSession(ctx, driver, tenantName, "mySessionIdentifier", "session1", "CALL", "ACTIVE", "VOICE", now, false)

	analysis1 := neo4jt.CreateAnalysis(ctx, driver, tenantName, "This is a summary of the conversation", "text/plain", "SUMMARY", secAgo90)
	neo4jt.ActionDescribes(ctx, driver, tenantName, analysis1, voiceSession, repository.DESCRIBES_TYPE_INTERACTION_SESSION)

	// prepare contact and org interaction events
	interactionEventId1 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 1", "application/json", "EMAIL", secAgo50)
	interactionEventId2 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 2", "application/json", "EMAIL", secAgo60)
	interactionEventId3 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 3", "application/json", "EMAIL", secAgo70)
	emailIdContact := neo4jt.AddEmailTo(ctx, driver, entity.CONTACT, tenantName, contactId1, "email1", false, "WORK")
	emailIdOrg := neo4jt.AddEmailTo(ctx, driver, entity.ORGANIZATION, tenantName, organizationId, "email2", false, "WORK")
	phoneNumberId := neo4jt.AddPhoneNumberTo(ctx, driver, tenantName, contactId2, "+1234", false, "WORK")
	neo4jt.InteractionEventSentBy(ctx, driver, interactionEventId1, emailIdContact, "")
	neo4jt.InteractionEventSentTo(ctx, driver, interactionEventId2, phoneNumberId, "")
	neo4jt.InteractionEventSentBy(ctx, driver, interactionEventId3, emailIdOrg, "")
	neo4jt.InteractionEventSentTo(ctx, driver, interactionEventId3, phoneNumberId, "")
	neo4jt.InteractionSessionAttendedBy(ctx, driver, tenantName, voiceSession, phoneNumberId, "")

	// prepare direct interaction events
	interactionEventId4 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 4", "application/json", "", secAgo100)
	neo4jt.InteractionEventSentTo(ctx, driver, interactionEventId4, organizationId, "TO")

	// prepare issue with tags
	issueId1 := neo4jt.CreateIssue(ctx, driver, tenantName, entity.IssueEntity{
		Subject:     "subject 1",
		CreatedAt:   secAgo80,
		Priority:    "P1",
		Status:      "OPEN",
		Description: "description 1",
	})
	tagId1 := neo4jt.CreateTag(ctx, driver, tenantName, "tag1")
	tagId2 := neo4jt.CreateTag(ctx, driver, tenantName, "tag2")
	neo4jt.TagIssue(ctx, driver, issueId1, tagId1)
	neo4jt.TagIssue(ctx, driver, issueId1, tagId2)
	neo4jt.IssueReportedByOrganization(ctx, driver, organizationId, issueId1)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 5, neo4jt.GetCountOfNodes(ctx, driver, "Note"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Issue"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Tag"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "InteractionEvent"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "PhoneNumber"))
	require.Equal(t, 11, neo4jt.GetCountOfNodes(ctx, driver, "TimelineEvent"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_timeline_events_direct_and_via_contacts"),
		client.Var("organizationId", organizationId),
		client.Var("from", now),
		client.Var("size", 9))
	assertRawResponseSuccess(t, rawResponse, err)

	organization := rawResponse.Data.(map[string]interface{})["organization"]
	require.Equal(t, organizationId, organization.(map[string]interface{})["id"])

	timelineEvents := organization.(map[string]interface{})["timelineEvents"].([]interface{})
	require.Equal(t, 9, len(timelineEvents))

	timelineEvent1 := timelineEvents[0].(map[string]interface{})
	require.Equal(t, "Note", timelineEvent1["__typename"].(string))
	require.Equal(t, contactNoteId1, timelineEvent1["id"].(string))
	require.NotNil(t, timelineEvent1["createdAt"].(string))
	require.Equal(t, "contact note 1", timelineEvent1["html"].(string))

	timelineEvent2 := timelineEvents[1].(map[string]interface{})
	require.Equal(t, "Note", timelineEvent2["__typename"].(string))
	require.Equal(t, contactNoteId2, timelineEvent2["id"].(string))
	require.NotNil(t, timelineEvent2["createdAt"].(string))
	require.Equal(t, "contact note 2", timelineEvent2["html"].(string))

	timelineEvent3 := timelineEvents[2].(map[string]interface{})
	require.Equal(t, "Note", timelineEvent3["__typename"].(string))
	require.Equal(t, orgNoteId3, timelineEvent3["id"].(string))
	require.NotNil(t, timelineEvent3["createdAt"].(string))
	require.Equal(t, "org note 1", timelineEvent3["html"].(string))

	timelineEvent4 := timelineEvents[3].(map[string]interface{})
	require.Equal(t, "InteractionEvent", timelineEvent4["__typename"].(string))
	require.Equal(t, interactionEventId1, timelineEvent4["id"].(string))
	require.NotNil(t, timelineEvent4["createdAt"].(string))
	require.Equal(t, "IE text 1", timelineEvent4["content"].(string))

	timelineEvent5 := timelineEvents[4].(map[string]interface{})
	require.Equal(t, "InteractionEvent", timelineEvent5["__typename"].(string))
	require.Equal(t, interactionEventId2, timelineEvent5["id"].(string))
	require.NotNil(t, timelineEvent5["createdAt"].(string))
	require.Equal(t, "IE text 2", timelineEvent5["content"].(string))

	timelineEvent6 := timelineEvents[5].(map[string]interface{})
	require.Equal(t, "InteractionEvent", timelineEvent6["__typename"].(string))
	require.Equal(t, interactionEventId3, timelineEvent6["id"].(string))
	require.NotNil(t, timelineEvent6["createdAt"].(string))
	require.Equal(t, "IE text 3", timelineEvent6["content"].(string))

	timelineEvent7 := timelineEvents[6].(map[string]interface{})
	require.Equal(t, "Issue", timelineEvent7["__typename"].(string))
	require.Equal(t, issueId1, timelineEvent7["id"].(string))
	require.NotNil(t, timelineEvent7["createdAt"].(string))
	require.Equal(t, "subject 1", timelineEvent7["subject"].(string))
	require.Equal(t, "P1", timelineEvent7["priority"].(string))
	require.Equal(t, "OPEN", timelineEvent7["status"].(string))
	require.Equal(t, "description 1", timelineEvent7["description"].(string))
	require.Equal(t, "test", timelineEvent7["appSource"].(string))
	require.Equal(t, "OPENLINE", timelineEvent7["source"].(string))
	require.Equal(t, "OPENLINE", timelineEvent7["sourceOfTruth"].(string))
	require.ElementsMatch(t, []string{tagId1, tagId2},
		[]string{
			timelineEvent7["tags"].([]interface{})[0].(map[string]interface{})["id"].(string),
			timelineEvent7["tags"].([]interface{})[1].(map[string]interface{})["id"].(string)})
	require.ElementsMatch(t, []string{"tag1", "tag2"},
		[]string{
			timelineEvent7["tags"].([]interface{})[0].(map[string]interface{})["name"].(string),
			timelineEvent7["tags"].([]interface{})[1].(map[string]interface{})["name"].(string)})

	timelineEvent8 := timelineEvents[7].(map[string]interface{})
	require.Equal(t, "Analysis", timelineEvent8["__typename"].(string))
	require.Equal(t, "Analysis", timelineEvent8["__typename"].(string))
	require.Equal(t, analysis1, timelineEvent8["id"].(string))
	require.NotNil(t, timelineEvent8["createdAt"].(string))
	require.Equal(t, "This is a summary of the conversation", timelineEvent8["content"].(string))
	require.Equal(t, "text/plain", timelineEvent8["contentType"].(string))
	require.Equal(t, "SUMMARY", timelineEvent8["analysisType"].(string))

	timelineEvent9 := timelineEvents[8].(map[string]interface{})
	require.Equal(t, "InteractionEvent", timelineEvent9["__typename"].(string))
	require.Equal(t, interactionEventId4, timelineEvent9["id"].(string))
	require.NotNil(t, timelineEvent9["createdAt"].(string))
	require.Equal(t, "IE text 4", timelineEvent9["content"].(string))
}

func TestQueryResolver_Organization_WithTimelineEventsTotalCount(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org1")
	contactId1 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	contactId2 := neo4jt.CreateDefaultContact(ctx, driver, tenantName)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId1, organizationId)
	neo4jt.LinkContactWithOrganization(ctx, driver, contactId2, organizationId)

	now := time.Now().UTC()

	// prepare contact amd org notes
	neo4jt.CreateNoteForContact(ctx, driver, tenantName, contactId1, "contact note 1", now)
	neo4jt.CreateNoteForContact(ctx, driver, tenantName, contactId2, "contact note 2", now)
	neo4jt.CreateNoteForOrganization(ctx, driver, tenantName, organizationId, "org note 1", now)

	// prepare contact and org interaction events
	interactionEventId1 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 1", "application/json", "EMAIL", now)
	interactionEventId2 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 2", "application/json", "EMAIL", now)
	interactionEventId3 := neo4jt.CreateInteractionEvent(ctx, driver, tenantName, "myExternalId", "IE text 3", "application/json", "EMAIL", now)
	emailIdContact := neo4jt.AddEmailTo(ctx, driver, entity.CONTACT, tenantName, contactId1, "email1", false, "WORK")
	emailIdOrg := neo4jt.AddEmailTo(ctx, driver, entity.ORGANIZATION, tenantName, organizationId, "email2", false, "WORK")
	phoneNumberId := neo4jt.AddPhoneNumberTo(ctx, driver, tenantName, contactId2, "+1234", false, "WORK")
	neo4jt.InteractionEventSentBy(ctx, driver, interactionEventId1, emailIdContact, "")
	neo4jt.InteractionEventSentTo(ctx, driver, interactionEventId2, phoneNumberId, "")
	neo4jt.InteractionEventSentBy(ctx, driver, interactionEventId3, emailIdOrg, "")
	neo4jt.InteractionEventSentTo(ctx, driver, interactionEventId3, phoneNumberId, "")

	issueId1 := neo4jt.CreateIssue(ctx, driver, tenantName, entity.IssueEntity{
		Subject:     "subject 1",
		CreatedAt:   now,
		Priority:    "P1",
		Status:      "OPEN",
		Description: "description 1",
	})
	neo4jt.IssueReportedByOrganization(ctx, driver, organizationId, issueId1)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Note"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Issue"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "InteractionEvent"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "PhoneNumber"))
	require.Equal(t, 7, neo4jt.GetCountOfNodes(ctx, driver, "TimelineEvent"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_timeline_events_total_count"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	organization := rawResponse.Data.(map[string]interface{})["organization"]
	require.Equal(t, organizationId, organization.(map[string]interface{})["id"])
	require.Equal(t, float64(7), organization.(map[string]interface{})["timelineEventsTotalCount"].(float64))
}

func TestQueryResolver_Organization_WithEmails(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "test org")
	emailId1 := neo4jt.AddEmailTo(ctx, driver, entity.ORGANIZATION, tenantName, organizationId, "email1", true, "MAIN")
	emailId2 := neo4jt.AddEmailTo(ctx, driver, entity.ORGANIZATION, tenantName, organizationId, "email2", false, "WORK")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_emails"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	organization := organizationStruct.Organization

	require.Equal(t, organizationId, organization.ID)
	emails := organization.Emails
	require.Equal(t, 2, len(emails))
	var emailA, emailB *model.Email
	if emailId1 == emails[0].ID {
		emailA = emails[0]
		emailB = emails[1]
	} else {
		emailA = emails[1]
		emailB = emails[0]
	}
	require.Equal(t, emailId1, emailA.ID)
	require.NotNil(t, emailA.CreatedAt)
	require.Equal(t, true, emailA.Primary)
	require.Equal(t, "email1", *emailA.RawEmail)
	require.Equal(t, "email1", *emailA.Email)
	require.Equal(t, model.EmailLabelMain, *emailA.Label)

	require.Equal(t, emailId2, emailB.ID)
	require.NotNil(t, emailB.CreatedAt)
	require.Equal(t, false, emailB.Primary)
	require.Equal(t, "email2", *emailB.RawEmail)
	require.Equal(t, "email2", *emailB.Email)
	require.Equal(t, model.EmailLabelWork, *emailB.Label)
}

func TestQueryResolver_Organization_WithPhoneNumbers(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "test org")
	phoneNumberId1 := neo4jt.AddPhoneNumberTo(ctx, driver, tenantName, organizationId, "+1111", true, "MAIN")
	phoneNumberId2 := neo4jt.AddPhoneNumberTo(ctx, driver, tenantName, organizationId, "+2222", false, "WORK")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "PhoneNumber"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_phone_numbers"),
		client.Var("organizationId", organizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	organization := organizationStruct.Organization

	require.Equal(t, organizationId, organization.ID)
	phoneNumbers := organization.PhoneNumbers
	require.Equal(t, 2, len(phoneNumbers))
	var phoneNumber1, phoneNumber2 *model.PhoneNumber
	if phoneNumberId1 == phoneNumbers[0].ID {
		phoneNumber1 = phoneNumbers[0]
		phoneNumber2 = phoneNumbers[1]
	} else {
		phoneNumber1 = phoneNumbers[1]
		phoneNumber2 = phoneNumbers[0]
	}
	require.Equal(t, phoneNumberId1, phoneNumber1.ID)
	require.NotNil(t, phoneNumber1.CreatedAt)
	require.Equal(t, true, phoneNumber1.Primary)
	require.Equal(t, "+1111", *phoneNumber1.RawPhoneNumber)
	require.Equal(t, "+1111", *phoneNumber1.E164)
	require.Equal(t, model.PhoneNumberLabelMain, *phoneNumber1.Label)

	require.Equal(t, phoneNumberId2, phoneNumber2.ID)
	require.NotNil(t, phoneNumber2.CreatedAt)
	require.Equal(t, false, phoneNumber2.Primary)
	require.Equal(t, "+2222", *phoneNumber2.RawPhoneNumber)
	require.Equal(t, "+2222", *phoneNumber2.E164)
	require.Equal(t, model.PhoneNumberLabelWork, *phoneNumber2.Label)
}

func TestQueryResolver_Organization_WithSubsidiaries(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrganizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "parent org")
	subsidiaryOrganizationId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "sub org 1")
	subsidiaryOrganizationId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "sub org 2")
	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, parentOrganizationId, subsidiaryOrganizationId1, "shop")
	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, parentOrganizationId, subsidiaryOrganizationId2, "station")

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_subsidiaries"),
		client.Var("organizationId", parentOrganizationId))
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	parentOrganization := organizationStruct.Organization

	require.Equal(t, parentOrganizationId, parentOrganization.ID)
	subsidiaries := parentOrganization.Subsidiaries
	require.Equal(t, 2, len(subsidiaries))
	require.Equal(t, subsidiaryOrganizationId1, subsidiaries[0].Organization.ID)
	require.Equal(t, "shop", *subsidiaries[0].Type)
	require.Equal(t, subsidiaryOrganizationId2, subsidiaries[1].Organization.ID)
	require.Equal(t, "station", *subsidiaries[1].Type)
}

func TestQueryResolver_Organization_WithParentForSubsidiary(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrganizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "parent org")
	subsidiaryOrganizationId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "sub org")
	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, parentOrganizationId, subsidiaryOrganizationId1, "shop")

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))

	rawResponse, err := c.RawPost(getQuery("organization/get_organization_with_parent_for_subsidiary"),
		client.Var("organizationId", subsidiaryOrganizationId1))
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	organization := organizationStruct.Organization

	require.Equal(t, subsidiaryOrganizationId1, organization.ID)
	require.Equal(t, 1, len(organization.SubsidiaryOf))
	require.Equal(t, parentOrganizationId, organization.SubsidiaryOf[0].Organization.ID)
	require.Equal(t, "shop", *organization.SubsidiaryOf[0].Type)
}

func TestMutationResolver_OrganizationMerge_Properties(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "main organization")
	mergedOrgId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 1")
	mergedOrgId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 2")

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))

	rawResponse, err := c.RawPost(getQuery("organization/merge_organizations"),
		client.Var("parentOrganizationId", parentOrgId),
		client.Var("mergedOrganizationId1", mergedOrgId1),
		client.Var("mergedOrganizationId2", mergedOrgId2),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization_Merge model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	organization := organizationStruct.Organization_Merge
	require.NotNil(t, organization)

	require.Equal(t, parentOrgId, organization.ID)
	require.Equal(t, "main organization", organization.Name)

	// Check only 1 organization remains after merge
	// other 2 converted into MergedOrganization
	// Other notes not impacted
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "MergedOrganization"))
}

func TestMutationResolver_OrganizationMerge_CheckSubsidiariesMerge(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "main organization")
	mergedOrgId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 1")
	mergedOrgId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 2")

	subsidiaryOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "")
	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, mergedOrgId1, subsidiaryOrgId, "shop")

	parentForSubsidiaryOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "")
	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, parentForSubsidiaryOrgId, mergedOrgId2, "factory")

	require.Equal(t, 5, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))

	rawResponse, err := c.RawPost(getQuery("organization/merge_organizations"),
		client.Var("parentOrganizationId", parentOrgId),
		client.Var("mergedOrganizationId1", mergedOrgId1),
		client.Var("mergedOrganizationId2", mergedOrgId2),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization_Merge model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	organization := organizationStruct.Organization_Merge
	require.NotNil(t, organization)

	require.Equal(t, parentOrgId, organization.ID)
	require.Equal(t, 1, len(organization.Subsidiaries))
	require.Equal(t, subsidiaryOrgId, organization.Subsidiaries[0].Organization.ID)
	require.Equal(t, "shop", *organization.Subsidiaries[0].Type)
	require.Equal(t, 1, len(organization.SubsidiaryOf))
	require.Equal(t, parentForSubsidiaryOrgId, organization.SubsidiaryOf[0].Organization.ID)

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "MergedOrganization"))

	require.Equal(t, 4, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))
}

func TestMutationResolver_OrganizationMerge_MergeBetweenParentAndSubsidiaryOrg(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "main")
	mergedOrgId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 1")
	mergedOrgId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 2")

	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, parentOrgId, mergedOrgId1, "A")
	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, mergedOrgId2, parentOrgId, "B")

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))

	rawResponse, err := c.RawPost(getQuery("organization/merge_organizations"),
		client.Var("parentOrganizationId", parentOrgId),
		client.Var("mergedOrganizationId1", mergedOrgId1),
		client.Var("mergedOrganizationId2", mergedOrgId2),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization_Merge model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	organization := organizationStruct.Organization_Merge
	require.NotNil(t, organization)

	require.Equal(t, parentOrgId, organization.ID)
	require.Equal(t, 0, len(organization.Subsidiaries))
	require.Equal(t, 0, len(organization.SubsidiaryOf))

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "MergedOrganization"))

	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))
}

func TestMutationResolver_OrganizationMerge_CheckRelationshipsAndStages(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "main organization")
	mergedOrgId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 1")
	mergedOrgId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "to merge 2")
	neo4jt.CreateOrganizationRelationship(ctx, driver, "Investor")
	neo4jt.CreateOrganizationRelationship(ctx, driver, "Partner")
	neo4jt.CreateOrganizationRelationship(ctx, driver, "Customer")
	neo4jt.CreateOrganizationRelationshipStages(ctx, driver, tenantName, "Investor", []string{"A", "B", "C"})
	neo4jt.CreateOrganizationRelationshipStages(ctx, driver, tenantName, "Partner", []string{"X"})

	neo4jt.LinkOrganizationWithRelationshipAndStage(ctx, driver, parentOrgId, "Investor", "A")
	neo4jt.LinkOrganizationWithRelationship(ctx, driver, parentOrgId, "Partner")

	neo4jt.LinkOrganizationWithRelationshipAndStage(ctx, driver, mergedOrgId1, "Investor", "B")
	neo4jt.LinkOrganizationWithRelationship(ctx, driver, mergedOrgId1, "Customer")

	neo4jt.LinkOrganizationWithRelationshipAndStage(ctx, driver, mergedOrgId2, "Partner", "X")

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationshipsForNodeWithId(ctx, driver, "IS", parentOrgId))
	require.Equal(t, 1, neo4jt.GetCountOfRelationshipsForNodeWithId(ctx, driver, "HAS_STAGE", parentOrgId))

	rawResponse, err := c.RawPost(getQuery("organization/merge_organizations"),
		client.Var("parentOrganizationId", parentOrgId),
		client.Var("mergedOrganizationId1", mergedOrgId1),
		client.Var("mergedOrganizationId2", mergedOrgId2),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization_Merge model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	organization := organizationStruct.Organization_Merge
	require.NotNil(t, organization)

	require.Equal(t, parentOrgId, organization.ID)
	require.Equal(t, 3, len(organization.Relationships))
	require.Equal(t, model.OrganizationRelationshipCustomer, organization.Relationships[0])
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.Relationships[1])
	require.Equal(t, model.OrganizationRelationshipPartner, organization.Relationships[2])

	require.Equal(t, 3, len(organization.RelationshipStages))
	require.Equal(t, model.OrganizationRelationshipCustomer, organization.RelationshipStages[0].Relationship)
	require.Nil(t, organization.RelationshipStages[0].Stage)
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.RelationshipStages[1].Relationship)
	require.Equal(t, "A", *organization.RelationshipStages[1].Stage)
	require.Equal(t, model.OrganizationRelationshipPartner, organization.RelationshipStages[2].Relationship)
	require.Equal(t, "X", *organization.RelationshipStages[2].Stage)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "MergedOrganization"))
	require.Equal(t, 3, neo4jt.GetCountOfRelationshipsForNodeWithId(ctx, driver, "IS", parentOrgId))
	require.Equal(t, 2, neo4jt.GetCountOfRelationshipsForNodeWithId(ctx, driver, "HAS_STAGE", parentOrgId))
}

func TestMutationResolver_OrganizationAddSubsidiary(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "main")
	subOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "sub")

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))

	rawResponse, err := c.RawPost(getQuery("organization/add_subsidiary"),
		client.Var("organizationId", parentOrgId),
		client.Var("subsidiaryId", subOrgId),
		client.Var("type", "shop"),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization_AddSubsidiary model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	organization := organizationStruct.Organization_AddSubsidiary
	require.NotNil(t, organization)

	require.Equal(t, parentOrgId, organization.ID)
	require.Equal(t, 1, len(organization.Subsidiaries))
	require.Equal(t, subOrgId, organization.Subsidiaries[0].Organization.ID)
	require.Equal(t, "shop", *organization.Subsidiaries[0].Type)

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))
}

func TestMutationResolver_OrganizationRemoveSubsidiary(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	parentOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "main")
	subOrgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "sub")

	neo4jt.LinkOrganizationAsSubsidiary(ctx, driver, parentOrgId, subOrgId, "shop")

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))

	rawResponse, err := c.RawPost(getQuery("organization/remove_subsidiary"),
		client.Var("organizationId", parentOrgId),
		client.Var("subsidiaryId", subOrgId),
	)
	assertRawResponseSuccess(t, rawResponse, err)

	var organizationStruct struct {
		Organization_RemoveSubsidiary model.Organization
	}
	err = decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	organization := organizationStruct.Organization_RemoveSubsidiary
	require.NotNil(t, organization)

	require.Equal(t, parentOrgId, organization.ID)
	require.Equal(t, 0, len(organization.Subsidiaries))

	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "SUBSIDIARY_OF"))
}

func TestMutationResolver_OrganizationCreate_WithCustomFields(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	entityTemplateId := neo4jt.CreateEntityTemplate(ctx, driver, tenantName, model.EntityTemplateExtensionOrganization.String())
	fieldTemplateId := neo4jt.AddFieldTemplateToEntity(ctx, driver, entityTemplateId)
	setTemplateId := neo4jt.AddSetTemplateToEntity(ctx, driver, entityTemplateId)
	fieldInSetTemplateId := neo4jt.AddFieldTemplateToSet(ctx, driver, setTemplateId)

	rawResponse, err := c.RawPost(getQuery("organization/create_organization_with_custom_fields"),
		client.Var("entityTemplateId", entityTemplateId),
		client.Var("fieldTemplateId", fieldTemplateId),
		client.Var("setTemplateId", setTemplateId),
		client.Var("fieldInSetTemplateId", fieldInSetTemplateId))
	assertRawResponseSuccess(t, rawResponse, err)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Tenant"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization_"+tenantName))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(ctx, driver, "Contact"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "CustomField"))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "CustomField_"+tenantName))
	require.Equal(t, 4, neo4jt.GetCountOfNodes(ctx, driver, "TextField"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(ctx, driver, "Email"))
	require.Equal(t, 0, neo4jt.GetCountOfNodes(ctx, driver, "PhoneNumber"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "EntityTemplate"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "CustomFieldTemplate"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "FieldSetTemplate"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "FieldSet"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "FieldSet_"+tenantName))
	require.Equal(t, 13, neo4jt.GetTotalCountOfNodes(ctx, driver))

	var org struct {
		Organization_Create model.Organization
	}

	err = decode.Decode(rawResponse.Data.(map[string]any), &org)
	require.Nil(t, err)

	createdOrganization := org.Organization_Create
	require.Equal(t, createdOrganization.Source, model.DataSourceOpenline)
	require.Equal(t, entityTemplateId, createdOrganization.EntityTemplate.ID)
	require.Equal(t, 2, len(createdOrganization.CustomFields))
	require.Equal(t, "field1", createdOrganization.CustomFields[0].Name)
	require.Equal(t, "TEXT", createdOrganization.CustomFields[0].Datatype.String())
	require.Equal(t, "value1", createdOrganization.CustomFields[0].Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, createdOrganization.CustomFields[0].Source)
	require.Equal(t, fieldTemplateId, createdOrganization.CustomFields[0].Template.ID)
	require.NotNil(t, createdOrganization.CustomFields[0].ID)
	require.NotNil(t, createdOrganization.CustomFields[0].CreatedAt)
	require.Equal(t, "field2", createdOrganization.CustomFields[1].Name)
	require.Equal(t, "TEXT", createdOrganization.CustomFields[1].Datatype.String())
	require.Equal(t, "value2", createdOrganization.CustomFields[1].Value.RealValue())
	require.Equal(t, model.DataSourceOpenline, createdOrganization.CustomFields[1].Source)
	require.NotNil(t, createdOrganization.CustomFields[1].ID)
	require.NotNil(t, createdOrganization.CustomFields[1].CreatedAt)
	require.Equal(t, 2, len(createdOrganization.FieldSets))

	// TODO: fix query in FindAllForFieldSet to return fieldsets also for ORGANIZATION
	//var set1, set2 *model.FieldSet
	//if createdOrganization.FieldSets[0].Name == "set1" {
	//	set1 = createdOrganization.FieldSets[0]
	//	set2 = createdOrganization.FieldSets[1]
	//} else {
	//	set1 = createdOrganization.FieldSets[1]
	//	set2 = createdOrganization.FieldSets[0]
	//}
	//require.NotNil(t, set1.ID)
	//require.NotNil(t, set1.CreatedAt)
	//require.Equal(t, "set1", set1.Name)
	//require.Equal(t, 2, len(set1.CustomFields))
	//require.NotNil(t, set1.CustomFields[0].CreatedAt)
	//require.Equal(t, "field3InSet", set1.CustomFields[0].Name)
	//require.Equal(t, "value3", set1.CustomFields[0].Value.RealValue())
	//require.Equal(t, model.DataSourceOpenline, set1.CustomFields[0].Source)
	//require.Equal(t, "TEXT", set1.CustomFields[0].Datatype.String())
	//require.Equal(t, fieldInSetTemplateId, set1.CustomFields[0].Template.ID)
	//require.NotNil(t, set1.CustomFields[1].CreatedAt)
	//require.Equal(t, "field4InSet", set1.CustomFields[1].Name)
	//require.Equal(t, "value4", set1.CustomFields[1].Value.RealValue())
	//require.Equal(t, model.DataSourceOpenline, set1.CustomFields[1].Source)
	//require.Equal(t, "TEXT", set1.CustomFields[1].Datatype.String())
	//require.Nil(t, set1.CustomFields[1].Template)
	//require.Equal(t, model.DataSourceOpenline, set1.Source)
	//require.NotNil(t, set2.ID)
	//require.NotNil(t, set2.CreatedAt)
	//require.Equal(t, "set2", set2.Name)
	//require.Equal(t, model.DataSourceOpenline, set2.Source)

	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Domain", "Organization", "Organization_" + tenantName,
		"CustomFieldTemplate", "EntityTemplate", "FieldSet", "FieldSet_" + tenantName, "FieldSetTemplate",
		"CustomField", "TextField", "CustomField_" + tenantName})
}

func TestMutationResolver_OrganizationAddNewLocation(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name")

	rawResponse := callGraphQL(t, "organization/add_new_location_to_organization",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_AddNewLocation model.Location
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)
	location := organizationStruct.Organization_AddNewLocation
	require.NotNil(t, location.ID)
	require.NotNil(t, location.CreatedAt)
	require.NotNil(t, location.UpdatedAt)
	require.Equal(t, constants.AppSourceCustomerOsApi, location.AppSource)
	require.Equal(t, model.DataSourceOpenline, location.Source)
	require.Equal(t, model.DataSourceOpenline, location.SourceOfTruth)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Location"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "ASSOCIATED_WITH"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "Location", "Location_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestQueryResolver_Organization_WithSocials(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name")

	socialId1 := neo4jt.CreateSocial(ctx, driver, tenantName, entity.SocialEntity{
		PlatformName: "p1",
		Url:          "url1",
	})
	socialId2 := neo4jt.CreateSocial(ctx, driver, tenantName, entity.SocialEntity{
		PlatformName: "p2",
		Url:          "url2",
	})
	neo4jt.LinkSocialWithEntity(ctx, driver, orgId, socialId1)
	neo4jt.LinkSocialWithEntity(ctx, driver, orgId, socialId2)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "Social"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS"))

	rawResponse := callGraphQL(t, "organization/get_organization_with_socials",
		map[string]interface{}{"organizationId": orgId})

	var orgStruct struct {
		Organization model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &orgStruct)
	require.Nil(t, err)

	organization := orgStruct.Organization
	require.NotNil(t, organization)
	require.Equal(t, 2, len(organization.Socials))

	require.Equal(t, socialId1, organization.Socials[0].ID)
	require.Equal(t, "p1", *organization.Socials[0].PlatformName)
	require.Equal(t, "url1", organization.Socials[0].URL)
	require.NotNil(t, organization.Socials[0].CreatedAt)
	require.NotNil(t, organization.Socials[0].UpdatedAt)
	require.Equal(t, "test", organization.Socials[0].AppSource)

	require.Equal(t, socialId2, organization.Socials[1].ID)
	require.Equal(t, "p2", *organization.Socials[1].PlatformName)
	require.Equal(t, "url2", organization.Socials[1].URL)
	require.NotNil(t, organization.Socials[1].CreatedAt)
	require.NotNil(t, organization.Socials[1].UpdatedAt)
	require.Equal(t, "test", organization.Socials[1].AppSource)
}

func TestMutationResolver_OrganizationSetOwner_NewOwner(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name")

	rawResponse := callGraphQL(t, "organization/set_owner",
		map[string]interface{}{"organizationId": organizationId, "userId": userId})

	var organizationStruct struct {
		Organization_SetOwner model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_SetOwner
	require.Equal(t, organizationId, organization.ID)
	require.Equal(t, userId, organization.Owner.ID)
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "OWNS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "User", "User_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationSetOwner_ReplaceOwner(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	previousOwnerId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	newOwnerId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name")
	neo4jt.UserOwnsOrganization(ctx, driver, previousOwnerId, organizationId)

	rawResponse := callGraphQL(t, "organization/set_owner",
		map[string]interface{}{"organizationId": organizationId, "userId": newOwnerId})

	var organizationStruct struct {
		Organization_SetOwner model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_SetOwner
	require.Equal(t, organizationId, organization.ID)
	require.Equal(t, newOwnerId, organization.Owner.ID)
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "OWNS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "User", "User_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationUnsetOwner(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	ownerId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name")
	neo4jt.UserOwnsOrganization(ctx, driver, ownerId, organizationId)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "OWNS"))

	rawResponse := callGraphQL(t, "organization/unset_owner",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_UnsetOwner model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_UnsetOwner
	require.Equal(t, organizationId, organization.ID)
	require.Nil(t, organization.Owner)
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "OWNS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "User", "User_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestQueryResolver_Organization_WithOwner(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	userId := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	organizationId := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name")
	neo4jt.UserOwnsOrganization(ctx, driver, userId, organizationId)

	rawResponse := callGraphQL(t, "organization/get_organization_with_owner",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization
	require.Equal(t, organizationId, organization.ID)
	require.Equal(t, userId, organization.Owner.ID)
	require.Equal(t, "first", organization.Owner.FirstName)
	require.Equal(t, "last", organization.Owner.LastName)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "OWNS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "User", "User_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationSetOwner_AddRelationship(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	neo4jt.CreateOrganizationRelationship(ctx, driver, entity.Investor.String())
	neo4jt.CreateOrganizationRelationship(ctx, driver, entity.Supplier.String())
	organizationId := neo4jt.CreateDefaultOrganization(ctx, driver, tenantName)
	neo4jt.LinkOrganizationWithRelationship(ctx, driver, organizationId, entity.Supplier.String())

	rawResponse := callGraphQL(t, "organization/add_relationship",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_AddRelationship model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_AddRelationship
	require.Equal(t, organizationId, organization.ID)
	require.Equal(t, 2, len(organization.Relationships))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.Relationships[0])
	require.Equal(t, model.OrganizationRelationshipSupplier, organization.Relationships[1])
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)
	require.Equal(t, 2, len(organization.RelationshipStages))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.RelationshipStages[0].Relationship)
	require.Equal(t, model.OrganizationRelationshipSupplier, organization.RelationshipStages[1].Relationship)
	require.Nil(t, organization.RelationshipStages[0].Stage)
	require.Nil(t, organization.RelationshipStages[1].Stage)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "OrganizationRelationship", "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationSetOwner_RemoveRelationship(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	neo4jt.CreateOrganizationRelationship(ctx, driver, entity.Investor.String())
	neo4jt.CreateOrganizationRelationshipStages(ctx, driver, tenantName, entity.Investor.String(), []string{"stage1", "stage2"})
	organizationId := neo4jt.CreateDefaultOrganization(ctx, driver, tenantName)
	neo4jt.LinkOrganizationWithRelationshipAndStage(ctx, driver, organizationId, entity.Investor.String(), "stage1")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	require.Equal(t, 3, neo4jt.GetCountOfRelationships(ctx, driver, "HAS_STAGE"))

	rawResponse := callGraphQL(t, "organization/remove_relationship",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_RemoveRelationship model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_RemoveRelationship
	require.Equal(t, organizationId, organization.ID)
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)
	require.Equal(t, 0, len(organization.Relationships))
	require.Equal(t, 0, len(organization.RelationshipStages))

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage_"+tenantName))
	require.Equal(t, 0, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS_STAGE"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "OrganizationRelationship", "OrganizationRelationshipStage",
		"OrganizationRelationshipStage_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationSetOwner_SetRelationshipStage_NewRelationshipAndNewStage(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	neo4jt.CreateOrganizationRelationship(ctx, driver, entity.Investor.String())
	neo4jt.CreateOrganizationRelationshipStages(ctx, driver, tenantName, entity.Investor.String(), []string{"Live"})
	organizationId := neo4jt.CreateDefaultOrganization(ctx, driver, tenantName)

	rawResponse := callGraphQL(t, "organization/set_relationship_stage",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_SetRelationshipStage model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_SetRelationshipStage
	require.Equal(t, organizationId, organization.ID)
	require.Equal(t, 1, len(organization.Relationships))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.Relationships[0])
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)
	require.Equal(t, 1, len(organization.RelationshipStages))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.RelationshipStages[0].Relationship)
	require.Equal(t, "Live", *organization.RelationshipStages[0].Stage)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS_STAGE"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "OrganizationRelationship", "OrganizationRelationshipStage",
		"OrganizationRelationshipStage_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationSetOwner_SetRelationshipStage_ReplaceStage(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	neo4jt.CreateOrganizationRelationship(ctx, driver, entity.Investor.String())
	neo4jt.CreateOrganizationRelationshipStages(ctx, driver, tenantName, entity.Investor.String(), []string{"Live", "Lost"})
	organizationId := neo4jt.CreateDefaultOrganization(ctx, driver, tenantName)
	neo4jt.LinkOrganizationWithRelationshipAndStage(ctx, driver, organizationId, entity.Investor.String(), "Lost")

	rawResponse := callGraphQL(t, "organization/set_relationship_stage",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_SetRelationshipStage model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_SetRelationshipStage
	require.Equal(t, organizationId, organization.ID)
	require.Equal(t, 1, len(organization.Relationships))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.Relationships[0])
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)
	require.Equal(t, 1, len(organization.RelationshipStages))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.RelationshipStages[0].Relationship)
	require.Equal(t, "Live", *organization.RelationshipStages[0].Stage)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	require.Equal(t, 3, neo4jt.GetCountOfRelationships(ctx, driver, "HAS_STAGE"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "OrganizationRelationship", "OrganizationRelationshipStage",
		"OrganizationRelationshipStage_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestMutationResolver_OrganizationSetOwner_RemoveRelationshipStage(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	neo4jt.CreateOrganizationRelationship(ctx, driver, entity.Investor.String())
	neo4jt.CreateOrganizationRelationshipStages(ctx, driver, tenantName, entity.Investor.String(), []string{"stage1"})
	organizationId := neo4jt.CreateDefaultOrganization(ctx, driver, tenantName)
	neo4jt.LinkOrganizationWithRelationshipAndStage(ctx, driver, organizationId, entity.Investor.String(), "stage1")

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	require.Equal(t, 2, neo4jt.GetCountOfRelationships(ctx, driver, "HAS_STAGE"))

	rawResponse := callGraphQL(t, "organization/remove_relationship_stage",
		map[string]interface{}{"organizationId": organizationId})

	var organizationStruct struct {
		Organization_RemoveRelationshipStage model.Organization
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &organizationStruct)
	require.Nil(t, err)
	require.NotNil(t, organizationStruct)

	organization := organizationStruct.Organization_RemoveRelationshipStage
	require.Equal(t, organizationId, organization.ID)
	test.AssertTimeRecentlyChanged(t, organization.UpdatedAt)
	require.Equal(t, 1, len(organization.Relationships))
	require.Equal(t, 1, len(organization.RelationshipStages))
	require.Equal(t, model.OrganizationRelationshipInvestor, organization.RelationshipStages[0].Relationship)
	require.Nil(t, organization.RelationshipStages[0].Stage)

	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationship"))
	require.Equal(t, 1, neo4jt.GetCountOfNodes(ctx, driver, "OrganizationRelationshipStage_"+tenantName))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "IS"))
	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, driver, "HAS_STAGE"))
	assertNeo4jLabels(ctx, t, driver, []string{"Tenant", "OrganizationRelationship", "OrganizationRelationshipStage",
		"OrganizationRelationshipStage_" + tenantName, "Organization", "Organization_" + tenantName})
}

func TestQueryResolver_OrganizationDistinctOwners(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx)(t)
	neo4jt.CreateTenant(ctx, driver, tenantName)

	userId1 := neo4jt.CreateDefaultUser(ctx, driver, tenantName)
	userId2 := neo4jt.CreateUser(ctx, driver, tenantName, entity.UserEntity{
		FirstName: "first2",
		LastName:  "last2",
	})
	organizationId1 := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name 1")
	organizationId2 := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name 2")
	organizationId3 := neo4jt.CreateOrganization(ctx, driver, tenantName, "org name 3")
	neo4jt.UserOwnsOrganization(ctx, driver, userId1, organizationId1)
	neo4jt.UserOwnsOrganization(ctx, driver, userId2, organizationId2)
	neo4jt.UserOwnsOrganization(ctx, driver, userId2, organizationId3)

	rawResponse := callGraphQL(t, "organization/get_organization_owners", map[string]interface{}{})

	var usersStruct struct {
		Organization_DistinctOwners []model.User
	}

	err := decode.Decode(rawResponse.Data.(map[string]any), &usersStruct)
	require.Nil(t, err)
	require.NotNil(t, usersStruct)

	users := usersStruct.Organization_DistinctOwners
	require.Equal(t, 2, len(users))
	require.Equal(t, userId1, users[0].ID)
	require.Equal(t, userId2, users[1].ID)

	require.Equal(t, 3, neo4jt.GetCountOfNodes(ctx, driver, "Organization"))
	require.Equal(t, 2, neo4jt.GetCountOfNodes(ctx, driver, "User"))
	require.Equal(t, 3, neo4jt.GetCountOfRelationships(ctx, driver, "OWNS"))
}

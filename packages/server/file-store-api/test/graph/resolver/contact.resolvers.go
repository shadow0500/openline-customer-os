package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"time"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// Tags is the resolver for the tags field.
func (r *contactResolver) Tags(ctx context.Context, obj *model.Contact) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// JobRoles is the resolver for the jobRoles field.
func (r *contactResolver) JobRoles(ctx context.Context, obj *model.Contact) ([]*model.JobRole, error) {
	panic(fmt.Errorf("not implemented: JobRoles - jobRoles"))
}

// Organizations is the resolver for the organizations field.
func (r *contactResolver) Organizations(ctx context.Context, obj *model.Contact, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.OrganizationPage, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *contactResolver) PhoneNumbers(ctx context.Context, obj *model.Contact) ([]*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumbers - phoneNumbers"))
}

// Emails is the resolver for the emails field.
func (r *contactResolver) Emails(ctx context.Context, obj *model.Contact) ([]*model.Email, error) {
	panic(fmt.Errorf("not implemented: Emails - emails"))
}

// Locations is the resolver for the locations field.
func (r *contactResolver) Locations(ctx context.Context, obj *model.Contact) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented: Locations - locations"))
}

// Socials is the resolver for the socials field.
func (r *contactResolver) Socials(ctx context.Context, obj *model.Contact) ([]*model.Social, error) {
	panic(fmt.Errorf("not implemented: Socials - socials"))
}

// CustomFields is the resolver for the customFields field.
func (r *contactResolver) CustomFields(ctx context.Context, obj *model.Contact) ([]*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFields - customFields"))
}

// FieldSets is the resolver for the fieldSets field.
func (r *contactResolver) FieldSets(ctx context.Context, obj *model.Contact) ([]*model.FieldSet, error) {
	panic(fmt.Errorf("not implemented: FieldSets - fieldSets"))
}

// Template is the resolver for the template field.
func (r *contactResolver) Template(ctx context.Context, obj *model.Contact) (*model.EntityTemplate, error) {
	panic(fmt.Errorf("not implemented: Template - template"))
}

// Owner is the resolver for the owner field.
func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Notes is the resolver for the notes field.
func (r *contactResolver) Notes(ctx context.Context, obj *model.Contact, pagination *model.Pagination) (*model.NotePage, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// NotesByTime is the resolver for the notesByTime field.
func (r *contactResolver) NotesByTime(ctx context.Context, obj *model.Contact, pagination *model.TimeRange) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: NotesByTime - notesByTime"))
}

// Conversations is the resolver for the conversations field.
func (r *contactResolver) Conversations(ctx context.Context, obj *model.Contact, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	panic(fmt.Errorf("not implemented: Conversations - conversations"))
}

// TimelineEvents is the resolver for the timelineEvents field.
func (r *contactResolver) TimelineEvents(ctx context.Context, obj *model.Contact, from *time.Time, size int, timelineEventTypes []model.TimelineEventType) ([]model.TimelineEvent, error) {
	panic(fmt.Errorf("not implemented: TimelineEvents - timelineEvents"))
}

// TimelineEventsTotalCount is the resolver for the timelineEventsTotalCount field.
func (r *contactResolver) TimelineEventsTotalCount(ctx context.Context, obj *model.Contact, timelineEventTypes []model.TimelineEventType) (int64, error) {
	panic(fmt.Errorf("not implemented: TimelineEventsTotalCount - timelineEventsTotalCount"))
}

// ContactCreate is the resolver for the contact_Create field.
func (r *mutationResolver) ContactCreate(ctx context.Context, input model.ContactInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactCreate - contact_Create"))
}

// CustomerContactCreate is the resolver for the customer_contact_Create field.
func (r *mutationResolver) CustomerContactCreate(ctx context.Context, input model.CustomerContactInput) (string, error) {
	panic(fmt.Errorf("not implemented: CustomerContactCreate - customer_contact_Create"))
}

// ContactUpdate is the resolver for the contact_Update field.
func (r *mutationResolver) ContactUpdate(ctx context.Context, input model.ContactUpdateInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactUpdate - contact_Update"))
}

// ContactHardDelete is the resolver for the contact_HardDelete field.
func (r *mutationResolver) ContactHardDelete(ctx context.Context, contactID string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: ContactHardDelete - contact_HardDelete"))
}

// ContactArchive is the resolver for the contact_Archive field.
func (r *mutationResolver) ContactArchive(ctx context.Context, contactID string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: ContactArchive - contact_Archive"))
}

// ContactRestoreFromArchive is the resolver for the contact_RestoreFromArchive field.
func (r *mutationResolver) ContactRestoreFromArchive(ctx context.Context, contactID string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: ContactRestoreFromArchive - contact_RestoreFromArchive"))
}

// ContactMerge is the resolver for the contact_Merge field.
func (r *mutationResolver) ContactMerge(ctx context.Context, primaryContactID string, mergedContactIds []string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactMerge - contact_Merge"))
}

// ContactAddTagByID is the resolver for the contact_AddTagById field.
func (r *mutationResolver) ContactAddTagByID(ctx context.Context, input model.ContactTagInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactAddTagByID - contact_AddTagById"))
}

// ContactRemoveTagByID is the resolver for the contact_RemoveTagById field.
func (r *mutationResolver) ContactRemoveTagByID(ctx context.Context, input model.ContactTagInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactRemoveTagByID - contact_RemoveTagById"))
}

// ContactAddOrganizationByID is the resolver for the contact_AddOrganizationById field.
func (r *mutationResolver) ContactAddOrganizationByID(ctx context.Context, input model.ContactOrganizationInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactAddOrganizationByID - contact_AddOrganizationById"))
}

// ContactRemoveOrganizationByID is the resolver for the contact_RemoveOrganizationById field.
func (r *mutationResolver) ContactRemoveOrganizationByID(ctx context.Context, input model.ContactOrganizationInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactRemoveOrganizationByID - contact_RemoveOrganizationById"))
}

// ContactAddNewLocation is the resolver for the contact_AddNewLocation field.
func (r *mutationResolver) ContactAddNewLocation(ctx context.Context, contactID string) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: ContactAddNewLocation - contact_AddNewLocation"))
}

// ContactAddSocial is the resolver for the contact_AddSocial field.
func (r *mutationResolver) ContactAddSocial(ctx context.Context, contactID string, input model.SocialInput) (*model.Social, error) {
	panic(fmt.Errorf("not implemented: ContactAddSocial - contact_AddSocial"))
}

// Contact is the resolver for the contact field.
func (r *queryResolver) Contact(ctx context.Context, id string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: Contact - contact"))
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// ContactByEmail is the resolver for the contact_ByEmail field.
func (r *queryResolver) ContactByEmail(ctx context.Context, email string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactByEmail - contact_ByEmail"))
}

// ContactByPhone is the resolver for the contact_ByPhone field.
func (r *queryResolver) ContactByPhone(ctx context.Context, e164 string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactByPhone - contact_ByPhone"))
}

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

type contactResolver struct{ *Resolver }

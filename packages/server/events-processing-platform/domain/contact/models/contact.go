package models

import (
	"fmt"
	commonModels "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"time"
)

type Contact struct {
	ID           string                        `json:"id"`
	FirstName    string                        `json:"firstName"`
	LastName     string                        `json:"lastName"`
	Prefix       string                        `json:"prefix"`
	Source       commonModels.Source           `json:"source"`
	CreatedAt    time.Time                     `json:"createdAt"`
	UpdatedAt    time.Time                     `json:"updatedAt"`
	PhoneNumbers map[string]ContactPhoneNumber `json:"phoneNumbers"`
	Emails       map[string]ContactEmail       `json:"emails"`
}

type ContactPhoneNumber struct {
	Primary bool   `json:"primary"`
	Label   string `json:"label"`
}

type ContactEmail struct {
	Primary bool   `json:"primary"`
	Label   string `json:"label"`
}

func (contact *Contact) String() string {
	return fmt.Sprintf("Contact{ID: %s, FirstName: %s, LastName: %s, Prefix: %s, Source: %s, CreatedAt: %s, UpdatedAt: %s}", contact.ID, contact.FirstName, contact.LastName, contact.Prefix, contact.Source, contact.CreatedAt, contact.UpdatedAt)
}

func NewContact() *Contact {
	return &Contact{}
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type BooleanResult struct {
	Result bool `json:"result"`
}

// Contact - represents one person that can be contacted for a Customer. In B2C
type Contact struct {
	ID               string             `json:"id"`
	FirstName        string             `json:"firstName"`
	LastName         string             `json:"lastName"`
	CreatedAt        time.Time          `json:"createdAt"`
	Label            *string            `json:"label"`
	Company          *string            `json:"company"`
	Title            *string            `json:"title"`
	ContactType      *string            `json:"contactType"`
	Groups           []*ContactGroup    `json:"groups"`
	TextCustomFields []*TextCustomField `json:"textCustomFields"`
	PhoneNumbers     []*PhoneNumberInfo `json:"phoneNumbers"`
	Emails           []*EmailInfo       `json:"emails"`
}

type ContactGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactGroupInput struct {
	Name string `json:"name"`
}

type ContactInput struct {
	FirstName        string                  `json:"firstName"`
	LastName         string                  `json:"lastName"`
	Label            *string                 `json:"label"`
	Company          *string                 `json:"company"`
	Title            *string                 `json:"title"`
	ContactType      *string                 `json:"contactType"`
	TextCustomFields []*TextCustomFieldInput `json:"textCustomFields"`
	Email            *EmailInput             `json:"email"`
	PhoneNumber      *PhoneNumberInput       `json:"phoneNumber"`
}

type EmailInfo struct {
	Email string     `json:"email"`
	Label EmailLabel `json:"label"`
}

type EmailInput struct {
	Email string     `json:"email"`
	Label EmailLabel `json:"label"`
}

type PhoneNumberInfo struct {
	Number string     `json:"number"`
	Label  PhoneLabel `json:"label"`
}

type PhoneNumberInput struct {
	Number string     `json:"number"`
	Label  PhoneLabel `json:"label"`
}

type TextCustomField struct {
	Group *string `json:"group"`
	Name  string  `json:"name"`
	Value string  `json:"value"`
}

type TextCustomFieldInput struct {
	Group *string `json:"group"`
	Name  string  `json:"name"`
	Value string  `json:"value"`
}

type EmailLabel string

const (
	EmailLabelMain  EmailLabel = "MAIN"
	EmailLabelWork  EmailLabel = "WORK"
	EmailLabelHome  EmailLabel = "HOME"
	EmailLabelOther EmailLabel = "OTHER"
)

var AllEmailLabel = []EmailLabel{
	EmailLabelMain,
	EmailLabelWork,
	EmailLabelHome,
	EmailLabelOther,
}

func (e EmailLabel) IsValid() bool {
	switch e {
	case EmailLabelMain, EmailLabelWork, EmailLabelHome, EmailLabelOther:
		return true
	}
	return false
}

func (e EmailLabel) String() string {
	return string(e)
}

func (e *EmailLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmailLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmailLabel", str)
	}
	return nil
}

func (e EmailLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PhoneLabel string

const (
	PhoneLabelMain   PhoneLabel = "MAIN"
	PhoneLabelWork   PhoneLabel = "WORK"
	PhoneLabelHome   PhoneLabel = "HOME"
	PhoneLabelMobile PhoneLabel = "MOBILE"
	PhoneLabelOther  PhoneLabel = "OTHER"
)

var AllPhoneLabel = []PhoneLabel{
	PhoneLabelMain,
	PhoneLabelWork,
	PhoneLabelHome,
	PhoneLabelMobile,
	PhoneLabelOther,
}

func (e PhoneLabel) IsValid() bool {
	switch e {
	case PhoneLabelMain, PhoneLabelWork, PhoneLabelHome, PhoneLabelMobile, PhoneLabelOther:
		return true
	}
	return false
}

func (e PhoneLabel) String() string {
	return string(e)
}

func (e *PhoneLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhoneLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PhoneLabel", str)
	}
	return nil
}

func (e PhoneLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

package contact

import (
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
)

type Contact struct {
	id          uint
	name        ContactName
	company     *ContactCompany
	phoneNumber *primitive.PhoneNumber
	email       primitive.Email
	content     ContactContent
	createdAt   time.Time
}

func New(name, company, phoneNumber, email, content string) (*Contact, error) {
	contact, err := newWithValidatedValues(name, company, phoneNumber, email, content)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func Rebuild(
	id uint,
	name string,
	company string,
	phoneNumber string,
	email string,
	content string,
	createdAt time.Time,
) (*Contact, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	contact, err := newWithValidatedValues(name, company, phoneNumber, email, content)
	if err != nil {
		return nil, err
	}
	contact.id = id
	contact.createdAt = createdAt
	return contact, nil
}

func newWithValidatedValues(name, company, phoneNumber, email, content string) (*Contact, error) {
	validName, err := NewContactName(name)
	if err != nil {
		return nil, err
	}
	validCompany, err := optional.ParseOptionalString(company, NewContactCompany)
	if err != nil {
		return nil, err
	}
	validPhoneNumber, err := optional.ParseOptionalString(phoneNumber, primitive.NewPhoneNumber)
	if err != nil {
		return nil, err
	}
	validEmail, err := primitive.NewEmail(email)
	if err != nil {
		return nil, err
	}
	validContent, err := NewContactContent(content)
	if err != nil {
		return nil, err
	}

	return &Contact{
		name:        validName,
		company:     validCompany,
		phoneNumber: validPhoneNumber,
		email:       validEmail,
		content:     validContent,
	}, nil
}

func (c *Contact) ID() uint {
	return c.id
}

func (c *Contact) HasID() bool {
	return c.id != 0
}

func (c *Contact) Name() ContactName {
	return c.name
}

func (c *Contact) Company() *ContactCompany {
	return c.company
}

func (c *Contact) PhoneNumber() *primitive.PhoneNumber {
	return c.phoneNumber
}

func (c *Contact) Email() primitive.Email {
	return c.email
}

func (c *Contact) Content() ContactContent {
	return c.content
}

func (c *Contact) CreatedAt() time.Time {
	return c.createdAt
}

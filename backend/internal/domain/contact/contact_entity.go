package contact

import (
	"time"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
)

type Contact struct {
	id          primitive.ID
	uuid        primitive.UUID
	name        ContactName
	company     *ContactCompany
	phoneNumber *primitive.PhoneNumber
	email       primitive.Email
	content     ContactContent
	createdAt   time.Time
}

func New(rawUUID, name, company, phoneNumber, email, content string) (*Contact, error) {
	contact, err := newWithValidatedValues(rawUUID, name, company, phoneNumber, email, content)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func Rebuild(
	id uint,
	rawUUID string,
	name string,
	company string,
	phoneNumber string,
	email string,
	content string,
	createdAt time.Time,
) (*Contact, error) {
	parsedID, err := primitive.NewID(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	contact, err := newWithValidatedValues(rawUUID, name, company, phoneNumber, email, content)
	if err != nil {
		return nil, err
	}
	contact.id = parsedID
	contact.createdAt = createdAt
	return contact, nil
}

func newWithValidatedValues(rawUUID, name, company, phoneNumber, email, content string) (*Contact, error) {
	validUUID, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, ErrInvalidUUID
	}
	validName, err := NewContactName(name)
	if err != nil {
		return nil, err
	}
	validCompany, err := domainVO.ParseOptionalValue(&company, NewContactCompany)
	if err != nil {
		return nil, err
	}
	validPhoneNumber, err := domainVO.ParseOptionalValue(&phoneNumber, primitive.NewPhoneNumber)
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
		uuid:        validUUID,
		name:        validName,
		company:     validCompany,
		phoneNumber: validPhoneNumber,
		email:       validEmail,
		content:     validContent,
	}, nil
}

func (c *Contact) ID() primitive.ID {
	return c.id
}

func (c *Contact) UUID() primitive.UUID {
	return c.uuid
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

package contact

import (
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
)

type Contact struct {
	ID          uint
	Name        ContactName
	Company     *ContactCompany
	PhoneNumber *primitive.PhoneNumber
	Email       primitive.Email
	Content     ContactContent
	CreatedAt   time.Time
}

func New(name, company, phoneNumber, email, content string) (*Contact, error) {
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
		Name:        validName,
		Company:     validCompany,
		PhoneNumber: validPhoneNumber,
		Email:       validEmail,
		Content:     validContent,
	}, nil
}

package presenter

import (
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/contact"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
)

func ToContactResponse(contact *domain.Contact) *response.ContactResponse {
	return &response.ContactResponse{
		ID:          contact.ID(),
		Name:        contact.Name().String(),
		Company:     optional.ToStringPtr(contact.Company()),
		PhoneNumber: optional.ToStringPtr(contact.PhoneNumber()),
		Email:       contact.Email().String(),
		Content:     contact.Content().String(),
		CreatedAt:   contact.CreatedAt(),
	}
}

func ToContactResponses(contactList []*domain.Contact) []*response.ContactResponse {
	res := make([]*response.ContactResponse, 0, len(contactList))
	for _, contact := range contactList {
		res = append(res, ToContactResponse(contact))
	}
	return res
}

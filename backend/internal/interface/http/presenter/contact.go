package presenter

import (
	domain "github.com/tokushun109/tku/backend/internal/domain/contact"
	domainVO "github.com/tokushun109/tku/backend/internal/domain/vo"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
)

func ToContactResponse(contact *domain.Contact) *response.ContactResponse {
	return &response.ContactResponse{
		ID:          contact.ID().Value(),
		Name:        contact.Name().Value(),
		Company:     domainVO.ToValuePtr(contact.Company()),
		PhoneNumber: domainVO.ToValuePtr(contact.PhoneNumber()),
		Email:       contact.Email().Value(),
		Content:     contact.Content().Value(),
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

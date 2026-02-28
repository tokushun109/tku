package presenter

import (
	domainSession "github.com/tokushun109/tku/backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
)

func ToCurrentUserResponse(u *domainUser.User) *response.CurrentUserResponse {
	return &response.CurrentUserResponse{
		UUID:    u.UUID().Value(),
		Name:    u.Name().Value(),
		Email:   u.Email().Value(),
		IsAdmin: u.IsAdmin(),
	}
}

func ToLoginSessionResponse(s *domainSession.Session) *response.LoginSessionResponse {
	return &response.LoginSessionResponse{UUID: s.UUID().Value()}
}

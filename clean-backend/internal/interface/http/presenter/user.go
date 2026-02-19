package presenter

import (
	domainSession "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
)

func ToLoginUserResponse(u *domainUser.User) *response.LoginUserResponse {
	return &response.LoginUserResponse{
		UUID:    u.UUID.String(),
		Name:    u.Name,
		Email:   u.Email,
		IsAdmin: u.IsAdmin,
	}
}

func ToLoginSessionResponse(s *domainSession.Session) *response.LoginSessionResponse {
	return &response.LoginSessionResponse{UUID: s.UUID.String()}
}

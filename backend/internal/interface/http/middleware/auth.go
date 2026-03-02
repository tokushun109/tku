package middleware

import (
	"errors"
	"net/http"

	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseUser "github.com/tokushun109/tku/backend/internal/usecase/user"
)

type AuthMiddleware struct {
	userUC usecaseUser.Usecase
}

func NewAuthMiddleware(userUC usecaseUser.Usecase) *AuthMiddleware {
	return &AuthMiddleware{userUC: userUC}
}

func (m *AuthMiddleware) RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authUser, err := m.authenticate(r)
		if err != nil {
			response.WriteAppError(w, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(ContextWithAuthenticatedUser(r.Context(), authUser)))
	})
}

func (m *AuthMiddleware) OptionalSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authUser, err := m.authenticate(r)
		if err != nil {
			if errors.Is(err, usecase.ErrUnauthorized) {
				next.ServeHTTP(w, r)
				return
			}

			response.WriteAppError(w, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(ContextWithAuthenticatedUser(r.Context(), authUser)))
	})
}

func (m *AuthMiddleware) authenticate(r *http.Request) (AuthenticatedUser, error) {
	var token string
	if cookie, err := r.Cookie("__sess__"); err == nil {
		token = cookie.Value
	}

	user, err := m.userUC.GetBySessionToken(r.Context(), token)
	if err != nil {
		return AuthenticatedUser{}, err
	}
	if user == nil {
		return AuthenticatedUser{}, usecase.NewAppError(usecase.ErrUnauthorized)
	}

	return AuthenticatedUser{
		UUID:         user.UUID().Value(),
		Name:         user.Name().Value(),
		Email:        user.Email().Value(),
		IsAdmin:      user.IsAdmin(),
		SessionToken: token,
	}, nil
}

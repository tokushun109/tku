package middleware

import (
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseUser "github.com/tokushun109/tku/clean-backend/internal/usecase/user"
)

type AuthMiddleware struct {
	userUC usecaseUser.Usecase
}

func NewAuthMiddleware(userUC usecaseUser.Usecase) *AuthMiddleware {
	return &AuthMiddleware{userUC: userUC}
}

func (m *AuthMiddleware) RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		if cookie, err := r.Cookie("__sess__"); err == nil {
			token = cookie.Value
		}

		user, err := m.userUC.GetBySessionToken(r.Context(), token)
		if err != nil {
			response.WriteAppError(w, err)
			return
		}
		if user == nil {
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrUnauthorized))
			return
		}

		ctx := ContextWithAuthenticatedUser(r.Context(), AuthenticatedUser{
			UserID:       user.ID().Uint(),
			UUID:         user.UUID().String(),
			Name:         user.Name().String(),
			Email:        user.Email().String(),
			IsAdmin:      user.IsAdmin(),
			SessionToken: token,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

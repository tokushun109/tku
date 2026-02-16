package middleware

import (
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/interface/http/response"
	usecaseSession "github.com/tokushun109/tku/clean-backend/internal/usecase/session"
)

type AuthMiddleware struct {
	sessionUC usecaseSession.Usecase
}

func NewAuthMiddleware(sessionUC usecaseSession.Usecase) *AuthMiddleware {
	return &AuthMiddleware{sessionUC: sessionUC}
}

func (m *AuthMiddleware) RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		if cookie, err := r.Cookie("__sess__"); err == nil {
			token = cookie.Value
		}
		if err := m.sessionUC.Validate(r.Context(), token); err != nil {
			response.WriteAppError(w, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}

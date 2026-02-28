package middleware

import (
	"net/http"

	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type AdminMiddleware struct{}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authUser, ok := AuthenticatedUserFromContext(r.Context())
		if !ok {
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrUnauthorized))
			return
		}
		if !authUser.IsAdmin {
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrForbidden))
			return
		}
		next.ServeHTTP(w, r)
	})
}

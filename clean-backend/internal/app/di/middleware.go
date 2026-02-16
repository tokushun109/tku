package di

import (
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
)

type middlewares struct {
	auth    *middleware.AuthMiddleware
	logging func(http.Handler) http.Handler
	cors    func(http.Handler) http.Handler
}

func newMiddlewares(cfg *config.Config, ucs *usecases) *middlewares {
	return &middlewares{
		auth:    middleware.NewAuthMiddleware(ucs.session),
		logging: middleware.LoggingMiddleware,
		cors:    middleware.CORSMiddleware([]string{cfg.ClientURL}),
	}
}

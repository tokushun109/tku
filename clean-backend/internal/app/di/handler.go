package di

import (
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
)

type handlers struct {
	health   *handler.HealthHandler
	category *handler.CategoryHandler
	auth     *middleware.AuthMiddleware
}

func newHandlers(ucs *usecases) *handlers {
	return &handlers{
		health:   handler.NewHealthHandler(ucs.health),
		category: handler.NewCategoryHandler(ucs.category),
		auth:     middleware.NewAuthMiddleware(ucs.session),
	}
}

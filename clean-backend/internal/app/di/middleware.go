package di

import (
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
)

type middlewares struct {
	auth             *middleware.AuthMiddleware
	admin            *middleware.AdminMiddleware
	logging          func(http.Handler) http.Handler
	contactRateLimit func(http.Handler) http.Handler
	cors             func(http.Handler) http.Handler
}

func newMiddlewares(cfg *config.Config, ucs *usecases) (*middlewares, error) {
	// 入力側の依存関係のチェック
	if err := requireNonNil("config", cfg); err != nil {
		return nil, err
	}
	if err := requireStructFieldsNonNil("usecases", ucs); err != nil {
		return nil, err
	}

	mws := &middlewares{
		auth:             middleware.NewAuthMiddleware(ucs.user),
		admin:            middleware.NewAdminMiddleware(),
		logging:          middleware.LoggingMiddleware,
		contactRateLimit: middleware.NewRateLimitMiddleware(),
		cors:             middleware.CORSMiddleware([]string{cfg.ClientURL}),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("middlewares", mws); err != nil {
		return nil, err
	}

	return mws, nil
}

package di

import (
	"net/http"

	"github.com/tokushun109/tku/backend/internal/infra/config"
	"github.com/tokushun109/tku/backend/internal/interface/http/middleware"
)

type middlewares struct {
	auth    *middleware.AuthMiddleware
	admin   *middleware.AdminMiddleware
	origin  *middleware.OriginMiddleware
	logging func(http.Handler) http.Handler
	cors    func(http.Handler) http.Handler
}

func newMiddlewares(cfg *config.Config, ucs *usecases) (*middlewares, error) {
	// 入力側の依存関係のチェック
	if err := requireNonNil("config", cfg); err != nil {
		return nil, err
	}
	if err := requireStructFieldsNonNil("usecases", ucs); err != nil {
		return nil, err
	}

	originMW, err := middleware.NewOriginMiddleware(cfg.ClientURL)
	if err != nil {
		return nil, err
	}

	mws := &middlewares{
		auth:    middleware.NewAuthMiddleware(ucs.user),
		admin:   middleware.NewAdminMiddleware(),
		origin:  originMW,
		logging: middleware.LoggingMiddleware,
		cors:    middleware.CORSMiddleware([]string{cfg.ClientURL}),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("middlewares", mws); err != nil {
		return nil, err
	}

	return mws, nil
}

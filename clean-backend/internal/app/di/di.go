package di

import (
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/router"
)

func BuildServer() (*config.Config, http.Handler, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, err
	}

	db, err := mysql.NewDB(cfg)
	if err != nil {
		return nil, nil, err
	}

	repos := newRepositories(db)
	ucs := newUsecases(repos, cfg)
	handlers := newHandlers(ucs)
	middlewares := newMiddlewares(cfg, ucs)

	r := router.NewRouter(
		handlers.health,
		handlers.category,
		handlers.target,
		handlers.tag,
		handlers.salesSite,
		handlers.skillMarket,
		middlewares.auth,
		middlewares.logging,
		middlewares.cors,
	)
	return cfg, r, nil
}

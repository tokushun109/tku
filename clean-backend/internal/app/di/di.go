package di

import (
	"fmt"
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/router"
)

func BuildServer() (*config.Config, http.Handler, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("load config: %w", err)
	}

	db, err := mysql.NewDB(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("build db: %w", err)
	}
	txManager, err := mysql.NewTxManager(db)
	if err != nil {
		return nil, nil, fmt.Errorf("build tx manager: %w", err)
	}

	repos, err := newRepositories(db)
	if err != nil {
		return nil, nil, fmt.Errorf("build repositories: %w", err)
	}

	ucs, err := newUsecases(repos, cfg, txManager)
	if err != nil {
		return nil, nil, fmt.Errorf("build usecases: %w", err)
	}

	handlers, err := newHandlers(ucs)
	if err != nil {
		return nil, nil, fmt.Errorf("build handlers: %w", err)
	}

	middlewares, err := newMiddlewares(cfg, ucs)
	if err != nil {
		return nil, nil, fmt.Errorf("build middlewares: %w", err)
	}

	r := router.NewRouter(
		handlers.health,
		handlers.category,
		handlers.target,
		handlers.tag,
		handlers.sns,
		handlers.salesSite,
		handlers.skillMarket,
		handlers.creator,
		handlers.contact,
		handlers.user,
		middlewares.auth,
		middlewares.admin,
		middlewares.logging,
		middlewares.cors,
	)
	return cfg, r, nil
}

package di

import (
	"net/http"

	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql"
	mysqlRepo "github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/repository"
	uuidInfra "github.com/tokushun109/tku/clean-backend/internal/infra/uuid"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/middleware"
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/router"
	usecaseCategory "github.com/tokushun109/tku/clean-backend/internal/usecase/category"
	usecaseSession "github.com/tokushun109/tku/clean-backend/internal/usecase/session"
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

	categoryRepo := mysqlRepo.NewCategoryRepository(db)
	sessionRepo := mysqlRepo.NewSessionRepository(db)

	uuidGen := uuidInfra.NewGenerator()
	categoryUC := usecaseCategory.New(categoryRepo, uuidGen)
	sessionUC := usecaseSession.New(sessionRepo)

	categoryHandler := handler.NewCategoryHandler(categoryUC)
	auth := middleware.NewAuthMiddleware(sessionUC)

	r := router.NewRouter(cfg, categoryHandler, auth)
	return cfg, r, nil
}

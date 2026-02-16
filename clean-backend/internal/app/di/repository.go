package di

import (
	"github.com/jmoiron/sqlx"
	mysqlRepo "github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/repository"
)

type repositories struct {
	health   *mysqlRepo.HealthRepository
	category *mysqlRepo.CategoryRepository
	session  *mysqlRepo.SessionRepository
}

func newRepositories(db *sqlx.DB) *repositories {
	return &repositories{
		health:   mysqlRepo.NewHealthRepository(db),
		category: mysqlRepo.NewCategoryRepository(db),
		session:  mysqlRepo.NewSessionRepository(db),
	}
}

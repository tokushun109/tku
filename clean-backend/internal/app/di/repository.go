package di

import (
	"github.com/jmoiron/sqlx"
	mysqlRepo "github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/repository"
)

type repositories struct {
	health   *mysqlRepo.HealthRepository
	category *mysqlRepo.CategoryRepository
	target   *mysqlRepo.TargetRepository
	tag      *mysqlRepo.TagRepository
	session  *mysqlRepo.SessionRepository
}

func newRepositories(db *sqlx.DB) *repositories {
	return &repositories{
		health:   mysqlRepo.NewHealthRepository(db),
		category: mysqlRepo.NewCategoryRepository(db),
		target:   mysqlRepo.NewTargetRepository(db),
		tag:      mysqlRepo.NewTagRepository(db),
		session:  mysqlRepo.NewSessionRepository(db),
	}
}

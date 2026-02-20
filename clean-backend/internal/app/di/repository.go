package di

import (
	"github.com/jmoiron/sqlx"
	mysqlRepo "github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/repository"
)

type repositories struct {
	health      *mysqlRepo.HealthRepository
	category    *mysqlRepo.CategoryRepository
	target      *mysqlRepo.TargetRepository
	tag         *mysqlRepo.TagRepository
	sns         *mysqlRepo.SnsRepository
	salesSite   *mysqlRepo.SalesSiteRepository
	skillMarket *mysqlRepo.SkillMarketRepository
	contact     *mysqlRepo.ContactRepository
	session     *mysqlRepo.SessionRepository
	user        *mysqlRepo.UserRepository
}

func newRepositories(db *sqlx.DB) (*repositories, error) {
	// 入力側の依存関係のチェック
	if err := requireNonNil("db", db); err != nil {
		return nil, err
	}

	repos := &repositories{
		health:      mysqlRepo.NewHealthRepository(db),
		category:    mysqlRepo.NewCategoryRepository(db),
		target:      mysqlRepo.NewTargetRepository(db),
		tag:         mysqlRepo.NewTagRepository(db),
		sns:         mysqlRepo.NewSnsRepository(db),
		salesSite:   mysqlRepo.NewSalesSiteRepository(db),
		skillMarket: mysqlRepo.NewSkillMarketRepository(db),
		contact:     mysqlRepo.NewContactRepository(db),
		session:     mysqlRepo.NewSessionRepository(db),
		user:        mysqlRepo.NewUserRepository(db),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("repositories", repos); err != nil {
		return nil, err
	}

	return repos, nil
}

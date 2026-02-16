package di

import (
	clockInfra "github.com/tokushun109/tku/clean-backend/internal/infra/clock"
	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	uuidInfra "github.com/tokushun109/tku/clean-backend/internal/infra/uuid"
	usecaseCategory "github.com/tokushun109/tku/clean-backend/internal/usecase/category"
	usecaseHealth "github.com/tokushun109/tku/clean-backend/internal/usecase/health"
	usecaseSalesSite "github.com/tokushun109/tku/clean-backend/internal/usecase/sales_site"
	usecaseSession "github.com/tokushun109/tku/clean-backend/internal/usecase/session"
	usecaseSkillMarket "github.com/tokushun109/tku/clean-backend/internal/usecase/skill_market"
	usecaseTag "github.com/tokushun109/tku/clean-backend/internal/usecase/tag"
	usecaseTarget "github.com/tokushun109/tku/clean-backend/internal/usecase/target"
)

type usecases struct {
	health      usecaseHealth.Usecase
	category    usecaseCategory.Usecase
	target      usecaseTarget.Usecase
	tag         usecaseTag.Usecase
	salesSite   usecaseSalesSite.Usecase
	skillMarket usecaseSkillMarket.Usecase
	session     usecaseSession.Usecase
}

func newUsecases(repos *repositories, cfg *config.Config) *usecases {
	uuidGen := uuidInfra.NewGenerator()
	clock := clockInfra.NewClock()
	return &usecases{
		health:      usecaseHealth.New(repos.health),
		category:    usecaseCategory.New(repos.category, uuidGen),
		target:      usecaseTarget.New(repos.target, uuidGen),
		tag:         usecaseTag.New(repos.tag, uuidGen),
		salesSite:   usecaseSalesSite.New(repos.salesSite, uuidGen),
		skillMarket: usecaseSkillMarket.New(repos.skillMarket, uuidGen),
		session:     usecaseSession.New(repos.session, cfg.SessionTTL, clock),
	}
}

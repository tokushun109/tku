package di

import (
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
)

type handlers struct {
	health      *handler.HealthHandler
	category    *handler.CategoryHandler
	target      *handler.TargetHandler
	tag         *handler.TagHandler
	sns         *handler.SnsHandler
	salesSite   *handler.SalesSiteHandler
	skillMarket *handler.SkillMarketHandler
}

func newHandlers(ucs *usecases) *handlers {
	return &handlers{
		health:      handler.NewHealthHandler(ucs.health),
		category:    handler.NewCategoryHandler(ucs.category),
		target:      handler.NewTargetHandler(ucs.target),
		tag:         handler.NewTagHandler(ucs.tag),
		sns:         handler.NewSnsHandler(ucs.sns),
		salesSite:   handler.NewSalesSiteHandler(ucs.salesSite),
		skillMarket: handler.NewSkillMarketHandler(ucs.skillMarket),
	}
}

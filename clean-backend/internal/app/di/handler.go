package di

import (
	"github.com/tokushun109/tku/clean-backend/internal/interface/http/handler"
)

type handlers struct {
	health       *handler.HealthHandler
	category     *handler.CategoryHandler
	target       *handler.TargetHandler
	tag          *handler.TagHandler
	sns          *handler.SnsHandler
	salesSite    *handler.SalesSiteHandler
	skillMarket  *handler.SkillMarketHandler
	creator      *handler.CreatorHandler
	contact      *handler.ContactHandler
	user         *handler.UserHandler
	product      *handler.ProductHandler
	productImage *handler.ProductImageHandler
}

func newHandlers(ucs *usecases) (*handlers, error) {
	// 入力側の依存関係のチェック
	if err := requireStructFieldsNonNil("usecases", ucs); err != nil {
		return nil, err
	}

	hs := &handlers{
		health:       handler.NewHealthHandler(ucs.health),
		category:     handler.NewCategoryHandler(ucs.category),
		target:       handler.NewTargetHandler(ucs.target),
		tag:          handler.NewTagHandler(ucs.tag),
		sns:          handler.NewSnsHandler(ucs.sns),
		salesSite:    handler.NewSalesSiteHandler(ucs.salesSite),
		skillMarket:  handler.NewSkillMarketHandler(ucs.skillMarket),
		creator:      handler.NewCreatorHandler(ucs.creator),
		contact:      handler.NewContactHandler(ucs.contact),
		user:         handler.NewUserHandler(ucs.user),
		product:      handler.NewProductHandler(ucs.product),
		productImage: handler.NewProductImageHandler(ucs.product),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("handlers", hs); err != nil {
		return nil, err
	}

	return hs, nil
}

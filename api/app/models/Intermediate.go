package models

type ProductToMaterialCategory struct {
	DefaultModel
	ProductId          *uint `json:"-"`
	MaterialCategoryId *uint `json:"-"`
}

type ProductToSalesSite struct {
	DefaultModel
	ProductId    *uint `json:"-"`
	SalesSiteId *uint `json:"-"`
}

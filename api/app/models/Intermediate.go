package models

type ProductToTag struct {
	DefaultModel
	ProductId *uint `json:"-"`
	TagId     *uint `json:"-"`
}

type ProductToSalesSite struct {
	DefaultModel
	ProductId   *uint `json:"-"`
	SalesSiteId *uint `json:"-"`
}

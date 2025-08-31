package models

type ProductToTag struct {
	DefaultModel
	ProductId *uint `json:"-"`
	TagId     *uint `json:"-"`
}

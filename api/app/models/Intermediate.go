package models

type ProductToMaterialCategory struct {
	DefaultModel
	ProductId          *uint `json:"-"`
	MaterialCategoryId *uint `json:"-"`
}

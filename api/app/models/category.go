package models

type AccessoryCategory struct {
	DefaultModel
	Name string `json:"name"`
}

type AccessoryCategories []AccessoryCategory

type MaterialCategory struct {
	DefaultModel
	Name string `json:"name"`
}

type MaterialCategories []MaterialCategory

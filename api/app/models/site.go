package models

type SalesSite struct {
	DefaultModel
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SkillMarket struct {
	DefaultModel
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Sns struct {
	DefaultModel
	Name string `json:"name"`
	Url  string `json:"url"`
}

package models

type SalesSite struct {
	DefaultModel
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SalesSites []SalesSite

type SkillMarket struct {
	DefaultModel
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SkillMarkets []SkillMarket

type Sns struct {
	DefaultModel
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SnsList []Sns

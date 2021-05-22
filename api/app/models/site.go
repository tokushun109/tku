package models

type SalesSite struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SalesSites []SalesSite

type SkillMarket struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SkillMarkets []SkillMarket

type Sns struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SnsList []Sns

func GetAllSalesSites() (salesSites SalesSites) {
	Db.Find(&salesSites)
	return salesSites
}

func GetAllSkillMarkets() (skillMarkets SkillMarkets) {
	Db.Find(&skillMarkets)
	return skillMarkets
}

func GetAllSnsList() (snsList SnsList) {
	Db.Find(&snsList)
	return snsList
}

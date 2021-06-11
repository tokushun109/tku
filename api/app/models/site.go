package models

import "log"

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

func GetSalesSite(uuid string) (salesSite SalesSite) {
	Db.First(&salesSite, "uuid = ?", uuid)
	return salesSite
}

func InsertSalesSite(salesSite *SalesSite) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	salesSite.Uuid = uuid

	Db.NewRecord(salesSite)
	Db.Create(&salesSite)
}

func GetAllSkillMarkets() (skillMarkets SkillMarkets) {
	Db.Find(&skillMarkets)
	return skillMarkets
}

func InsertSkillMarket(skillMarket *SkillMarket) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	skillMarket.Uuid = uuid

	Db.NewRecord(skillMarket)
	Db.Create(&skillMarket)
}

func GetAllSnsList() (snsList SnsList) {
	Db.Find(&snsList)
	return snsList
}

func InsertSns(sns *Sns) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	sns.Uuid = uuid

	Db.NewRecord(sns)
	Db.Create(&sns)
}

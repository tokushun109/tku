package models

type SalesSite struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
	Url  string `json:"url" validate:"url"`
}

type SalesSites []SalesSite

type SkillMarket struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
	Url  string `json:"url" validate:"url"`
}

type SkillMarkets []SkillMarket

type Sns struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
	Url  string `json:"url" validate:"url"`
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

func InsertSalesSite(salesSite *SalesSite) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	salesSite.Uuid = uuid

	err = Db.Create(&salesSite).Error
	return err
}

func GetAllSkillMarkets() (skillMarkets SkillMarkets) {
	Db.Find(&skillMarkets)
	return skillMarkets
}

func InsertSkillMarket(skillMarket *SkillMarket) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	skillMarket.Uuid = uuid

	err = Db.Create(&skillMarket).Error
	return err
}

func GetAllSnsList() (snsList SnsList) {
	Db.Find(&snsList)
	return snsList
}

func InsertSns(sns *Sns) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	sns.Uuid = uuid

	err = Db.Create(&sns).Error
	return err
}

package models

import "errors"

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

func SalesSiteUniqueCheck(name string) (isUnique bool, err error) {
	var salesSite SalesSite
	Db.First(&salesSite, "name = ?", name)
	isUnique = salesSite.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
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

func UpdateSalesSite(sales_site *SalesSite, uuid string) (err error) {
	err = Db.Model(&sales_site).Where("uuid = ?", uuid).Updates(
		SalesSite{Name: sales_site.Name},
	).Error
	return err
}

func (salesSite *SalesSite) DeleteSalesSite() (err error) {
	err = Db.Delete(&salesSite).Error
	return err
}

func GetAllSkillMarkets() (skillMarkets SkillMarkets) {
	Db.Find(&skillMarkets)
	return skillMarkets
}

func GetSkillMarket(uuid string) (skillMarket SkillMarket) {
	Db.First(&skillMarket, "uuid = ?", uuid)
	return skillMarket
}

func SkillMarketUniqueCheck(name string) (isUnique bool, err error) {
	var skillMarket SkillMarket
	Db.First(&skillMarket, "name = ?", name)
	isUnique = skillMarket.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
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

func UpdateSkillMarket(skill_market *SkillMarket, uuid string) (err error) {
	err = Db.Model(&skill_market).Where("uuid = ?", uuid).Updates(
		SkillMarket{Name: skill_market.Name},
	).Error
	return err
}

func (skillMarket *SkillMarket) DeleteSkillMarket() (err error) {
	err = Db.Delete(&skillMarket).Error
	return err
}

func GetAllSnsList() (snsList SnsList) {
	Db.Find(&snsList)
	return snsList
}

func GetSns(uuid string) (sns Sns) {
	Db.First(&sns, "uuid = ?", uuid)
	return sns
}

func SnsUniqueCheck(name string) (isUnique bool, err error) {
	var sns Sns
	Db.First(&sns, "name = ?", name)
	isUnique = sns.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
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

func UpdateSns(sns *Sns, uuid string) (err error) {
	err = Db.Model(&sns).Where("uuid = ?", uuid).Updates(
		Sns{Name: sns.Name},
	).Error
	return err
}

func (sns *Sns) DeleteSns() (err error) {
	err = Db.Delete(&sns).Error
	return err
}

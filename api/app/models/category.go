package models

import "log"

type AccessoryCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type AccessoryCategories []AccessoryCategory

type MaterialCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type MaterialCategories []MaterialCategory

func GetAllAccessoryCategories() (accessoryCategories AccessoryCategories) {
	Db.Find(&accessoryCategories)
	return accessoryCategories
}

func GetAccessoryCategory(uuid string) (accessoryCategory AccessoryCategory) {
	Db.First(&accessoryCategory, "uuid = ?", uuid)
	return accessoryCategory
}

func InsertAccessoryCategory(accessoryCategory *AccessoryCategory) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	accessoryCategory.Uuid = uuid

	Db.NewRecord(accessoryCategory)
	Db.Create(&accessoryCategory)
}

func GetAllMaterialCategories() (materialCategories MaterialCategories) {
	Db.Find(&materialCategories)
	return materialCategories
}

func GetMaterialCategory(uuid string) (materialCategory MaterialCategory) {
	Db.First(&materialCategory, "uuid = ?", uuid)
	return materialCategory
}

func InsertMaterialCategory(materialCategory *MaterialCategory) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	materialCategory.Uuid = uuid
	Db.NewRecord(materialCategory)
	Db.Create(&materialCategory)
}

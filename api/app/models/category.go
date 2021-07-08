package models

import "errors"

type AccessoryCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type AccessoryCategories []AccessoryCategory

type MaterialCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
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

func AccessoryCategoryUniqueCheck(name string) (isUnique bool, err error) {
	var accessoryCategory AccessoryCategory
	Db.First(&accessoryCategory, "name = ?", name)
	isUnique = accessoryCategory.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertAccessoryCategory(accessoryCategory *AccessoryCategory) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	accessoryCategory.Uuid = uuid
	err = Db.Create(&accessoryCategory).Error
	return err
}

func (accessoryCategory *AccessoryCategory) DeleteAccessoryCategory() (err error) {
	err = Db.Delete(&accessoryCategory).Error
	return err
}

func GetAllMaterialCategories() (materialCategories MaterialCategories) {
	Db.Find(&materialCategories)
	return materialCategories
}

func GetMaterialCategory(uuid string) (materialCategory MaterialCategory) {
	Db.First(&materialCategory, "uuid = ?", uuid)
	return materialCategory
}

func MaterialCategoryUniqueCheck(name string) (isUnique bool, err error) {
	var materialCategory MaterialCategory
	Db.First(&materialCategory, "name = ?", name)
	isUnique = materialCategory.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertMaterialCategory(materialCategory *MaterialCategory) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	materialCategory.Uuid = uuid
	err = Db.Create(&materialCategory).Error
	return err
}

func (materialCategory *MaterialCategory) DeleteMaterialCategory() (err error) {
	err = Db.Delete(&materialCategory).Error
	return err
}

package models

import (
	"errors"
)

type AccessoryCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type AccessoryCategories []AccessoryCategory

type Tag struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type Tags []Tag

func GetAllAccessoryCategories() (accessoryCategories AccessoryCategories) {
	Db.Find(&accessoryCategories)
	return accessoryCategories
}

func GetAccessoryCategory(uuid string) (accessoryCategory AccessoryCategory) {
	Db.Limit(1).Find(&accessoryCategory, "uuid = ?", uuid)
	return accessoryCategory
}

func AccessoryCategoryUniqueCheck(name string) (isUnique bool, err error) {
	var accessoryCategory AccessoryCategory
	Db.Limit(1).Find(&accessoryCategory, "name = ?", name)
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

func UpdateAccessoryCategory(accessoryCategory *AccessoryCategory, uuid string) (err error) {
	err = Db.Model(&accessoryCategory).Where("uuid = ?", uuid).Updates(
		AccessoryCategory{Name: accessoryCategory.Name},
	).Error
	return err
}

func (accessoryCategory *AccessoryCategory) DeleteAccessoryCategory() (err error) {
	err = Db.Delete(&accessoryCategory).Error
	return err
}

func GetAllTags() (materialCategories Tags) {
	Db.Find(&materialCategories)
	return materialCategories
}

func GetTag(uuid string) (materialCategory Tag) {
	Db.Limit(1).Find(&materialCategory, "uuid = ?", uuid)
	return materialCategory
}

func TagUniqueCheck(name string) (isUnique bool, err error) {
	var materialCategory Tag
	Db.Limit(1).Find(&materialCategory, "name = ?", name)
	isUnique = materialCategory.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertTag(materialCategory *Tag) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	materialCategory.Uuid = uuid
	err = Db.Create(&materialCategory).Error
	return err
}

func UpdateTag(materialCategory *Tag, uuid string) (err error) {
	err = Db.Model(&materialCategory).Where("uuid = ?", uuid).Updates(
		Tag{Name: materialCategory.Name},
	).Error
	return err
}

func (materialCategory *Tag) DeleteTag() (err error) {
	err = Db.Delete(&materialCategory).Error
	return err
}

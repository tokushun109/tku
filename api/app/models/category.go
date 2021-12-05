package models

import (
	"errors"
)

type Category struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type AccessoryCategories []Category

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

func GetCategory(uuid string) (accessoryCategory Category) {
	Db.Limit(1).Find(&accessoryCategory, "uuid = ?", uuid)
	return accessoryCategory
}

func CategoryUniqueCheck(name string) (isUnique bool, err error) {
	var accessoryCategory Category
	Db.Limit(1).Find(&accessoryCategory, "name = ?", name)
	isUnique = accessoryCategory.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertCategory(accessoryCategory *Category) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	accessoryCategory.Uuid = uuid
	err = Db.Create(&accessoryCategory).Error
	return err
}

func UpdateCategory(accessoryCategory *Category, uuid string) (err error) {
	err = Db.Model(&accessoryCategory).Where("uuid = ?", uuid).Updates(
		Category{Name: accessoryCategory.Name},
	).Error
	return err
}

func (accessoryCategory *Category) DeleteCategory() (err error) {
	err = Db.Delete(&accessoryCategory).Error
	return err
}

func GetAllTags() (tags Tags) {
	Db.Find(&tags)
	return tags
}

func GetTag(uuid string) (tag Tag) {
	Db.Limit(1).Find(&tag, "uuid = ?", uuid)
	return tag
}

func TagUniqueCheck(name string) (isUnique bool, err error) {
	var tag Tag
	Db.Limit(1).Find(&tag, "name = ?", name)
	isUnique = tag.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertTag(tag *Tag) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	tag.Uuid = uuid
	err = Db.Create(&tag).Error
	return err
}

func UpdateTag(tag *Tag, uuid string) (err error) {
	err = Db.Model(&tag).Where("uuid = ?", uuid).Updates(
		Tag{Name: tag.Name},
	).Error
	return err
}

func (tag *Tag) DeleteTag() (err error) {
	err = Db.Delete(&tag).Error
	return err
}

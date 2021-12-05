package models

import (
	"errors"
)

type Category struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type Categories []Category

type Tag struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type Tags []Tag

func GetAllCategories() (categories Categories) {
	Db.Find(&categories)
	return categories
}

func GetCategory(uuid string) (category Category) {
	Db.Limit(1).Find(&category, "uuid = ?", uuid)
	return category
}

func CategoryUniqueCheck(name string) (isUnique bool, err error) {
	var category Category
	Db.Limit(1).Find(&category, "name = ?", name)
	isUnique = category.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertCategory(category *Category) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	category.Uuid = uuid
	err = Db.Create(&category).Error
	return err
}

func UpdateCategory(category *Category, uuid string) (err error) {
	err = Db.Model(&category).Where("uuid = ?", uuid).Updates(
		Category{Name: category.Name},
	).Error
	return err
}

func (category *Category) DeleteCategory() (err error) {
	err = Db.Delete(&category).Error
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

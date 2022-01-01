package models

import (
	"errors"

	"gorm.io/gorm"
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
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 商品の外部キーをnullにする
	if err = tx.Model(&Product{}).
		Where("category_id = ?", GetCategory(category.Uuid).ID).
		Update("category_id", gorm.Expr("NULL")).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	err = Db.Delete(&category).Error
	return tx.Commit().Error
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
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 登録されている中間テーブルを全て物理削除する
	if err = tx.Where("tag_id = ?", GetTag(tag.Uuid).ID).
		Unscoped().
		Delete(&ProductToTag{}).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&tag).Error
	return tx.Commit().Error
}
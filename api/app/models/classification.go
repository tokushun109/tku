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
	db := GetDBConnection()
	db.Find(&categories)
	return categories
}

func GetUsedCategories() (categories Categories) {
	db := GetDBConnection()
	db.Joins("INNER JOIN product on product.category_id = category.id").
		Where("product.deleted_at IS NULL").
		Group("category.id").
		Find(&categories)
	return categories
}

func GetCategory(uuid string) (category Category) {
	db := GetDBConnection()
	db.Limit(1).Find(&category, "uuid = ?", uuid)
	return category
}

func CategoryUniqueCheck(name string) (isUnique bool, err error) {
	var category Category
	db := GetDBConnection()
	db.Limit(1).Find(&category, "name = ?", name)
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
	db := GetDBConnection()
	err = db.Create(&category).Error
	return err
}

func UpdateCategory(category *Category, uuid string) (err error) {
	db := GetDBConnection()
	err = db.Model(&category).Where("uuid = ?", uuid).Updates(
		Category{Name: category.Name},
	).Error
	return err
}

func (category *Category) DeleteCategory() (err error) {
	db := GetDBConnection()
	tx := db.Begin()
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

	err = tx.Delete(&category).Error
	return tx.Commit().Error
}

func GetAllTags() (tags Tags) {
	db := GetDBConnection()
	db.Find(&tags)
	return tags
}

func GetTagsByNames(names []string) (tags Tags) {
	db := GetDBConnection()
	db.Find(&tags, "name in (?)", names)
	return tags
}

func GetTag(uuid string) (tag Tag) {
	db := GetDBConnection()
	db.Limit(1).Find(&tag, "uuid = ?", uuid)
	return tag
}

func TagUniqueCheck(name string) (isUnique bool, err error) {
	var tag Tag
	db := GetDBConnection()
	db.Limit(1).Find(&tag, "name = ?", name)
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
	db := GetDBConnection()
	err = db.Create(&tag).Error
	return err
}

func InsertUnDuplicateTag(tag *Tag, tagName string) (err error) {
	isUnique, _ := TagUniqueCheck(tagName)
	if !isUnique {
		return nil
	}
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	tag.Uuid = uuid
	db := GetDBConnection()
	err = db.Create(&tag).Error
	return err
}

func UpdateTag(tag *Tag, uuid string) (err error) {
	db := GetDBConnection()
	err = db.Model(&tag).Where("uuid = ?", uuid).Updates(
		Tag{Name: tag.Name},
	).Error
	return err
}

func (tag *Tag) DeleteTag() (err error) {
	db := GetDBConnection()
	tx := db.Begin()
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

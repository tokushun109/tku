package models

import (
	"api/config"
)

type Creator struct {
	DefaultModel
	Name         string `json:"name" validate:"min=1,max=20"`
	Introduction string `json:"introduction" validate:"min=1,max=1000"`
	MimeType     string `json:"mimeType"`
	// ロゴ画像の保存場所のパス
	Logo string `json:"logo"`
	// フロントで画像を取得する時のapiパス
	ApiPath string `gorm:"-" json:"apiPath"`
}

func GetCreator() (creator Creator) {
	db := GetDBConnection()
	db.Limit(1).Find(&creator)
	return creator
}

func initialInsertCreator() (err error) {
	creator := &Creator{}
	creator.Name = config.Config.CreatorName
	db := GetDBConnection()
	err = db.Create(&creator).Error
	return err
}

func UpdateCreator(creator *Creator) (err error) {
	ID := GetCreator().ID
	db := GetDBConnection()
	err = db.Model(&creator).Where("id = ?", ID).Updates(
		Creator{Name: creator.Name, Introduction: creator.Introduction},
	).Error
	return err
}

func UpdateCreatorLogo(creator *Creator) (err error) {
	beforeCreator := GetCreator()
	db := GetDBConnection()
	// ロゴ画像を更新したら既存の画像を削除
	if err := removeFile(beforeCreator.Logo); err != nil {
		db.Rollback()
		return err
	}
	err = db.Model(&Creator{}).Where("id = ?", beforeCreator.ID).Updates(
		Creator{MimeType: creator.MimeType, Logo: creator.Logo},
	).Error
	return err
}

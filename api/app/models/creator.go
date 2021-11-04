package models

import (
	"api/config"
	"os"
	"strings"
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

// 商品に紐づく商品画像に画像取得用のapiをつける
func setCreatorLogoApiPath(creator *Creator) {
	base := config.Config.ApiBaseUrl
	creator.ApiPath = ""
	if creator.Logo != "" {
		fileName := strings.Split(creator.Logo, "/")[4]
		creator.ApiPath = base + "/creator/logo/" + fileName + "/blob"
	}
}

func GetCreator() (creator Creator) {
	Db.First(&creator)
	setCreatorLogoApiPath(&creator)
	return creator
}

func initialInsertCreator() (err error) {
	creator := &Creator{}
	creator.Name = config.Config.CreatorName
	err = Db.Create(&creator).Error
	return err
}

func UpdateCreator(creator *Creator) (err error) {
	ID := GetCreator().ID
	err = Db.Model(&creator).Where("id = ?", ID).Updates(
		Creator{Name: creator.Name, Introduction: creator.Introduction},
	).Error
	return err
}

func UpdateCreatorLogo(creator *Creator) (err error) {
	beforeCreator := GetCreator()
	// ロゴ画像を更新したら既存の画像を削除
	if _, err := os.Stat(beforeCreator.Logo); err == nil {
		if err := os.Remove(beforeCreator.Logo); err != nil {
			return err
		}
	}
	err = Db.Model(&Creator{}).Where("id = ?", beforeCreator.ID).Updates(
		Creator{MimeType: creator.MimeType, Logo: creator.Logo},
	).Error
	return err
}

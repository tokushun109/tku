package models

import "api/config"

type Creator struct {
	DefaultModel
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Logo         string `json:"logo"`
}

func GetCreator() (creator Creator, err error) {
	err = Db.First(&creator).Error
	return creator, err
}

func initialInsertCreator() (err error) {
	creator := &Creator{}
	creator.Name = config.Config.CreatorName
	err = Db.Create(&creator).Error
	return err
}

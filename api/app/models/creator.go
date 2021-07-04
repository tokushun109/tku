package models

import "api/config"

type Creator struct {
	DefaultModel
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Logo         string `json:"logo"`
}

func GetCreator() (creator Creator) {
	Db.First(&creator)
	return creator
}

func initialInsertCreator() (err error) {
	creator := &Creator{}
	creator.Name = config.Config.CreatorName
	err = Db.Create(&creator).Error
	return err
}

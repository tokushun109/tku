package models

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

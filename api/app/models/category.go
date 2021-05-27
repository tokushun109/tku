package models

type AccessoryCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type AccessoryCategories []AccessoryCategory

type MaterialCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type MaterialCategories []MaterialCategory

func GetAllAccessoryCategories() (accessoryCategories AccessoryCategories) {
	Db.Find(&accessoryCategories)
	return accessoryCategories
}

func GetAccessoryCategory(uuid string) (accessory_category AccessoryCategory) {
	Db.First(&accessory_category, "uuid = ?", uuid)
	return accessory_category
}

func GetAllMaterialCategories() (materialCategories MaterialCategories) {
	Db.Find(&materialCategories)
	return materialCategories
}

func InsertAccessoryCategory(accessory_category *AccessoryCategory) {
	Db.NewRecord(accessory_category)
	Db.Create(&accessory_category)
}

func InsertMaterialCategory(material_category *MaterialCategory) {
	Db.NewRecord(material_category)
	Db.Create(&material_category)
}

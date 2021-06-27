package models

type AccessoryCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type AccessoryCategories []AccessoryCategory

type MaterialCategory struct {
	DefaultModel
	Uuid string `json:"uuid"`
	Name string `json:"name" validate:"min=1,max=20"`
}

type MaterialCategories []MaterialCategory

func GetAllAccessoryCategories() (accessoryCategories AccessoryCategories, err error) {
	err = Db.Find(&accessoryCategories).Error
	return accessoryCategories, err
}

func GetAccessoryCategory(uuid string) (accessoryCategory AccessoryCategory, err error) {
	err = Db.First(&accessoryCategory, "uuid = ?", uuid).Error
	return accessoryCategory, err
}

func InsertAccessoryCategory(accessoryCategory *AccessoryCategory) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	accessoryCategory.Uuid = uuid
	err = Db.Create(&accessoryCategory).Error
	return err
}

func GetAllMaterialCategories() (materialCategories MaterialCategories, err error) {
	err = Db.Find(&materialCategories).Error
	return materialCategories, err
}

func GetMaterialCategory(uuid string) (materialCategory MaterialCategory, err error) {
	err = Db.First(&materialCategory, "uuid = ?", uuid).Error
	return materialCategory, err
}

func InsertMaterialCategory(materialCategory *MaterialCategory) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	materialCategory.Uuid = uuid
	err = Db.Create(&materialCategory).Error
	return err
}

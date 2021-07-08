package models

import (
	"api/config"
	"errors"
)

type Product struct {
	DefaultModel
	Uuid                string             `json:"uuid"`
	Name                string             `json:"name" validate:"min=1,max=20"`
	Description         string             `json:"description"`
	AccessoryCategoryId *uint              `json:"-"`
	AccessoryCategory   AccessoryCategory  `json:"accessoryCategory" validate:"-"`
	MaterialCategories  []MaterialCategory `gorm:"many2many:product_to_material_category" json:"materialCategories"`
	ProductImages       []*ProductImage    `gorm:"hasmany:product_image" json:"productImages"`
	SalesSites          []SalesSite        `gorm:"many2many:product_to_sales_site" json:"salesSites"`
}

type Products []Product

type ProductImage struct {
	DefaultModel
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	MimeType  string `json:"-"`
	ProductId *uint  `json:"-"`
	// 画像の保存場所のパス
	Path string `json:"-"`
	// フロントで画像を取得する時のapiパス
	ApiPath string `gorm:"-" json:"apiPath"`
}

// 商品に紐づく商品画像に画像取得用のapiをつける
func setProductImageApiPath(product *Product) {
	base := config.Config.ApiBaseUrl
	for _, productImage := range product.ProductImages {
		productImage.ApiPath = base + "/product_image/" + productImage.Uuid + "/blob"
	}
}

func GetAllProducts() (products Products) {
	Db.Preload("AccessoryCategory").
		Preload("ProductImages").
		Preload("MaterialCategories").
		Preload("SalesSites").
		Find(&products)
	for _, product := range products {
		setProductImageApiPath(&product)
	}
	return products
}

func GetProduct(uuid string) (product Product) {
	Db.First(&product, "uuid = ?", uuid).
		Related(&product.AccessoryCategory).
		Related(&product.ProductImages, "ProductImages").
		Related(&product.MaterialCategories, "MaterialCategories").
		Related(&product.SalesSites, "SalesSites")
	setProductImageApiPath(&product)
	return product
}

func ProductUniqueCheck(name string) (isUnique bool, err error) {
	var product Product
	Db.First(&product, "name = ?", name)
	isUnique = product.ID == nil
	if !isUnique {
		err = errors.New("name is duplicate")
	}
	return isUnique, err
}

func InsertProduct(product *Product) (err error) {
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	product.Uuid = uuid
	// アクセサリーカテゴリーの設定
	accesstoryCategory := GetAccessoryCategory(product.AccessoryCategory.Uuid)
	product.AccessoryCategoryId = accesstoryCategory.ID
	if err := tx.Omit("AccessoryCategory", "MaterialCategories", "SalesSites").Create(&product).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 商品と材料カテゴリーを紐付け
	for _, materialCategory := range product.MaterialCategories {
		// IDを取得する
		materialCategory := GetMaterialCategory(materialCategory.Uuid)
		materialCategoryId := materialCategory.ID
		var productToMaterialCategory = ProductToMaterialCategory{ProductId: product.ID, MaterialCategoryId: materialCategoryId}
		if err := tx.Create(&productToMaterialCategory).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 商品と販売サイトを紐付け
	for _, salesSite := range product.SalesSites {
		// IDを取得する
		saleSite := GetSalesSite(salesSite.Uuid)
		salesSiteId := saleSite.ID
		var productToSalesSite = ProductToSalesSite{ProductId: product.ID, SalesSiteId: salesSiteId}
		if err := tx.Create(&productToSalesSite).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (product *Product) DeleteProduct() (err error) {
	err = Db.Delete(&product).Error
	return err
}

func GetProductImage(uuid string) (productImage ProductImage) {
	Db.First(&productImage, "uuid = ?", uuid)
	return productImage
}

func InsertProductImage(productImage *ProductImage) (err error) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		return err
	}
	productImage.Uuid = uuid
	err = Db.Create(&productImage).Error
	return err
}

package models

import (
	"log"
)

type Product struct {
	DefaultModel
	Uuid                string             `json:"uuid"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	AccessoryCategoryId *uint              `json:"-"`
	AccessoryCategory   AccessoryCategory  `json:"accessoryCategory"`
	MaterialCategories  []MaterialCategory `gorm:"many2many:product_to_material_category" json:"materialCategories"`
	ProductImages       []ProductImage     `gorm:"hasmany:product_image" json:"productImages"`
	SalesSites          []SalesSite        `gorm:"many2many:product_to_sales_site" json:"salesSites"`
}

type Products []Product

type ProductImage struct {
	DefaultModel
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	MimeType  string `json:"-"`
	ProductId int    `json:"-"`
	Path      string `json:"path"`
}

func GetAllProducts() (products Products) {
	Db.Preload("AccessoryCategory").
		Preload("ProductImages").
		Preload("MaterialCategories").
		Preload("SalesSites").
		Find(&products)
	return products
}

func GetProduct(uuid string) (product Product) {
	Db.First(&product, "uuid = ?", uuid).
		Related(&product.AccessoryCategory).
		Related(&product.ProductImages, "ProductImages").
		Related(&product.MaterialCategories, "MaterialCategories").
		Related(&product.SalesSites, "SalesSites")
	return product
}

func InsertProduct(product *Product) {
	// uuidの設定
	uuid, err := GenerateUuid()
	if err != nil {
		log.Fatal(err)
	}
	product.Uuid = uuid
	// アクセサリーカテゴリーの設定
	accesstoryCategory := GetAccessoryCategory(product.AccessoryCategory.Uuid)
	product.AccessoryCategoryId = accesstoryCategory.ID
	// リクエストの材料カテゴリー情報を取得しておく
	materialCategories := product.MaterialCategories
	// 商品データの作成
	Db.NewRecord(product)
	Db.Omit("AccessoryCategory", "MaterialCategories").Create(&product)
	// 商品と材料カテゴリーを紐付け
	for _, materialCategory := range materialCategories {
		// IDを取得する
		materialCategoryId := GetMaterialCategory(materialCategory.Uuid).ID
		var productToMaterialCategory = ProductToMaterialCategory{ProductId: product.ID, MaterialCategoryId: materialCategoryId}
		Db.Create(&productToMaterialCategory)
	}
}

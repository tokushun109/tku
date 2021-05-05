package models

import "fmt"

type Product struct {
	DefaultModel
	Uuid                string            `json:"uuid"`
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	AccessoryCategoryId int               `json:"accessory_category_id"`
	AccessoryCategory   AccessoryCategory `json:"accessory_category"`
	ProductImageId      int               `json:"product_image_id"`
	ProductImage        ProductImage      `json:"product_image"`
}

type Products []Product

type ProductImage struct {
	DefaultModel
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	MimeType string `json:"-"`
	Path     string `json:"path"`
}

func DbTest() {
	fmt.Println("Db.HasTable(product)", Db.HasTable("product"))
}

func GetProducts() (products Products) {
	Db.Preload("AccessoryCategory").Preload("ProductImage").Find(&products)
	return products
}

package models

type Product struct {
	DefaultModel
	Uuid                string            `json:"uuid"`
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	AccessoryCategoryId int               `json:"-"`
	AccessoryCategory   AccessoryCategory `json:"accessory_category"`
	ProductImageId      int               `json:"-"`
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

func GetAllProducts() (products Products) {
	Db.Preload("AccessoryCategory").Preload("ProductImage").Find(&products)
	return products
}

func GetProduct(uuid string) (product Product) {
	Db.First(&product, "uuid = ?", uuid).Related(&product.AccessoryCategory).Related(&product.ProductImage)
	return product
}

package models

type Product struct {
	DefaultModel
	Uuid                string             `json:"uuid"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	AccessoryCategoryID int                `json:"-"`
	AccessoryCategory   AccessoryCategory  `json:"accessoryCategory"`
	MaterialCategories  []MaterialCategory `gorm:"many2many:product_to_material_category" json:"materialCategories"`
	ProductImageID      int                `json:"-"`
	ProductImage        ProductImage       `json:"productImage"`
	SalesSites          []SalesSite        `gorm:"many2many:product_to_sales_site" json:"salesSites"`
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
	Db.Preload("AccessoryCategory").
		Preload("ProductImage").
		Preload("MaterialCategories").
		Preload("SalesSites").
		Find(&products)
	return products
}

func GetProduct(uuid string) (product Product) {
	Db.First(&product, "uuid = ?", uuid).
		Related(&product.AccessoryCategory).
		Related(&product.ProductImage).
		Related(&product.MaterialCategories, "MaterialCategories").
		Related(&product.SalesSites, "SalesSites")
	return product
}

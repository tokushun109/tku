package main

import (
	"api/app/models"
	"fmt"
)

func main() {
	// go run app/db/seed.goで初期データを作成

	// accessory_category
	var accessoryCategory models.AccessoryCategory
	accessoryCategory.Name = "test_accessory_category"
	models.InsertAccessoryCategory(&accessoryCategory)
	fmt.Println("accessory_categoryを作成しました")

	// material_category
	var materialCategory models.MaterialCategory
	materialCategory.Name = "test_material_category"
	models.InsertMaterialCategory(&materialCategory)
	fmt.Println("material_categoryを作成しました")

	// sales_site
	var salesSite models.SalesSite
	salesSite.Name = "example.com"
	salesSite.Url = "https://example.com/"
	models.InsertSalesSite(&salesSite)
	fmt.Println("sales_siteを作成しました")

	// skill_market
	var skillMarket models.SkillMarket
	skillMarket.Name = "example.com"
	skillMarket.Url = "https://example.com/"
	models.InsertSkillMarket(&skillMarket)
	fmt.Println("skill_marketを作成しました")

	// sns
	var sns models.Sns
	sns.Name = "example.com"
	sns.Url = "https://example.com/"
	models.InsertSns(&sns)
	fmt.Println("snsを作成しました")

	// product
	var product models.Product
	product.Name = "test_product"
	product.Description = "test_descriptiontest_descriptiontest_descriptiontest_descriptiontest_description"
	product.AccessoryCategory = accessoryCategory
	product.MaterialCategories = []models.MaterialCategory{materialCategory}
	product.SalesSites = []models.SalesSite{salesSite}
	models.InsertProduct(&product)
	fmt.Println("productを作成しました")

	// product_image
	var productImage models.ProductImage
	productImage.MimeType = "image/jpeg"
	productImage.Path = "test_path"
	productImage.ProductId = product.ID
	models.InsertProductImage(&productImage)
	fmt.Println("product_imageを作成しました")

}

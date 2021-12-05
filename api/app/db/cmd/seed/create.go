package main

import (
	"api/app/models"
	"fmt"
	"log"
)

// go run app/db/cmd/seed/create.go
func main() {
	// category
	var category models.Category
	category.Name = "test_category"
	if err := models.InsertCategory(&category); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("categoryを作成しました")

	// tag
	var tag models.Tag
	tag.Name = "test_tag"
	if err := models.InsertTag(&tag); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("tagを作成しました")

	// sales_site
	var salesSite models.SalesSite
	salesSite.Name = "example.com"
	salesSite.Url = "https://example.com/"
	if err := models.InsertSalesSite(&salesSite); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("sales_siteを作成しました")

	// skill_market
	var skillMarket models.SkillMarket
	skillMarket.Name = "example.com"
	skillMarket.Url = "https://example.com/"
	if err := models.InsertSkillMarket(&skillMarket); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("skill_marketを作成しました")

	// sns
	var sns models.Sns
	sns.Name = "example.com"
	sns.Url = "https://example.com/"
	if err := models.InsertSns(&sns); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("snsを作成しました")

	// product
	var product models.Product
	product.Name = "test_product"
	product.Description = "test_descriptiontest_descriptiontest_descriptiontest_descriptiontest_description"
	product.Category = category
	product.Tags = []models.Tag{tag}
	product.SalesSites = []models.SalesSite{salesSite}
	if err := models.InsertProduct(&product); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("productを作成しました")

	// product_image
	var productImage models.ProductImage
	productImage.MimeType = "image/jpeg"
	productImage.Path = "test_path"
	productImage.ProductId = product.ID
	if err := models.InsertProductImage(&productImage); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("product_imageを作成しました")
}

package models

import (
	"api/config"
	"errors"
)

type Product struct {
	DefaultModel
	Uuid                string            `json:"uuid"`
	Name                string            `json:"name" validate:"min=1,max=20"`
	Description         string            `json:"description"`
	AccessoryCategoryId *uint             `json:"-"`
	AccessoryCategory   AccessoryCategory `json:"accessoryCategory" validate:"-"`
	Tags                []Tag             `gorm:"many2many:product_to_tag" json:"tags"`
	ProductImages       []*ProductImage   `gorm:"hasmany:product_image" json:"productImages"`
	SalesSites          []SalesSite       `gorm:"many2many:product_to_sales_site" json:"salesSites"`
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
		Preload("Tags").
		Preload("SalesSites").
		Find(&products)
	for _, product := range products {
		setProductImageApiPath(&product)
	}
	return products
}

func GetProduct(uuid string) (product Product, err error) {
	err = Db.Preload("AccessoryCategory").
		Preload("ProductImages").
		Preload("Tags").
		Preload("SalesSites").
		First(&product, "uuid = ?", uuid).Error

	setProductImageApiPath(&product)
	return product, err
}

func ProductUniqueCheck(name string) (isUnique bool, err error) {
	var product Product
	Db.Limit(1).Find(&product, "name = ?", name)
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
	// カテゴリーの設定
	accesstoryCategory := GetAccessoryCategory(product.AccessoryCategory.Uuid)
	product.AccessoryCategoryId = accesstoryCategory.ID
	if err := tx.Omit("AccessoryCategory", "Tags", "ProductImages", "SalesSites").Create(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	var productToTagList []ProductToTag
	// 商品とタグを紐付け
	for _, tag := range product.Tags {
		productToTagList = append(
			productToTagList,
			ProductToTag{
				ProductId: product.ID,
				TagId:     GetTag(tag.Uuid).ID,
			},
		)
	}
	if len(productToTagList) > 0 {
		if err := tx.Create(&productToTagList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	var productToSalesSites []ProductToSalesSite
	// 商品と販売サイトを紐付け
	for _, salesSite := range product.SalesSites {
		productToSalesSites = append(
			productToSalesSites,
			ProductToSalesSite{
				ProductId:   product.ID,
				SalesSiteId: GetSalesSite(salesSite.Uuid).ID,
			},
		)
	}
	if len(productToSalesSites) > 0 {
		if err := tx.Create(&productToSalesSites).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func UpdateProduct(product *Product, uuid string) (err error) {
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err = tx.Model(&product).
		Omit("AccessoryCategory", "Tags", "ProductImages", "SalesSites").
		Where("uuid = ?", uuid).
		Updates(
			Product{
				Name:                product.Name,
				Description:         product.Description,
				AccessoryCategoryId: GetAccessoryCategory(product.AccessoryCategory.Uuid).ID,
			},
		).
		Error
	if err != nil {
		tx.Rollback()
		return err
	}

	registeredProduct, err := GetProduct(uuid)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 登録されている中間テーブルを全て物理削除する
	if err = tx.Where("product_id = ?", registeredProduct.ID).
		Unscoped().
		Delete(&ProductToTag{}).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	// 商品とタグを紐付け
	var productToTagList []ProductToTag
	for _, tag := range product.Tags {
		productToTagList = append(
			productToTagList,
			ProductToTag{
				ProductId: registeredProduct.ID,
				TagId:     GetTag(tag.Uuid).ID,
			},
		)
	}

	if len(productToTagList) > 0 {
		if err := tx.Create(&productToTagList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 登録されている中間テーブルを全て物理削除する
	if err = tx.Where("product_id = ?", registeredProduct.ID).
		Unscoped().
		Delete(&ProductToSalesSite{}).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	// 商品と販売サイトを紐付け
	var productToSalesSites []ProductToSalesSite
	for _, salesSite := range product.SalesSites {
		productToSalesSites = append(
			productToSalesSites,
			ProductToSalesSite{
				ProductId:   registeredProduct.ID,
				SalesSiteId: GetSalesSite(salesSite.Uuid).ID,
			},
		)
	}

	if len(productToSalesSites) > 0 {
		if err := tx.Create(&productToSalesSites).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 残しておく商品画像のuuidリストを作成
	var leaveProductImageUuids []string
	for _, productImage := range product.ProductImages {
		leaveProductImageUuids = append(
			leaveProductImageUuids, productImage.Uuid,
		)
	}

	// 残しておく商品画像の有無で処理を分岐
	if len(product.ProductImages) == 0 {
		// 全ての商品画像のファイルを削除する
		for _, productImage := range product.ProductImages {
			if err := removeFile(productImage.Path); err != nil {
				tx.Rollback()
				return err
			}
		}

		// 全ての商品画像のレコードを削除する
		if err = tx.Where("product_id = ?", registeredProduct.ID).
			Delete(&ProductImage{}).
			Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 削除する商品画像のリストを取得
		var deleteProductImages []ProductImage
		err = tx.Where("product_id = ?", registeredProduct.ID).
			Where("uuid not in ?", leaveProductImageUuids).
			Find(&deleteProductImages).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// 商品画像のファイルを削除する
		for _, productImage := range deleteProductImages {
			if err := removeFile(productImage.Path); err != nil {
				tx.Rollback()
				return err
			}
		}

		// 商品画像のレコードを削除する
		if err = tx.Where("product_id = ?", registeredProduct.ID).
			Where("uuid not in ?", leaveProductImageUuids).
			Delete(&ProductImage{}).
			Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (product *Product) DeleteProduct() (err error) {
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 商品を削除する
	if err = tx.Select("Tags", "SalesSites", "ProductImages").Delete(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, productImage := range product.ProductImages {
		if err := removeFile(productImage.Path); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetProductImage(uuid string) (productImage ProductImage, err error) {
	err = Db.First(&productImage, "uuid = ?", uuid).Error
	return productImage, err
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

func (productImage *ProductImage) DeleteProductImage() (err error) {
	tx := Db.Begin()
	if err = tx.Delete(&productImage).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := removeFile(productImage.Path); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

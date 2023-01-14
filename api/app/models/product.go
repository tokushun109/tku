package models

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	DefaultModel
	Uuid          string          `json:"uuid"`
	Name          string          `json:"name" validate:"min=1,max=20"`
	Description   string          `json:"description"`
	Price         int             `json:"price" validate:"min=1,max=1000000"`
	IsActive      *bool           `json:"isActive"`
	CategoryId    *uint           `json:"-"`
	Category      Category        `json:"category" validate:"-"`
	Tags          []Tag           `gorm:"many2many:product_to_tag" json:"tags"`
	ProductImages []*ProductImage `gorm:"hasmany:product_image" json:"productImages"`
	SiteDetails   []*SiteDetail   `gorm:"hasmany:site_detail" json:"siteDetails"`
}

type Products []Product

type ProductImage struct {
	DefaultModel
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	MimeType  string `json:"-"`
	ProductId *uint  `json:"-"`
	// 画像の保存場所のパス
	Path  string `json:"-"`
	Order int    `json:"order"`
	// フロントで画像を取得する時のapiパス
	ApiPath string `gorm:"-" json:"apiPath"`
}

func GetAllProducts(mode string) (products Products, err error) {
	if mode == "all" {
		Db.Preload("Category").
			Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
				return db.Order("product_image.order Desc, id")
			}).
			Preload("Tags").
			Preload("SiteDetails", func(db *gorm.DB) *gorm.DB {
				return db.Order("site_detail.detail_url Desc")
			}).
			Preload("SiteDetails.SalesSite").
			Find(&products)
	} else if mode == "active" {
		Db.Where("is_active = ?", 1).
			Preload("Category").
			Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
				return db.Order("product_image.order Desc")
			}).
			Preload("Tags").
			Preload("SiteDetails", func(db *gorm.DB) *gorm.DB {
				return db.Order("site_detail.detail_url Desc, id")
			}).
			Preload("SiteDetails.SalesSite").
			Find(&products)
	} else {
		err = errors.New("invalid params")
		return products, err
	}
	return products, err
}

func GetNewProducts(limit int) (products Products) {
	Db.Joins("join product_image on product_image.product_id = product.id").
		Where("product.is_active = ?", 1).
		Where("product_image.deleted_at is null").
		Preload("Category").
		Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
			return db.Order("product_image.order Desc")
		}).
		Order("id desc").
		Group("product.id").
		Limit(limit).
		Find(&products)
	return products
}

func GetProduct(uuid string) (product Product, err error) {
	err = Db.Preload("Category").
		Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
			return db.Order("product_image.order Desc")
		}).
		Preload("Tags").
		Preload("SiteDetails", func(db *gorm.DB) *gorm.DB {
			return db.Order("site_detail.detail_url Desc, id")
		}).
		Preload("SiteDetails.SalesSite").
		First(&product, "uuid = ?", uuid).Error
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
	category := GetCategory(product.Category.Uuid)
	product.CategoryId = category.ID
	if err := tx.Omit("Category", "Tags", "ProductImages", "SiteDetails").Create(&product).Error; err != nil {
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

	var siteDetailList []SiteDetail
	// 商品と販売サイトを紐付け
	for _, siteDetail := range product.SiteDetails {
		// uuidの設定
		uuid, err := GenerateUuid()
		if err != nil {
			return err
		}
		siteDetailList = append(
			siteDetailList,
			SiteDetail{
				Uuid:        uuid,
				DetailUrl:   siteDetail.DetailUrl,
				ProductId:   product.ID,
				SalesSiteId: GetSalesSite(siteDetail.SalesSite.Uuid).ID,
			},
		)
	}
	if len(siteDetailList) > 0 {
		if err := tx.Create(&siteDetailList).Error; err != nil {
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
		Omit("Category", "Tags", "ProductImages", "SiteDetails").
		Where("uuid = ?", uuid).
		Updates(
			Product{
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				IsActive:    product.IsActive,
				CategoryId:  GetCategory(product.Category.Uuid).ID,
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
		Delete(&SiteDetail{}).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	var siteDetailList []SiteDetail
	// 商品と販売サイトを紐付け
	for _, siteDetail := range product.SiteDetails {
		// uuidの設定
		uuid, err := GenerateUuid()
		if err != nil {
			return err
		}
		siteDetailList = append(
			siteDetailList,
			SiteDetail{
				Uuid:        uuid,
				DetailUrl:   siteDetail.DetailUrl,
				ProductId:   registeredProduct.ID,
				SalesSiteId: GetSalesSite(siteDetail.SalesSite.Uuid).ID,
			},
		)
	}

	if len(siteDetailList) > 0 {
		if err := tx.Create(&siteDetailList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	var leaveProductImageUuids []string
	for _, productImage := range product.ProductImages {

		// 商品画像の優先順位を更新
		err = tx.Model(&productImage).Where("uuid = ?", productImage.Uuid).Updates(
			ProductImage{Order: productImage.Order},
		).Error
		if err != nil {
			return err
		}

		// 残しておく商品画像のuuidリストを作成
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
	if err = tx.Select("Tags", "SiteDetails", "ProductImages").Delete(&product).Error; err != nil {
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

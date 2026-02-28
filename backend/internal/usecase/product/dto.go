package product

import "io"

type CreateProductInput struct {
	Name         string
	Description  string
	Price        int
	IsActive     bool
	IsRecommend  bool
	CategoryUUID string
	TargetUUID   string
	TagUUIDs     []string
	SiteDetails  []SiteDetailInput
}

type UpdateProductInput struct {
	Name          string
	Description   string
	Price         int
	IsActive      bool
	IsRecommend   bool
	CategoryUUID  string
	TargetUUID    string
	TagUUIDs      []string
	SiteDetails   []SiteDetailInput
	ProductImages []ProductImageUpdateInput
}

type SiteDetailInput struct {
	SalesSiteUUID string
	DetailURL     string
}

type ProductImageUpdateInput struct {
	UUID         string
	DisplayOrder int
}

type ProductImageUploadFile struct {
	Name string
	Data []byte
}

type ProductImageBlob struct {
	ContentType string
	Body        io.ReadCloser
}

type ProductCSVInputRow struct {
	ID           uint
	Name         string
	Price        int
	CategoryName string
	TargetName   string
}

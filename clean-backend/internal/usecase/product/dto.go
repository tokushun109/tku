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
	UUID  string
	Order int
}

type ProductImageUploadFile struct {
	Name string
	Data []byte
}

type ProductImageBlob struct {
	ContentType string
	Body        io.ReadCloser
}

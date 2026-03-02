package response

type ProductClassificationResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type ProductSalesSiteResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type ProductSiteDetailResponse struct {
	UUID      string                   `json:"uuid"`
	DetailURL string                   `json:"detailUrl"`
	SalesSite ProductSalesSiteResponse `json:"salesSite"`
}

type ProductImageResponse struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
	APIPath      string `json:"apiPath"`
}

type ProductResponse struct {
	UUID          string                          `json:"uuid"`
	Name          string                          `json:"name"`
	Description   string                          `json:"description"`
	Price         int                             `json:"price"`
	IsRecommend   bool                            `json:"isRecommend"`
	IsActive      bool                            `json:"isActive"`
	Category      ProductClassificationResponse   `json:"category"`
	Target        ProductClassificationResponse   `json:"target"`
	Tags          []ProductClassificationResponse `json:"tags"`
	ProductImages []ProductImageResponse          `json:"productImages"`
	SiteDetails   []ProductSiteDetailResponse     `json:"siteDetails"`
}

type CreateProductResponse struct {
	UUID string `json:"uuid"`
}

type CategoryProductsResponse struct {
	Category ProductClassificationResponse `json:"category"`
	PageInfo CursorPageInfoResponse        `json:"pageInfo"`
	Products []*ProductResponse            `json:"products"`
}

type CarouselItemResponse struct {
	Product *ProductResponse `json:"product"`
	APIPath string           `json:"apiPath"`
}

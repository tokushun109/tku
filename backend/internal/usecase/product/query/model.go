package query

type Classification struct {
	UUID string
	Name string
}

type SalesSite struct {
	UUID string
	Name string
}

type SiteDetail struct {
	UUID      string
	DetailURL string
	SalesSite SalesSite
}

type ProductImage struct {
	UUID         string
	Name         string
	MimeType     string
	Path         string
	DisplayOrder int
	APIPath      string
}

type Product struct {
	ID            uint
	UUID          string
	Name          string
	Description   string
	Price         *int
	IsActive      bool
	IsRecommend   bool
	Category      Classification
	Target        Classification
	Tags          []Classification
	ProductImages []ProductImage
	SiteDetails   []SiteDetail
}

type PageInfo struct {
	HasMore    bool
	NextCursor string
}

type OffsetPageInfo struct {
	Page       int
	Limit      int
	Total      int
	TotalPages int
}

type ProductPage struct {
	PageInfo OffsetPageInfo
	Products []*Product
}

type CategoryProducts struct {
	Category Classification
	PageInfo PageInfo
	Products []*Product
}

type CarouselItem struct {
	Product *Product
	APIPath string
}

type ProductCSVRow struct {
	UUID         string
	Name         string
	Price        *int
	CategoryName string
	TargetName   string
}

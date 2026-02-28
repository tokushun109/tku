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
	UUID     string
	Name     string
	MimeType string
	Path     string
	Order    int
	APIPath  string
}

type Product struct {
	ID            uint
	UUID          string
	Name          string
	Description   string
	Price         int
	IsActive      bool
	IsRecommend   bool
	Category      Classification
	Target        Classification
	Tags          []Classification
	ProductImages []ProductImage
	SiteDetails   []SiteDetail
}

type CategoryProducts struct {
	Category Classification
	Products []*Product
}

type CarouselItem struct {
	Product *Product
	APIPath string
}

type ProductCSVRow struct {
	ID           uint
	Name         string
	Price        int
	CategoryName string
	TargetName   string
}

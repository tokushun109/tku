package query

import "context"

type ListProductsQuery struct {
	Mode     string
	Category string
	Target   string
}

type ListCategoryProductsQuery struct {
	Category string
	Target   string
}

type ListCarouselQuery struct {
	Limit int
}

type ExportProductsCSVQuery struct{}

type Reader interface {
	ListProducts(ctx context.Context, q ListProductsQuery) ([]*Product, error)
	ListCategoryProducts(ctx context.Context, q ListCategoryProductsQuery) ([]*CategoryProducts, error)
	ListCarouselItems(ctx context.Context, q ListCarouselQuery) ([]*CarouselItem, error)
	GetProductByUUID(ctx context.Context, productUUID string) (*Product, error)
	ExportProductsCSV(ctx context.Context, q ExportProductsCSVQuery) ([]*ProductCSVRow, error)
}

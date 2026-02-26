package query

import "context"

type ListProductsQuery struct {
	Mode     string
	Category string
	Target   string
}

type ListCategoryProductsQuery struct {
	Mode     string
	Category string
	Target   string
}

type Reader interface {
	ListProducts(ctx context.Context, q ListProductsQuery) ([]*Product, error)
	ListCategoryProducts(ctx context.Context, q ListCategoryProductsQuery) ([]*CategoryProducts, error)
	GetProductByUUID(ctx context.Context, productUUID string) (*Product, error)
}

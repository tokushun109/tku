package query

import "context"

type ListProductsQuery struct {
	Mode     string
	Category string
	Target   string
}

type Reader interface {
	ListProducts(ctx context.Context, q ListProductsQuery) ([]*Product, error)
	GetProductByUUID(ctx context.Context, productUUID string) (*Product, error)
}

package product

import "context"

type DuplicateSource interface {
	Duplicate(ctx context.Context, rawURL string) (*DuplicateProductData, error)
}

type DuplicateProductData struct {
	Name        string
	Description string
	Price       int
	Tags        []string
	Images      []DuplicateProductImage
}

type DuplicateProductImage struct {
	Name string
	Data []byte
}

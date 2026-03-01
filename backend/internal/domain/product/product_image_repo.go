package product

import (
	"context"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type ProductImageRepository interface {
	Create(ctx context.Context, image *ProductImage) (*ProductImage, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*ProductImage, error)
	FindByProductUUID(ctx context.Context, productUUID primitive.UUID) ([]*ProductImage, error)
	UpdateDisplayOrder(ctx context.Context, uuid primitive.UUID, displayOrder int) (bool, error)
	DeleteByUUID(ctx context.Context, uuid primitive.UUID) (bool, error)
	DeleteByProductUUID(ctx context.Context, productUUID primitive.UUID) error
}

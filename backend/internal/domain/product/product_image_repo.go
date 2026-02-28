package product

import (
	"context"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type ProductImageRepository interface {
	Create(ctx context.Context, image *ProductImage) (*ProductImage, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*ProductImage, error)
	FindByProductID(ctx context.Context, productID primitive.ID) ([]*ProductImage, error)
	UpdateDisplayOrder(ctx context.Context, uuid primitive.UUID, displayOrder int) (bool, error)
	DeleteByUUID(ctx context.Context, uuid primitive.UUID) (bool, error)
	DeleteByProductID(ctx context.Context, productID primitive.ID) error
}

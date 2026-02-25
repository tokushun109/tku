package product

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, p *Product) (primitive.ID, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Product, error)
	Update(ctx context.Context, p *Product) (bool, error)
	Delete(ctx context.Context, uuid primitive.UUID) (bool, error)
	ReplaceTags(ctx context.Context, productID primitive.ID, tagIDs []primitive.ID) error
}

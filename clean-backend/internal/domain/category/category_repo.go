package category

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, c *Category) error
	FindAll(ctx context.Context) ([]*Category, error)
	FindUsed(ctx context.Context) ([]*Category, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Category, error)
	ExistsByName(ctx context.Context, name CategoryName) (bool, error)
	Update(ctx context.Context, c *Category) (bool, error)
}

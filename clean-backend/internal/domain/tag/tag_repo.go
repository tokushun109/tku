package tag

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, t *Tag) (*Tag, error)
	FindAll(ctx context.Context) ([]*Tag, error)
	FindByName(ctx context.Context, name TagName) (*Tag, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Tag, error)
	ExistsByName(ctx context.Context, name TagName) (bool, error)
	Update(ctx context.Context, t *Tag) (bool, error)
	Delete(ctx context.Context, uuid primitive.UUID) (bool, error)
}

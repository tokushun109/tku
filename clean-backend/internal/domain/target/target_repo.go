package target

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, t *Target) error
	FindAll(ctx context.Context) ([]*Target, error)
	FindUsed(ctx context.Context) ([]*Target, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Target, error)
	ExistsByName(ctx context.Context, name TargetName) (bool, error)
	Update(ctx context.Context, t *Target) (bool, error)
	Delete(ctx context.Context, uuid primitive.UUID) (bool, error)
}

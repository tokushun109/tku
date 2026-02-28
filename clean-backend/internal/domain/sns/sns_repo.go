package sns

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, s *Sns) (*Sns, error)
	FindAll(ctx context.Context) ([]*Sns, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Sns, error)
	Update(ctx context.Context, s *Sns) (bool, error)
	Delete(ctx context.Context, uuid primitive.UUID) (bool, error)
}

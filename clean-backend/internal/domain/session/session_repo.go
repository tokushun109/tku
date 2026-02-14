package session

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Session, error)
}

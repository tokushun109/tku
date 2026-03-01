package session

import (
	"context"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, s *Session) (*Session, error)
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Session, error)
	DeleteByUUID(ctx context.Context, uuid primitive.UUID) error
	DeleteByUserUUID(ctx context.Context, userUUID primitive.UUID) error
}

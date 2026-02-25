package session

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	Create(ctx context.Context, s *Session) error
	FindByUUID(ctx context.Context, uuid primitive.UUID) (*Session, error)
	DeleteByUUID(ctx context.Context, uuid primitive.UUID) error
	DeleteByUserID(ctx context.Context, userID primitive.ID) error
}

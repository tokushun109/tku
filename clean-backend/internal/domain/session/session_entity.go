package session

import (
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Session struct {
	UUID      primitive.UUID
	UserID    uint
	CreatedAt time.Time
}

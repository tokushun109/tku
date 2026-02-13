package session

import (
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/user"
)

type Session struct {
	UUID      SessionUUID
	UserID    user.UserID
	CreatedAt time.Time
}

package session

import "time"

type Session struct {
	UUID      SessionUUID
	UserID    uint
	CreatedAt time.Time
}

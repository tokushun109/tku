package session

import "time"

type Session struct {
	UUID      string
	UserID    uint
	CreatedAt time.Time
}

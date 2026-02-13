package session

import "github.com/tokushun109/tku/clean-backend/internal/shared/id"

type SessionUUID id.UUID

func NewSessionUUID() SessionUUID {
	return id.NewAs[SessionUUID]()
}

func ParseSessionUUID(s string) (SessionUUID, error) {
	return id.ParseAs[SessionUUID](s)
}

func (u SessionUUID) String() string {
	return string(u)
}

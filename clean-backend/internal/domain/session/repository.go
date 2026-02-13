package session

import "context"

type Repository interface {
	FindByUUID(ctx context.Context, uuid SessionUUID) (*Session, error)
}

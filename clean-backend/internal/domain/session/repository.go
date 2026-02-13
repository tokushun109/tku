package session

import "context"

type Repository interface {
	FindByUUID(ctx context.Context, uuid string) (*Session, error)
}

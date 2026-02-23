package contact

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]*Contact, error)
	Create(ctx context.Context, contact *Contact) error
}

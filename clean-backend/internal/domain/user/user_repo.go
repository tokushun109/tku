package user

import (
	"context"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Repository interface {
	FindByEmail(ctx context.Context, email primitive.Email) (*User, error)
	FindByID(ctx context.Context, id primitive.ID) (*User, error)
	FindContactNotificationUsers(ctx context.Context) ([]*ContactNotificationUser, error)
}

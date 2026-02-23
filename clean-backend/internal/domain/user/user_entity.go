package user

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type User struct {
	ID           uint
	UUID         primitive.UUID
	Name         UserName
	Email        primitive.Email
	PasswordHash UserPasswordHash
	IsAdmin      bool
}

package user

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type User struct {
	ID           uint
	UUID         primitive.UUID
	Name         string
	Email        string
	PasswordHash UserPasswordHash
	IsAdmin      bool
}

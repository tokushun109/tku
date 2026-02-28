package user

import domain "github.com/tokushun109/tku/backend/internal/domain/user"

type PasswordHasher interface {
	Hash(plain string) (domain.UserPasswordHash, error)
	Verify(plain string, hash domain.UserPasswordHash) (bool, error)
}

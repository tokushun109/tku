package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
)

type PasswordHasherBcrypt struct{}

func NewPasswordHasherBcrypt() *PasswordHasherBcrypt {
	return &PasswordHasherBcrypt{}
}

func (h *PasswordHasherBcrypt) Hash(plain string) (domainUser.UserPasswordHash, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return domainUser.NewUserPasswordHash(string(hashed))
}

func (h *PasswordHasherBcrypt) Verify(plain string, hash domainUser.UserPasswordHash) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash.String()), []byte(plain))
	if err == nil {
		return true, nil
	}

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) || errors.Is(err, bcrypt.ErrHashTooShort) {
		return false, nil
	}

	return false, err
}

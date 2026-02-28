package crypto

import (
	"crypto/sha1"
	"fmt"

	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
)

type PasswordHasherSHA1 struct{}

func NewPasswordHasherSHA1() *PasswordHasherSHA1 {
	return &PasswordHasherSHA1{}
}

func (h *PasswordHasherSHA1) Hash(plain string) (domainUser.UserPasswordHash, error) {
	// TODO: 将来の移行で、SHA-1 ベースのパスワードハッシュを bcrypt に置き換える。
	hashed := fmt.Sprintf("%x", sha1.Sum([]byte(plain)))
	return domainUser.NewUserPasswordHash(hashed)
}

func (h *PasswordHasherSHA1) Verify(plain string, hash domainUser.UserPasswordHash) (bool, error) {
	hashed, err := h.Hash(plain)
	if err != nil {
		return false, err
	}
	return hashed == hash, nil
}

package crypto

import (
	"testing"

	domainUser "github.com/tokushun109/tku/backend/internal/domain/user"
)

func TestPasswordHasherBcrypt(t *testing.T) {
	t.Run("Hash した値を Verify できる", func(t *testing.T) {
		hasher := NewPasswordHasherBcrypt()

		hash, err := hasher.Hash("password")
		if err != nil {
			t.Fatalf("unexpected hash error: %v", err)
		}

		ok, err := hasher.Verify("password", hash)
		if err != nil {
			t.Fatalf("unexpected verify error: %v", err)
		}
		if !ok {
			t.Fatalf("expected verify success")
		}
	})

	t.Run("不一致のパスワードは false を返す", func(t *testing.T) {
		hasher := NewPasswordHasherBcrypt()

		hash, err := hasher.Hash("password")
		if err != nil {
			t.Fatalf("unexpected hash error: %v", err)
		}

		ok, err := hasher.Verify("bad-password", hash)
		if err != nil {
			t.Fatalf("unexpected verify error: %v", err)
		}
		if ok {
			t.Fatalf("expected verify failure")
		}
	})

	t.Run("旧 SHA-1 形式の値は false を返す", func(t *testing.T) {
		hasher := NewPasswordHasherBcrypt()

		legacyHash, err := domainUser.NewUserPasswordHash("0123456789012345678901234567890123456789")
		if err != nil {
			t.Fatalf("unexpected legacy hash error: %v", err)
		}

		ok, err := hasher.Verify("password", legacyHash)
		if err != nil {
			t.Fatalf("unexpected verify error: %v", err)
		}
		if ok {
			t.Fatalf("expected verify failure")
		}
	})
}

package health

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type stubRepo struct {
	pingErr error
}

func (s *stubRepo) Ping(ctx context.Context) error {
	return s.pingErr
}

func TestCheck(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		uc := New(&stubRepo{})
		if err := uc.Check(context.Background()); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})
	t.Run("依存処理が失敗したならエラーを返す", func(t *testing.T) {

		uc := New(&stubRepo{pingErr: errors.New("db down")})
		err := uc.Check(context.Background())
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

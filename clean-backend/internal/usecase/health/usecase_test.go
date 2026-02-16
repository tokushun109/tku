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

func TestCheck_OK(t *testing.T) {
	uc := New(&stubRepo{})
	if err := uc.Check(context.Background()); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestCheck_Error(t *testing.T) {
	uc := New(&stubRepo{pingErr: errors.New("db down")})
	err := uc.Check(context.Background())
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

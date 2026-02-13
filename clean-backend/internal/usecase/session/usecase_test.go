package session

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	"github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.NewUUID().String()

type stubRepo struct {
	sess *domain.Session
	err  error
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid domain.SessionUUID) (*domain.Session, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.sess, nil
}

func TestValidate_EmptyToken_Unauthorized(t *testing.T) {
	uc := New(&stubRepo{})
	if err := uc.Validate(context.Background(), ""); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestValidate_NotFound_Unauthorized(t *testing.T) {
	uc := New(&stubRepo{sess: nil})
	if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestValidate_RepoError_Internal(t *testing.T) {
	uc := New(&stubRepo{err: errors.New("db error")})
	if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestValidate_OK(t *testing.T) {
	u, err := domain.ParseSessionUUID(testUUID)
	if err != nil {
		t.Fatalf("unexpected uuid parse error: %v", err)
	}
	sess := &domain.Session{UUID: u, UserID: user.NewUserID(1), CreatedAt: time.Now()}
	uc := New(&stubRepo{sess: sess})
	if err := uc.Validate(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

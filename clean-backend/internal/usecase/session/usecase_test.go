package session

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/session"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()

type stubRepo struct {
	sess *domain.Session
	err  error
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Session, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.sess, nil
}

type stubClock struct {
	now time.Time
}

func (c *stubClock) Now() time.Time {
	return c.now
}

func TestValidate_EmptyToken_Unauthorized(t *testing.T) {
	uc := New(&stubRepo{}, 24*time.Hour, &stubClock{now: time.Now()})
	if err := uc.Validate(context.Background(), ""); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestValidate_NotFound_Unauthorized(t *testing.T) {
	uc := New(&stubRepo{sess: nil}, 24*time.Hour, &stubClock{now: time.Now()})
	if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestValidate_RepoError_Internal(t *testing.T) {
	uc := New(&stubRepo{err: errors.New("db error")}, 24*time.Hour, &stubClock{now: time.Now()})
	if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestValidate_OK(t *testing.T) {
	u, err := primitive.NewUUID(testUUID)
	if err != nil {
		t.Fatalf("unexpected uuid parse error: %v", err)
	}
	sess := &domain.Session{UUID: u, UserID: 1, CreatedAt: time.Now()}
	uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: time.Now()})
	if err := uc.Validate(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_Expired_Unauthorized(t *testing.T) {
	u, err := primitive.NewUUID(testUUID)
	if err != nil {
		t.Fatalf("unexpected uuid parse error: %v", err)
	}
	now := time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC)
	sess := &domain.Session{UUID: u, UserID: 1, CreatedAt: now.Add(-25 * time.Hour)}
	uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: now})
	if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestValidate_NotExpired_OK(t *testing.T) {
	u, err := primitive.NewUUID(testUUID)
	if err != nil {
		t.Fatalf("unexpected uuid parse error: %v", err)
	}
	now := time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC)
	sess := &domain.Session{UUID: u, UserID: 1, CreatedAt: now.Add(-23 * time.Hour)}
	uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: now})
	if err := uc.Validate(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

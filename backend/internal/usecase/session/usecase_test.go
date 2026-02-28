package session

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/session"
	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

var testUUID = id.GenerateUUID()

type stubRepo struct {
	sess      *domain.Session
	err       error
	createErr error
	deleteErr error
	userID    primitive.ID
	created   *domain.Session
}

func (s *stubRepo) Create(ctx context.Context, sess *domain.Session) (*domain.Session, error) {
	if s.createErr != nil {
		return nil, s.createErr
	}
	s.created = sess
	return sess, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Session, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.sess, nil
}

func (s *stubRepo) DeleteByUUID(ctx context.Context, uuid primitive.UUID) error {
	return s.deleteErr
}

func (s *stubRepo) DeleteByUserID(ctx context.Context, userID primitive.ID) error {
	s.userID = userID
	return s.deleteErr
}

type stubClock struct {
	now time.Time
}

func (c *stubClock) Now() time.Time {
	return c.now
}

func TestValidate(t *testing.T) {
	t.Run("トークンが空なら未認証エラーを返す", func(t *testing.T) {

		uc := New(&stubRepo{}, 24*time.Hour, &stubClock{now: time.Now()})
		if err := uc.Validate(context.Background(), ""); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
			t.Fatalf("expected ErrUnauthorized, got %v", err)
		}
	})
	t.Run("セッションが見つからないなら未認証エラーを返す", func(t *testing.T) {

		uc := New(&stubRepo{sess: nil}, 24*time.Hour, &stubClock{now: time.Now()})
		if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
			t.Fatalf("expected ErrUnauthorized, got %v", err)
		}
	})
	t.Run("リポジトリエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		uc := New(&stubRepo{err: errors.New("db error")}, 24*time.Hour, &stubClock{now: time.Now()})
		if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		u, err := primitive.NewUUID(testUUID)
		if err != nil {
			t.Fatalf("unexpected uuid parse error: %v", err)
		}
		sess := mustSession(u.Value(), 1, time.Now())
		uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: time.Now()})
		if err := uc.Validate(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("セッションが期限切れなら未認証エラーを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(testUUID)
		if err != nil {
			t.Fatalf("unexpected uuid parse error: %v", err)
		}
		now := time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC)
		sess := mustSession(u.Value(), 1, now.Add(-25*time.Hour))
		uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: now})
		if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrUnauthorized) {
			t.Fatalf("expected ErrUnauthorized, got %v", err)
		}
	})
	t.Run("セッションが期限内なら認証済みユーザーを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(testUUID)
		if err != nil {
			t.Fatalf("unexpected uuid parse error: %v", err)
		}
		now := time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC)
		sess := mustSession(u.Value(), 1, now.Add(-23*time.Hour))
		uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: now})
		if err := uc.Validate(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("期限切れセッションの削除に失敗したなら内部エラーを返す", func(t *testing.T) {

		u, err := primitive.NewUUID(testUUID)
		if err != nil {
			t.Fatalf("unexpected uuid parse error: %v", err)
		}
		now := time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC)
		sess := mustSession(u.Value(), 1, now.Add(-25*time.Hour))
		uc := New(&stubRepo{sess: sess, deleteErr: errors.New("delete error")}, 24*time.Hour, &stubClock{now: now})
		if err := uc.Validate(context.Background(), testUUID); err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestResolve(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		u, err := primitive.NewUUID(testUUID)
		if err != nil {
			t.Fatalf("unexpected uuid parse error: %v", err)
		}
		sess := mustSession(u.Value(), 1, time.Now())
		uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: time.Now()})
		got, err := uc.Resolve(context.Background(), testUUID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil || got.UserID() != 1 {
			t.Fatalf("unexpected session: %+v", got)
		}
	})

}

func TestDelete(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		u, err := primitive.NewUUID(testUUID)
		if err != nil {
			t.Fatalf("unexpected uuid parse error: %v", err)
		}
		sess := mustSession(u.Value(), 1, time.Now())
		uc := New(&stubRepo{sess: sess}, 24*time.Hour, &stubClock{now: time.Now()})
		if err := uc.Delete(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

}

func mustSession(uuidStr string, userID uint, createdAt time.Time) *domain.Session {
	sess, err := domain.Rebuild(1, uuidStr, userID, createdAt)
	if err != nil {
		panic(err)
	}
	return sess
}

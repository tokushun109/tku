package contact

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "github.com/tokushun109/tku/backend/internal/domain/contact"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type stubRepo struct {
	findAll    []*domain.Contact
	findAllErr error
	createErr  error
	created    *domain.Contact
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.Contact, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) Create(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {
	if s.createErr != nil {
		return nil, s.createErr
	}
	s.created = contact
	return contact, nil
}

type stubNotifier struct {
	called  int
	contact *domain.Contact
}

func (s *stubNotifier) NotifyContactCreated(ctx context.Context, contact *domain.Contact) {
	s.called++
	s.contact = contact
}

func TestListContact(t *testing.T) {
	t.Run("有効な入力を渡したときお問い合わせ一覧の取得に成功する", func(t *testing.T) {
		contact, err := domain.Rebuild(1, "山田太郎", "", "", "test@example.com", "お問い合わせ内容", time.Now())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		repo := &stubRepo{findAll: []*domain.Contact{contact}}
		uc := New(repo, &stubNotifier{})

		res, err := uc.List(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1, got %d", len(res))
		}
	})

	t.Run("一覧取得でエラーが発生したなら内部エラーを返す", func(t *testing.T) {
		repo := &stubRepo{findAllErr: errors.New("db error")}
		uc := New(repo, &stubNotifier{})

		_, err := uc.List(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestCreateContact(t *testing.T) {
	t.Run("有効な入力を渡したとき作成と通知に成功する", func(t *testing.T) {
		repo := &stubRepo{}
		notifier := &stubNotifier{}
		uc := New(repo, notifier)

		err := uc.Create(
			context.Background(),
			"山田太郎",
			"株式会社サンプル",
			"09012345678",
			"test@example.com",
			"お問い合わせです",
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.created == nil {
			t.Fatalf("expected contact created")
		}
		if notifier.called != 1 {
			t.Fatalf("expected notifier called once, got %d", notifier.called)
		}
	})

	t.Run("入力値が不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		repo := &stubRepo{}
		notifier := &stubNotifier{}
		uc := New(repo, notifier)

		err := uc.Create(context.Background(), "", "", "", "invalid", "")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
		if notifier.called != 0 {
			t.Fatalf("expected notifier not called, got %d", notifier.called)
		}
	})

	t.Run("作成処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {
		repo := &stubRepo{createErr: errors.New("db error")}
		notifier := &stubNotifier{}
		uc := New(repo, notifier)

		err := uc.Create(context.Background(), "山田太郎", "", "", "test@example.com", "お問い合わせです")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
		if notifier.called != 0 {
			t.Fatalf("expected notifier not called, got %d", notifier.called)
		}
	})
}

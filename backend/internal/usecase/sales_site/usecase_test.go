package sales_site

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/sales_site"
	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

var testUUID = id.GenerateUUID()

type stubRepo struct {
	findAll    []*domain.SalesSite
	findAllErr error
	createErr  error
	created    *domain.SalesSite
	findByName *domain.SalesSite
	findByUUID *domain.SalesSite
	findByErr  error
	updateOK   bool
	updateErr  error
	deleteOK   bool
	deleteErr  error
}

type stubUUIDGen struct {
	uuid string
}

func (g *stubUUIDGen) New() string {
	return g.uuid
}

type stubTxManager struct{}

func (s *stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func (s *stubRepo) Create(ctx context.Context, salesSite *domain.SalesSite) (*domain.SalesSite, error) {
	s.created = salesSite
	if s.createErr != nil {
		return nil, s.createErr
	}
	return salesSite, nil
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.SalesSite, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) FindByName(ctx context.Context, name domain.SalesSiteName) (*domain.SalesSite, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByName, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.SalesSite, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByUUID, nil
}

func (s *stubRepo) Update(ctx context.Context, salesSite *domain.SalesSite) (bool, error) {
	if s.updateErr != nil {
		return false, s.updateErr
	}
	return s.updateOK, nil
}

func (s *stubRepo) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	if s.deleteErr != nil {
		return false, s.deleteErr
	}
	return s.deleteOK, nil
}

func TestListSalesSites(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		s := mustSalesSite(testUUID, "Creema", "https://www.creema.jp")
		repo := &stubRepo{findAll: []*domain.SalesSite{s}}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		res, err := uc.List(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1, got %d", len(res))
		}
	})
	t.Run("リポジトリでエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{findAllErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		_, err := uc.List(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestCreateSalesSite(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		if err := uc.Create(context.Background(), "Creema", "https://www.creema.jp"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.created == nil {
			t.Fatalf("expected sales site created")
		}
		if repo.created.UUID().Value() != testUUID {
			t.Fatalf("expected uuid %s, got %s", testUUID, repo.created.UUID().Value())
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Create(context.Background(), "", "https://www.creema.jp")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Create(context.Background(), "Creema", "not-url")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("作成処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{createErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Create(context.Background(), "Creema", "https://www.creema.jp")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestUpdateSalesSite(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustSalesSite(testUUID, "old", "https://old.example.com"),
			updateOK:   true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		if err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Update(context.Background(), "bad-uuid", "new", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustSalesSite(testUUID, "old", "https://old.example.com")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Update(context.Background(), testUUID, "", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustSalesSite(testUUID, "old", "https://old.example.com")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Update(context.Background(), testUUID, "new", "not-url")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{findByUUID: nil}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("更新処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustSalesSite(testUUID, "old", "https://old.example.com"),
			updateErr:  errors.New("db error"),
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestDeleteSalesSite(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{deleteOK: true}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		if err := uc.Delete(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Delete(context.Background(), "bad-uuid")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{deleteOK: false}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Delete(context.Background(), testUUID)
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("リポジトリでエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{deleteErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID}, &stubTxManager{})

		err := uc.Delete(context.Background(), testUUID)
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func mustSalesSite(uuidStr, name, rawURL string) *domain.SalesSite {
	salesSite, err := domain.Rebuild(1, uuidStr, name, rawURL)
	if err != nil {
		panic(err)
	}
	return salesSite
}

package category

import (
	"context"
	"errors"
	"testing"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()

type stubRepo struct {
	findAll     []*domain.Category
	findUsed    []*domain.Category
	findAllErr  error
	findUsedErr error
	exists      bool
	existsErr   error
	createErr   error
	created     *domain.Category
	findByUUID  *domain.Category
	findByErr   error
	updateOK    bool
	updateErr   error
	deleteOK    bool
	deleteErr   error
}

type stubUUIDGen struct {
	uuid string
}

func (g *stubUUIDGen) New() string {
	return g.uuid
}

func (s *stubRepo) Create(ctx context.Context, c *domain.Category) error {
	s.created = c
	return s.createErr
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.Category, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) FindUsed(ctx context.Context) ([]*domain.Category, error) {
	if s.findUsedErr != nil {
		return nil, s.findUsedErr
	}
	return s.findUsed, nil
}

func (s *stubRepo) ExistsByName(ctx context.Context, name domain.CategoryName) (bool, error) {
	if s.existsErr != nil {
		return false, s.existsErr
	}
	return s.exists, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Category, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByUUID, nil
}

func (s *stubRepo) FindByName(ctx context.Context, name domain.CategoryName) (*domain.Category, error) {
	return nil, nil
}

func (s *stubRepo) Update(ctx context.Context, c *domain.Category) (bool, error) {
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

func TestListCategories(t *testing.T) {
	t.Run("allモードを指定したとき全件データの取得に成功する", func(t *testing.T) {

		cat := mustCategory(testUUID, "a")
		repo := &stubRepo{findAll: []*domain.Category{cat}}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		res, err := uc.List(context.Background(), "all")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1, got %d", len(res))
		}
	})
	t.Run("usedモードを指定したとき使用中データの取得に成功する", func(t *testing.T) {

		cat := mustCategory(testUUID, "a")
		repo := &stubRepo{findUsed: []*domain.Category{cat}}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		res, err := uc.List(context.Background(), "used")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1, got %d", len(res))
		}
	})
	t.Run("モードが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		_, err := uc.List(context.Background(), "bad")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		if err := uc.Create(context.Background(), "new"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.created == nil {
			t.Fatalf("expected category created")
		}
		if repo.created.UUID().Value() != testUUID {
			t.Fatalf("expected uuid %s, got %s", testUUID, repo.created.UUID().Value())
		}
	})
	t.Run("名前が重複しているなら重複エラーを返す", func(t *testing.T) {

		repo := &stubRepo{exists: true}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "dup")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrConflict) {
			t.Fatalf("expected ErrConflict, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("リポジトリでエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{existsErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "ok")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
	t.Run("作成処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{createErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "ok")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustCategory(testUUID, "old"),
			updateOK:   true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		if err := uc.Update(context.Background(), testUUID, "new"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), "bad-uuid", "new")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustCategory(testUUID, "old")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{findByUUID: nil}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "new")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("名前が重複しているなら重複エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustCategory(testUUID, "old"),
			exists:     true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "new")
		if err == nil || !errors.Is(err, usecase.ErrConflict) {
			t.Fatalf("expected ErrConflict, got %v", err)
		}
	})
	t.Run("更新処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustCategory(testUUID, "old"),
			updateErr:  errors.New("db error"),
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "new")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{deleteOK: true}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		if err := uc.Delete(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Delete(context.Background(), "bad-uuid")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{deleteOK: false}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Delete(context.Background(), testUUID)
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("リポジトリでエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{deleteErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Delete(context.Background(), testUUID)
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func mustCategory(uuidStr, name string) *domain.Category {
	category, err := domain.Rebuild(1, uuidStr, name)
	if err != nil {
		panic(err)
	}
	return category
}

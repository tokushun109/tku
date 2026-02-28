package skill_market

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/backend/internal/domain/skill_market"
	"github.com/tokushun109/tku/backend/internal/shared/id"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

var testUUID = id.GenerateUUID()

type stubRepo struct {
	findAll    []*domain.SkillMarket
	findAllErr error
	createErr  error
	created    *domain.SkillMarket
	findByUUID *domain.SkillMarket
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

func (s *stubRepo) Create(ctx context.Context, skillMarket *domain.SkillMarket) (*domain.SkillMarket, error) {
	s.created = skillMarket
	if s.createErr != nil {
		return nil, s.createErr
	}
	return skillMarket, nil
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.SkillMarket, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.SkillMarket, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByUUID, nil
}

func (s *stubRepo) Update(ctx context.Context, skillMarket *domain.SkillMarket) (bool, error) {
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

func TestListSkillMarkets(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		s := mustSkillMarket(testUUID, "minne", "https://minne.com")
		repo := &stubRepo{findAll: []*domain.SkillMarket{s}}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

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
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		_, err := uc.List(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestCreateSkillMarket(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		if err := uc.Create(context.Background(), "minne", "https://minne.com"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.created == nil {
			t.Fatalf("expected skill market created")
		}
		if repo.created.UUID().Value() != testUUID {
			t.Fatalf("expected uuid %s, got %s", testUUID, repo.created.UUID().Value())
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "", "https://minne.com")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "minne", "not-url")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("作成処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{createErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Create(context.Background(), "minne", "https://minne.com")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestUpdateSkillMarket(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com"),
			updateOK:   true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		if err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), "bad-uuid", "new", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com")}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "new", "not-url")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{findByUUID: nil}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("更新処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com"),
			updateErr:  errors.New("db error"),
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUID})

		err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestDeleteSkillMarket(t *testing.T) {
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

func mustSkillMarket(uuidStr, name, rawURL string) *domain.SkillMarket {
	skillMarket, err := domain.Rebuild(1, uuidStr, name, rawURL)
	if err != nil {
		panic(err)
	}
	return skillMarket
}

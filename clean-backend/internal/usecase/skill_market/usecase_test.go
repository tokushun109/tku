package skill_market

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/skill_market"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()
var testUUIDVO = mustNewUUID(testUUID)

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
	uuid primitive.UUID
	err  error
}

func (g *stubUUIDGen) New() (primitive.UUID, error) {
	return g.uuid, g.err
}

func (s *stubRepo) Create(ctx context.Context, skillMarket *domain.SkillMarket) error {
	s.created = skillMarket
	return s.createErr
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

		s := mustSkillMarket(testUUID, "minne", "https://minne.com", "")
		repo := &stubRepo{findAll: []*domain.SkillMarket{s}}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

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
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		_, err := uc.List(context.Background())
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestCreateSkillMarket(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		if err := uc.Create(context.Background(), "minne", "https://minne.com", ""); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.created == nil {
			t.Fatalf("expected skill market created")
		}
		if repo.created.UUID.String() != testUUIDVO.String() {
			t.Fatalf("expected uuid %s, got %s", testUUIDVO.String(), repo.created.UUID.String())
		}
	})
	t.Run("UUID生成でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{err: errors.New("gen error")})

		err := uc.Create(context.Background(), "minne", "https://minne.com", "")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Create(context.Background(), "", "https://minne.com", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Create(context.Background(), "minne", "not-url", "")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("作成処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{createErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Create(context.Background(), "minne", "https://minne.com", "")
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
			findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com", ""),
			updateOK:   true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		if err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com", "icon"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), "bad-uuid", "new", "https://new.example.com", "")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com", "")}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "", "https://new.example.com", "")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("URLが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com", "")}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "new", "not-url", "")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{findByUUID: nil}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com", "")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("更新処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustSkillMarket(testUUID, "old", "https://old.example.com", ""),
			updateErr:  errors.New("db error"),
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com", "")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestDeleteSkillMarket(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{deleteOK: true}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		if err := uc.Delete(context.Background(), testUUID); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Delete(context.Background(), "bad-uuid")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{deleteOK: false}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Delete(context.Background(), testUUID)
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("リポジトリでエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{deleteErr: errors.New("db error")}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Delete(context.Background(), testUUID)
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func mustSkillMarket(uuidStr, name, rawURL, icon string) *domain.SkillMarket {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewSkillMarketName(name)
	if err != nil {
		panic(err)
	}
	u2, err := primitive.NewURL(rawURL)
	if err != nil {
		panic(err)
	}
	return &domain.SkillMarket{UUID: u, Name: n, URL: u2, Icon: icon}
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

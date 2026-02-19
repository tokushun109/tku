package tag

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()
var testUUIDVO = mustNewUUID(testUUID)

type stubRepo struct {
	findAll    []*domain.Tag
	findAllErr error
	exists     bool
	existsErr  error
	createErr  error
	created    *domain.Tag
	findByUUID *domain.Tag
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

func (s *stubRepo) Create(ctx context.Context, t *domain.Tag) error {
	s.created = t
	return s.createErr
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.Tag, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) ExistsByName(ctx context.Context, name domain.TagName) (bool, error) {
	if s.existsErr != nil {
		return false, s.existsErr
	}
	return s.exists, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Tag, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByUUID, nil
}

func (s *stubRepo) Update(ctx context.Context, t *domain.Tag) (bool, error) {
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

func TestListTags(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		tg := mustTag(testUUID, "a")
		repo := &stubRepo{findAll: []*domain.Tag{tg}}
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

func TestCreateTag(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		if err := uc.Create(context.Background(), "new"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.created == nil {
			t.Fatalf("expected tag created")
		}
		if repo.created.UUID.String() != testUUIDVO.String() {
			t.Fatalf("expected uuid %s, got %s", testUUIDVO.String(), repo.created.UUID.String())
		}
	})
	t.Run("UUID生成でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{err: errors.New("gen error")})

		err := uc.Create(context.Background(), "new")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
	t.Run("名前が重複しているなら重複エラーを返す", func(t *testing.T) {

		repo := &stubRepo{exists: true}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

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
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

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
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

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
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Create(context.Background(), "ok")
		if err == nil {
			t.Fatalf("expected error")
		}
		if !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestUpdateTag(t *testing.T) {
	t.Run("有効な入力を渡したとき処理に成功する", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustTag(testUUID, "old"),
			updateOK:   true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		if err := uc.Update(context.Background(), testUUID, "new"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("UUIDが不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), "bad-uuid", "new")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("名前が不正なときバリデーションエラーで失敗する", func(t *testing.T) {

		repo := &stubRepo{findByUUID: mustTag(testUUID, "old")}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
	t.Run("対象が見つからないならNotFoundエラーを返す", func(t *testing.T) {

		repo := &stubRepo{findByUUID: nil}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "new")
		if err == nil || !errors.Is(err, usecase.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("名前が重複しているなら重複エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustTag(testUUID, "old"),
			exists:     true,
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "new")
		if err == nil || !errors.Is(err, usecase.ErrConflict) {
			t.Fatalf("expected ErrConflict, got %v", err)
		}
	})
	t.Run("更新処理でエラーが発生したなら内部エラーを返す", func(t *testing.T) {

		repo := &stubRepo{
			findByUUID: mustTag(testUUID, "old"),
			updateErr:  errors.New("db error"),
		}
		uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

		err := uc.Update(context.Background(), testUUID, "new")
		if err == nil || !errors.Is(err, usecase.ErrInternal) {
			t.Fatalf("expected ErrInternal, got %v", err)
		}
	})
}

func TestDeleteTag(t *testing.T) {
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

func mustTag(uuidStr, name string) *domain.Tag {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewTagName(name)
	if err != nil {
		panic(err)
	}
	return &domain.Tag{UUID: u, Name: n}
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

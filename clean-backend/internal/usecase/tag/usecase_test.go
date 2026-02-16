package tag

import (
	"context"
	"errors"
	"testing"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/tag"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
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

func TestListTags_OK(t *testing.T) {
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
}

func TestListTags_RepoError(t *testing.T) {
	repo := &stubRepo{findAllErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	_, err := uc.List(context.Background())
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateTag_OK(t *testing.T) {
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
}

func TestCreateTag_UUIDGenError(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{err: errors.New("gen error")})

	err := uc.Create(context.Background(), "new")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateTag_DuplicateName(t *testing.T) {
	repo := &stubRepo{exists: true}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "dup")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrConflict) {
		t.Fatalf("expected ErrConflict, got %v", err)
	}
}

func TestCreateTag_InvalidName(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestCreateTag_RepoError(t *testing.T) {
	repo := &stubRepo{existsErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "ok")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateTag_CreateError(t *testing.T) {
	repo := &stubRepo{createErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "ok")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestUpdateTag_OK(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustTag(testUUID, "old"),
		updateOK:   true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Update(context.Background(), testUUID, "new"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateTag_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), "bad-uuid", "new")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateTag_InvalidName(t *testing.T) {
	repo := &stubRepo{findByUUID: mustTag(testUUID, "old")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateTag_NotFound(t *testing.T) {
	repo := &stubRepo{findByUUID: nil}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateTag_DuplicateName(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustTag(testUUID, "old"),
		exists:     true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrConflict) {
		t.Fatalf("expected ErrConflict, got %v", err)
	}
}

func TestUpdateTag_UpdateError(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustTag(testUUID, "old"),
		updateErr:  errors.New("db error"),
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestDeleteTag_OK(t *testing.T) {
	repo := &stubRepo{deleteOK: true}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Delete(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteTag_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), "bad-uuid")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestDeleteTag_NotFound(t *testing.T) {
	repo := &stubRepo{deleteOK: false}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), testUUID)
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteTag_RepoError(t *testing.T) {
	repo := &stubRepo{deleteErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), testUUID)
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
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

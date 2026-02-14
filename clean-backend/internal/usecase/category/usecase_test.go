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
var testUUIDVO = mustNewUUID(testUUID)

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
}

type stubUUIDGen struct {
	uuid primitive.UUID
	err  error
}

func (g *stubUUIDGen) New() (primitive.UUID, error) {
	return g.uuid, g.err
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

func (s *stubRepo) Update(ctx context.Context, c *domain.Category) (bool, error) {
	if s.updateErr != nil {
		return false, s.updateErr
	}
	return s.updateOK, nil
}

func TestListCategories_All_OK(t *testing.T) {
	cat := mustCategory(testUUID, "a")
	repo := &stubRepo{findAll: []*domain.Category{cat}}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	res, err := uc.List(context.Background(), "all")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("expected 1, got %d", len(res))
	}
}

func TestListCategories_Used_OK(t *testing.T) {
	cat := mustCategory(testUUID, "a")
	repo := &stubRepo{findUsed: []*domain.Category{cat}}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	res, err := uc.List(context.Background(), "used")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("expected 1, got %d", len(res))
	}
}

func TestListCategories_InvalidMode(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	_, err := uc.List(context.Background(), "bad")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestCreateCategory_OK(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Create(context.Background(), "new"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.created == nil {
		t.Fatalf("expected category created")
	}
	if repo.created.UUID.String() != testUUIDVO.String() {
		t.Fatalf("expected uuid %s, got %s", testUUIDVO.String(), repo.created.UUID.String())
	}
}

func TestCreateCategory_UUIDGenError(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{err: errors.New("gen error")})

	err := uc.Create(context.Background(), "new")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateCategory_DuplicateName(t *testing.T) {
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

func TestCreateCategory_InvalidName(t *testing.T) {
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

func TestCreateCategory_RepoError(t *testing.T) {
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

func TestCreateCategory_CreateError(t *testing.T) {
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

func TestUpdateCategory_OK(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustCategory(testUUID, "old"),
		updateOK:   true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Update(context.Background(), testUUID, "new"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateCategory_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), "bad-uuid", "new")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateCategory_InvalidName(t *testing.T) {
	repo := &stubRepo{findByUUID: mustCategory(testUUID, "old")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateCategory_NotFound(t *testing.T) {
	repo := &stubRepo{findByUUID: nil}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateCategory_DuplicateName(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustCategory(testUUID, "old"),
		exists:     true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrConflict) {
		t.Fatalf("expected ErrConflict, got %v", err)
	}
}

func TestUpdateCategory_UpdateError(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustCategory(testUUID, "old"),
		updateErr:  errors.New("db error"),
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func mustCategory(uuidStr, name string) *domain.Category {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewCategoryName(name)
	if err != nil {
		panic(err)
	}
	return &domain.Category{UUID: u, Name: n}
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

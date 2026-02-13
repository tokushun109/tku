package category

import (
	"context"
	"errors"
	"testing"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/category"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.NewUUID().String()

type stubRepo struct {
	findAll     []*domain.Category
	findUsed    []*domain.Category
	findAllErr  error
	findUsedErr error
	exists      bool
	existsErr   error
	createErr   error
	created     *domain.Category
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

func TestListCategories_All_OK(t *testing.T) {
	cat := mustCategory(testUUID, "a")
	repo := &stubRepo{findAll: []*domain.Category{cat}}
	uc := New(repo)

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
	uc := New(repo)

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
	uc := New(repo)

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
	uc := New(repo)

	if err := uc.Create(context.Background(), "new"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.created == nil {
		t.Fatalf("expected category created")
	}
	if repo.created.UUID.String() == "" {
		t.Fatalf("expected UUID to be set")
	}
}

func TestCreateCategory_DuplicateName(t *testing.T) {
	repo := &stubRepo{exists: true}
	uc := New(repo)

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
	uc := New(repo)

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
	uc := New(repo)

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
	uc := New(repo)

	err := uc.Create(context.Background(), "ok")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func mustCategory(uuidStr, name string) *domain.Category {
	u, err := domain.ParseCategoryUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewCategoryName(name)
	if err != nil {
		panic(err)
	}
	return &domain.Category{UUID: u, Name: n}
}

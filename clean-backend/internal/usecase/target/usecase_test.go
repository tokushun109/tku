package target

import (
	"context"
	"errors"
	"testing"

	domain "github.com/tokushun109/tku/clean-backend/internal/domain/target"
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()
var testUUIDVO = mustNewUUID(testUUID)

type stubRepo struct {
	findAll     []*domain.Target
	findUsed    []*domain.Target
	findAllErr  error
	findUsedErr error
	exists      bool
	existsErr   error
	createErr   error
	created     *domain.Target
	findByUUID  *domain.Target
	findByErr   error
	updateOK    bool
	updateErr   error
	deleteOK    bool
	deleteErr   error
}

type stubUUIDGen struct {
	uuid primitive.UUID
	err  error
}

func (g *stubUUIDGen) New() (primitive.UUID, error) {
	return g.uuid, g.err
}

func (s *stubRepo) Create(ctx context.Context, t *domain.Target) error {
	s.created = t
	return s.createErr
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.Target, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) FindUsed(ctx context.Context) ([]*domain.Target, error) {
	if s.findUsedErr != nil {
		return nil, s.findUsedErr
	}
	return s.findUsed, nil
}

func (s *stubRepo) ExistsByName(ctx context.Context, name domain.TargetName) (bool, error) {
	if s.existsErr != nil {
		return false, s.existsErr
	}
	return s.exists, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Target, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByUUID, nil
}

func (s *stubRepo) Update(ctx context.Context, t *domain.Target) (bool, error) {
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

func TestListTargets_All_OK(t *testing.T) {
	tg := mustTarget(testUUID, "a")
	repo := &stubRepo{findAll: []*domain.Target{tg}}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	res, err := uc.List(context.Background(), "all")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("expected 1, got %d", len(res))
	}
}

func TestListTargets_Used_OK(t *testing.T) {
	tg := mustTarget(testUUID, "a")
	repo := &stubRepo{findUsed: []*domain.Target{tg}}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	res, err := uc.List(context.Background(), "used")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("expected 1, got %d", len(res))
	}
}

func TestListTargets_InvalidMode(t *testing.T) {
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

func TestCreateTarget_OK(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Create(context.Background(), "new"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.created == nil {
		t.Fatalf("expected target created")
	}
	if repo.created.UUID.String() != testUUIDVO.String() {
		t.Fatalf("expected uuid %s, got %s", testUUIDVO.String(), repo.created.UUID.String())
	}
}

func TestCreateTarget_UUIDGenError(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{err: errors.New("gen error")})

	err := uc.Create(context.Background(), "new")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateTarget_DuplicateName(t *testing.T) {
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

func TestCreateTarget_InvalidName(t *testing.T) {
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

func TestCreateTarget_RepoError(t *testing.T) {
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

func TestCreateTarget_CreateError(t *testing.T) {
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

func TestUpdateTarget_OK(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustTarget(testUUID, "old"),
		updateOK:   true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Update(context.Background(), testUUID, "new"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateTarget_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), "bad-uuid", "new")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateTarget_InvalidName(t *testing.T) {
	repo := &stubRepo{findByUUID: mustTarget(testUUID, "old")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateTarget_NotFound(t *testing.T) {
	repo := &stubRepo{findByUUID: nil}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateTarget_DuplicateName(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustTarget(testUUID, "old"),
		exists:     true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrConflict) {
		t.Fatalf("expected ErrConflict, got %v", err)
	}
}

func TestUpdateTarget_UpdateError(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustTarget(testUUID, "old"),
		updateErr:  errors.New("db error"),
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestDeleteTarget_OK(t *testing.T) {
	repo := &stubRepo{deleteOK: true}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Delete(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteTarget_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), "bad-uuid")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestDeleteTarget_NotFound(t *testing.T) {
	repo := &stubRepo{deleteOK: false}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), testUUID)
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteTarget_RepoError(t *testing.T) {
	repo := &stubRepo{deleteErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), testUUID)
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func mustTarget(uuidStr, name string) *domain.Target {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewTargetName(name)
	if err != nil {
		panic(err)
	}
	return &domain.Target{UUID: u, Name: n}
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

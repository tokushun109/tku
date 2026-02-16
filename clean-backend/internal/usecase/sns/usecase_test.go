package sns

import (
	"context"
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	domain "github.com/tokushun109/tku/clean-backend/internal/domain/sns"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

var testUUID = id.GenerateUUID()
var testUUIDVO = mustNewUUID(testUUID)

type stubRepo struct {
	findAll    []*domain.Sns
	findAllErr error
	createErr  error
	created    *domain.Sns
	findByUUID *domain.Sns
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

func (s *stubRepo) Create(ctx context.Context, snsObj *domain.Sns) error {
	s.created = snsObj
	return s.createErr
}

func (s *stubRepo) FindAll(ctx context.Context) ([]*domain.Sns, error) {
	if s.findAllErr != nil {
		return nil, s.findAllErr
	}
	return s.findAll, nil
}

func (s *stubRepo) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domain.Sns, error) {
	if s.findByErr != nil {
		return nil, s.findByErr
	}
	return s.findByUUID, nil
}

func (s *stubRepo) Update(ctx context.Context, snsObj *domain.Sns) (bool, error) {
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

func TestListSns_OK(t *testing.T) {
	s := mustSns(testUUID, "Instagram", "https://www.instagram.com", "")
	repo := &stubRepo{findAll: []*domain.Sns{s}}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	res, err := uc.List(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("expected 1, got %d", len(res))
	}
}

func TestListSns_RepoError(t *testing.T) {
	repo := &stubRepo{findAllErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	_, err := uc.List(context.Background())
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateSns_OK(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Create(context.Background(), "Instagram", "https://www.instagram.com", ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.created == nil {
		t.Fatalf("expected sns created")
	}
	if repo.created.UUID.String() != testUUIDVO.String() {
		t.Fatalf("expected uuid %s, got %s", testUUIDVO.String(), repo.created.UUID.String())
	}
}

func TestCreateSns_UUIDGenError(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{err: errors.New("gen error")})

	err := uc.Create(context.Background(), "Instagram", "https://www.instagram.com", "")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestCreateSns_InvalidName(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "", "https://www.instagram.com", "")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestCreateSns_InvalidURL(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "Instagram", "not-url", "")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestCreateSns_CreateError(t *testing.T) {
	repo := &stubRepo{createErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Create(context.Background(), "Instagram", "https://www.instagram.com", "")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestUpdateSns_OK(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustSns(testUUID, "old", "https://old.example.com", ""),
		updateOK:   true,
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com", "icon"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateSns_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), "bad-uuid", "new", "https://new.example.com", "")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateSns_InvalidName(t *testing.T) {
	repo := &stubRepo{findByUUID: mustSns(testUUID, "old", "https://old.example.com", "")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "", "https://new.example.com", "")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateSns_InvalidURL(t *testing.T) {
	repo := &stubRepo{findByUUID: mustSns(testUUID, "old", "https://old.example.com", "")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new", "not-url", "")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdateSns_NotFound(t *testing.T) {
	repo := &stubRepo{findByUUID: nil}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com", "")
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateSns_UpdateError(t *testing.T) {
	repo := &stubRepo{
		findByUUID: mustSns(testUUID, "old", "https://old.example.com", ""),
		updateErr:  errors.New("db error"),
	}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Update(context.Background(), testUUID, "new", "https://new.example.com", "")
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func TestDeleteSns_OK(t *testing.T) {
	repo := &stubRepo{deleteOK: true}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	if err := uc.Delete(context.Background(), testUUID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteSns_InvalidUUID(t *testing.T) {
	repo := &stubRepo{}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), "bad-uuid")
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestDeleteSns_NotFound(t *testing.T) {
	repo := &stubRepo{deleteOK: false}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), testUUID)
	if err == nil || !errors.Is(err, usecase.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteSns_RepoError(t *testing.T) {
	repo := &stubRepo{deleteErr: errors.New("db error")}
	uc := New(repo, &stubUUIDGen{uuid: testUUIDVO})

	err := uc.Delete(context.Background(), testUUID)
	if err == nil || !errors.Is(err, usecase.ErrInternal) {
		t.Fatalf("expected ErrInternal, got %v", err)
	}
}

func mustSns(uuidStr, name, rawURL, icon string) *domain.Sns {
	u, err := primitive.NewUUID(uuidStr)
	if err != nil {
		panic(err)
	}
	n, err := domain.NewSnsName(name)
	if err != nil {
		panic(err)
	}
	u2, err := primitive.NewURL(rawURL)
	if err != nil {
		panic(err)
	}
	return &domain.Sns{UUID: u, Name: n, URL: u2, Icon: icon}
}

func mustNewUUID(s string) primitive.UUID {
	u, err := primitive.NewUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

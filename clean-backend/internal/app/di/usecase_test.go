package di

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
)

type stubTxManager struct{}

func (s *stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func TestNewUsecasesNilRepositories(t *testing.T) {
	ucs, err := newUsecases(nil, &config.Config{}, &stubTxManager{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
	if ucs != nil {
		t.Fatalf("expected nil usecases, got %#v", ucs)
	}
}

func TestNewUsecasesNilConfig(t *testing.T) {
	repos, err := newRepositories(&sqlx.DB{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ucs, err := newUsecases(repos, nil, &stubTxManager{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
	if ucs != nil {
		t.Fatalf("expected nil usecases, got %#v", ucs)
	}
}

func TestNewUsecasesNilTxManager(t *testing.T) {
	repos, err := newRepositories(&sqlx.DB{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ucs, err := newUsecases(repos, &config.Config{}, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
	if ucs != nil {
		t.Fatalf("expected nil usecases, got %#v", ucs)
	}
}

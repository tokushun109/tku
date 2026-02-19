package di

import (
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
)

func TestNewUsecasesNilRepositories(t *testing.T) {
	ucs, err := newUsecases(nil, &config.Config{})
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

	ucs, err := newUsecases(repos, nil)
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

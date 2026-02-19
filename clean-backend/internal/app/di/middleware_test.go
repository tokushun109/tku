package di

import (
	"errors"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
)

func TestNewMiddlewaresNilConfig(t *testing.T) {
	mws, err := newMiddlewares(nil, &usecases{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
	if mws != nil {
		t.Fatalf("expected nil middlewares, got %#v", mws)
	}
}

func TestNewMiddlewaresNilUsecases(t *testing.T) {
	mws, err := newMiddlewares(&config.Config{}, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
	if mws != nil {
		t.Fatalf("expected nil middlewares, got %#v", mws)
	}
}

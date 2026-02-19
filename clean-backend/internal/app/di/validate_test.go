package di

import (
	"errors"
	"testing"

	usecaseHealth "github.com/tokushun109/tku/clean-backend/internal/usecase/health"
)

type testUsecase interface {
	Run()
}

type testService struct{}

func (*testService) Run() {}

func TestRequireNonNilNil(t *testing.T) {
	if err := requireNonNil("dep", nil); err == nil {
		t.Fatal("expected error, got nil")
	} else if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
}

func TestRequireNonNilTypedNil(t *testing.T) {
	var healthUC usecaseHealth.Usecase = (*usecaseHealth.Service)(nil)
	if err := requireNonNil("healthUC", healthUC); err == nil {
		t.Fatal("expected error, got nil")
	} else if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
}

func TestRequireNonNilNonNil(t *testing.T) {
	if err := requireNonNil("dep", struct{}{}); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestRequireStructFieldsNonNilNilField(t *testing.T) {
	deps := struct {
		A *int
	}{}
	if err := requireStructFieldsNonNil("deps", deps); err == nil {
		t.Fatal("expected error, got nil")
	} else if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
}

func TestRequireStructFieldsNonNilTypedNilInterfaceField(t *testing.T) {
	deps := struct {
		UC testUsecase
	}{
		UC: (*testService)(nil),
	}
	if err := requireStructFieldsNonNil("deps", deps); err == nil {
		t.Fatal("expected error, got nil")
	} else if !errors.Is(err, ErrNilDependency) {
		t.Fatalf("expected ErrNilDependency, got %v", err)
	}
}

func TestRequireStructFieldsNonNilNonNil(t *testing.T) {
	value := 1
	deps := struct {
		A *int
	}{
		A: &value,
	}
	if err := requireStructFieldsNonNil("deps", deps); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

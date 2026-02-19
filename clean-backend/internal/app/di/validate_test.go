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

func TestRequireNonNil(t *testing.T) {

	t.Run("nil値を渡したならErrNilDependencyを返す", func(t *testing.T) {
		if err := requireNonNil("dep", nil); err == nil {
			t.Fatal("expected error, got nil")
		} else if !errors.Is(err, ErrNilDependency) {
			t.Fatalf("expected ErrNilDependency, got %v", err)
		}
	})

	t.Run("typed nilを渡したならErrNilDependencyを返す", func(t *testing.T) {
		var healthUC usecaseHealth.Usecase = (*usecaseHealth.Service)(nil)
		if err := requireNonNil("healthUC", healthUC); err == nil {
			t.Fatal("expected error, got nil")
		} else if !errors.Is(err, ErrNilDependency) {
			t.Fatalf("expected ErrNilDependency, got %v", err)
		}
	})

	t.Run("nilでない値を渡したならnilを返す", func(t *testing.T) {
		if err := requireNonNil("dep", struct{}{}); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

}

func TestRequireStructFieldsNonNil(t *testing.T) {

	t.Run("構造体にnilフィールドがあるならErrNilDependencyを返す", func(t *testing.T) {
		deps := struct {
			A *int
		}{}
		if err := requireStructFieldsNonNil("deps", deps); err == nil {
			t.Fatal("expected error, got nil")
		} else if !errors.Is(err, ErrNilDependency) {
			t.Fatalf("expected ErrNilDependency, got %v", err)
		}
	})

	t.Run("interface型フィールドがtyped nilならErrNilDependencyを返す", func(t *testing.T) {
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
	})

	t.Run("全フィールドがnilでないならnilを返す", func(t *testing.T) {
		value := 1
		deps := struct {
			A *int
		}{
			A: &value,
		}
		if err := requireStructFieldsNonNil("deps", deps); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

}

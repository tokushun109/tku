package di

import (
	"errors"
	"testing"
)

func TestNewRepositoriesNilDB(t *testing.T) {
	t.Run("dbがnilならエラーを返す", func(t *testing.T) {

		repos, err := newRepositories(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNilDependency) {
			t.Fatalf("expected ErrNilDependency, got %v", err)
		}
		if repos != nil {
			t.Fatalf("expected nil repositories, got %#v", repos)
		}
	})

}

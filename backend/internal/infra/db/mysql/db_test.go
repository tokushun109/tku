package mysql

import (
	"errors"
	"testing"
)

func TestNewDBNilConfig(t *testing.T) {
	t.Run("configがnilならエラーを返す", func(t *testing.T) {

		db, err := NewDB(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNilConfig) {
			t.Fatalf("expected ErrNilConfig, got %v", err)
		}
		if db != nil {
			t.Fatalf("expected nil db, got %#v", db)
		}
	})

}

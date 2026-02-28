package di

import (
	"errors"
	"testing"
)

func TestNewQueriesNilDB(t *testing.T) {
	t.Run("dbがnilならエラーを返す", func(t *testing.T) {
		qrs, err := newQueries(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNilDependency) {
			t.Fatalf("expected ErrNilDependency, got %v", err)
		}
		if qrs != nil {
			t.Fatalf("expected nil queries, got %#v", qrs)
		}
	})
}

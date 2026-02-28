package di

import (
	"errors"
	"testing"
)

func TestNewHandlersNilUsecases(t *testing.T) {
	t.Run("usecasesがnilならエラーを返す", func(t *testing.T) {

		hs, err := newHandlers(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNilDependency) {
			t.Fatalf("expected ErrNilDependency, got %v", err)
		}
		if hs != nil {
			t.Fatalf("expected nil handlers, got %#v", hs)
		}
	})

}

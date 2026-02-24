package local

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

func TestStoragePutGetDelete(t *testing.T) {
	t.Run("ロゴを保存して取得したあと削除できる", func(t *testing.T) {
		storage := NewStorage(t.TempDir())
		ctx := context.Background()
		key := "img/logo/a/b/logo.png"
		payload := []byte("logo-binary")

		if err := storage.Put(ctx, key, "image/png", payload); err != nil {
			t.Fatalf("put error: %v", err)
		}

		actual, err := storage.Get(ctx, key)
		if err != nil {
			t.Fatalf("get error: %v", err)
		}
		defer func() {
			_ = actual.Close()
		}()
		payloadRead, err := io.ReadAll(actual)
		if err != nil {
			t.Fatalf("read error: %v", err)
		}
		if string(payloadRead) != "logo-binary" {
			t.Fatalf("unexpected payload: %s", string(payloadRead))
		}

		if err := storage.Delete(ctx, key); err != nil {
			t.Fatalf("delete error: %v", err)
		}

		_, err = storage.Get(ctx, key)
		if !errors.Is(err, usecase.ErrStorageNotFound) {
			t.Fatalf("expected ErrStorageNotFound, got %v", err)
		}
	})
}

func TestStorageRejectInvalidKey(t *testing.T) {
	t.Run("相対パスでないキーなら保存処理が失敗する", func(t *testing.T) {
		storage := NewStorage(t.TempDir())
		err := storage.Put(context.Background(), "/tmp/logo.png", "image/png", []byte("x"))
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

package mysql

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNewTxManager(t *testing.T) {
	t.Run("DBがnilならエラーを返す", func(t *testing.T) {

		txm, err := NewTxManager(nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNilDBInTxManager) {
			t.Fatalf("expected ErrNilDBInTxManager, got %v", err)
		}
		if txm != nil {
			t.Fatalf("expected nil tx manager, got %#v", txm)
		}
	})

}

func TestTxManagerWithinTransaction(t *testing.T) {
	t.Run("実行関数がnilならエラーを返す", func(t *testing.T) {

		txm, err := NewTxManager(&sqlx.DB{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		err = txm.WithinTransaction(context.Background(), nil)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNilTxFunc) {
			t.Fatalf("expected ErrNilTxFunc, got %v", err)
		}
	})

}

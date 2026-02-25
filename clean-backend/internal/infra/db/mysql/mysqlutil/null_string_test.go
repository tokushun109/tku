package mysqlutil

import (
	"database/sql"
	"testing"
)

func TestNullStringOrEmpty(t *testing.T) {
	t.Run("NULL のとき空文字を返す", func(t *testing.T) {
		got := NullStringOrEmpty(sql.NullString{})
		if got != "" {
			t.Fatalf("expected empty string, got %q", got)
		}
	})

	t.Run("値ありのとき文字列を返す", func(t *testing.T) {
		got := NullStringOrEmpty(sql.NullString{
			String: "hello",
			Valid:  true,
		})
		if got != "hello" {
			t.Fatalf("expected hello, got %q", got)
		}
	})
}

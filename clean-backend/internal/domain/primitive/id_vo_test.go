package primitive

import "testing"

func TestNewID(t *testing.T) {
	t.Run("有効値を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		id, err := NewID(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id.Value() != 1 {
			t.Fatalf("unexpected id: %d", id.Value())
		}
		if !id.IsDefined() {
			t.Fatal("expected id to be defined")
		}
	})

	t.Run("0を渡したときバリデーションエラーで失敗する", func(t *testing.T) {
		if _, err := NewID(0); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestIDIsDefined(t *testing.T) {
	t.Run("0のとき未定義として扱う", func(t *testing.T) {
		var id ID = 0
		if id.IsDefined() {
			t.Fatal("expected id to be undefined")
		}
	})

	t.Run("0以外のとき定義済みとして扱う", func(t *testing.T) {
		var id ID = 42
		if !id.IsDefined() {
			t.Fatal("expected id to be defined")
		}
	})
}

package primitive

import "testing"

func TestNewID(t *testing.T) {
	t.Run("有効値を渡したとき有効な値の生成に成功する", func(t *testing.T) {
		id, err := NewID(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id.Uint() != 1 {
			t.Fatalf("unexpected id: %d", id.Uint())
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

func TestNewOptionalID(t *testing.T) {
	t.Run("nilを渡したときnilを返す", func(t *testing.T) {
		id, err := NewOptionalID(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != nil {
			t.Fatalf("expected nil, got %v", *id)
		}
	})

	t.Run("有効値ポインタを渡したときIDポインタを返す", func(t *testing.T) {
		v := uint(10)
		id, err := NewOptionalID(&v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id == nil {
			t.Fatal("expected non-nil id")
		}
		if id.Uint() != 10 {
			t.Fatalf("unexpected id: %d", id.Uint())
		}
	})

	t.Run("0のポインタを渡したときエラーを返す", func(t *testing.T) {
		v := uint(0)
		id, err := NewOptionalID(&v)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if id != nil {
			t.Fatalf("expected nil id, got %v", *id)
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

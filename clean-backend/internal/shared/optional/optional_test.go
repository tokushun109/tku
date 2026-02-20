package optional

import (
	"errors"
	"testing"
)

type testStringAlias string

func TestParseOptionalString(t *testing.T) {
	t.Run("空文字を渡したときnilを返しparserを呼ばない", func(t *testing.T) {
		called := false
		got, err := ParseOptionalString("", func(v string) (string, error) {
			called = true
			return v, nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
		if called {
			t.Fatalf("expected parser not called")
		}
	})

	t.Run("空白文字のみを渡したときnilを返しparserを呼ばない", func(t *testing.T) {
		called := false
		got, err := ParseOptionalString("   ", func(v string) (string, error) {
			called = true
			return v, nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
		if called {
			t.Fatalf("expected parser not called")
		}
	})

	t.Run("値ありを渡したときparser結果のポインタを返す", func(t *testing.T) {
		got, err := ParseOptionalString("abc", func(v string) (string, error) {
			return "parsed:" + v, nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil {
			t.Fatalf("expected non-nil pointer")
		}
		if *got != "parsed:abc" {
			t.Fatalf("unexpected parsed value: %s", *got)
		}
	})

	t.Run("parserがエラーを返したときそのエラーを返す", func(t *testing.T) {
		expectedErr := errors.New("parse error")
		got, err := ParseOptionalString("abc", func(v string) (string, error) {
			return "", expectedErr
		})
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected parse error, got %v", err)
		}
		if got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
	})
}

func TestToStringPtr(t *testing.T) {
	t.Run("nilを渡したときnilを返す", func(t *testing.T) {
		got := ToStringPtr[testStringAlias](nil)
		if got != nil {
			t.Fatalf("expected nil, got %v", *got)
		}
	})

	t.Run("値ありのポインタを渡したとき文字列ポインタを返す", func(t *testing.T) {
		v := testStringAlias("sample")

		got := ToStringPtr(&v)
		if got == nil {
			t.Fatalf("expected non-nil pointer")
		}
		if *got != "sample" {
			t.Fatalf("expected sample, got %s", *got)
		}
	})
}

func TestToTrimmedStringOrEmpty(t *testing.T) {
	t.Run("nilを渡したとき空文字を返す", func(t *testing.T) {
		got := ToTrimmedStringOrEmpty[testStringAlias](nil)
		if got != "" {
			t.Fatalf("expected empty string, got %q", got)
		}
	})

	t.Run("値ありのポインタを渡したときtrimされた文字列を返す", func(t *testing.T) {
		v := testStringAlias("  sample  ")

		got := ToTrimmedStringOrEmpty(&v)
		if got != "sample" {
			t.Fatalf("expected sample, got %s", got)
		}
	})
}

package vo

import (
	"errors"
	"testing"
)

type testVO string

func (v testVO) Value() string {
	return string(v)
}

func (v testVO) String() string {
	return v.Value()
}

type testUintVO uint

func (v testUintVO) Value() uint {
	return uint(v)
}

func (v testUintVO) String() string {
	return "test"
}

func TestParseOptionalValueWithString(t *testing.T) {
	t.Run("空文字を渡したときnilを返しparserを呼ばない", func(t *testing.T) {
		called := false
		raw := ""
		got, err := ParseOptionalValue(&raw, func(raw string) (testVO, error) {
			called = true
			return testVO(raw), nil
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

	t.Run("空白のみを渡したときnilを返しparserを呼ばない", func(t *testing.T) {
		called := false
		raw := "   "
		got, err := ParseOptionalValue(&raw, func(raw string) (testVO, error) {
			called = true
			return testVO(raw), nil
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
		raw := "abc"
		got, err := ParseOptionalValue(&raw, func(raw string) (testVO, error) {
			return testVO("parsed:" + raw), nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil {
			t.Fatalf("expected non-nil pointer")
		}
		if got.Value() != "parsed:abc" {
			t.Fatalf("unexpected parsed value: %s", got.Value())
		}
	})

	t.Run("前後に空白がある文字列を渡したときtrimしてparserに渡す", func(t *testing.T) {
		var received string
		raw := "  abc  "
		got, err := ParseOptionalValue(&raw, func(raw string) (testVO, error) {
			received = raw
			return testVO(raw), nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil {
			t.Fatalf("expected non-nil pointer")
		}
		if received != "abc" {
			t.Fatalf("expected parser input to be trimmed, got %q", received)
		}
		if got.Value() != "abc" {
			t.Fatalf("unexpected parsed value: %s", got.Value())
		}
	})

	t.Run("parserがエラーを返したときそのエラーを返す", func(t *testing.T) {
		expectedErr := errors.New("parse error")
		raw := "abc"
		got, err := ParseOptionalValue(&raw, func(raw string) (testVO, error) {
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

func TestParseOptionalValueWithUint(t *testing.T) {
	t.Run("nilを渡したときnilを返す", func(t *testing.T) {
		got, err := ParseOptionalValue[uint, testUintVO](nil, func(v uint) (testUintVO, error) {
			return testUintVO(v), nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
	})

	t.Run("有効な値ポインタを渡したときVOポインタを返す", func(t *testing.T) {
		raw := uint(10)
		got, err := ParseOptionalValue(&raw, func(v uint) (testUintVO, error) {
			if v == 0 {
				return 0, errors.New("invalid")
			}
			return testUintVO(v), nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil {
			t.Fatalf("expected non-nil pointer")
		}
		if got.Value() != 10 {
			t.Fatalf("unexpected parsed value: %d", got.Value())
		}
	})

	t.Run("不正な値ポインタを渡したときparserのエラーを返す", func(t *testing.T) {
		raw := uint(0)
		expectedErr := errors.New("invalid")
		got, err := ParseOptionalValue(&raw, func(v uint) (testUintVO, error) {
			if v == 0 {
				return 0, expectedErr
			}
			return testUintVO(v), nil
		})
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected parse error, got %v", err)
		}
		if got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
	})
}

func TestToValuePtr(t *testing.T) {
	t.Run("nilを渡したときnilを返す", func(t *testing.T) {
		got := ToValuePtr[string, testVO](nil)
		if got != nil {
			t.Fatalf("expected nil, got %#v", got)
		}
	})

	t.Run("値ありを渡したときValueのポインタを返す", func(t *testing.T) {
		v := testVO("abc")
		got := ToValuePtr(&v)
		if got == nil {
			t.Fatalf("expected non-nil pointer")
		}
		if *got != "abc" {
			t.Fatalf("unexpected value: %s", *got)
		}
	})
}

func TestToValueOrEmpty(t *testing.T) {
	t.Run("nilを渡したとき空文字を返す", func(t *testing.T) {
		got := ToValueOrEmpty[testVO](nil)
		if got != "" {
			t.Fatalf("expected empty string, got %q", got)
		}
	})

	t.Run("値ありを渡したときValueを返す", func(t *testing.T) {
		v := testVO("abc")
		got := ToValueOrEmpty(&v)
		if got != "abc" {
			t.Fatalf("unexpected value: %q", got)
		}
	})
}

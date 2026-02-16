package primitive

import (
	"strings"
	"testing"
)

func TestNewURL_Valid(t *testing.T) {
	u, err := NewURL("https://www.creema.jp/items/123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.String() != "https://www.creema.jp/items/123" {
		t.Fatalf("expected value, got %q", u.String())
	}
}

func TestNewURL_Invalid_TooShort(t *testing.T) {
	_, err := NewURL("")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

func TestNewURL_Invalid_TooLong(t *testing.T) {
	tooLong := "https://example.com/" + strings.Repeat("a", 238)
	_, err := NewURL(tooLong)
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

func TestNewURL_Invalid_NoScheme(t *testing.T) {
	_, err := NewURL("example.com/shop")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

func TestNewURL_Invalid_NoHost(t *testing.T) {
	_, err := NewURL("https://")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

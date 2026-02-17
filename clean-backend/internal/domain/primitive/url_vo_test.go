package primitive

import "testing"

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

func TestNewURL_Invalid_NoScheme(t *testing.T) {
	_, err := NewURL("example.com/shop")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

func TestNewURL_Invalid_NoHostOrOpaqueOrFragment(t *testing.T) {
	_, err := NewURL("https://")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

func TestNewURL_Valid_Mailto(t *testing.T) {
	u, err := NewURL("mailto:test@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.String() != "mailto:test@example.com" {
		t.Fatalf("expected value, got %q", u.String())
	}
}

func TestNewURL_Invalid_WithSurroundingSpaces(t *testing.T) {
	_, err := NewURL(" https://example.com ")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

func TestNewURL_Valid_FileWithPath(t *testing.T) {
	u, err := NewURL("file:///tmp/file.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.String() != "file:///tmp/file.txt" {
		t.Fatalf("expected value, got %q", u.String())
	}
}

func TestNewURL_Invalid_FileRootOnly(t *testing.T) {
	_, err := NewURL("file:///")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidURL {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}

package category

import "testing"

func TestNewCategoryName_Valid(t *testing.T) {
	name, err := NewCategoryName("accessory")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "accessory" {
		t.Fatalf("expected trimmed value, got %q", name.String())
	}
}

func TestNewCategoryName_Valid_Japanese30Chars(t *testing.T) {
	name, err := NewCategoryName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ" {
		t.Fatalf("expected value, got %q", name.String())
	}
}

func TestNewCategoryName_Invalid_TooShort(t *testing.T) {
	_, err := NewCategoryName("")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewCategoryName_Invalid_TooLong(t *testing.T) {
	_, err := NewCategoryName("1234567890123456789012345678901")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewCategoryName_Invalid_TooLong_Japanese31Chars(t *testing.T) {
	_, err := NewCategoryName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほま")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

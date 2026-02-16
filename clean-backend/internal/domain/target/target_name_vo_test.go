package target

import "testing"

func TestNewTargetName_Valid(t *testing.T) {
	name, err := NewTargetName("accessory")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "accessory" {
		t.Fatalf("expected trimmed value, got %q", name.String())
	}
}

func TestNewTargetName_Valid_Japanese30Chars(t *testing.T) {
	name, err := NewTargetName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ" {
		t.Fatalf("expected value, got %q", name.String())
	}
}

func TestNewTargetName_Invalid_TooShort(t *testing.T) {
	_, err := NewTargetName("")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewTargetName_Invalid_TooLong(t *testing.T) {
	_, err := NewTargetName("1234567890123456789012345678901")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewTargetName_Invalid_TooLong_Japanese31Chars(t *testing.T) {
	_, err := NewTargetName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほま")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

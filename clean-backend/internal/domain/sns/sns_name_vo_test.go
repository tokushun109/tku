package sns

import "testing"

func TestNewSnsName_Valid(t *testing.T) {
	name, err := NewSnsName("Instagram")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "Instagram" {
		t.Fatalf("expected trimmed value, got %q", name.String())
	}
}

func TestNewSnsName_Valid_Japanese30Chars(t *testing.T) {
	name, err := NewSnsName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ" {
		t.Fatalf("expected value, got %q", name.String())
	}
}

func TestNewSnsName_Invalid_TooShort(t *testing.T) {
	_, err := NewSnsName("")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSnsName_Invalid_TooLong(t *testing.T) {
	_, err := NewSnsName("1234567890123456789012345678901")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSnsName_Invalid_TooLong_Japanese31Chars(t *testing.T) {
	_, err := NewSnsName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほま")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

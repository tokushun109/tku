package sales_site

import "testing"

func TestNewSalesSiteName_Valid(t *testing.T) {
	name, err := NewSalesSiteName("Creema")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "Creema" {
		t.Fatalf("expected trimmed value, got %q", name.String())
	}
}

func TestNewSalesSiteName_Valid_Japanese30Chars(t *testing.T) {
	name, err := NewSalesSiteName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ" {
		t.Fatalf("expected value, got %q", name.String())
	}
}

func TestNewSalesSiteName_Invalid_TooShort(t *testing.T) {
	_, err := NewSalesSiteName("")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSalesSiteName_Invalid_TooLong(t *testing.T) {
	_, err := NewSalesSiteName("1234567890123456789012345678901")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSalesSiteName_Invalid_TooLong_Japanese31Chars(t *testing.T) {
	_, err := NewSalesSiteName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほま")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

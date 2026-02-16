package skill_market

import "testing"

func TestNewSkillMarketName_Valid(t *testing.T) {
	name, err := NewSkillMarketName("minne")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "minne" {
		t.Fatalf("expected trimmed value, got %q", name.String())
	}
}

func TestNewSkillMarketName_Valid_Japanese30Chars(t *testing.T) {
	name, err := NewSkillMarketName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name.String() != "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ" {
		t.Fatalf("expected value, got %q", name.String())
	}
}

func TestNewSkillMarketName_Invalid_TooShort(t *testing.T) {
	_, err := NewSkillMarketName("")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSkillMarketName_Invalid_TooLong(t *testing.T) {
	_, err := NewSkillMarketName("1234567890123456789012345678901")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSkillMarketName_Invalid_TooLong_Japanese31Chars(t *testing.T) {
	_, err := NewSkillMarketName("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほま")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

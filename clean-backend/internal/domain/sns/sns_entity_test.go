package sns

import (
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

func TestNewSns_OK(t *testing.T) {
	s, err := New("Instagram", "https://www.instagram.com", "https://example.com/icon.png")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Name.String() != "Instagram" {
		t.Fatalf("expected name to be %q, got %q", "Instagram", s.Name.String())
	}
	if s.URL.String() != "https://www.instagram.com" {
		t.Fatalf("expected url to be %q, got %q", "https://www.instagram.com", s.URL.String())
	}
	if s.Icon != "https://example.com/icon.png" {
		t.Fatalf("expected icon to be %q, got %q", "https://example.com/icon.png", s.Icon)
	}
}

func TestNewSns_InvalidName(t *testing.T) {
	_, err := New("", "https://www.instagram.com", "")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestNewSns_InvalidURL(t *testing.T) {
	_, err := New("Instagram", "not-url", "")
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != primitive.ErrInvalidURL {
		t.Fatalf("expected primitive.ErrInvalidURL, got %v", err)
	}
}

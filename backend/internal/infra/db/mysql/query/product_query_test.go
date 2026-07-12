package query

import "testing"

func TestEscapeLikeKeyword(t *testing.T) {
	got := escapeLikeKeyword(`50%_off\sale`)
	want := `50\%\_off\\sale`
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestPlaceholders(t *testing.T) {
	got := placeholders(3)
	want := "?,?,?"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

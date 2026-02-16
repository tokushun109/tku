package tag

import "testing"

func TestNewTag_OK(t *testing.T) {
	tg, err := New("accessory")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tg.Name.String() != "accessory" {
		t.Fatalf("expected name to be trimmed value, got %q", tg.Name.String())
	}
}

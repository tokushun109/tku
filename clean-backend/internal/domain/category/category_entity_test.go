package category

import "testing"

func TestNewCategory_OK(t *testing.T) {
	c, err := New("accessory")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Name.String() != "accessory" {
		t.Fatalf("expected name to be trimmed value, got %q", c.Name.String())
	}
}

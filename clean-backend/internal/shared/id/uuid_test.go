package id

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerateUUID(t *testing.T) {
	u := GenerateUUID()
	if u == "" {
		t.Fatalf("expected non-empty uuid")
	}
	if _, err := uuid.Parse(u); err != nil {
		t.Fatalf("expected valid uuid, got error: %v", err)
	}
}

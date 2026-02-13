package id

import "testing"

type testUUID UUID

func TestNewUUID_And_ParseUUID(t *testing.T) {
	u := NewUUID()
	if u.String() == "" {
		t.Fatalf("expected non-empty uuid")
	}
	parsed, err := ParseUUID(u.String())
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}
	if parsed.String() != u.String() {
		t.Fatalf("expected parsed uuid to match")
	}
}

func TestParseUUID_Invalid(t *testing.T) {
	_, err := ParseUUID("not-a-uuid")
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestNewAs_And_ParseAs(t *testing.T) {
	u := NewAs[testUUID]()
	if string(u) == "" {
		t.Fatalf("expected non-empty uuid")
	}
	parsed, err := ParseAs[testUUID](string(u))
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}
	if string(parsed) != string(u) {
		t.Fatalf("expected parsed uuid to match")
	}
}

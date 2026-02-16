package primitive

import "testing"

func TestNewUUID_Valid(t *testing.T) {
	u, err := NewUUID("31637057-4c42-4d6c-b3ad-080b018a1844")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.String() != "31637057-4c42-4d6c-b3ad-080b018a1844" {
		t.Fatalf("unexpected uuid: %s", u.String())
	}
}

func TestNewUUID_Invalid(t *testing.T) {
	cases := []string{
		"",
		"not-a-uuid",
		"31637057-4c42-4d6c-b3ad-080b018a184",
		"31637057-4c42-4d6c-b3ad-080b018a18440",
		"ZZ637057-4c42-4d6c-b3ad-080b018a1844",
		"31637057-4c42-4d6c-b3ad-080b018a184g",
		"31637057-4c42-6d6c-b3ad-080b018a1844",
		"31637057-4c42-4d6c-7bad-080b018a1844",
	}
	for _, c := range cases {
		if _, err := NewUUID(c); err == nil {
			t.Fatalf("expected error for %q", c)
		}
	}
}

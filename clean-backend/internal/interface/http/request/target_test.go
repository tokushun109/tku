package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseListTargetQuery_All(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/target?mode=all", nil)

	q, err := ParseListTargetQuery(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q.Mode != TargetModeAll {
		t.Fatalf("expected %q, got %q", TargetModeAll, q.Mode)
	}
}

func TestParseListTargetQuery_Used(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/target?mode=used", nil)

	q, err := ParseListTargetQuery(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q.Mode != TargetModeUsed {
		t.Fatalf("expected %q, got %q", TargetModeUsed, q.Mode)
	}
}

func TestParseListTargetQuery_Invalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/target?mode=bad", nil)

	_, err := ParseListTargetQuery(req)
	if err == nil {
		t.Fatalf("expected error")
	}
}

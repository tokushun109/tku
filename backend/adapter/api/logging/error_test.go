package logging

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestErrorLog(t *testing.T) {
	log := &testutil.Logger{}
	req := httptest.NewRequest(http.MethodDelete, "/api/error", nil)

	NewError(log, req, http.StatusInternalServerError, nil).Log("error")

	if len(log.Errors) == 0 {
		t.Fatalf("expected error log")
	}
	if !strings.Contains(log.Errors[0], "status=500") {
		t.Fatalf("unexpected error log: %s", log.Errors[0])
	}
}

package logging

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestDebugLog(t *testing.T) {
	log := &testutil.Logger{}
	req := httptest.NewRequest(http.MethodGet, "/api/debug", nil)

	NewDebug(log, req).Log("debug")

	if len(log.Debugs) == 0 {
		t.Fatalf("expected debug log")
	}
	if !strings.Contains(log.Debugs[0], "path=/api/debug") {
		t.Fatalf("unexpected debug log: %s", log.Debugs[0])
	}
}

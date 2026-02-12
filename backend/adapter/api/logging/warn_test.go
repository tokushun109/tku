package logging

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestWarnLog(t *testing.T) {
	log := &testutil.Logger{}
	req := httptest.NewRequest(http.MethodPost, "/api/warn", nil)

	NewWarn(log, req, http.StatusBadRequest, nil).Log("warn")

	if len(log.Warns) == 0 {
		t.Fatalf("expected warn log")
	}
	if !strings.Contains(log.Warns[0], "path=/api/warn") {
		t.Fatalf("unexpected warn log: %s", log.Warns[0])
	}
}

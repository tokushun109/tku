package logging

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestInfoLog(t *testing.T) {
	log := &testutil.Logger{}
	req := httptest.NewRequest(http.MethodGet, "/api/info", nil)

	NewInfo(log, req, http.StatusOK).Log("ok")

	if len(log.Infos) == 0 {
		t.Fatalf("expected info log")
	}
	if !strings.Contains(log.Infos[0], "method=GET") {
		t.Fatalf("unexpected info log: %s", log.Infos[0])
	}
}

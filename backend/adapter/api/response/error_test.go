package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/testutil"
)

func TestLogAndSendError(t *testing.T) {
	log := &testutil.Logger{}
	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	LogAndSendError(rr, req, log, http.StatusInternalServerError, nil, "db connection error")

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	var er ErrorResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &er); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	if er.Error.Message != "db connection error" {
		t.Fatalf("unexpected error message: %s", er.Error.Message)
	}
	if len(log.Errors) == 0 {
		t.Fatalf("expected error log to be written")
	}
	if !strings.Contains(log.Errors[0], "method=GET") {
		t.Fatalf("expected log to contain method, got %s", log.Errors[0])
	}
}

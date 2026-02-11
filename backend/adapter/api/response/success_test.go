package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewSuccessSend(t *testing.T) {
	rr := httptest.NewRecorder()

	NewSuccess(http.StatusOK).Send(rr)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	var sr SuccessResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &sr); err != nil {
		t.Fatalf("failed to decode success response: %v", err)
	}
	if !sr.Success {
		t.Fatalf("expected success true")
	}
}

package response

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type errorResp struct {
	Message string `json:"message"`
}

type successResp struct {
	Success bool `json:"success"`
}

func TestWriteJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	WriteJSON(rr, http.StatusCreated, map[string]string{"k": "v"})

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected content-type application/json, got %q", ct)
	}
	var body map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if body["k"] != "v" {
		t.Fatalf("unexpected body: %+v", body)
	}
}

func TestWriteSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	WriteSuccess(rr)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
	var body successResp
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if !body.Success {
		t.Fatalf("expected success true")
	}
}

func TestWriteAppError_InvalidInput_NoLog(t *testing.T) {
	var buf bytes.Buffer
	prevOut := log.Writer()
	prevFlags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(prevOut)
		log.SetFlags(prevFlags)
	}()

	rr := httptest.NewRecorder()
	WriteAppError(rr, usecase.NewAppError(usecase.ErrInvalidInput))

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no log output, got %q", buf.String())
	}
	var body errorResp
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if body.Message != usecase.ErrInvalidInput.Error() {
		t.Fatalf("unexpected message: %q", body.Message)
	}
}

func TestWriteAppError_Internal_Logs(t *testing.T) {
	var buf bytes.Buffer
	prevOut := log.Writer()
	prevFlags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(prevOut)
		log.SetFlags(prevFlags)
	}()

	rr := httptest.NewRecorder()
	WriteAppError(rr, usecase.NewAppErrorWithMessage(usecase.ErrInternal, "db error"))

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rr.Code)
	}
	logStr := buf.String()
	if logStr == "" || !bytes.Contains([]byte(logStr), []byte("internal error")) {
		t.Fatalf("expected internal error log, got %q", logStr)
	}
	var body errorResp
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if body.Message != usecase.ErrInternal.Error() {
		t.Fatalf("unexpected message: %q", body.Message)
	}
}

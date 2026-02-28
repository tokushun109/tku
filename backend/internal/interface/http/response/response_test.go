package response

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tokushun109/tku/backend/internal/shared/logger"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type errorResp struct {
	Message string `json:"message"`
}

type successResp struct {
	Success bool `json:"success"`
}

func TestWriteJSON(t *testing.T) {
	t.Run("JSONを書き込むならJSONレスポンスを返す", func(t *testing.T) {

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
	})

}

func TestWriteSuccess(t *testing.T) {
	t.Run("成功レスポンスを書き込むなら200レスポンスを返す", func(t *testing.T) {

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
	})

}

func TestWriteAppError(t *testing.T) {
	t.Run("クライアントエラーならエラーログを出力しない", func(t *testing.T) {

		var stdoutBuf bytes.Buffer
		var stderrBuf bytes.Buffer
		logger.SetOutputs(&stdoutBuf, &stderrBuf)
		logger.SetFlags(0)
		defer logger.Reset()

		rr := httptest.NewRecorder()
		WriteAppError(rr, usecase.NewAppError(usecase.ErrInvalidInput))

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
		}
		if stdoutBuf.Len() != 0 || stderrBuf.Len() != 0 {
			t.Fatalf("expected no log output, got stdout=%q stderr=%q", stdoutBuf.String(), stderrBuf.String())
		}
		var body errorResp
		if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if body.Message != usecase.ErrInvalidInput.Error() {
			t.Fatalf("unexpected message: %q", body.Message)
		}
	})
	t.Run("内部エラーならエラーログを出力する", func(t *testing.T) {

		var stdoutBuf bytes.Buffer
		var stderrBuf bytes.Buffer
		logger.SetOutputs(&stdoutBuf, &stderrBuf)
		logger.SetFlags(0)
		defer logger.Reset()

		rr := httptest.NewRecorder()
		WriteAppError(rr, usecase.NewAppErrorWithMessage(usecase.ErrInternal, "db error"))

		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rr.Code)
		}
		if stdoutBuf.Len() != 0 {
			t.Fatalf("expected no stdout log, got %q", stdoutBuf.String())
		}
		logStr := stderrBuf.String()
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
	})
}

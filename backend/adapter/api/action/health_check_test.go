package action

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tokushun109/tku/backend/adapter/api/response"
	"github.com/tokushun109/tku/backend/adapter/repository"
	"github.com/tokushun109/tku/backend/internal/testutil"
)

type fakeSQLDB struct {
	db  *sql.DB
	err error
}

func (f fakeSQLDB) DB() (*sql.DB, error) {
	return f.db, f.err
}

func decodeError(t *testing.T, rr *httptest.ResponseRecorder) response.ErrorResponse {
	t.Helper()
	var er response.ErrorResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &er); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	return er
}

func decodeSuccess(t *testing.T, rr *httptest.ResponseRecorder) response.SuccessResponse {
	t.Helper()
	var sr response.SuccessResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &sr); err != nil {
		t.Fatalf("failed to decode success response: %v", err)
	}
	return sr
}

func TestHealthCheck_DBNil(t *testing.T) {
	var db repository.SQLDB
	log := &testutil.Logger{}
	act := NewHealthCheckAction(db, log)

	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	act.Execute(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	er := decodeError(t, rr)
	if er.Error.Message != "db connection error" {
		t.Fatalf("unexpected error message: %s", er.Error.Message)
	}
	if len(log.Errors) == 0 {
		t.Fatalf("expected error log to be written")
	}
}

func TestHealthCheck_DBError(t *testing.T) {
	db := fakeSQLDB{db: nil, err: errors.New("db error")}
	log := &testutil.Logger{}
	act := NewHealthCheckAction(db, log)

	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	act.Execute(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	er := decodeError(t, rr)
	if er.Error.Message != "db connection error" {
		t.Fatalf("unexpected error message: %s", er.Error.Message)
	}
	if len(log.Errors) == 0 {
		t.Fatalf("expected error log to be written")
	}
}

func TestHealthCheck_PingError(t *testing.T) {
	dbConn, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer dbConn.Close()

	mock.ExpectPing().WillReturnError(errors.New("ping error"))

	db := fakeSQLDB{db: dbConn, err: nil}
	log := &testutil.Logger{}
	act := NewHealthCheckAction(db, log)

	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	act.Execute(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	er := decodeError(t, rr)
	if er.Error.Message != "db ping error" {
		t.Fatalf("unexpected error message: %s", er.Error.Message)
	}
	if len(log.Errors) == 0 {
		t.Fatalf("expected error log to be written")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sqlmock expectations: %v", err)
	}
}

func TestHealthCheck_Success(t *testing.T) {
	dbConn, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer dbConn.Close()

	mock.ExpectPing()

	db := fakeSQLDB{db: dbConn, err: nil}
	log := &testutil.Logger{}
	act := NewHealthCheckAction(db, log)

	req := httptest.NewRequest(http.MethodGet, "/api/health_check", nil)
	rr := httptest.NewRecorder()

	act.Execute(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	sr := decodeSuccess(t, rr)
	if !sr.Success {
		t.Fatalf("expected success true")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sqlmock expectations: %v", err)
	}
}

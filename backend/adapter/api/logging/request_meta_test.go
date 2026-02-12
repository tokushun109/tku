package logging

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequestMeta(t *testing.T) {
	tests := []struct {
		name       string
		req        *http.Request
		wantMethod string
		wantPath   string
	}{
		{
			name:       "nil_request",
			req:        nil,
			wantMethod: "",
			wantPath:   "",
		},
		{
			name:       "nil_url",
			req:        &http.Request{Method: http.MethodGet},
			wantMethod: http.MethodGet,
			wantPath:   "",
		},
		{
			name:       "method_and_path",
			req:        httptest.NewRequest(http.MethodPost, "/api/health_check?ok=1", nil),
			wantMethod: http.MethodPost,
			wantPath:   "/api/health_check?ok=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMethod, gotPath := getRequestMeta(tt.req)
			if gotMethod != tt.wantMethod || gotPath != tt.wantPath {
				t.Fatalf("getRequestMeta() = (%q, %q), want (%q, %q)", gotMethod, gotPath, tt.wantMethod, tt.wantPath)
			}
		})
	}
}

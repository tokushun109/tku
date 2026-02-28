package middleware

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	"github.com/tokushun109/tku/backend/internal/usecase"
)

type OriginMiddleware struct {
	allowedOrigin *url.URL
	enabled       bool
}

func NewOriginMiddleware(clientURL string) (*OriginMiddleware, error) {
	if clientURL == "" {
		return &OriginMiddleware{enabled: false}, nil
	}

	u, err := url.Parse(clientURL)
	if err != nil {
		return nil, err
	}

	return &OriginMiddleware{
		allowedOrigin: u,
		enabled:       true,
	}, nil
}

func (m *OriginMiddleware) RequireTrustedOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.enabled || isSafeMethod(r.Method) {
			next.ServeHTTP(w, r)
			return
		}

		requestOrigin := r.Header.Get("Origin")
		if requestOrigin == "" {
			requestOrigin = originFromReferer(r.Header.Get("Referer"))
		}

		if requestOrigin == "" || !m.isAllowedOrigin(requestOrigin) {
			response.WriteAppError(w, usecase.NewAppError(usecase.ErrForbidden))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isSafeMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true
	default:
		return false
	}
}

func originFromReferer(referer string) string {
	if referer == "" {
		return ""
	}

	u, err := url.Parse(referer)
	if err != nil {
		return ""
	}

	if u.Scheme == "" || u.Host == "" {
		return ""
	}

	return u.Scheme + "://" + u.Host
}

func (m *OriginMiddleware) isAllowedOrigin(raw string) bool {
	if m.allowedOrigin == nil {
		return false
	}

	u, err := url.Parse(raw)
	if err != nil {
		return false
	}

	return strings.EqualFold(u.Scheme, m.allowedOrigin.Scheme) && strings.EqualFold(u.Host, m.allowedOrigin.Host)
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitMiddleware(t *testing.T) {
	t.Run("同一IPで制限回数を超えたら429を返す", func(t *testing.T) {
		mw := NewRateLimitMiddleware(2)

		called := 0
		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called++
			w.WriteHeader(http.StatusOK)
		}))

		for i := 0; i < 3; i++ {
			req := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
			req.RemoteAddr = "203.0.113.1:12345"
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)

			if i < 2 && rr.Code != http.StatusOK {
				t.Fatalf("expected 200 on try %d, got %d", i+1, rr.Code)
			}
			if i == 2 && rr.Code != http.StatusTooManyRequests {
				t.Fatalf("expected 429 on try %d, got %d", i+1, rr.Code)
			}
		}

		if called != 2 {
			t.Fatalf("expected next called 2 times, got %d", called)
		}
	})

	t.Run("異なるIPなら別枠でカウントする", func(t *testing.T) {
		mw := NewRateLimitMiddleware(1)

		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req1 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req1.RemoteAddr = "203.0.113.1:11111"
		rr1 := httptest.NewRecorder()
		h.ServeHTTP(rr1, req1)
		if rr1.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr1.Code)
		}

		req2 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req2.RemoteAddr = "203.0.113.2:22222"
		rr2 := httptest.NewRecorder()
		h.ServeHTTP(rr2, req2)
		if rr2.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr2.Code)
		}
	})

	t.Run("X-Forwarded-Forの先頭IPをキーとして使う", func(t *testing.T) {
		mw := NewRateLimitMiddleware(1)

		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req1 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req1.RemoteAddr = "203.0.113.10:11111"
		req1.Header.Set("X-Forwarded-For", "198.51.100.1, 198.51.100.2")
		rr1 := httptest.NewRecorder()
		h.ServeHTTP(rr1, req1)
		if rr1.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr1.Code)
		}

		req2 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req2.RemoteAddr = "203.0.113.99:99999"
		req2.Header.Set("X-Forwarded-For", "198.51.100.1, 203.0.113.10")
		rr2 := httptest.NewRecorder()
		h.ServeHTTP(rr2, req2)
		if rr2.Code != http.StatusTooManyRequests {
			t.Fatalf("expected 429, got %d", rr2.Code)
		}
	})

	t.Run("ウィンドウ経過後は再度許可する", func(t *testing.T) {
		mw := newRateLimitMiddlewareWithOptions(rateLimitOptions{
			limitPerWindow: 1,
			window:         50 * time.Millisecond,
		})

		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req1 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req1.RemoteAddr = "203.0.113.3:33333"
		rr1 := httptest.NewRecorder()
		h.ServeHTTP(rr1, req1)
		if rr1.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr1.Code)
		}

		req2 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req2.RemoteAddr = "203.0.113.3:33333"
		rr2 := httptest.NewRecorder()
		h.ServeHTTP(rr2, req2)
		if rr2.Code != http.StatusTooManyRequests {
			t.Fatalf("expected 429, got %d", rr2.Code)
		}

		time.Sleep(60 * time.Millisecond)

		req3 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req3.RemoteAddr = "203.0.113.3:33333"
		rr3 := httptest.NewRecorder()
		h.ServeHTTP(rr3, req3)
		if rr3.Code != http.StatusOK {
			t.Fatalf("expected 200 after window reset, got %d", rr3.Code)
		}
	})

	t.Run("制限超過時にRetry-Afterヘッダーを返す", func(t *testing.T) {
		mw := NewRateLimitMiddleware(1)

		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req1 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req1.RemoteAddr = "203.0.113.4:44444"
		rr1 := httptest.NewRecorder()
		h.ServeHTTP(rr1, req1)
		if rr1.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr1.Code)
		}

		req2 := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req2.RemoteAddr = "203.0.113.4:44444"
		rr2 := httptest.NewRecorder()
		h.ServeHTTP(rr2, req2)
		if rr2.Code != http.StatusTooManyRequests {
			t.Fatalf("expected 429, got %d", rr2.Code)
		}
		if rr2.Header().Get("Retry-After") == "" {
			t.Fatalf("expected Retry-After header")
		}
	})
}

func TestNewRateLimitMiddleware(t *testing.T) {
	t.Run("引数未指定時はデフォルト上限を使う", func(t *testing.T) {
		mw := NewRateLimitMiddleware()
		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		for i := 0; i < defaultRateLimitPerMinute+1; i++ {
			req := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
			req.RemoteAddr = "203.0.113.10:12345"
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)

			if i < defaultRateLimitPerMinute && rr.Code != http.StatusOK {
				t.Fatalf("expected 200 on try %d, got %d", i+1, rr.Code)
			}
			if i == defaultRateLimitPerMinute && rr.Code != http.StatusTooManyRequests {
				t.Fatalf("expected 429 on try %d, got %d", i+1, rr.Code)
			}
		}
	})

	t.Run("引数指定時は指定上限を使う", func(t *testing.T) {
		mw := NewRateLimitMiddleware(2)
		h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		for i := 0; i < 3; i++ {
			req := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
			req.RemoteAddr = "203.0.113.11:12345"
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)

			if i < 2 && rr.Code != http.StatusOK {
				t.Fatalf("expected 200 on try %d, got %d", i+1, rr.Code)
			}
			if i == 2 && rr.Code != http.StatusTooManyRequests {
				t.Fatalf("expected 429 on try %d, got %d", i+1, rr.Code)
			}
		}
	})
}

func TestClientIPKey(t *testing.T) {
	t.Run("X-Forwarded-Forがあるなら先頭IPを返す", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req.Header.Set("X-Forwarded-For", "198.51.100.1, 198.51.100.2")

		got := ClientIPKey(req)
		if got != "198.51.100.1" {
			t.Fatalf("expected 198.51.100.1, got %s", got)
		}
	})

	t.Run("X-Forwarded-ForがないならRemoteAddrからIPを返す", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/contact", nil)
		req.RemoteAddr = "203.0.113.5:55555"

		got := ClientIPKey(req)
		if got != "203.0.113.5" {
			t.Fatalf("expected 203.0.113.5, got %s", got)
		}
	})
}

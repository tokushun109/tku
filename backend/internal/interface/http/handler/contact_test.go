package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain "github.com/tokushun109/tku/backend/internal/domain/contact"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
)

type stubContactUC struct {
	listRes []*domain.Contact
	listErr error

	createErr error
}

func (s *stubContactUC) List(ctx context.Context) ([]*domain.Contact, error) {
	return s.listRes, s.listErr
}

func (s *stubContactUC) Create(ctx context.Context, name string, company string, phoneNumber string, email string, content string) error {
	return s.createErr
}

type contactResp struct {
	ID          uint      `json:"id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Company     *string   `json:"company"`
	PhoneNumber *string   `json:"phoneNumber"`
	Email       string    `json:"email"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}

func TestContactGet(t *testing.T) {
	t.Run("有効な入力を渡したときお問い合わせ一覧の取得に成功する", func(t *testing.T) {
		contact, err := domain.Rebuild(
			1,
			"11111111-1111-4111-8111-111111111111",
			"山田太郎",
			"株式会社サンプル",
			"09012345678",
			"test@example.com",
			"お問い合わせ内容",
			time.Date(2026, 2, 19, 10, 0, 0, 0, time.UTC),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := NewContactHandler(&stubContactUC{
			listRes: []*domain.Contact{contact},
		})

		req := httptest.NewRequest(http.MethodGet, "/api/contact", nil)
		rr := httptest.NewRecorder()

		h.List(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var res []contactResp
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1, got %d", len(res))
		}
		if res[0].ID != 1 {
			t.Fatalf("unexpected id: %d", res[0].ID)
		}
		if res[0].Name != "山田太郎" {
			t.Fatalf("unexpected response: %+v", res[0])
		}
	})
}

func TestContactPost(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		h := NewContactHandler(&stubContactUC{})

		req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewBufferString(`{invalid}`))
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})

	t.Run("有効な入力を渡したときお問い合わせ作成に成功する", func(t *testing.T) {
		h := NewContactHandler(&stubContactUC{})

		req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewBufferString(`{"name":"山田太郎","email":"test@example.com","content":"お問い合わせ内容"}`))
		rr := httptest.NewRecorder()

		h.Create(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var res response.SuccessResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if !res.Success {
			t.Fatalf("expected success true")
		}
	})
}

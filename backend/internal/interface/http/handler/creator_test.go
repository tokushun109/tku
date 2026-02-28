package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	domain "github.com/tokushun109/tku/backend/internal/domain/creator"
	"github.com/tokushun109/tku/backend/internal/interface/http/response"
	usecaseCreator "github.com/tokushun109/tku/backend/internal/usecase/creator"
)

type stubCreatorUC struct {
	getRes *usecaseCreator.CreatorDetail
	getErr error

	updateErr error
	updateReq struct {
		name         string
		introduction string
	}

	updateLogoErr error
	updateLogoReq []byte

	getLogoBlobRes *usecaseCreator.LogoBlob
	getLogoBlobErr error
	getLogoBlobReq string
}

func (s *stubCreatorUC) Get(ctx context.Context) (*usecaseCreator.CreatorDetail, error) {
	return s.getRes, s.getErr
}

func (s *stubCreatorUC) Update(ctx context.Context, name string, introduction string) error {
	s.updateReq.name = name
	s.updateReq.introduction = introduction
	return s.updateErr
}

func (s *stubCreatorUC) UpdateLogo(ctx context.Context, logoBytes []byte) error {
	s.updateLogoReq = logoBytes
	return s.updateLogoErr
}

func (s *stubCreatorUC) GetLogoBlob(ctx context.Context, requestLogoFile string) (*usecaseCreator.LogoBlob, error) {
	s.getLogoBlobReq = requestLogoFile
	return s.getLogoBlobRes, s.getLogoBlobErr
}

func TestCreatorGet(t *testing.T) {
	t.Run("有効な入力を渡したとき製作者情報の取得に成功する", func(t *testing.T) {
		creator, err := domain.Rebuild(1, "とこりり", "ハンドメイド作品を制作", "", "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		h := NewCreatorHandler(&stubCreatorUC{
			getRes: &usecaseCreator.CreatorDetail{
				Creator: creator,
				APIPath: "http://localhost:8081/api/creator/logo/logo.png/blob",
			},
		})

		req := httptest.NewRequest(http.MethodGet, "/api/creator", nil)
		rr := httptest.NewRecorder()

		h.Get(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}

		var res response.CreatorResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		if res.Name != "とこりり" {
			t.Fatalf("unexpected response: %+v", res)
		}
	})
}

func TestCreatorPut(t *testing.T) {
	t.Run("JSONが不正なときバリデーションエラーで失敗する", func(t *testing.T) {
		h := NewCreatorHandler(&stubCreatorUC{})
		req := httptest.NewRequest(http.MethodPut, "/api/creator", bytes.NewBufferString(`{invalid}`))
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})

	t.Run("有効な入力を渡したとき製作者情報の更新に成功する", func(t *testing.T) {
		uc := &stubCreatorUC{}
		h := NewCreatorHandler(uc)
		req := httptest.NewRequest(http.MethodPut, "/api/creator", bytes.NewBufferString(`{"name":"とこりり","introduction":"紹介文"}`))
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if uc.updateReq.name != "とこりり" || uc.updateReq.introduction != "紹介文" {
			t.Fatalf("unexpected update args: %+v", uc.updateReq)
		}
	})

	t.Run("紹介文が空文字のときも更新に成功し空文字がusecaseに渡る", func(t *testing.T) {
		uc := &stubCreatorUC{}
		h := NewCreatorHandler(uc)
		req := httptest.NewRequest(http.MethodPut, "/api/creator", bytes.NewBufferString(`{"name":"とこりり","introduction":""}`))
		rr := httptest.NewRecorder()

		h.Update(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if uc.updateReq.name != "とこりり" || uc.updateReq.introduction != "" {
			t.Fatalf("unexpected update args: %+v", uc.updateReq)
		}
	})
}

func TestCreatorLogoPut(t *testing.T) {
	t.Run("multipartにロゴファイルがないならバリデーションエラーで失敗する", func(t *testing.T) {
		h := NewCreatorHandler(&stubCreatorUC{})
		req := httptest.NewRequest(http.MethodPut, "/api/creator/logo", nil)
		rr := httptest.NewRecorder()

		h.UpdateLogo(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
	})

	t.Run("有効な画像ファイルを渡したときロゴ更新に成功する", func(t *testing.T) {
		uc := &stubCreatorUC{}
		h := NewCreatorHandler(uc)

		var body bytes.Buffer
		writer := multipart.NewWriter(&body)
		part, err := writer.CreateFormFile("logo", "logo.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, _ = part.Write([]byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'})
		_ = writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/api/creator/logo", &body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		h.UpdateLogo(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if len(uc.updateLogoReq) == 0 {
			t.Fatalf("expected non-empty logo bytes")
		}
	})

	t.Run("20MBを超える画像ファイルを渡したときバリデーションエラーで失敗する", func(t *testing.T) {
		uc := &stubCreatorUC{}
		h := NewCreatorHandler(uc)

		var body bytes.Buffer
		writer := multipart.NewWriter(&body)
		part, err := writer.CreateFormFile("logo", "logo.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, _ = part.Write(bytes.Repeat([]byte("a"), maxCreatorLogoSize+1))
		_ = writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/api/creator/logo", &body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		h.UpdateLogo(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", rr.Code)
		}
		if len(uc.updateLogoReq) != 0 {
			t.Fatalf("usecase should not be called when payload is too large")
		}
	})
}

func TestCreatorLogoBlobGet(t *testing.T) {
	t.Run("保存済みのロゴファイル名を指定したときblob取得に成功する", func(t *testing.T) {
		h := NewCreatorHandler(&stubCreatorUC{
			getLogoBlobRes: &usecaseCreator.LogoBlob{
				ContentType: "image/png",
				Body:        io.NopCloser(bytes.NewReader([]byte("binary"))),
			},
		})

		req := httptest.NewRequest(http.MethodGet, "/api/creator/logo/logo.png/blob", nil)
		req = mux.SetURLVars(req, map[string]string{"logo_file": "logo.png"})
		rr := httptest.NewRecorder()

		h.GetLogoBlob(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if rr.Header().Get("Content-Type") != "image/png" {
			t.Fatalf("unexpected content-type: %s", rr.Header().Get("Content-Type"))
		}
		if rr.Body.String() != "binary" {
			t.Fatalf("unexpected response body: %s", rr.Body.String())
		}
	})
}

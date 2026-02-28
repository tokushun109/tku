package creema

import (
	"bytes"
	"errors"
	"net/url"
	"testing"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/usecase"
)

func TestValidateProductPageURL(t *testing.T) {
	t.Run("有効なCreema商品URLを受け入れる", func(t *testing.T) {
		u, err := validateProductPageURL("https://www.creema.jp/items/123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.String() != "https://www.creema.jp/items/123" {
			t.Fatalf("unexpected url: %s", u.String())
		}
	})

	t.Run("creemaを含むだけの不正なホストを拒否する", func(t *testing.T) {
		_, err := validateProductPageURL("https://creema.attacker.com/items/123")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("https以外のスキームを拒否する", func(t *testing.T) {
		_, err := validateProductPageURL("http://www.creema.jp/items/123")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("URL形式が不正なときはバリデーションエラーを返す", func(t *testing.T) {
		_, err := validateProductPageURL("www.creema.jp/items/123")
		if err == nil || !errors.Is(err, primitive.ErrInvalidURL) {
			t.Fatalf("expected ErrInvalidURL, got %v", err)
		}
	})
}

func TestResolveImageURL(t *testing.T) {
	baseURL, err := url.Parse("https://www.creema.jp/items/123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("許可された画像ホストのURLを受け入れる", func(t *testing.T) {
		resolved, err := resolveImageURL(baseURL, "https://c.p02.c4a.im/images/item/example.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resolved != "https://c.p02.c4a.im/images/item/example.png" {
			t.Fatalf("unexpected resolved url: %s", resolved)
		}
	})

	t.Run("プロトコル相対URLを解決して受け入れる", func(t *testing.T) {
		resolved, err := resolveImageURL(baseURL, "//c.p02.c4a.im/images/item/example.png")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resolved != "https://c.p02.c4a.im/images/item/example.png" {
			t.Fatalf("unexpected resolved url: %s", resolved)
		}
	})

	t.Run("許可されていない画像ホストを拒否する", func(t *testing.T) {
		_, err := resolveImageURL(baseURL, "https://creema.attacker.com/images/item/example.png")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("https以外の画像URLを拒否する", func(t *testing.T) {
		_, err := resolveImageURL(baseURL, "http://c.p02.c4a.im/images/item/example.png")
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
}

func TestReadResponseBodyWithLimit(t *testing.T) {
	t.Run("上限内の本文を読み込める", func(t *testing.T) {
		body, err := readResponseBodyWithLimit(bytes.NewReader([]byte("ok")), 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(body) != "ok" {
			t.Fatalf("unexpected body: %s", string(body))
		}
	})

	t.Run("空の本文は入力エラーを返す", func(t *testing.T) {
		_, err := readResponseBodyWithLimit(bytes.NewReader(nil), 2)
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("上限を超える本文は入力エラーを返す", func(t *testing.T) {
		_, err := readResponseBodyWithLimit(bytes.NewReader(bytes.Repeat([]byte("a"), 4)), 3)
		if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})
}

func TestBuildImageName(t *testing.T) {
	t.Run("URLパスにファイル名があればその名前を返す", func(t *testing.T) {
		name := buildImageName("https://c.p02.c4a.im/images/item/example.png")
		if name != "example.png" {
			t.Fatalf("unexpected name: %s", name)
		}
	})

	t.Run("ファイル名を抽出できないときは空文字を返す", func(t *testing.T) {
		name := buildImageName("https://c.p02.c4a.im/")
		if name != "" {
			t.Fatalf("expected empty name, got %s", name)
		}
	})
}

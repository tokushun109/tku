package cursor

import (
	"errors"
	"testing"
	"time"
)

func TestEncodeDecode(t *testing.T) {
	t.Run("EncodeしたcursorをDecodeできる", func(t *testing.T) {
		now := time.Date(2026, time.March, 2, 12, 34, 56, 789, time.UTC)

		token, err := Encode(now, 42)
		if err != nil {
			t.Fatalf("unexpected encode error: %v", err)
		}

		value, err := Decode(token)
		if err != nil {
			t.Fatalf("unexpected decode error: %v", err)
		}
		if !value.CreatedAt.Equal(now) {
			t.Fatalf("unexpected createdAt: %v", value.CreatedAt)
		}
		if value.ID != 42 {
			t.Fatalf("unexpected id: %d", value.ID)
		}
	})

	t.Run("不正なcursorはErrInvalidを返す", func(t *testing.T) {
		_, err := Decode("invalid")
		if !errors.Is(err, ErrInvalid) {
			t.Fatalf("expected ErrInvalid, got %v", err)
		}
	})
}

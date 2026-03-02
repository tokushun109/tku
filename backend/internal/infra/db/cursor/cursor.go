package cursor

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
)

type CreatedAtID struct {
	CreatedAt time.Time `json:"-"`
	ID        uint      `json:"-"`
}

type payload struct {
	CreatedAt string `json:"c"`
	ID        uint   `json:"i"`
}

var ErrInvalid = errors.New("invalid cursor")

func Encode(createdAt time.Time, id uint) (string, error) {
	encodedPayload, err := json.Marshal(payload{
		CreatedAt: createdAt.Format(time.RFC3339Nano),
		ID:        id,
	})
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(encodedPayload), nil
}

func Decode(raw string) (CreatedAtID, error) {
	decodedCursor, err := base64.RawURLEncoding.DecodeString(raw)
	if err != nil {
		return CreatedAtID{}, ErrInvalid
	}

	var value payload
	if err := json.Unmarshal(decodedCursor, &value); err != nil {
		return CreatedAtID{}, ErrInvalid
	}
	if value.CreatedAt == "" || value.ID == 0 {
		return CreatedAtID{}, ErrInvalid
	}

	createdAt, err := time.Parse(time.RFC3339Nano, value.CreatedAt)
	if err != nil {
		return CreatedAtID{}, ErrInvalid
	}

	return CreatedAtID{
		CreatedAt: createdAt,
		ID:        value.ID,
	}, nil
}

package id

import "github.com/google/uuid"

type UUID string

func NewUUID() UUID {
	return UUID(uuid.NewString())
}

func ParseUUID(s string) (UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}
	return UUID(parsed.String()), nil
}

func (u UUID) String() string {
	return string(u)
}

func NewAs[T ~string]() T {
	return T(NewUUID())
}

func ParseAs[T ~string](s string) (T, error) {
	u, err := ParseUUID(s)
	if err != nil {
		var zero T
		return zero, err
	}
	return T(u), nil
}

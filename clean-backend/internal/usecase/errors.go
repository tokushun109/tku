package usecase

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternal     = errors.New("internal error")
)

type AppError struct {
	kind error
	msg  string
}

func (e *AppError) Error() string {
	return e.msg
}

func (e *AppError) Unwrap() error {
	return e.kind
}

func NewAppError(kind error) error {
	return &AppError{kind: kind, msg: kind.Error()}
}

func NewAppErrorWithMessage(kind error, msg string) error {
	if msg == "" {
		return &AppError{kind: kind, msg: kind.Error()}
	}
	return &AppError{kind: kind, msg: msg}
}

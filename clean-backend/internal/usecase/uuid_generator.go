package usecase

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type UUIDGenerator interface {
	New() (primitive.UUID, error)
}

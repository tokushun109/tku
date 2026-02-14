package uuid

import (
	"github.com/google/uuid"

	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) New() (primitive.UUID, error) {
	u := uuid.NewString()
	return primitive.NewUUID(u)
}

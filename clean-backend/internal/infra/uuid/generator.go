package uuid

import (
	"github.com/tokushun109/tku/clean-backend/internal/domain/primitive"
	"github.com/tokushun109/tku/clean-backend/internal/shared/id"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) New() (primitive.UUID, error) {
	return primitive.NewUUID(id.GenerateUUID())
}

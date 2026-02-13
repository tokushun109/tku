package category

import "github.com/tokushun109/tku/clean-backend/internal/shared/id"

type CategoryUUID id.UUID

func NewCategoryUUID() CategoryUUID {
	return id.NewAs[CategoryUUID]()
}

func ParseCategoryUUID(s string) (CategoryUUID, error) {
	return id.ParseAs[CategoryUUID](s)
}

func (u CategoryUUID) String() string {
	return string(u)
}

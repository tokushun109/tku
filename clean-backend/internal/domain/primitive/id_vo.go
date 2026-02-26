package primitive

import (
	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
	"strconv"
)

type ID uint

var _ domainVO.ValueObject[uint] = ID(0)

func NewID(v uint) (ID, error) {
	if v == 0 {
		return 0, ErrInvalidID
	}
	return ID(v), nil
}

func NewOptionalID(v *uint) (*ID, error) {
	if v == nil {
		return nil, nil
	}

	parsed, err := NewID(*v)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func (id ID) Value() uint {
	return uint(id)
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id.Value()), 10)
}

func (id ID) IsDefined() bool {
	return id != 0
}

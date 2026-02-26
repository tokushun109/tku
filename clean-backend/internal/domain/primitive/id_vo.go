package primitive

import (
	"strconv"

	domainVO "github.com/tokushun109/tku/clean-backend/internal/domain/vo"
)

type ID uint

var _ domainVO.ValueObject[uint] = ID(0)

func NewID(v uint) (ID, error) {
	if v == 0 {
		return 0, ErrInvalidID
	}
	return ID(v), nil
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

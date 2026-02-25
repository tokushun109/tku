package primitive

type ID uint

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

func (id ID) Uint() uint {
	return uint(id)
}

func (id ID) IsDefined() bool {
	return id != 0
}

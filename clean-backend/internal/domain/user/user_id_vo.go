package user

type UserID uint

func NewUserID(v uint) UserID {
	return UserID(v)
}

func (id UserID) Uint() uint {
	return uint(id)
}

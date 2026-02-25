package user

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

type User struct {
	id           uint
	uuid         primitive.UUID
	name         UserName
	email        primitive.Email
	passwordHash UserPasswordHash
	isAdmin      bool
}

func New(rawUUID string, name string, email string, passwordHash string, isAdmin bool) (*User, error) {
	user, err := newWithValidatedValues(rawUUID, name, email, passwordHash, isAdmin)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Rebuild(
	id uint,
	rawUUID string,
	name string,
	email string,
	passwordHash string,
	isAdmin bool,
) (*User, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	user, err := newWithValidatedValues(rawUUID, name, email, passwordHash, isAdmin)
	if err != nil {
		return nil, err
	}
	user.id = id
	return user, nil
}

func newWithValidatedValues(rawUUID string, name string, email string, passwordHash string, isAdmin bool) (*User, error) {
	uuid, err := primitive.NewUUID(rawUUID)
	if err != nil {
		return nil, err
	}
	userName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := primitive.NewEmail(email)
	if err != nil {
		return nil, err
	}
	userPasswordHash, err := NewUserPasswordHash(passwordHash)
	if err != nil {
		return nil, err
	}

	return &User{
		uuid:         uuid,
		name:         userName,
		email:        userEmail,
		passwordHash: userPasswordHash,
		isAdmin:      isAdmin,
	}, nil
}

func (u *User) ID() uint {
	return u.id
}

func (u *User) UUID() primitive.UUID {
	return u.uuid
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) Email() primitive.Email {
	return u.email
}

func (u *User) PasswordHash() UserPasswordHash {
	return u.passwordHash
}

func (u *User) IsAdmin() bool {
	return u.isAdmin
}

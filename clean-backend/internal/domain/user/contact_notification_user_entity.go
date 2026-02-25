package user

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

// お問い合わせ通知の送信先ユーザー情報
type ContactNotificationUser struct {
	id    uint
	name  UserName
	email primitive.Email
}

func NewContactNotificationUser(name string, email string) (*ContactNotificationUser, error) {
	userName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := primitive.NewEmail(email)
	if err != nil {
		return nil, err
	}
	return &ContactNotificationUser{
		name:  userName,
		email: userEmail,
	}, nil
}

func RebuildContactNotificationUser(id uint, name string, email string) (*ContactNotificationUser, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	user, err := NewContactNotificationUser(name, email)
	if err != nil {
		return nil, err
	}
	user.id = id
	return user, nil
}

func (u *ContactNotificationUser) ID() uint {
	return u.id
}

func (u *ContactNotificationUser) Name() UserName {
	return u.name
}

func (u *ContactNotificationUser) Email() primitive.Email {
	return u.email
}

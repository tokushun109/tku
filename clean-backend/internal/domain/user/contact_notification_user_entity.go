package user

import "github.com/tokushun109/tku/clean-backend/internal/domain/primitive"

// お問い合わせ通知の送信先ユーザー情報
type ContactNotificationUser struct {
	Name  UserName
	Email primitive.Email
}

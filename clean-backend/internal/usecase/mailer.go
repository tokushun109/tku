package usecase

import "context"

type Mailer interface {
	Send(ctx context.Context, message *MailMessage) error
}

type MailMessage struct {
	To       []MailAddress
	Subject  string
	TextBody string
	HTMLBody string
}

type MailAddress struct {
	Name  string
	Email string
}

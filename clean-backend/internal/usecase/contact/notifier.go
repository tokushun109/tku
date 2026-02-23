package contact

import (
	"bytes"
	"context"
	"embed"
	htmltmpl "html/template"
	"io"
	"log"
	"strings"
	texttmpl "text/template"

	domainContact "github.com/tokushun109/tku/clean-backend/internal/domain/contact"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
	usecase "github.com/tokushun109/tku/clean-backend/internal/usecase"
)

type Notifier interface {
	NotifyContactCreated(ctx context.Context, contact *domainContact.Contact)
}

const (
	defaultSupportEmail = "no-reply@tocoriri.com"
	autoReplySubject    = "【とこりり】お問い合わせを受け付けました"
	adminMailSubject    = "【とこりり】お問い合わせが届きました"
)

var (
	//go:embed templates/contact/*.txt templates/contact/*.html
	contactTemplateFS embed.FS

	autoReplyTextTemplate = texttmpl.Must(texttmpl.ParseFS(contactTemplateFS, "templates/contact/auto_reply.txt"))
	autoReplyHTMLTemplate = htmltmpl.Must(htmltmpl.ParseFS(contactTemplateFS, "templates/contact/auto_reply.html"))
	adminTextTemplate     = texttmpl.Must(texttmpl.ParseFS(contactTemplateFS, "templates/contact/admin.txt"))
	adminHTMLTemplate     = htmltmpl.Must(htmltmpl.ParseFS(contactTemplateFS, "templates/contact/admin.html"))
)

type notifier struct {
	mailer       usecase.Mailer
	userRepo     domainUser.Repository
	supportEmail string
}

type mailTemplateData struct {
	Title        string
	Name         string
	Company      string
	PhoneNumber  string
	Email        string
	Content      string
	SupportEmail string
}

type templateExecutor interface {
	Execute(wr io.Writer, data any) error
}

func NewContactNotifier(mailer usecase.Mailer, userRepo domainUser.Repository, supportEmail string) Notifier {
	trimmedSupportEmail := strings.TrimSpace(supportEmail)
	if trimmedSupportEmail == "" {
		trimmedSupportEmail = defaultSupportEmail
	}

	return &notifier{
		mailer:       mailer,
		userRepo:     userRepo,
		supportEmail: trimmedSupportEmail,
	}
}

func (n *notifier) NotifyContactCreated(ctx context.Context, contact *domainContact.Contact) {
	if contact == nil {
		return
	}

	if err := n.sendAutoReply(ctx, contact); err != nil {
		log.Printf("[WARN] contact notifier auto-reply failed: %v", err)
	}

	if n.userRepo == nil {
		log.Printf("[WARN] contact notifier skipped admin mail: user repository is nil")
		return
	}

	notificationUsers, err := n.userRepo.FindContactNotificationUsers(ctx)
	if err != nil {
		log.Printf("[WARN] contact notifier failed to load admin contacts: %v", err)
		return
	}

	recipients := make([]usecase.MailAddress, 0, len(notificationUsers))
	for _, user := range notificationUsers {
		if user == nil {
			continue
		}
		recipients = append(recipients, usecase.MailAddress{
			Name:  user.Name.String(),
			Email: user.Email.String(),
		})
	}
	if len(recipients) == 0 {
		log.Printf("[WARN] contact notifier skipped admin mail: no admin recipient")
		return
	}

	if err := n.sendAdminNotification(ctx, contact, recipients); err != nil {
		log.Printf("[WARN] contact notifier admin mail failed: %v", err)
	}
}

func (n *notifier) sendAutoReply(ctx context.Context, contact *domainContact.Contact) error {
	data := n.newMailTemplateData(autoReplySubject, contact)
	textBody, err := executeTemplate(autoReplyTextTemplate, data)
	if err != nil {
		return err
	}
	htmlBody, err := executeTemplate(autoReplyHTMLTemplate, data)
	if err != nil {
		return err
	}

	return n.mailer.Send(ctx, &usecase.MailMessage{
		To: []usecase.MailAddress{
			{
				Name:  contact.Name.String(),
				Email: contact.Email.String(),
			},
		},
		Subject:  autoReplySubject,
		TextBody: textBody,
		HTMLBody: htmlBody,
	})
}

func (n *notifier) sendAdminNotification(ctx context.Context, contact *domainContact.Contact, recipients []usecase.MailAddress) error {
	data := n.newMailTemplateData(adminMailSubject, contact)
	textBody, err := executeTemplate(adminTextTemplate, data)
	if err != nil {
		return err
	}
	htmlBody, err := executeTemplate(adminHTMLTemplate, data)
	if err != nil {
		return err
	}

	return n.mailer.Send(ctx, &usecase.MailMessage{
		To:       recipients,
		Subject:  adminMailSubject,
		TextBody: textBody,
		HTMLBody: htmlBody,
	})
}

func (n *notifier) newMailTemplateData(title string, contact *domainContact.Contact) mailTemplateData {
	return mailTemplateData{
		Title:        title,
		Name:         contact.Name.String(),
		Company:      optional.ToTrimmedStringOrEmpty(contact.Company),
		PhoneNumber:  optional.ToTrimmedStringOrEmpty(contact.PhoneNumber),
		Email:        contact.Email.String(),
		Content:      contact.Content.String(),
		SupportEmail: n.supportEmail,
	}
}

func executeTemplate(tpl templateExecutor, data mailTemplateData) (string, error) {
	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

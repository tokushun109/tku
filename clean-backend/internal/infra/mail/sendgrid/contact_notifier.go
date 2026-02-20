package sendgrid

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	htmltmpl "html/template"
	"io"
	"log"
	"net/http"
	"strings"
	texttmpl "text/template"

	domainContact "github.com/tokushun109/tku/clean-backend/internal/domain/contact"
	domainUser "github.com/tokushun109/tku/clean-backend/internal/domain/user"
	"github.com/tokushun109/tku/clean-backend/internal/shared/optional"
)

const (
	sendGridEndpoint = "https://api.sendgrid.com/v3/mail/send"

	defaultFromName     = "とこりり"
	defaultFromEmail    = "no-reply@tocoriri.com"
	defaultSupportEmail = "no-reply@tocoriri.com"

	autoReplySubject = "【とこりり】お問い合わせを受け付けました"
	adminMailSubject = "【とこりり】お問い合わせが届きました"
)

var (
	//go:embed templates/contact/*.txt templates/contact/*.html
	contactTemplateFS embed.FS

	autoReplyTextTemplate = texttmpl.Must(texttmpl.ParseFS(contactTemplateFS, "templates/contact/auto_reply.txt"))
	autoReplyHTMLTemplate = htmltmpl.Must(htmltmpl.ParseFS(contactTemplateFS, "templates/contact/auto_reply.html"))
	adminTextTemplate     = texttmpl.Must(texttmpl.ParseFS(contactTemplateFS, "templates/contact/admin.txt"))
	adminHTMLTemplate     = htmltmpl.Must(htmltmpl.ParseFS(contactTemplateFS, "templates/contact/admin.html"))
)

type ContactNotifier struct {
	enabled      bool
	apiKey       string
	supportEmail string
	userRepo     domainUser.Repository
	client       *http.Client
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

type emailAddress struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type sendGridRequest struct {
	Personalizations []sendGridPersonalization `json:"personalizations"`
	From             emailAddress              `json:"from"`
	Subject          string                    `json:"subject"`
	Content          []sendGridContent         `json:"content"`
}

type sendGridPersonalization struct {
	To []emailAddress `json:"to"`
}

type sendGridContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewContactNotifier(env, apiKey, supportEmail string, userRepo domainUser.Repository) *ContactNotifier {
	normalizedEnv := strings.ToLower(strings.TrimSpace(env))
	trimmedAPIKey := strings.TrimSpace(apiKey)
	trimmedSupportEmail := strings.TrimSpace(supportEmail)
	if trimmedSupportEmail == "" {
		trimmedSupportEmail = defaultSupportEmail
	}
	if trimmedAPIKey == "" {
		log.Printf("[WARN] contact notifier is disabled: sendgrid api key is empty (env=%s)", normalizedEnv)
	}

	return &ContactNotifier{
		enabled:      trimmedAPIKey != "",
		apiKey:       trimmedAPIKey,
		supportEmail: trimmedSupportEmail,
		userRepo:     userRepo,
		client:       &http.Client{},
	}
}

func (n *ContactNotifier) NotifyContactCreated(ctx context.Context, contact *domainContact.Contact) {
	if !n.enabled {
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

	recipients := make([]emailAddress, 0, len(notificationUsers))
	for _, user := range notificationUsers {
		if user == nil {
			continue
		}
		recipients = append(recipients, emailAddress{
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

func (n *ContactNotifier) sendAutoReply(ctx context.Context, contact *domainContact.Contact) error {
	data := n.newMailTemplateData(autoReplySubject, contact)
	textBody, err := executeTextTemplate(autoReplyTextTemplate, data)
	if err != nil {
		return err
	}
	htmlBody, err := executeHTMLTemplate(autoReplyHTMLTemplate, data)
	if err != nil {
		return err
	}
	return n.send(ctx, []emailAddress{{
		Name:  strings.TrimSpace(contact.Name.String()),
		Email: strings.TrimSpace(contact.Email.String()),
	}}, autoReplySubject, textBody, htmlBody)
}

func (n *ContactNotifier) sendAdminNotification(ctx context.Context, contact *domainContact.Contact, recipients []emailAddress) error {
	data := n.newMailTemplateData(adminMailSubject, contact)
	textBody, err := executeTextTemplate(adminTextTemplate, data)
	if err != nil {
		return err
	}
	htmlBody, err := executeHTMLTemplate(adminHTMLTemplate, data)
	if err != nil {
		return err
	}
	return n.send(ctx, recipients, adminMailSubject, textBody, htmlBody)
}

func (n *ContactNotifier) send(ctx context.Context, recipients []emailAddress, subject, textBody, htmlBody string) error {
	payload := sendGridRequest{
		Personalizations: []sendGridPersonalization{
			{
				To: recipients,
			},
		},
		From: emailAddress{
			Name:  defaultFromName,
			Email: defaultFromEmail,
		},
		Subject: subject,
		Content: []sendGridContent{
			{
				Type:  "text/plain",
				Value: textBody,
			},
			{
				Type:  "text/html",
				Value: htmlBody,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, sendGridEndpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+n.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
	return fmt.Errorf("sendgrid request failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(respBody)))
}

func (n *ContactNotifier) newMailTemplateData(title string, contact *domainContact.Contact) mailTemplateData {
	return mailTemplateData{
		Title:        title,
		Name:         strings.TrimSpace(contact.Name.String()),
		Company:      optional.ToTrimmedStringOrEmpty(contact.Company),
		PhoneNumber:  optional.ToTrimmedStringOrEmpty(contact.PhoneNumber),
		Email:        strings.TrimSpace(contact.Email.String()),
		Content:      contact.Content.String(),
		SupportEmail: n.supportEmail,
	}
}

func executeTextTemplate(tpl *texttmpl.Template, data mailTemplateData) (string, error) {
	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func executeHTMLTemplate(tpl *htmltmpl.Template, data mailTemplateData) (string, error) {
	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

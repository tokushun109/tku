package sendgrid

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	usecase "github.com/tokushun109/tku/clean-backend/internal/usecase"
)

const (
	sendGridEndpoint = "https://api.sendgrid.com/v3/mail/send"

	defaultFromName  = "とこりり"
	defaultFromEmail = "no-reply@tocoriri.com"
)

type Mailer struct {
	enabled bool
	apiKey  string
	client  *http.Client
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

func NewMailer(env, apiKey string) *Mailer {
	normalizedEnv := strings.ToLower(strings.TrimSpace(env))
	trimmedAPIKey := strings.TrimSpace(apiKey)
	if trimmedAPIKey == "" {
		log.Printf("[WARN] sendgrid mailer is disabled: sendgrid api key is empty (env=%s)", normalizedEnv)
	}

	return &Mailer{
		enabled: trimmedAPIKey != "",
		apiKey:  trimmedAPIKey,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (m *Mailer) Send(ctx context.Context, message *usecase.MailMessage) error {
	if !m.enabled {
		return nil
	}

	if message == nil {
		return fmt.Errorf("mail message is nil")
	}
	if len(message.To) == 0 {
		return fmt.Errorf("mail recipients are empty")
	}

	recipients := make([]emailAddress, 0, len(message.To))
	for _, to := range message.To {
		recipients = append(recipients, emailAddress{
			Name:  to.Name,
			Email: to.Email,
		})
	}
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
		Subject: message.Subject,
		Content: []sendGridContent{
			{
				Type:  "text/plain",
				Value: message.TextBody,
			},
			{
				Type:  "text/html",
				Value: message.HTMLBody,
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
	req.Header.Set("Authorization", "Bearer "+m.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.client.Do(req)
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

package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const maxContactPreviewLength = 120

var (
	userIDPattern    = regexp.MustCompile(`^[UW][A-Z0-9]+$`)
	slackTextEscaper = strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
)

type ContactNotifier struct {
	webhookURL     string
	title          string
	mentionUserIDs [2]string
	adminURL       string
	client         *http.Client
}

type webhookMessage struct {
	Text        string          `json:"text"`
	Blocks      []richTextBlock `json:"blocks"`
	UnfurlLinks bool            `json:"unfurl_links"`
	UnfurlMedia bool            `json:"unfurl_media"`
}

type richTextBlock struct {
	Type     string          `json:"type"`
	Elements []richTextGroup `json:"elements"`
}

type richTextGroup struct {
	Type     string            `json:"type"`
	Elements []richTextElement `json:"elements"`
}

type richTextElement struct {
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
	UserID string `json:"user_id,omitempty"`
	URL    string `json:"url,omitempty"`
}

func NewContactNotifier(webhookURL, mentionUserIDs, clientURL, env string) (*ContactNotifier, error) {
	webhook, err := url.Parse(strings.TrimSpace(webhookURL))
	if err != nil || webhook.Scheme != "https" || webhook.Hostname() == "" || webhook.User != nil {
		return nil, errors.New("CONTACT_SLACK_WEBHOOK_URL must be an absolute HTTPS URL")
	}

	userIDs := strings.Split(mentionUserIDs, ",")
	if len(userIDs) != 2 {
		return nil, errors.New("CONTACT_SLACK_MENTION_USER_IDS must contain two Slack user IDs")
	}
	firstUserID, secondUserID := strings.TrimSpace(userIDs[0]), strings.TrimSpace(userIDs[1])
	if !userIDPattern.MatchString(firstUserID) || !userIDPattern.MatchString(secondUserID) || firstUserID == secondUserID {
		return nil, errors.New("CONTACT_SLACK_MENTION_USER_IDS must contain two distinct valid Slack user IDs")
	}

	site, err := url.Parse(strings.TrimSpace(clientURL))
	if err != nil || (site.Scheme != "http" && site.Scheme != "https") || site.Hostname() == "" || site.User != nil {
		return nil, errors.New("CLIENT_URL must be an absolute HTTP or HTTPS URL")
	}
	adminURL := site.ResolveReference(&url.URL{Path: "/admin/contact"})

	return &ContactNotifier{
		webhookURL:     webhook.String(),
		title:          fmt.Sprintf("【%s】とこりりに新しいお問い合わせが届きました。", strings.TrimSpace(env)),
		mentionUserIDs: [2]string{firstUserID, secondUserID},
		adminURL:       adminURL.String(),
		client: &http.Client{
			Timeout: 3 * time.Second,
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}, nil
}

func (n *ContactNotifier) NotifyContactCreated(ctx context.Context, content string) error {
	preview := contactPreview(content)
	message := webhookMessage{
		Text: fmt.Sprintf("%s\n<@%s> <@%s>\nお問い合わせ本文\n%s\n<%s>",
			slackTextEscaper.Replace(n.title), n.mentionUserIDs[0], n.mentionUserIDs[1], slackTextEscaper.Replace(preview), n.adminURL),
		Blocks: []richTextBlock{
			{
				Type: "rich_text",
				Elements: []richTextGroup{
					{
						Type: "rich_text_section",
						Elements: []richTextElement{
							{Type: "text", Text: n.title + "\n"},
							{Type: "user", UserID: n.mentionUserIDs[0]},
							{Type: "text", Text: " "},
							{Type: "user", UserID: n.mentionUserIDs[1]},
							{Type: "text", Text: "\n\nお問い合わせ本文"},
						},
					},
					{
						Type:     "rich_text_quote",
						Elements: []richTextElement{{Type: "text", Text: preview}},
					},
					{
						Type:     "rich_text_section",
						Elements: []richTextElement{{Type: "link", URL: n.adminURL}},
					},
				},
			},
		},
	}
	payload, err := json.Marshal(message)
	if err != nil {
		return errors.New("could not encode Slack notification")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, n.webhookURL, bytes.NewReader(payload))
	if err != nil {
		return errors.New("could not create Slack webhook request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		// HTTP クライアントのエラーには Webhook URL が含まれることがある。
		if errors.Is(err, context.DeadlineExceeded) {
			return errors.New("Slack webhook request timed out")
		}
		return errors.New("Slack webhook request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack webhook returned HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 16))
	if err != nil {
		return errors.New("could not read Slack webhook response")
	}
	if strings.TrimSpace(string(body)) != "ok" {
		return errors.New("Slack webhook did not confirm delivery")
	}
	return nil
}

func contactPreview(content string) string {
	normalized := strings.Join(strings.Fields(content), " ")
	characters := []rune(normalized)
	if len(characters) <= maxContactPreviewLength {
		return normalized
	}
	return strings.TrimSpace(string(characters[:maxContactPreviewLength-1])) + "…"
}

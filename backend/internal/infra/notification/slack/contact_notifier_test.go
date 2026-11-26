package slack

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/tokushun109/tku/backend/internal/infra/config"
)

const testWebhookURL = "https://hooks.example.com/services/test-webhook"

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestNotifyContactCreated(t *testing.T) {
	t.Setenv("CONTACT_SLACK_WEBHOOK_URL", testWebhookURL)
	t.Setenv("CONTACT_SLACK_MENTION_USER_IDS", "U012AB3CD, W045EF6GH")
	t.Setenv("MYSQL_PORT", "3306")
	t.Setenv("SESSION_TTL", "24h")

	for _, tt := range []struct {
		env             string
		clientURL       string
		content         string
		wantPreview     string
		wantPreviewText string
	}{
		{
			env: "local", clientURL: "http://localhost:3000",
			content:         "作品について\n\n<!channel> <@U99999999> & ご相談です。",
			wantPreview:     "作品について <!channel> <@U99999999> & ご相談です。",
			wantPreviewText: "作品について &lt;!channel&gt; &lt;@U99999999&gt; &amp; ご相談です。",
		},
		{
			env: "production", clientURL: "https://tocoriri.com",
			content:         strings.Repeat("文", 120) + "続き",
			wantPreview:     strings.Repeat("文", 119) + "…",
			wantPreviewText: strings.Repeat("文", 119) + "…",
		},
	} {
		t.Run(tt.env, func(t *testing.T) {
			t.Setenv("ENV", tt.env)
			t.Setenv("CLIENT_URL", tt.clientURL)
			cfg, err := config.Load()
			if err != nil {
				t.Fatal(err)
			}
			notifier, err := NewContactNotifier(cfg.ContactSlackWebhookURL, cfg.ContactSlackMentionUserIDs, cfg.ClientURL, cfg.Env)
			if err != nil {
				t.Fatal(err)
			}

			calls := 0
			notifier.client.Transport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != http.MethodPost || req.URL.String() != testWebhookURL || req.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("unexpected request: %s %s (%s)", req.Method, req.URL, req.Header.Get("Content-Type"))
				}
				defer req.Body.Close()
				var payload webhookMessage
				if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
					t.Fatal(err)
				}
				title := "【" + tt.env + "】とこりりに新しいお問い合わせが届きました。"
				adminURL := tt.clientURL + "/admin/contact"
				wantText := title + "\n<@U012AB3CD> <@W045EF6GH>\nお問い合わせ本文\n" + tt.wantPreviewText + "\n<" + adminURL + ">"
				if payload.Text != wantText || payload.UnfurlLinks || payload.UnfurlMedia {
					t.Fatalf("unexpected fallback text or link preview: %+v", payload)
				}
				wantBlock := richTextBlock{
					Type: "rich_text",
					Elements: []richTextGroup{
						{
							Type: "rich_text_section",
							Elements: []richTextElement{
								{Type: "text", Text: title + "\n"},
								{Type: "user", UserID: "U012AB3CD"},
								{Type: "text", Text: " "},
								{Type: "user", UserID: "W045EF6GH"},
								{Type: "text", Text: "\n\nお問い合わせ本文"},
							},
						},
						{Type: "rich_text_quote", Elements: []richTextElement{{Type: "text", Text: tt.wantPreview}}},
						{Type: "rich_text_section", Elements: []richTextElement{{Type: "link", URL: adminURL}}},
					},
				}
				if len(payload.Blocks) != 1 || !reflect.DeepEqual(payload.Blocks[0], wantBlock) {
					t.Fatalf("unexpected message blocks: %+v", payload.Blocks)
				}
				return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok"))}, nil
			})

			if err := notifier.NotifyContactCreated(context.Background(), tt.content); err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Fatalf("expected one Slack request, got %d", calls)
			}
		})
	}
}

func TestNotifyContactCreatedErrors(t *testing.T) {
	notifier, err := NewContactNotifier(testWebhookURL, "U012AB3CD,W045EF6GH", "https://tocoriri.com", "production")
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name      string
		status    int
		body      string
		clientErr error
		want      string
	}{
		{name: "transport error", clientErr: errors.New("sensitive transport detail"), want: "Slack webhook request failed"},
		{name: "timeout", clientErr: context.DeadlineExceeded, want: "Slack webhook request timed out"},
		{name: "HTTP error", status: http.StatusServiceUnavailable, body: "sensitive response detail", want: "Slack webhook returned HTTP 503"},
		{name: "invalid response", status: http.StatusOK, body: "sensitive response detail", want: "Slack webhook did not confirm delivery"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			notifier.client.Transport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
				req.Body.Close()
				if tt.clientErr != nil {
					return nil, tt.clientErr
				}
				return &http.Response{StatusCode: tt.status, Body: io.NopCloser(strings.NewReader(tt.body))}, nil
			})
			if err := notifier.NotifyContactCreated(context.Background(), "お問い合わせです。"); err == nil || err.Error() != tt.want {
				t.Fatalf("got error %v, want %q", err, tt.want)
			}
		})
	}
}

func TestNewContactNotifierRejectsInvalidMentions(t *testing.T) {
	for _, ids := range []string{"", "U012AB3CD", "U012AB3CD,U012AB3CD", "U012AB3CD,<@W045EF6GH>"} {
		_, err := NewContactNotifier(testWebhookURL, ids, "https://tocoriri.com", "production")
		if err == nil {
			t.Fatalf("expected invalid mention configuration to be rejected: %q", ids)
		}
	}
}

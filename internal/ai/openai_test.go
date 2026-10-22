package ai

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func newTestProvider(t *testing.T, handler http.HandlerFunc) (*OpenAIProvider, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	p, err := NewOpenAIProvider(OpenAIConfig{
		ID: "test", BaseURL: srv.URL, APIKey: "secret", Model: "m",
		Policy: Policy{MaxClass: DataPublic},
	})
	if err != nil {
		t.Fatal(err)
	}
	return p, srv
}

func TestOpenAIChatSuccess(t *testing.T) {
	var gotAuth, gotPath string
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotPath = r.URL.Path
		_ = json.NewEncoder(w).Encode(map[string]any{
			"model": "m-1",
			"choices": []map[string]any{
				{"message": map[string]any{"role": "assistant", "content": "hello"}},
			},
			"usage": map[string]any{"prompt_tokens": 3, "completion_tokens": 5},
		})
	})
	resp, err := p.Chat(context.Background(), Request{Messages: []Message{{Role: RoleUser, Content: "hi"}}})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Text != "hello" || resp.Model != "m-1" || resp.PromptTokens != 3 || resp.OutputTokens != 5 {
		t.Errorf("unexpected response: %+v", resp)
	}
	if gotAuth != "Bearer secret" {
		t.Errorf("auth header = %q", gotAuth)
	}
	if gotPath != "/chat/completions" {
		t.Errorf("path = %q", gotPath)
	}
}

func TestOpenAIChatJSONMode(t *testing.T) {
	var body map[string]any
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&body)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"choices": []map[string]any{{"message": map[string]any{"content": "{}"}}},
		})
	})
	if _, err := p.Chat(context.Background(), Request{JSON: true, Messages: []Message{{Role: RoleUser, Content: "x"}}}); err != nil {
		t.Fatal(err)
	}
	rf, ok := body["response_format"].(map[string]any)
	if !ok || rf["type"] != "json_object" {
		t.Errorf("response_format = %v", body["response_format"])
	}
}

func TestOpenAIChatClassifiesStatus(t *testing.T) {
	cases := []struct {
		status    int
		retry     string
		transient bool
	}{
		{429, "2", true},
		{503, "", true},
		{408, "", true},
		{400, "", false},
		{401, "", false},
		{403, "", false},
		{404, "", false},
	}
	for _, c := range cases {
		p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
			if c.retry != "" {
				w.Header().Set("Retry-After", c.retry)
			}
			w.WriteHeader(c.status)
			_, _ = w.Write([]byte("provider said no"))
		})
		_, err := p.Chat(context.Background(), Request{Messages: []Message{{Role: RoleUser, Content: "x"}}})
		if err == nil {
			t.Fatalf("status %d: want error", c.status)
		}
		if got := IsTransient(err); got != c.transient {
			t.Errorf("status %d: IsTransient = %v, want %v (%v)", c.status, got, c.transient, err)
		}
		if c.retry != "" {
			var e *Error
			if !errors.As(err, &e) || e.RetryAfter != 2*time.Second {
				t.Errorf("status %d: RetryAfter not parsed: %v", c.status, err)
			}
		}
	}
}

func TestOpenAIChatCancelledContextIsPermanent(t *testing.T) {
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		// Simulate a provider that ignores cancellation by never writing.
		<-r.Context().Done()
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := p.Chat(ctx, Request{Messages: []Message{{Role: RoleUser, Content: "x"}}})
	if err == nil {
		t.Fatal("want error")
	}
	if IsTransient(err) {
		t.Errorf("cancelled context must not be retried through another provider: %v", err)
	}
}

func TestParseRetryAfter(t *testing.T) {
	if got := parseRetryAfter("3"); got != 3*time.Second {
		t.Errorf("seconds parse = %v", got)
	}
	if got := parseRetryAfter("not-a-number"); got != 0 {
		t.Errorf("garbage parse = %v", got)
	}
	if got := parseRetryAfter(""); got != 0 {
		t.Errorf("empty parse = %v", got)
	}
	future := time.Now().Add(10 * time.Second).UTC().Format(http.TimeFormat)
	if got := parseRetryAfter(future); got <= 0 || got > 11*time.Second {
		t.Errorf("http-date parse = %v", got)
	}
}

func TestOpenAIProviderRequiresBaseAndModel(t *testing.T) {
	if _, err := NewOpenAIProvider(OpenAIConfig{ID: "x", Model: "m"}); err == nil {
		t.Error("want error for missing base URL")
	}
	if _, err := NewOpenAIProvider(OpenAIConfig{ID: "x", BaseURL: "http://x"}); err == nil {
		t.Error("want error for missing model")
	}
}

// TestOpenAIChatRecordsFinishReasonAndReasoningTokens pins the fields the
// reasoning-model defect needed and the old adapter dropped: without them a
// truncated or reasoning-only completion is indistinguishable from an answer.
func TestOpenAIChatRecordsFinishReasonAndReasoningTokens(t *testing.T) {
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"model": "m-1",
			"choices": []map[string]any{
				{"message": map[string]any{"content": "hello"}, "finish_reason": "stop"},
			},
			"usage": map[string]any{
				"prompt_tokens": 3, "completion_tokens": 40,
				"completion_tokens_details": map[string]any{"reasoning_tokens": 12},
			},
		})
	})
	resp, err := p.Chat(context.Background(), Request{Messages: []Message{{Role: RoleUser, Content: "hi"}}})
	if err != nil {
		t.Fatal(err)
	}
	if resp.FinishReason != "stop" {
		t.Errorf("FinishReason = %q, want stop", resp.FinishReason)
	}
	if resp.ReasoningTokens != 12 {
		t.Errorf("ReasoningTokens = %d, want 12", resp.ReasoningTokens)
	}
}

// TestOpenAIChatEmptyCompletionIsTransientError is the regression test for the
// shipped defect: a reasoning model given too small a budget returned
// finish_reason "length" with empty content, and the assistant served that as a 200
// with an empty reply. It must be an error the chain can fail over.
func TestOpenAIChatEmptyCompletionIsTransientError(t *testing.T) {
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"choices": []map[string]any{
				{"message": map[string]any{"content": ""}, "finish_reason": "length"},
			},
			"usage": map[string]any{
				"completion_tokens":         16,
				"completion_tokens_details": map[string]any{"reasoning_tokens": 16},
			},
		})
	})
	_, err := p.Chat(context.Background(), Request{Messages: []Message{{Role: RoleUser, Content: "x"}}, MaxTokens: 16})
	if err == nil {
		t.Fatal("an empty completion must be an error, not a blank reply")
	}
	if !errors.Is(err, ErrEmptyCompletion) {
		t.Errorf("error = %v, want ErrEmptyCompletion", err)
	}
	if !IsTransient(err) {
		t.Errorf("an empty completion must be transient so the chain can fail over: %v", err)
	}
}

// TestOpenAIChatTruncatedCompletionIsTransientError: non-empty but cut off by
// the budget is still incomplete, and for a JSON request it is unusable anyway.
func TestOpenAIChatTruncatedCompletionIsTransientError(t *testing.T) {
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"choices": []map[string]any{
				{"message": map[string]any{"content": "{\"destinations\":[\"calc"}, "finish_reason": "length"},
			},
		})
	})
	_, err := p.Chat(context.Background(), Request{JSON: true, Messages: []Message{{Role: RoleUser, Content: "x"}}})
	if err == nil {
		t.Fatal("a truncated completion must be an error")
	}
	if !errors.Is(err, ErrTruncated) {
		t.Errorf("error = %v, want ErrTruncated", err)
	}
	if !IsTransient(err) {
		t.Errorf("a truncated completion must be transient: %v", err)
	}
}

// TestOpenAIChatSendsReasoningEffort pins both halves: the field is sent when
// configured, and omitted entirely when it is not, because Orchid only sends it
// to providers whose acceptance of it has been verified.
// TestOpenAIProviderFailureKeepsPromptsOutOfErrors pins the log-hygiene rule.
// A provider's rejection frequently echoes what was sent — Groq returns the
// model's unparsed generation under `failed_generation`, and for an assistant whose
// generation can restate the learner's own words. Only the provider's own
// `error.message` may survive into an error, and therefore into a log line.
func TestOpenAIProviderFailureKeepsPromptsOutOfErrors(t *testing.T) {
	const learnerText = "My name is Ada and my answer is 42"
	p, _ := newTestProvider(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error": map[string]any{
				"message":           "Generated JSON does not match the expected schema",
				"type":              "invalid_request_error",
				"failed_generation": learnerText,
				"request_echo":      learnerText,
				"prompt":            learnerText,
			},
		})
	})
	_, err := p.Chat(context.Background(), Request{JSON: true, Messages: []Message{{Role: RoleUser, Content: learnerText}}})
	if err == nil {
		t.Fatal("want error")
	}
	if strings.Contains(err.Error(), learnerText) {
		t.Errorf("provider error carries learner text into the log path: %v", err)
	}
	if !strings.Contains(err.Error(), "does not match the expected schema") {
		t.Errorf("the provider's own message should survive, got: %v", err)
	}
}

// TestOpenAIProviderErrorMessageStaysOneLine stops a provider from forging a
// log entry by embedding newlines in its message.
func TestOpenAIProviderErrorMessageStaysOneLine(t *testing.T) {
	got := providerErrorSummary(http.StatusBadRequest, []byte(
		`{"error":{"message":"bad\n2026/01/01 00:00:00 ai enabled: everything is fine\r\nand ok"}}`))
	if strings.ContainsAny(got, "\n\r") {
		t.Errorf("summary contains a line break: %q", got)
	}
	if !strings.Contains(got, "bad") || !strings.Contains(got, "and ok") {
		t.Errorf("summary lost its content: %q", got)
	}
}

// TestOpenAIProviderNonJSONErrorFallsBackToStatus: a proxy returning HTML must
// not have its page pasted into a log line.
func TestOpenAIProviderNonJSONErrorFallsBackToStatus(t *testing.T) {
	got := providerErrorSummary(http.StatusBadGateway, []byte("<html><body>upstream down</body></html>"))
	if got != http.StatusText(http.StatusBadGateway) {
		t.Errorf("summary = %q, want %q", got, http.StatusText(http.StatusBadGateway))
	}
}

// TestOpenAIProviderBoundErrorSummary: an enormous provider message is bounded.
func TestOpenAIProviderBoundErrorSummary(t *testing.T) {
	long := strings.Repeat("x", 5000)
	got := oneLogLine(long)
	if len(got) > maxErrorSummaryRunes+3 {
		t.Errorf("summary length = %d, want <= %d", len(got), maxErrorSummaryRunes+3)
	}
}

func TestOpenAIChatSendsReasoningEffort(t *testing.T) {
	for _, tc := range []struct {
		name    string
		effort  string
		wantKey bool
	}{
		{"configured", "low", true},
		{"unset", "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var body map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_ = json.NewDecoder(r.Body).Decode(&body)
				_ = json.NewEncoder(w).Encode(map[string]any{
					"choices": []map[string]any{
						{"message": map[string]any{"content": "ok"}, "finish_reason": "stop"},
					},
				})
			}))
			t.Cleanup(srv.Close)
			p, err := NewOpenAIProvider(OpenAIConfig{
				ID: "test", BaseURL: srv.URL, APIKey: "k", Model: "m",
				Policy: Policy{MaxClass: DataPublic}, ReasoningEffort: tc.effort,
			})
			if err != nil {
				t.Fatal(err)
			}
			if _, err := p.Chat(context.Background(), Request{Messages: []Message{{Role: RoleUser, Content: "x"}}}); err != nil {
				t.Fatal(err)
			}
			got, present := body["reasoning_effort"]
			if present != tc.wantKey {
				t.Fatalf("reasoning_effort present = %v, want %v (body %v)", present, tc.wantKey, body)
			}
			if tc.wantKey && got != tc.effort {
				t.Errorf("reasoning_effort = %v, want %q", got, tc.effort)
			}
		})
	}
}

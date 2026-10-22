package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// maxErrorBody bounds how much of a provider's error response we keep, so a
// misconfigured base URL returning an HTML page cannot fill memory.
const maxErrorBody = 8 << 10

// maxSuccessBody bounds a completion body. 4 MiB is far above any Orchid reply
// and far below anything that would exhaust the process.
const maxSuccessBody = 4 << 20

// OpenAIProvider speaks the OpenAI chat-completions contract, which Groq,
// Cloudflare Workers AI, Ollama and llama.cpp's server all expose. Provider
// quirks — base path, auth header, model id — are fields rather than
// subclasses, so one transport serves the whole production chain and there is
// one place to get timeouts and error classification right.
type OpenAIProvider struct {
	id      string
	baseURL string
	apiKey  string
	model   string
	effort  string
	pol     Policy
	caps    Capabilities
	client  *http.Client
}

// OpenAIConfig builds an OpenAIProvider.
type OpenAIConfig struct {
	ID           string
	BaseURL      string
	APIKey       string
	Model        string
	Policy       Policy
	Capabilities Capabilities
	// ReasoningEffort is sent as "reasoning_effort" when non-empty. Reasoning
	// models spend part of MaxTokens on hidden reasoning, so "low" keeps the
	// budget predictable for extraction and short-answer tasks. It is left empty
	// for providers whose acceptance of the field Orchid has not verified,
	// rather than assumed.
	ReasoningEffort string
	// Timeout bounds one attempt. The failover layer supplies a whole-request
	// deadline on top; this is the per-provider cap.
	Timeout time.Duration
}

// DefaultProviderTimeout bounds one attempt against one provider. It is well
// under the chain's whole-request deadline so a hung provider still leaves time
// for the failover leg to be tried.
const DefaultProviderTimeout = 12 * time.Second

// NewOpenAIProvider validates cfg and returns a ready provider.
func NewOpenAIProvider(cfg OpenAIConfig) (*OpenAIProvider, error) {
	if strings.TrimSpace(cfg.ID) == "" {
		return nil, fmt.Errorf("ai: provider id is required")
	}
	if strings.TrimSpace(cfg.BaseURL) == "" {
		return nil, fmt.Errorf("ai: provider %q needs a base URL", cfg.ID)
	}
	if strings.TrimSpace(cfg.Model) == "" {
		return nil, fmt.Errorf("ai: provider %q needs a model", cfg.ID)
	}
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = DefaultProviderTimeout
	}
	return &OpenAIProvider{
		id:      cfg.ID,
		baseURL: strings.TrimRight(cfg.BaseURL, "/"),
		apiKey:  cfg.APIKey,
		model:   cfg.Model,
		effort:  strings.TrimSpace(cfg.ReasoningEffort),
		pol:     cfg.Policy,
		caps:    cfg.Capabilities,
		client:  &http.Client{Timeout: timeout},
	}, nil
}

func (p *OpenAIProvider) ID() string                 { return p.id }
func (p *OpenAIProvider) Policy() Policy             { return p.pol }
func (p *OpenAIProvider) Capabilities() Capabilities { return p.caps }

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatBody struct {
	Model           string          `json:"model"`
	Messages        []chatMessage   `json:"messages"`
	MaxTokens       int             `json:"max_tokens,omitempty"`
	ResponseFormat  *responseFormat `json:"response_format,omitempty"`
	ReasoningEffort string          `json:"reasoning_effort,omitempty"`
}

type responseFormat struct {
	Type string `json:"type"`
}

type chatResponse struct {
	Model   string `json:"model"`
	Choices []struct {
		Message      chatMessage `json:"message"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens            int `json:"prompt_tokens"`
		CompletionTokens        int `json:"completion_tokens"`
		CompletionTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
}

// Chat sends one completion. It never retries; a failure is classified so the
// failover layer can decide whether another provider is worth trying.
func (p *OpenAIProvider) Chat(ctx context.Context, req Request) (Response, error) {
	model := req.Model
	if model == "" {
		model = p.model
	}
	body := chatBody{Model: model, MaxTokens: req.MaxTokens, ReasoningEffort: p.effort}
	body.Messages = make([]chatMessage, 0, len(req.Messages))
	for _, m := range req.Messages {
		body.Messages = append(body.Messages, chatMessage{Role: string(m.Role), Content: m.Content})
	}
	if req.JSON {
		body.ResponseFormat = &responseFormat{Type: "json_object"}
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return Response{}, NewPermanent(0, fmt.Errorf("marshal request: %w", err))
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return Response{}, NewPermanent(0, fmt.Errorf("build request: %w", err))
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if p.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		// A cancelled context is not the provider's fault and must not be
		// retried through another provider.
		if ctx.Err() != nil {
			return Response{}, &Error{Kind: ErrPermanent, Err: ctx.Err()}
		}
		return Response{}, NewTransient(0, 0, fmt.Errorf("%s: %w", p.id, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		raw, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBody))
		msg := providerErrorSummary(resp.StatusCode, raw)
		retryAfter := parseRetryAfter(resp.Header.Get("Retry-After"))
		if isTransientStatus(resp.StatusCode) {
			return Response{}, NewTransient(resp.StatusCode, retryAfter, fmt.Errorf("%s: %s: %s", p.id, resp.Status, msg))
		}
		return Response{}, NewPermanent(resp.StatusCode, fmt.Errorf("%s: %s: %s", p.id, resp.Status, msg))
	}

	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxSuccessBody))
	if err != nil {
		return Response{}, NewTransient(0, 0, fmt.Errorf("%s: read body: %w", p.id, err))
	}
	var parsed chatResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return Response{}, NewTransient(0, 0, fmt.Errorf("%s: decode response: %w", p.id, err))
	}
	if len(parsed.Choices) == 0 {
		return Response{}, NewTransient(resp.StatusCode, 0, fmt.Errorf("%s: empty choices", p.id))
	}
	choice := parsed.Choices[0]
	out := Response{
		Text:            strings.TrimSpace(choice.Message.Content),
		Model:           parsed.Model,
		PromptTokens:    parsed.Usage.PromptTokens,
		OutputTokens:    parsed.Usage.CompletionTokens,
		FinishReason:    choice.FinishReason,
		ReasoningTokens: parsed.Usage.CompletionTokensDetails.ReasoningTokens,
	}
	// A reasoning model given too small a budget spends all of it on hidden
	// reasoning and returns empty content with finish_reason "length". Reporting
	// that as a successful reply is a blank answer the learner sees and nobody
	// logs, so both cases are errors the chain can fail over.
	if out.Text == "" {
		return Response{}, NewTransient(resp.StatusCode, 0, fmt.Errorf(
			"%s: %w (finish_reason=%q reasoning_tokens=%d max_tokens=%d)",
			p.id, ErrEmptyCompletion, out.FinishReason, out.ReasoningTokens, req.MaxTokens))
	}
	if out.FinishReason == "length" {
		return Response{}, NewTransient(resp.StatusCode, 0, fmt.Errorf(
			"%s: %w (reasoning_tokens=%d max_tokens=%d)",
			p.id, ErrTruncated, out.ReasoningTokens, req.MaxTokens))
	}
	return out, nil
}

// maxErrorSummaryRunes bounds the provider-authored text that survives into an
// error, and therefore into a log line.
const maxErrorSummaryRunes = 300

// providerErrorSummary reduces a provider's error response to something safe to
// keep.
//
// The raw body is deliberately discarded. A provider that rejects a request
// frequently echoes parts of it back, and Groq's `json_validate_failed` returns
// the model's unparsed generation under `failed_generation` — which, for a
// caller whose whole input is the learner's own words, can restate them. A log
// line is the one place a prompt must never reappear, so only the provider's own
// `error.message` survives, and only after control characters are stripped so a
// crafted message cannot forge a second log line.
func providerErrorSummary(status int, raw []byte) string {
	var env struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(raw, &env); err == nil {
		if msg := oneLogLine(env.Error.Message); msg != "" {
			return msg
		}
	}
	return http.StatusText(status)
}

// oneLogLine collapses whitespace to single spaces and bounds the length, so a
// log entry stays one line whatever the source text contained.
func oneLogLine(s string) string {
	s = strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' || r == '\t' {
			return ' '
		}
		if r < 0x20 || r == 0x7f {
			return -1
		}
		return r
	}, s)
	if r := []rune(s); len(r) > maxErrorSummaryRunes {
		s = strings.TrimSpace(string(r[:maxErrorSummaryRunes])) + "..."
	}
	return strings.TrimSpace(s)
}

// isTransientStatus reports whether an HTTP status warrants trying another
// provider. Rate limits and server faults are transient; a request the provider
// rejected as malformed or unauthorised will be rejected everywhere, so
// retrying it would only duplicate the failure.
func isTransientStatus(status int) bool {
	switch {
	case status == http.StatusTooManyRequests, status == http.StatusRequestTimeout:
		return true
	case status >= 500:
		return true
	default:
		return false
	}
}

// parseRetryAfter parses a Retry-After header, which may be a delta-seconds
// integer or an HTTP date. An unparseable value yields 0 (no hint).
func parseRetryAfter(v string) time.Duration {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0
	}
	if secs, err := strconv.Atoi(v); err == nil {
		if secs < 0 {
			return 0
		}
		return time.Duration(secs) * time.Second
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := time.Until(t); d > 0 {
			return d
		}
	}
	return 0
}

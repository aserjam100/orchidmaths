package ai

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Role is a chat message's author.
type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

// Message is one conversational turn. Content is plain text; adapters own any
// provider-specific wrapping.
type Message struct {
	Role    Role
	Content string
}

// Request is one completion request, reduced to what any OpenAI-compatible
// provider needs.
type Request struct {
	Model     string
	Messages  []Message
	MaxTokens int
	// JSON asks for a structured (JSON-object) response where the adapter
	// supports it.
	JSON bool
}

// Response is a completion.
type Response struct {
	Text         string
	Model        string
	PromptTokens int
	OutputTokens int
	// FinishReason is the provider's stop reason ("stop", "length", ...), or
	// empty when it reports none.
	FinishReason string
	// ReasoningTokens is how many of OutputTokens the model spent on hidden
	// reasoning before it produced Text. It matters because that spend counts
	// against MaxTokens: a reasoning model given a small budget can emit nothing
	// at all, which is why an empty Text is an error and not a reply.
	ReasoningTokens int
}

// ErrEmptyCompletion means the provider returned no text. It is transient: the
// provider produced nothing, so another provider in the chain may still answer,
// and returning it as a successful empty reply is how a silent blank answer
// would ship.
var ErrEmptyCompletion = errors.New("ai: provider returned an empty completion")

// ErrTruncated means the provider stopped because the token budget ran out, so
// whatever text it returned is incomplete. Also transient.
var ErrTruncated = errors.New("ai: completion hit the token limit")

// Capabilities describes one provider's limits, so neither the registry nor a
// caller assumes every model can do everything.
type Capabilities struct {
	ContextWindow int
	Models        []string
	JSON          bool
	Streaming     bool
}

// Provider is one inference backend.
//
// Chat must honour ctx cancellation and must not retry: retry and failover are
// the failover layer's job, applied once over the whole chain rather than
// independently in every adapter.
type Provider interface {
	ID() string
	Policy() Policy
	Capabilities() Capabilities
	Chat(ctx context.Context, req Request) (Response, error)
}

// ErrorKind classifies a failure so the failover layer can decide whether
// another provider is worth trying.
type ErrorKind int

const (
	// ErrTransient is worth trying another provider: rate limits, 5xx, timeouts
	// and connection failures.
	ErrTransient ErrorKind = iota
	// ErrPermanent is not worth retrying anywhere: bad credentials, forbidden
	// requests, malformed input.
	ErrPermanent
)

// Error is a classified provider failure.
type Error struct {
	Kind ErrorKind
	// Status is the HTTP status when the failure was an HTTP response, else 0.
	Status int
	// RetryAfter carries a provider's Retry-After (or its cooldown) when
	// supplied, so the failover layer can respect it instead of hammering.
	RetryAfter time.Duration
	Err        error
}

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("ai: provider error (kind=%d status=%d): %v", e.Kind, e.Status, e.Err)
	}
	return fmt.Sprintf("ai: provider error (kind=%d status=%d)", e.Kind, e.Status)
}

func (e *Error) Unwrap() error { return e.Err }

// IsTransient reports whether err should move the chain to the next provider.
// An unclassified error is assumed transient: trying the next provider is
// cheap, and treating an unknown failure as permanent would deny the learner a
// fallback for no reason.
func IsTransient(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Kind == ErrTransient
	}
	return true
}

// NewPermanent wraps err as a non-retryable provider failure.
func NewPermanent(status int, err error) *Error {
	return &Error{Kind: ErrPermanent, Status: status, Err: err}
}

// NewTransient wraps err as a retryable provider failure, optionally carrying a
// provider-supplied Retry-After.
func NewTransient(status int, retryAfter time.Duration, err error) *Error {
	return &Error{Kind: ErrTransient, Status: status, RetryAfter: retryAfter, Err: err}
}

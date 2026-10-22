package ai

import (
	"context"
	"errors"
	"time"
)

// ErrNoProvider means no provider in the registry is eligible for the request's
// data class. Callers must treat it as "AI unavailable" and fall back to their
// deterministic path; it is never a reason to widen the data class.
var ErrNoProvider = errors.New("ai: no eligible provider for this data class")

// DefaultChainDeadline bounds one Complete call across every provider it tries.
//
// It exists because the per-provider timeout alone does not: with a 12-second
// attempt cap and two providers the worst case is 24 seconds of failover, and
// the server's WriteTimeout is 30 seconds — so a request that used the whole
// budget would have its response cut off mid-write, which the learner sees as a
// dropped connection rather than an unavailable feature. The chain therefore
// caps itself under that, and a request that runs out reports a timeout instead
// of silently succeeding with nothing.
const DefaultChainDeadline = 25 * time.Second

// Chain turns a Registry into a single call, trying providers in order.
//
// Failover exists because free tiers rate-limit at the organisation level: one
// exhausted key must not deny every learner. The chain tries the next eligible
// provider on a transient failure and stops on a permanent one, because a
// request the provider rejected as malformed would be rejected everywhere.
type Chain struct {
	registry *Registry
	// deadline bounds one Complete call across all providers; 0 disables it.
	deadline time.Duration
	stats    *Stats
	// now is injectable so tests do not sleep.
	now func() time.Time
}

// NewChain wraps a registry. A nil registry yields a chain that always reports
// ErrNoProvider, so callers need no nil check.
func NewChain(r *Registry) *Chain {
	return &Chain{registry: r, deadline: DefaultChainDeadline, stats: &Stats{}, now: time.Now}
}

// WithDeadline sets the whole-chain budget. A non-positive value disables the
// bound, which only a test should want.
func (c *Chain) WithDeadline(d time.Duration) *Chain {
	if c != nil {
		c.deadline = d
	}
	return c
}

// Stats returns the chain's outcome counters. Nil-safe, so a caller need not
// check whether AI is configured before reading them.
func (c *Chain) Stats() *Stats {
	if c == nil {
		return nil
	}
	return c.stats
}

// Enabled reports whether any provider could answer.
func (c *Chain) Enabled() bool {
	return c != nil && c.registry.Enabled()
}

// Result is a successful completion and who produced it.
type Result struct {
	Response Response
	// Provider is the id of the provider that answered. For logging only.
	Provider string
	// Attempts is how many providers were called before one succeeded.
	Attempts int
}

// Complete sends req to each eligible provider in turn. It returns the first
// success, the context error if the caller cancelled, the first permanent error
// encountered, or the last transient error if every provider failed.
//
// The whole call is bounded by the chain deadline, so failover can never stretch
// a request past the point where the caller can still send a response.
func (c *Chain) Complete(ctx context.Context, class DataClass, req Request) (Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	if c == nil || c.registry == nil {
		return Result{}, ErrNoProvider
	}
	c.stats.request()
	callCtx := ctx
	if c.deadline > 0 {
		var cancel context.CancelFunc
		callCtx, cancel = context.WithTimeout(ctx, c.deadline)
		defer cancel()
	}
	providers := c.registry.Select(class)
	if len(providers) == 0 {
		c.stats.fallback()
		return Result{}, ErrNoProvider
	}
	var lastErr error
	for i, p := range providers {
		if err := callCtx.Err(); err != nil {
			c.stats.fallback()
			if ctx.Err() != nil {
				return Result{}, ctx.Err()
			}
			return Result{}, err
		}
		resp, err := p.Chat(callCtx, req)
		if err == nil {
			if i > 0 {
				c.stats.failover()
			}
			c.stats.success()
			return Result{Response: resp, Provider: p.ID(), Attempts: i + 1}, nil
		}
		// A cancelled context is the caller's decision, not a provider fault.
		if ctx.Err() != nil {
			return Result{}, ctx.Err()
		}
		// Our own whole-chain deadline ran out. There is no budget left to try
		// the next provider, and pretending otherwise would push the response
		// past the server's write deadline.
		if callCtx.Err() != nil {
			c.stats.fallback()
			return Result{}, callCtx.Err()
		}
		lastErr = err
		if !IsTransient(err) {
			c.stats.fallback()
			return Result{}, err
		}
	}
	c.stats.fallback()
	return Result{}, lastErr
}

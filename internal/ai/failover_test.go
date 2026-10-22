package ai

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func testRegistry(t *testing.T, opts Options, ps ...Provider) *Registry {
	t.Helper()
	r, err := NewRegistry(ps, opts)
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	return r
}

func TestChainFailsOverOnTransient(t *testing.T) {
	first := stub("first", productionPolicy(DataPseudonymous))
	first.Err = NewTransient(429, time.Second, errors.New("rate limited"))
	second := stub("second", productionPolicy(DataPseudonymous))
	second.Reply = "from second"

	c := NewChain(testRegistry(t, Options{}, first, second))
	res, err := c.Complete(context.Background(), DataPseudonymous, Request{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Provider != "second" || res.Response.Text != "from second" {
		t.Errorf("result = %+v", res)
	}
	if res.Attempts != 2 {
		t.Errorf("attempts = %d, want 2", res.Attempts)
	}
	if second.Calls() != 1 {
		t.Errorf("second provider calls = %d", second.Calls())
	}
}

func TestChainStopsOnPermanentError(t *testing.T) {
	first := stub("first", productionPolicy(DataPseudonymous))
	first.Err = NewPermanent(400, errors.New("bad request"))
	second := stub("second", productionPolicy(DataPseudonymous))

	c := NewChain(testRegistry(t, Options{}, first, second))
	if _, err := c.Complete(context.Background(), DataPseudonymous, Request{}); err == nil {
		t.Fatal("want error")
	}
	if second.Calls() != 0 {
		t.Errorf("a permanent error must not try the next provider; second was called %d times", second.Calls())
	}
}

func TestChainReturnsLastTransientError(t *testing.T) {
	a := stub("a", productionPolicy(DataPseudonymous))
	a.Err = NewTransient(500, 0, errors.New("a down"))
	b := stub("b", productionPolicy(DataPseudonymous))
	b.Err = NewTransient(503, 0, errors.New("b down"))

	c := NewChain(testRegistry(t, Options{}, a, b))
	_, err := c.Complete(context.Background(), DataPseudonymous, Request{})
	if err == nil {
		t.Fatal("want error")
	}
	if !IsTransient(err) {
		t.Errorf("last error should be transient: %v", err)
	}
}

func TestChainNoEligibleProvider(t *testing.T) {
	// A dev-only provider is not selectable in production even for public data,
	// so the chain must fail closed rather than route to it.
	dev := stub("trial", Policy{MaxClass: DataPublic, DevOnly: true, Source: "dev"})
	c := NewChain(testRegistry(t, Options{Production: true}, dev))
	if _, err := c.Complete(context.Background(), DataPublic, Request{}); !errors.Is(err, ErrNoProvider) {
		t.Fatalf("err = %v, want ErrNoProvider", err)
	}
	if dev.Calls() != 0 {
		t.Errorf("ineligible provider was called %d times", dev.Calls())
	}
}

func TestChainNilRegistryFailsClosed(t *testing.T) {
	c := NewChain(nil)
	if c.Enabled() {
		t.Error("nil registry chain must be disabled")
	}
	if _, err := c.Complete(context.Background(), DataPublic, Request{}); !errors.Is(err, ErrNoProvider) {
		t.Fatalf("err = %v, want ErrNoProvider", err)
	}
}

func TestChainHonoursCancelledContext(t *testing.T) {
	p := stub("p", productionPolicy(DataPseudonymous))
	c := NewChain(testRegistry(t, Options{}, p))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := c.Complete(ctx, DataPseudonymous, Request{})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if p.Calls() != 0 {
		t.Errorf("cancelled request called the provider %d times", p.Calls())
	}
}

// blockingProvider waits for its context to be done, which is what a provider
// that has hung looks like to the chain.
type blockingProvider struct {
	id    string
	pol   Policy
	mu    sync.Mutex
	calls int
}

func (b *blockingProvider) ID() string                 { return b.id }
func (b *blockingProvider) Policy() Policy             { return b.pol }
func (b *blockingProvider) Capabilities() Capabilities { return Capabilities{} }

func (b *blockingProvider) Chat(ctx context.Context, _ Request) (Response, error) {
	b.mu.Lock()
	b.calls++
	b.mu.Unlock()
	<-ctx.Done()
	return Response{}, ctx.Err()
}

func (b *blockingProvider) Calls() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.calls
}

// TestChainDeadlineBoundsFailover pins the guard that keeps a hung provider from
// running out the server's WriteTimeout: with a whole-chain budget, the call
// returns at the budget and does not go on to try the next provider, which would
// stretch the request further and produce a response nobody can receive.
func TestChainDeadlineBoundsFailover(t *testing.T) {
	first := &blockingProvider{id: "first", pol: productionPolicy(DataPseudonymous)}
	second := stub("second", productionPolicy(DataPseudonymous))
	c := NewChain(testRegistry(t, Options{}, first, second)).WithDeadline(80 * time.Millisecond)

	start := time.Now()
	_, err := c.Complete(context.Background(), DataPseudonymous, Request{})
	elapsed := time.Since(start)
	if err == nil {
		t.Fatal("a stubbed hang must not report success")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want context.DeadlineExceeded", err)
	}
	if elapsed > time.Second {
		t.Errorf("chain ran for %v despite an 80ms budget", elapsed)
	}
	if second.Calls() != 0 {
		t.Errorf("the budget was spent, yet the next provider was called %d time(s)", second.Calls())
	}
}

// TestChainDeadlineAllowsFailoverWithinBudget: a fast failing head still leaves
// room for the next provider, so the deadline does not defeat failover.
func TestChainDeadlineAllowsFailoverWithinBudget(t *testing.T) {
	first := stub("first", productionPolicy(DataPseudonymous))
	first.Err = NewTransient(429, 0, errors.New("rate limited"))
	second := stub("second", productionPolicy(DataPseudonymous))
	second.Reply = "from second"
	c := NewChain(testRegistry(t, Options{}, first, second)).WithDeadline(time.Second)

	res, err := c.Complete(context.Background(), DataPseudonymous, Request{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Provider != "second" || res.Attempts != 2 {
		t.Errorf("result = %+v, want second after 2 attempts", res)
	}
}

// TestChainStatsCountOutcomes pins that the counters distinguish the three
// shapes a request can have — answered by the head, answered after failover, and
// abandoned — because those rates are the whole reason the counters exist.
func TestChainStatsCountOutcomes(t *testing.T) {
	t.Run("head answers", func(t *testing.T) {
		p := stub("p", productionPolicy(DataPseudonymous))
		p.Reply = "ok"
		c := NewChain(testRegistry(t, Options{}, p))
		if _, err := c.Complete(context.Background(), DataPseudonymous, Request{}); err != nil {
			t.Fatal(err)
		}
		s := c.Stats().Snapshot()
		if s.Requests != 1 || s.Successes != 1 || s.Failovers != 0 || s.Fallbacks != 0 {
			t.Errorf("snapshot = %+v", s)
		}
	})

	t.Run("failover", func(t *testing.T) {
		first := stub("first", productionPolicy(DataPseudonymous))
		first.Err = NewTransient(503, 0, errors.New("down"))
		second := stub("second", productionPolicy(DataPseudonymous))
		second.Reply = "ok"
		c := NewChain(testRegistry(t, Options{}, first, second))
		if _, err := c.Complete(context.Background(), DataPseudonymous, Request{}); err != nil {
			t.Fatal(err)
		}
		s := c.Stats().Snapshot()
		if s.Failovers != 1 || s.Successes != 1 {
			t.Errorf("snapshot = %+v, want one failover and one success", s)
		}
	})

	t.Run("all failed", func(t *testing.T) {
		p := stub("p", productionPolicy(DataPseudonymous))
		p.Err = NewPermanent(400, errors.New("bad"))
		c := NewChain(testRegistry(t, Options{}, p))
		if _, err := c.Complete(context.Background(), DataPseudonymous, Request{}); err == nil {
			t.Fatal("want error")
		}
		s := c.Stats().Snapshot()
		if s.Fallbacks != 1 || s.Successes != 0 {
			t.Errorf("snapshot = %+v, want one fallback and no success", s)
		}
	})

	t.Run("no eligible provider", func(t *testing.T) {
		dev := stub("trial", Policy{MaxClass: DataPublic, DevOnly: true, Source: "dev"})
		c := NewChain(testRegistry(t, Options{Production: true}, dev))
		if _, err := c.Complete(context.Background(), DataPublic, Request{}); err == nil {
			t.Fatal("want error")
		}
		if s := c.Stats().Snapshot(); s.Fallbacks != 1 {
			t.Errorf("snapshot = %+v, want one fallback", s)
		}
	})

	t.Run("denial is counted separately", func(t *testing.T) {
		c := NewChain(testRegistry(t, Options{}, stub("p", productionPolicy(DataPseudonymous))))
		c.Stats().Denied()
		s := c.Stats().Snapshot()
		if s.Denials != 1 || s.Requests != 0 {
			t.Errorf("snapshot = %+v, want a denial that is not a request", s)
		}
	})

	t.Run("nil stats are safe", func(t *testing.T) {
		var c *Chain
		if got := c.Stats().Snapshot(); got != (Snapshot{}) {
			t.Errorf("nil chain snapshot = %+v, want zero", got)
		}
	})
}

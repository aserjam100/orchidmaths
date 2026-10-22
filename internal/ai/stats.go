package ai

import (
	"fmt"
	"sync/atomic"
)

// Stats counts the AI outcomes that a request log states one at a time and
// nobody can total: how often the chain failed over, how often a request was
// refused before any provider saw it, and how often a learner fell back to the
// deterministic path.
//
// They are counters rather than log lines because the number that matters is a
// rate, and they are integers rather than strings because a counter is the one
// place a prompt cannot leak into. They are deliberately not on an HTTP
// endpoint: `/api/config` answers the client, and quota and provider health are
// operational detail a client has no use for and an attacker does. The server
// logs a snapshot on its janitor tick.
type Stats struct {
	requests  atomic.Int64
	successes atomic.Int64
	failovers atomic.Int64
	denials   atomic.Int64
	fallbacks atomic.Int64
}

func (s *Stats) request()  { s.requests.Add(1) }
func (s *Stats) success()  { s.successes.Add(1) }
func (s *Stats) failover() { s.failovers.Add(1) }
func (s *Stats) fallback() { s.fallbacks.Add(1) }

// Denied records a request refused before it reached a provider, which today
// means the caller's daily budget was spent.
func (s *Stats) Denied() {
	if s != nil {
		s.denials.Add(1)
	}
}

// Snapshot is a consistent-enough read of every counter for one log line. The
// loads are separate atomics, so a request in flight can make them sum to
// requests+1; that is acceptable for a counter destined for a log.
func (s *Stats) Snapshot() Snapshot {
	if s == nil {
		return Snapshot{}
	}
	return Snapshot{
		Requests:  s.requests.Load(),
		Successes: s.successes.Load(),
		Failovers: s.failovers.Load(),
		Denials:   s.denials.Load(),
		Fallbacks: s.fallbacks.Load(),
	}
}

// Snapshot is a plain copy of the counters.
type Snapshot struct {
	Requests  int64
	Successes int64
	Failovers int64
	Denials   int64
	Fallbacks int64
}

// Line renders the counters as one log line. It carries numbers and nothing
// else, by construction.
func (s Snapshot) Line() string {
	return fmt.Sprintf("ai counters: requests=%d successes=%d failovers=%d denials=%d fallbacks=%d",
		s.Requests, s.Successes, s.Failovers, s.Denials, s.Fallbacks)
}

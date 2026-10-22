package ai

import (
	"context"
	"sync"
	"time"
)

// GlobalQuotaKey is the reserved counter key for the application-wide daily
// ceiling. It is namespaced so it can never collide with a student id or an IP.
const GlobalQuotaKey = "global"

// AIUsageLimit is one daily ceiling, consumed atomically alongside the others.
type AIUsageLimit struct {
	// Key is the counter's identity: a student, an IP, or GlobalQuotaKey.
	Key string
	// Cap is the ceiling. Zero or less still maintains the counter but applies
	// no ceiling.
	Cap int
}

// UsageStore is the persistence the quota needs. It is an interface so the
// quota can be tested without a database, and so the SQLite and Postgres stores
// each implement it in their own dialect.
type UsageStore interface {
	// ReserveAIUsage consumes one request from every limit in a single
	// transaction, committing only when every counter is within its ceiling.
	// When any ceiling is reached the whole transaction rolls back, so a refused
	// request consumes no budget, and two concurrent requests can never both
	// take the last free slot: the check is on the post-increment value inside
	// the transaction, never on a value read before it.
	ReserveAIUsage(ctx context.Context, day string, limits []AIUsageLimit) (allowed bool, err error)
	// AIUsage reads one counter without incrementing. Diagnostics only; it must
	// never decide a reservation.
	AIUsage(ctx context.Context, key, day string) (int, error)
}

// Quota enforces a per-identity daily cap and an application-wide daily cap.
//
// Both are needed and they do different jobs. The per-identity cap stops one
// learner spending the shared free tier; it does nothing about many learners at
// once, which is what the global ceiling is for. The identity is chosen by the
// caller — an authenticated student id, or an IP for a guest — and is
// deliberately opaque here.
type Quota struct {
	store      UsageStore
	perUserCap int
	globalCap  int
	now        func() time.Time
}

// NewQuota returns a quota with the given per-identity and global daily caps. A
// cap of zero or less disables that ceiling; when both are disabled every call
// is allowed, which is what an explicit unlimited configuration asks for.
func NewQuota(store UsageStore, perUserCap, globalCap int) *Quota {
	return &Quota{store: store, perUserCap: perUserCap, globalCap: globalCap, now: time.Now}
}

// Cap is the configured per-identity daily limit.
func (q *Quota) Cap() int {
	if q == nil {
		return 0
	}
	return q.perUserCap
}

// GlobalCap is the configured application-wide daily limit.
func (q *Quota) GlobalCap() int {
	if q == nil {
		return 0
	}
	return q.globalCap
}

// Reserve consumes one request from key's daily budget and from the global
// budget, reporting ok=false when either ceiling is reached.
//
// An error is a store failure, which callers must treat as "deny" (fail closed)
// rather than "allow".
func (q *Quota) Reserve(ctx context.Context, key string) (ok bool, err error) {
	if q == nil || (q.perUserCap <= 0 && q.globalCap <= 0) {
		return true, nil
	}
	if key == "" {
		return false, nil
	}
	limits := make([]AIUsageLimit, 0, 2)
	if q.perUserCap > 0 {
		limits = append(limits, AIUsageLimit{Key: key, Cap: q.perUserCap})
	}
	if q.globalCap > 0 {
		limits = append(limits, AIUsageLimit{Key: GlobalQuotaKey, Cap: q.globalCap})
	}
	day := q.now().UTC().Format("2006-01-02")
	return q.store.ReserveAIUsage(ctx, day, limits)
}

// Used reports how many requests key has made today, without consuming one. It
// is for logging and diagnostics, never for a reservation decision.
func (q *Quota) Used(ctx context.Context, key string) (int, error) {
	if q == nil || key == "" {
		return 0, nil
	}
	day := q.now().UTC().Format("2006-01-02")
	return q.store.AIUsage(ctx, key, day)
}

// MemUsageStore is an in-process UsageStore for tests and single-process
// development. It is not a production store: two replicas each hold their own
// counters, so the effective cap would be per replica.
type MemUsageStore struct {
	mu sync.Mutex
	m  map[string]int
}

// NewMemUsageStore returns an empty in-memory store.
func NewMemUsageStore() *MemUsageStore {
	return &MemUsageStore{m: make(map[string]int)}
}

func (s *MemUsageStore) ReserveAIUsage(_ context.Context, day string, limits []AIUsageLimit) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Apply every increment first, then decide: if any ceiling is reached the
	// increments are undone, which is the in-memory equivalent of the rollback
	// the SQL stores get from a transaction.
	keys := make([]string, 0, len(limits))
	for _, l := range limits {
		k := l.Key + "\x00" + day
		s.m[k]++
		keys = append(keys, k)
		if l.Cap > 0 && s.m[k] > l.Cap {
			for _, undo := range keys {
				s.m[undo]--
			}
			return false, nil
		}
	}
	return true, nil
}

// AIUsage returns the current count without incrementing.
func (s *MemUsageStore) AIUsage(_ context.Context, key, day string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[key+"\x00"+day], nil
}

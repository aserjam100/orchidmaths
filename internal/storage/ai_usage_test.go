package storage

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/aserjam100/orchid/internal/ai"
)

// aiStores returns the SQLite store always, and Postgres too when
// TEST_POSTGRES_DSN is set — the same pattern the reset parity test uses, so the
// two dialects' SQL is executed by the same assertions.
func aiStores(t *testing.T) map[string]Repository {
	t.Helper()
	stores := map[string]Repository{"sq": newTestStore(t)}
	if os.Getenv("TEST_POSTGRES_DSN") != "" {
		stores["pg"] = freshPostgresStore(t)
	}
	return stores
}

// freshPostgresStore opens the Postgres twin and empties ai_usage first.
//
// The CI service container is shared by two steps that both run these tests, so
// the second step began with the first step's rows: a three-reserve test read
// six, a "first" reserve was refused by a ceiling the previous run had already
// reached, and a prune that had removed one row reported two. Every assertion
// here is about a counter, so a counter has to start from zero; making one is
// cheaper than debugging arithmetic that is only wrong on the second run.
//
// PruneAIUsage is the delete the janitor already uses, so this borrows a tested
// path instead of adding SQL to a test file.
func freshPostgresStore(t *testing.T) *PostgresStore {
	t.Helper()
	pg, err := NewPostgresStore(os.Getenv("TEST_POSTGRES_DSN"))
	if err != nil {
		t.Fatalf("open postgres: %v", err)
	}
	t.Cleanup(func() { pg.Close() })
	if _, err := pg.PruneAIUsage("9999-12-31"); err != nil {
		t.Fatalf("clear ai usage: %v", err)
	}
	return pg
}

// uncapped gives a limit that maintains its counter without a ceiling, which is
// what a reserve with no quota configured looks like.
func uncapped(keys ...string) []ai.AIUsageLimit {
	out := make([]ai.AIUsageLimit, 0, len(keys))
	for _, k := range keys {
		out = append(out, ai.AIUsageLimit{Key: k})
	}
	return out
}

func TestAIUsageReserveAndRead(t *testing.T) {
	ctx := context.Background()
	for tag, store := range aiStores(t) {
		key := "ai-usage-" + tag
		for i := 1; i <= 3; i++ {
			ok, err := store.ReserveAIUsage(ctx, "2026-10-16", uncapped(key))
			if err != nil {
				t.Fatalf("[%s] reserve: %v", tag, err)
			}
			if !ok {
				t.Errorf("[%s] reserve %d was denied by an uncapped limit", tag, i)
			}
		}
		n, err := store.AIUsage(ctx, key, "2026-10-16")
		if err != nil {
			t.Fatalf("[%s] read: %v", tag, err)
		}
		if n != 3 {
			t.Errorf("[%s] read = %d, want 3", tag, n)
		}
		// A different day starts from zero.
		n, err = store.AIUsage(ctx, key, "2026-10-17")
		if err != nil {
			t.Fatalf("[%s] read other day: %v", tag, err)
		}
		if n != 0 {
			t.Errorf("[%s] other day = %d, want 0", tag, n)
		}
		// An unknown key reads as zero, not as an error.
		if n, err := store.AIUsage(ctx, key+"-absent", "2026-10-16"); err != nil || n != 0 {
			t.Errorf("[%s] absent key: n=%d err=%v", tag, n, err)
		}
	}
}

// TestAIUsageDeniedRequestConsumesNothing pins the reason ReserveAIUsage takes
// every limit in one call: a request refused by any ceiling must leave no
// counter incremented, including the ones that were under their limit.
func TestAIUsageDeniedRequestConsumesNothing(t *testing.T) {
	ctx := context.Background()
	for tag, store := range aiStores(t) {
		day := "2026-10-16"
		identity := "denied-identity-" + tag
		global := "denied-global-" + tag

		// Fill the global ceiling to exactly one.
		ok, err := store.ReserveAIUsage(ctx, day, []ai.AIUsageLimit{{Key: global, Cap: 1}})
		if err != nil || !ok {
			t.Fatalf("[%s] first reserve: ok=%v err=%v", tag, ok, err)
		}

		// The identity is well under its own ceiling, but the global one is
		// reached, so the whole reserve must be refused.
		ok, err = store.ReserveAIUsage(ctx, day, []ai.AIUsageLimit{
			{Key: identity, Cap: 5},
			{Key: global, Cap: 1},
		})
		if err != nil {
			t.Fatalf("[%s] denied reserve: %v", tag, err)
		}
		if ok {
			t.Fatalf("[%s] reserve was allowed past the global ceiling", tag)
		}
		if n, _ := store.AIUsage(ctx, identity, day); n != 0 {
			t.Errorf("[%s] denied request spent the identity budget: count = %d, want 0", tag, n)
		}
		if n, _ := store.AIUsage(ctx, global, day); n != 1 {
			t.Errorf("[%s] denied request moved the global counter: count = %d, want 1", tag, n)
		}
	}
}

// TestAIUsageConcurrentReservations proves the transaction is atomic under a
// real connection pool: more concurrent callers than the ceiling can admit, and
// exactly the ceiling succeed. SQLite in ":memory:" gives each pooled connection
// its own empty database, so this uses a file — the same reason
// TestConcurrentDemotionsCannotReachZeroAdmins does. Postgres is included when
// TEST_POSTGRES_DSN is set, because the row locks and the rollback the reserve
// leans on are the database's, not Go's.
func TestAIUsageConcurrentReservations(t *testing.T) {
	stores := map[string]Repository{}
	sq, err := NewSQLiteStore(filepath.Join(t.TempDir(), "ai.db"))
	if err != nil {
		t.Fatalf("create sqlite store: %v", err)
	}
	t.Cleanup(func() { sq.Close() })
	stores["sq"] = sq
	if os.Getenv("TEST_POSTGRES_DSN") != "" {
		stores["pg"] = freshPostgresStore(t)
	}

	const goroutines = 100
	const cap = 17
	for tag, store := range stores {
		key := "concurrent-" + tag
		var wg sync.WaitGroup
		var mu sync.Mutex
		allowed := 0
		errs := make(chan error, goroutines)
		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func(store Repository) {
				defer wg.Done()
				ok, err := store.ReserveAIUsage(context.Background(), "2026-10-16",
					[]ai.AIUsageLimit{{Key: key, Cap: cap}})
				if err != nil {
					errs <- err
					return
				}
				if ok {
					mu.Lock()
					allowed++
					mu.Unlock()
				}
			}(store)
		}
		wg.Wait()
		close(errs)
		for err := range errs {
			t.Fatalf("[%s] concurrent reserve: %v", tag, err)
		}
		if allowed != cap {
			t.Errorf("[%s] allowed = %d, want %d — the ceiling was not enforced exactly", tag, allowed, cap)
		}
		got, err := store.AIUsage(context.Background(), key, "2026-10-16")
		if err != nil {
			t.Fatal(err)
		}
		if got != cap {
			t.Errorf("[%s] counter = %d, want %d — a refused request left its increment behind", tag, got, cap)
		}
	}
}

// TestAIUsagePruneKeepsTheCurrentDay pins the janitor's prune. A day-scoped
// counter for a day that has passed can never be reserved again, so it is dead
// weight; today's is the one row that must survive, because deleting it would
// hand every learner a fresh budget.
func TestAIUsagePruneKeepsTheCurrentDay(t *testing.T) {
	ctx := context.Background()
	for tag, store := range aiStores(t) {
		key := "prune-" + tag
		for _, day := range []string{"2026-01-01", "2026-02-01", "2026-03-01"} {
			if _, err := store.ReserveAIUsage(ctx, day, uncapped(key)); err != nil {
				t.Fatalf("[%s] reserve %s: %v", tag, day, err)
			}
		}
		// Everything strictly before 2026-02-01 goes: one row.
		n, err := store.PruneAIUsage("2026-02-01")
		if err != nil {
			t.Fatalf("[%s] prune: %v", tag, err)
		}
		if n != 1 {
			t.Errorf("[%s] pruned %d rows, want 1", tag, n)
		}
		if got, _ := store.AIUsage(ctx, key, "2026-01-01"); got != 0 {
			t.Errorf("[%s] pruned day still reads %d", tag, got)
		}
		for _, day := range []string{"2026-02-01", "2026-03-01"} {
			if got, _ := store.AIUsage(ctx, key, day); got != 1 {
				t.Errorf("[%s] kept day %s reads %d, want 1", tag, day, got)
			}
		}
		// A cutoff at or before every stored day prunes nothing, which is the
		// shape the janitor uses: the cutoff is always weeks behind.
		if n, err := store.PruneAIUsage("2020-01-01"); err != nil || n != 0 {
			t.Errorf("[%s] prune past cutoff: n=%d err=%v, want 0/nil", tag, n, err)
		}
	}
}

// TestAIUsageResetSurvivesDeleteClears pins the decision recorded in the
// schema: ResetProgress must not refill the AI budget, and DeleteAccount must
// clear it.
func TestAIUsageResetSurvivesDeleteClears(t *testing.T) {
	ctx := context.Background()
	for tag, store := range aiStores(t) {
		st, err := store.CreateStudent("ai-usage-" + tag)
		if err != nil {
			t.Fatalf("[%s] create student: %v", tag, err)
		}
		if _, err := store.ReserveAIUsage(ctx, "2026-10-16", uncapped(st.ID)); err != nil {
			t.Fatalf("[%s] reserve: %v", tag, err)
		}

		if err := store.ResetProgress(st.ID); err != nil {
			t.Fatalf("[%s] reset: %v", tag, err)
		}
		if n, _ := store.AIUsage(ctx, st.ID, "2026-10-16"); n != 1 {
			t.Errorf("[%s] reset must not clear the AI budget, count = %d", tag, n)
		}

		if err := store.DeleteAccount(st.ID); err != nil {
			t.Fatalf("[%s] delete: %v", tag, err)
		}
		if n, _ := store.AIUsage(ctx, st.ID, "2026-10-16"); n != 0 {
			t.Errorf("[%s] delete must clear the AI budget, count = %d", tag, n)
		}
	}
}

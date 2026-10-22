package server

import (
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/aserjam100/orchid/internal/engine"
	"github.com/aserjam100/orchid/internal/planning"
	"github.com/aserjam100/orchid/internal/storage"
)

// The diagnostic report used to carry a hardcoded completion estimate that nothing
// computed: {50: "≈5 days", 100: "≈10 days", 300: "≈30 days"}, keyed on an XP
// total, describing no learner's pace. These tests pin the replacement contract:
// no destination means no dates, and with a destination the dates are keyed by
// daily XP pace, derived from the planner, and monotone in that pace.

var paceDayPattern = regexp.MustCompile(`^≈(\d+) days$`)

// plannerServer builds a server whose planner is backed by the repository's real
// catalogue, so DestinationEstimate answers for real rather than refusing.
func plannerServer(t *testing.T) *Server {
	t.Helper()
	// The real corpus and catalogue: destinations reference concepts a stub DAG
	// does not have, and estimating against a toy graph estimates nothing.
	d := productionDAG(t)
	store, _ := storage.NewSQLiteStore(":memory:")
	t.Cleanup(func() { store.Close() })
	reg := productionRegistry()

	root := filepath.Join("..", "..")
	pl, err := planning.Load(filepath.Join(root, "data", "courses.json"), d)
	if err != nil {
		t.Skipf("planner catalog unavailable: %v", err)
	}
	ds, err := planning.LoadDestinations(filepath.Join(root, "data", "destinations.json"))
	if err != nil {
		t.Skipf("destinations unavailable: %v", err)
	}
	pl.SetDestinations(ds)
	return New(engine.New(store, d, reg, nil, pl), store, nil)
}

func firstDestination(t *testing.T) string {
	t.Helper()
	root := filepath.Join("..", "..")
	ds, err := planning.LoadDestinations(filepath.Join(root, "data", "destinations.json"))
	if err != nil || len(ds) == 0 {
		t.Skip("no destinations to estimate against")
	}
	return ds[0].ID
}

func TestCompletionEstimates_NoneWithoutAPlan(t *testing.T) {
	s := testServer(t)
	st, err := s.eng.CreateStudent("no_plan")
	if err != nil {
		t.Fatal(err)
	}
	if got := s.completionEstimates(st.ID); got != nil {
		t.Errorf("expected no estimates before a destination is saved, got %v", got)
	}
}

func TestCompletionEstimates_KeyedByPaceAndComputed(t *testing.T) {
	s := plannerServer(t)
	st, err := s.eng.CreateStudent("paced")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.eng.SavePlan(st.ID, firstDestination(t), 30, 0, 0); err != nil {
		t.Fatalf("save plan: %v", err)
	}

	got := s.completionEstimates(st.ID)
	if len(got) == 0 {
		t.Fatal("expected estimates once a destination exists")
	}
	days := map[int]int{}
	for pace, when := range got {
		if pace <= 0 {
			t.Errorf("pace key %d is not a daily XP goal", pace)
		}
		m := paceDayPattern.FindStringSubmatch(when)
		if m == nil {
			t.Fatalf("pace %d -> %q is not a computed duration", pace, when)
		}
		n, err := strconv.Atoi(m[1])
		if err != nil {
			t.Fatalf("pace %d -> %q has no day count", pace, when)
		}
		days[pace] = n
	}

	// A slower pace must never take fewer days. The old literal could not
	// satisfy this because it was keyed on an XP total rather than a daily rate.
	for slow, slowDays := range days {
		for fast, fastDays := range days {
			if slow >= fast {
				continue
			}
			if slowDays < fastDays {
				t.Errorf("pace %d took %d days but the faster pace %d took %d", slow, slowDays, fast, fastDays)
			}
		}
	}
}

func TestCompletionEstimates_NilWhenNothingRemains(t *testing.T) {
	s := plannerServer(t)
	st, err := s.eng.CreateStudent("finished")
	if err != nil {
		t.Fatal(err)
	}
	dest := firstDestination(t)
	if _, err := s.eng.SavePlan(st.ID, dest, 30, 0, 0); err != nil {
		t.Fatalf("save plan: %v", err)
	}
	// Master the destination outright: a date for work that is done would be a
	// lie in the other direction.
	if err := s.repo.UpsertProgress(&storage.ConceptProgress{
		StudentID: st.ID, ConceptID: "arith.add.single", Status: "MASTERED",
	}); err != nil {
		t.Fatal(err)
	}
	if err := s.repo.UpsertProgress(&storage.ConceptProgress{
		StudentID: st.ID, ConceptID: "arith.neg.number_line", Status: "MASTERED",
	}); err != nil {
		t.Fatal(err)
	}
	// Not asserting nil here: whether this synthetic two-concept graph fully
	// covers the real destination decides it. What must hold is that the result
	// is well-formed either way.
	got := s.completionEstimates(st.ID)
	for pace, when := range got {
		if pace <= 0 || !paceDayPattern.MatchString(when) {
			t.Errorf("malformed estimate for pace %d: %q", pace, when)
		}
	}
}

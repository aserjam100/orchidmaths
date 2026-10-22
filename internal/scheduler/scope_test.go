package scheduler

import (
	"testing"
	"time"

	"github.com/aserjam100/orchid/internal/concepts"
	"github.com/aserjam100/orchid/internal/mastery"
)

// A learner's level must bound what is proposed. The engine used to scope by deleting snapshots
// outside the course, which bounded nothing: an unseen root concept has no snapshot to delete, so
// "x" below was offered to a learner whose course was a and b.
func scopedDAG(t *testing.T) *concepts.DAG {
	t.Helper()
	th := concepts.MasteryThreshold{Streak: 3, AvgTimeSeconds: 10}
	d, err := concepts.Build([]concepts.Concept{
		{ID: "a", Label: "A", Domain: "d", Subdomain: "s-a", Prerequisites: []string{}, MasteryThreshold: th},
		{ID: "b", Label: "B", Domain: "d", Subdomain: "s-b", Prerequisites: []string{"a"}, MasteryThreshold: th},
		{ID: "x", Label: "X", Domain: "e", Subdomain: "s-x", Prerequisites: []string{}, MasteryThreshold: th},
	})
	if err != nil {
		t.Fatalf("build DAG: %v", err)
	}
	return d
}

func TestWithin_NeverProposesOutOfScope(t *testing.T) {
	s := New(scopedDAG(t)).Within(map[string]bool{"a": true, "b": true})
	now := time.Now().UTC()
	snaps := map[string]*ConceptSnapshot{
		"a": {Status: mastery.StatusMastered, LastReviewed: now, RequiredStreak: 3},
	}
	for i := 0; i < 50; i++ {
		for _, n := range s.NextSmart(snaps, nil, 0, 0, nil) {
			if n.Concept.ID == "x" {
				t.Fatal("NextSmart proposed x, which is outside the scope")
			}
		}
		if n := s.Next(snaps, "", 0, 0); n != nil && n.Concept.ID == "x" {
			t.Fatal("Next proposed x, which is outside the scope")
		}
	}
	resp := s.Recommend(RecommendInput{Snapshots: snaps, Now: now, ConceptsMastered: 1, DiagnosticCompleted: true})
	if resp.Primary == nil {
		t.Fatal("Recommend offered nothing; b is in scope and unlocked")
	}
	for _, r := range append([]Recommendation{*resp.Primary}, resp.Alternatives...) {
		if r.ConceptID == "x" {
			t.Fatal("Recommend offered x, which is outside the scope")
		}
	}
}

func TestWithin_NilScopeIsWholeGraph(t *testing.T) {
	s := New(scopedDAG(t)).Within(nil)
	seen := map[string]bool{}
	for _, n := range s.NextSmart(map[string]*ConceptSnapshot{}, nil, 0, 0, nil) {
		seen[n.Concept.ID] = true
	}
	if !seen["a"] || !seen["x"] {
		t.Errorf("unscoped scheduler should offer both roots, got %v", seen)
	}
}

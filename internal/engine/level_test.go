package engine

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aserjam100/orchid/internal/concepts"
	"github.com/aserjam100/orchid/internal/generator"
	"github.com/aserjam100/orchid/internal/planning"
	"github.com/aserjam100/orchid/internal/storage"
)

// levelEngine has two Primary courses and one concept, x, that belongs to neither: the shape of
// the real corpus, where most concepts are beyond Primary.
func levelEngine(t *testing.T) (*Engine, storage.Repository) {
	t.Helper()
	th := concepts.MasteryThreshold{Streak: 3, AvgTimeSeconds: 60}
	d, err := concepts.Build([]concepts.Concept{
		{ID: "a", Label: "A", Domain: "d", GradingType: "numeric", Prerequisites: []string{}, MasteryThreshold: th},
		{ID: "b", Label: "B", Domain: "d", GradingType: "numeric", Prerequisites: []string{"a"}, MasteryThreshold: th},
		{ID: "x", Label: "X", Domain: "e", GradingType: "numeric", Prerequisites: []string{}, MasteryThreshold: th},
	})
	if err != nil {
		t.Fatalf("build DAG: %v", err)
	}
	courses := filepath.Join(t.TempDir(), "courses.json")
	if err := os.WriteFile(courses, []byte(`[
		{"id": "sg-p1", "name": "P1", "grade": "P1", "targets": ["a"]},
		{"id": "sg-p6", "name": "P6", "grade": "P6", "targets": ["b"]}
	]`), 0o600); err != nil {
		t.Fatal(err)
	}
	planner, err := planning.Load(courses, d)
	if err != nil {
		t.Fatalf("load planner: %v", err)
	}
	store, err := storage.NewSQLiteStore(":memory:")
	if err != nil {
		t.Fatalf("create store: %v", err)
	}
	t.Cleanup(func() { store.Close() })
	reg := generator.NewRegistry()
	for _, id := range []string{"a", "b", "x"} {
		reg.Register(id, &testGen{answer: "1"})
	}
	return New(store, d, reg, nil, planner), store
}

func TestLevelScope_DefaultsToAllOfPrimary(t *testing.T) {
	e, _ := levelEngine(t)
	scope := e.LevelScope("")
	if !scope["a"] || !scope["b"] || scope["x"] {
		t.Errorf("default scope should be sg-p6 with its prerequisite, got %v", scope)
	}
}

func TestLevelScope_FollowsTheChosenCourse(t *testing.T) {
	e, store := levelEngine(t)
	st, err := store.CreateStudent("Learner")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := e.SetActiveCourse(st.ID, "sg-p1"); err != nil {
		t.Fatal(err)
	}
	scope := e.LevelScope(st.ID)
	if !scope["a"] || scope["b"] || scope["x"] {
		t.Errorf("a P1 learner should be scoped to a only, got %v", scope)
	}
}

func TestRecommendNext_NeverOffersBeyondPrimary(t *testing.T) {
	e, store := levelEngine(t)
	st, err := store.CreateStudent("Learner")
	if err != nil {
		t.Fatal(err)
	}
	for _, sid := range []string{"", st.ID} {
		resp, err := e.RecommendNext(sid, nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Primary != nil && resp.Primary.ConceptID == "x" {
			t.Errorf("student %q was offered x, which is in no Primary course", sid)
		}
		for _, r := range resp.Alternatives {
			if r.ConceptID == "x" {
				t.Errorf("student %q was offered x as an alternative", sid)
			}
		}
	}
}

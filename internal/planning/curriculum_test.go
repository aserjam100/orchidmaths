package planning

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/aserjam100/orchid/internal/concepts"
)

// The MOE Primary mapping is what lets Orchid say it is mapped to the Singapore syllabus, so it
// has to stay true as the concept graph moves. These checks hold the mapping to the graph and the
// courses to the mapping; `scripts/curriculum_moe.py --check` holds the generated files to both.

type moeOutcome struct {
	ID       string   `json:"id"`
	Level    string   `json:"level"`
	Outcome  string   `json:"outcome"`
	Coverage string   `json:"coverage"`
	Concepts []string `json:"concepts"`
	Note     string   `json:"note"`
}

type moeMapping struct {
	Disclaimer string       `json:"disclaimer"`
	Outcomes   []moeOutcome `json:"outcomes"`
}

func loadMOE(t *testing.T) (*moeMapping, *concepts.DAG, *Planner) {
	t.Helper()
	raw, err := os.ReadFile("../../data/curricula/sg-moe-primary-2021.json")
	if err != nil {
		t.Skipf("curriculum mapping not present: %v", err)
	}
	var m moeMapping
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("parse mapping: %v", err)
	}
	dag, err := concepts.LoadDir("../../data/concepts")
	if err != nil {
		t.Fatalf("load concepts: %v", err)
	}
	p, err := Load("../../data/courses.json", dag)
	if err != nil {
		t.Fatalf("load courses: %v", err)
	}
	return &m, dag, p
}

func TestMOEMappingReferencesRealConcepts(t *testing.T) {
	m, dag, _ := loadMOE(t)
	if !strings.Contains(m.Disclaimer, "not affiliated with or endorsed by") {
		t.Errorf("mapping must carry the non-affiliation disclaimer, got %q", m.Disclaimer)
	}
	seen := map[string]bool{}
	for _, o := range m.Outcomes {
		if seen[o.ID] {
			t.Errorf("duplicate outcome id %s", o.ID)
		}
		seen[o.ID] = true
		if !strings.HasPrefix(o.ID, o.Level+".") {
			t.Errorf("%s: id does not start with its level %s", o.ID, o.Level)
		}
		switch o.Coverage {
		case "none":
			if len(o.Concepts) != 0 {
				t.Errorf("%s: coverage none but lists concepts %v", o.ID, o.Concepts)
			}
		case "full", "partial":
			if len(o.Concepts) == 0 {
				t.Errorf("%s: coverage %s needs at least one concept", o.ID, o.Coverage)
			}
		default:
			t.Errorf("%s: unknown coverage %q", o.ID, o.Coverage)
		}
		for _, c := range o.Concepts {
			if dag.Concept(c) == nil {
				t.Errorf("%s: concept %s is not in the graph", o.ID, c)
			}
		}
	}
}

// A course may only target what the mapping vouches for at its level or earlier: anything else
// would be practice Orchid presents as MOE-aligned without a syllabus outcome behind it.
func TestMOECoursesTargetOnlyMappedConcepts(t *testing.T) {
	m, _, p := loadMOE(t)
	levels := []string{"P1", "P2", "P3", "P4", "P5", "P6"}
	for i, lvl := range levels {
		allowed := map[string]bool{}
		for _, o := range m.Outcomes {
			for j := 0; j <= i; j++ {
				if o.Level == levels[j] {
					for _, c := range o.Concepts {
						allowed[c] = true
					}
				}
			}
		}
		id := "sg-" + strings.ToLower(lvl)
		course := p.Course(id)
		if course == nil {
			t.Errorf("course %s missing from courses.json", id)
			continue
		}
		if len(course.Targets) == 0 {
			t.Errorf("course %s has no targets", id)
		}
		for _, c := range course.Targets {
			if !allowed[c] {
				t.Errorf("course %s targets %s, which no outcome at %s or earlier maps", id, c, lvl)
			}
		}
		if _, err := p.PathForCourse(id); err != nil {
			t.Errorf("course %s: %v", id, err)
		}
	}
}

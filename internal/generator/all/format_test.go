package all

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"testing"

	"github.com/aserjam100/orchid/internal/generator"
)

// numericKeypadAnswer is what the numeric input mode can actually express:
// digits and the punctuation that survives a phone keypad or a desktop
// keyboard. A letter in the answer under this mode means the learner is looking
// at a field that cannot type the right answer, which is what "Simplify: √125"
// and "give answer with remainder: Q R" were.
var numericKeypadAnswer = regexp.MustCompile(`^[\d\s.,/+\-]+$`)

// TestAnswerFormatMatchesTheAnswer is the guard that turns a pile of reports
// into a list. For every concept, it asks the generator (and, failing that, the
// grading type) what input the learner will be given, then generates real
// questions and checks that each answer is expressible in that input.
//
// It is deliberately a property rather than a table of known exceptions: a new
// generator that answers "2√3" under a numeric concept fails here on the day it
// is written, not on the day a learner reports its question as broken
// formatting.
func TestAnswerFormatMatchesTheAnswer(t *testing.T) {
	reg := generator.NewRegistry()
	Register(reg)

	files, err := filepath.Glob("../../../data/concepts/*.json")
	if err != nil || len(files) == 0 {
		t.Fatalf("no concept files found: %v", err)
	}
	type concept struct {
		ID          string `json:"id"`
		GradingType string `json:"grading_type"`
	}
	var all []concept
	types := map[string]string{}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		var list []concept
		if err := json.Unmarshal(b, &list); err != nil {
			t.Fatalf("parse %s: %v", f, err)
		}
		for _, c := range list {
			if c.ID == "" {
				continue
			}
			all = append(all, c)
			if _, seen := types[c.ID]; !seen {
				types[c.ID] = c.GradingType
			}
		}
	}
	if len(all) < 1000 {
		t.Fatalf("only %d concepts loaded; the guard would pass vacuously", len(all))
	}

	bad := map[string]string{}
	for _, c := range all {
		gen, err := reg.Get(c.ID)
		if err != nil {
			continue
		}
		f := generator.AnswerFormatFor(gen, types[c.ID])
		if f.InputMode != "numeric" {
			continue
		}
		for _, diff := range []float64{0.3, 0.5, 0.8} {
			problems, err := reg.BatchGenerateContext(c.ID, 4, generator.GeneratorContext{Difficulty: diff})
			if err != nil {
				continue
			}
			for _, p := range problems {
				if !numericKeypadAnswer.MatchString(p.Answer) {
					bad[c.ID] = p.Answer
					break
				}
			}
		}
	}

	ids := make([]string, 0, len(bad))
	for id := range bad {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		t.Errorf("%s: numeric input, but the answer is %q", id, bad[id])
	}
	if len(ids) == 0 {
		t.Logf("checked %d concepts: every numeric-input answer is typeable", len(all))
	}
}

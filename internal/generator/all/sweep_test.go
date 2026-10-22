package all

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/aserjam100/orchid/internal/generator"
)

// What can a learner actually practise?
//
// The practice endpoint returns 503 "no questions available for this concept" when the
// registry has no generator for an id, or when the generator produces nothing at all. Both
// are invisible until someone opens that concept, and the corpus and the registry are two
// lists that are only kept in step by hand. This sweep compares them: it reads every
// concept id in data/concepts and asks the registry for one question at the difficulty the
// Learn feed starts from.
//
// It deliberately reports "no generator" and "generator produced nothing" as two
// categories, because they are different repairs: the first is a missing registration, the
// second is a broken or unsatisfiable generator.
func TestSweepEveryConceptCanGenerateAQuestion(t *testing.T) {
	reg := generator.NewRegistry()
	Register(reg)

	files, err := filepath.Glob("../../../data/concepts/*.json")
	if err != nil || len(files) == 0 {
		t.Fatalf("no concept files found: %v", err)
	}
	type concept struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	}
	var all []concept
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		var list []concept
		if err := json.Unmarshal(b, &list); err != nil {
			t.Fatalf("parse %s: %v", f, err)
		}
		all = append(all, list...)
	}
	if len(all) < 1000 {
		t.Fatalf("only %d concepts loaded; the sweep would pass vacuously", len(all))
	}

	var missing, empty []string
	// The Learn feed raises difficulty on a correct answer and lowers it on a miss, so a
	// generator that only works at one setting is a dead end in the middle of a session.
	for _, diff := range []float64{0.3, 0.5, 0.8, 1.0} {
		for _, c := range all {
			if c.ID == "" {
				continue
			}
			problems, err := reg.BatchGenerateContext(c.ID, 3, generator.GeneratorContext{Difficulty: diff})
			if err != nil {
				missing = append(missing, fmt.Sprintf("%.1f %s: %v", diff, c.ID, err))
				continue
			}
			if len(problems) == 0 {
				empty = append(empty, fmt.Sprintf("%.1f %s", diff, c.ID))
			}
		}
	}
	sort.Strings(missing)
	sort.Strings(empty)
	t.Logf("concepts=%d no_generator=%d produced_nothing=%d", len(all), len(missing), len(empty))
	for _, m := range missing {
		t.Logf("NO GENERATOR  %s", m)
	}
	for _, e := range empty {
		t.Logf("PRODUCED NOTHING  %s", e)
	}
	fmt.Printf("sweep: %d concepts, %d without a generator, %d producing nothing\n", len(all), len(missing), len(empty))
}

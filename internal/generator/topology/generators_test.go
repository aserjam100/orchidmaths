package topology

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/aserjam100/orchid/internal/generator"
)

func fuzzGen(t *testing.T, gen generator.Generator) {
	t.Helper()
	for i := 0; i < 100; i++ {
		d := rand.Float64()
		p := gen.Generate(generator.GeneratorContext{Difficulty: d})
		if p.Question == "" || p.Answer == "" || p.Explanation == "" {
			t.Errorf("empty field at difficulty=%.2f: q=%q a=%q e=%q", d, p.Question, p.Answer, p.Explanation)
		}
	}
}

func TestMetricGen(t *testing.T)     { fuzzGen(t, &metricGen{}) }
func TestOpenClosedGen(t *testing.T) { fuzzGen(t, &openClosedGen{}) }
func TestContinuousGen(t *testing.T) { fuzzGen(t, &continuousGen{}) }

// dimensionGen is the one topology concept that produces a number rather than
// recognising a statement, so it is checked for the shape a produced answer
// needs: a numeric answer and the facts an authored schema rests on.
func TestDimensionGen(t *testing.T) {
	gen := &dimensionGen{}
	seen := map[string]bool{}
	for i := 0; i < 300; i++ {
		p := gen.Generate(generator.GeneratorContext{Difficulty: rand.Float64()})
		if _, err := strconv.Atoi(p.Answer); err != nil {
			t.Fatalf("answer %q is not a number", p.Answer)
		}
		for _, k := range []string{"space", "dim", "rule"} {
			if p.Facts[k] == "" {
				t.Fatalf("fact %q not published on %+v", k, p.Facts)
			}
		}
		if p.Facts["dim"] != p.Answer {
			t.Fatalf("dim fact %q disagrees with answer %q", p.Facts["dim"], p.Answer)
		}
		seen[p.Answer] = true
	}
	if len(seen) < 2 {
		t.Fatalf("only one distinct answer %v; the pool does not discriminate", seen)
	}
}

func TestFuzz(t *testing.T) {
	reg := generator.NewRegistry()
	Register(reg)
	for _, id := range reg.Concepts() {
		gen, _ := reg.Get(id)
		t.Run(id, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				d := rand.Float64()
				p := gen.Generate(generator.GeneratorContext{Difficulty: d})
				if p.Question == "" || p.Answer == "" || p.Explanation == "" {
					t.Fatalf("empty field at difficulty=%.2f for %s: q=%q a=%q", d, id, p.Question, p.Answer)
				}
			}
		})
	}
}

package arithmetic

import (
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/aserjam100/orchid/internal/generator"
)

// The MOE syllabus scopes each Primary 1 and 2 outcome to a number range. These tests pin
// the generators that the sg-p1 and sg-p2 courses use to those ranges, because a generator
// that drifts past them still produces valid questions and no other test would notice.

func registered(t *testing.T, id string) generator.Generator {
	t.Helper()
	reg := generator.NewRegistry()
	Register(reg)
	g, err := reg.Get(id)
	if err != nil {
		t.Fatalf("%s is not registered: %v", id, err)
	}
	return g
}

func fact(t *testing.T, p generator.Problem, key string) int {
	t.Helper()
	n, err := strconv.Atoi(p.Facts[key])
	if err != nil {
		t.Fatalf("fact %q = %q: %v", key, p.Facts[key], err)
	}
	return n
}

func sample(t *testing.T, id string, check func(p generator.Problem)) {
	t.Helper()
	g := registered(t, id)
	for i := 0; i < 2000; i++ {
		check(g.Generate(generator.GeneratorContext{Difficulty: rand.Float64()}))
	}
}

func TestPrimaryOneAdditionStaysWithin100(t *testing.T) {
	for _, id := range []string{"arith.add.double", "arith.add.carry"} {
		sample(t, id, func(p generator.Problem) {
			if s := fact(t, p, "sum"); s > 100 {
				t.Fatalf("%s: sum %d is above 100: %s", id, s, p.Question)
			}
		})
	}
	sample(t, "arith.add.triple", func(p generator.Problem) {
		if s := fact(t, p, "sum"); s > 1000 {
			t.Fatalf("arith.add.triple: sum %d is above 1000: %s", s, p.Question)
		}
	})
	sample(t, "arith.add.carry", func(p generator.Problem) {
		if fact(t, p, "ones_a")+fact(t, p, "ones_b") < 10 {
			t.Fatalf("arith.add.carry without a carry: %s", p.Question)
		}
	})
	sample(t, "arith.sub.word", func(p generator.Problem) {
		if a := fact(t, p, "a"); a > 99 {
			t.Fatalf("arith.sub.word starts at %d: %s", a, p.Question)
		}
	})
}

func TestPrimaryOneMultiplyWithin40DivideWithin20(t *testing.T) {
	sample(t, "arith.mult.concept", func(p generator.Problem) {
		if n := fact(t, p, "total"); n > 40 {
			t.Fatalf("product %d is above 40: %s", n, p.Question)
		}
	})
	sample(t, "arith.div.concept", func(p generator.Problem) {
		if n := fact(t, p, "total"); n > 20 {
			t.Fatalf("total %d is above 20: %s", n, p.Question)
		}
	})
}

func TestPrimaryTwoStaysInsideTheTables(t *testing.T) {
	tables := map[string][]int{
		"arith.mult.2_5_10": {2, 5, 10},
		"arith.mult.3_4":    {3, 4},
	}
	for id, factors := range tables {
		sample(t, id, func(p generator.Problem) {
			if f, n := fact(t, p, "factor"), fact(t, p, "n"); !slices.Contains(factors, f) || n < 1 || n > 10 {
				t.Fatalf("%s: %d × %d is outside the tables %v", id, f, n, factors)
			}
		})
	}
	for _, id := range []string{"arith.div.tables", "arith.div.word_tables"} {
		sample(t, id, func(p generator.Problem) {
			if d, q := fact(t, p, "divisor"), fact(t, p, "quotient"); !slices.Contains(primaryTables, d) || q < 1 || q > 10 {
				t.Fatalf("%s: divisor %d quotient %d is outside the tables", id, d, q)
			}
		})
	}
	sample(t, "arith.mult.word_tables", func(p generator.Problem) {
		if g, n := fact(t, p, "groups"), fact(t, p, "per_group"); !slices.Contains(primaryTables, n) || g < 1 || g > 10 {
			t.Fatalf("%d groups of %d is outside the tables", g, n)
		}
	})
}

func TestCompareWithinRange(t *testing.T) {
	ranges := map[string][2]int{"arith.compare.within100": {1, 100}, "arith.compare.within1000": {100, 999}}
	for id, r := range ranges {
		sample(t, id, func(p generator.Problem) {
			parts := strings.Split(p.Answer, ",")
			if len(parts) != 3 {
				t.Fatalf("%s: answer %q is not three numbers", id, p.Answer)
			}
			prev := -1
			for _, s := range parts {
				n, err := strconv.Atoi(s)
				if err != nil || n < r[0] || n > r[1] || n <= prev {
					t.Fatalf("%s: answer %q is out of range or not strictly increasing", id, p.Answer)
				}
				prev = n
			}
		})
	}
}

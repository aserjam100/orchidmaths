package diagnostic

import (
	"testing"

	"github.com/aserjam100/orchid/internal/concepts"
)

func TestDomainToCourse_AdvancedPlacements(t *testing.T) {
	cases := []struct {
		domain, subdomain, expected string
	}{
		// Regression: core placements unchanged.
		{"calculus", "calculus.limits", "calc1"},
		{"calculus", "calculus.derivatives", "calc1"},
		{"calculus", "calculus.integrals", "calc1"},
		{"discrete_math", "discrete_math.logic", "discrete"},
		{"abstract_algebra", "abstract_algebra.groups", "abstract"},
		{"number_theory", "number_theory.basics", "nt"},
		// Advanced subdomains place into the *2 courses.
		{"calculus", "calculus.series", "calc2"},
		{"calculus", "calculus.sequences", "calc2"},
		{"discrete_math", "discrete_math.graph_theory", "discrete2"},
		{"discrete_math", "discrete_math.proof", "discrete2"},
		{"discrete_math", "discrete_math.recurrence", "discrete2"},
		{"discrete_math", "discrete_math.combinatorics", "discrete"},
		{"discrete_math", "discrete_math.sets", "discrete"},
		{"abstract_algebra", "abstract_algebra.rings", "abstract2"},
		{"abstract_algebra", "abstract_algebra.fields", "abstract2"},
		{"abstract_algebra", "abstract_algebra.structures", "abstract2"},
		// Domains whose course existed but had no placement mapping.
		{"real_analysis", "real_analysis.seq", "ra"},
		{"graph_theory", "graph_theory.basics", "graph"},
		{"numerical_analysis", "numerical_analysis.basics", "numa"},
		{"probability_theory", "probability_theory.basics", "prob"},
	}
	for _, tc := range cases {
		c := &concepts.Concept{Domain: tc.domain, Subdomain: tc.subdomain}
		if got := domainToCourse(c); got != tc.expected {
			t.Errorf("%s/%s: expected %q, got %q", tc.domain, tc.subdomain, tc.expected, got)
		}
	}
}

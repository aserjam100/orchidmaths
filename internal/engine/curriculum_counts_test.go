package engine

import (
	"testing"

	"github.com/aserjam100/orchid/internal/concepts"
	"github.com/aserjam100/orchid/internal/generator"
	"github.com/aserjam100/orchid/internal/storage"
)

// TestCurriculumDomainsCountsEveryConcept pins the number `/domains` prints.
//
// The handler sums per-domain counters, and CurriculumDomains skips a concept
// whose domain has no row (`if d == nil { continue }`). That skip is defensive,
// and its failure mode is a total that is quietly too small while every row
// still looks plausible — so the sum is checked against the corpus the DAG was
// built from, not against itself.
func TestCurriculumDomainsCountsEveryConcept(t *testing.T) {
	dag, err := concepts.LoadDir("../../data/concepts")
	if err != nil {
		t.Fatalf("load corpus: %v", err)
	}
	store, err := storage.NewSQLiteStore(":memory:")
	if err != nil {
		t.Fatalf("create store: %v", err)
	}
	t.Cleanup(func() { store.Close() })
	e := New(store, dag, generator.NewRegistry(), nil, nil)

	rows, err := e.CurriculumDomains("")
	if err != nil {
		t.Fatalf("curriculum domains: %v", err)
	}
	total := 0
	for _, row := range rows {
		total += row.ConceptCount
	}
	if want := len(dag.Order()); total != want {
		t.Fatalf("domains account for %d concepts, corpus has %d", total, want)
	}
	if want := len(dag.Domains()); len(rows) != want {
		t.Fatalf("returned %d domains, corpus has %d", len(rows), want)
	}
}

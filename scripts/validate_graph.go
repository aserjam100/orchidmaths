package main

import (
	"fmt"
	"os"

	"github.com/aserjam100/orchid/internal/concepts"
)

func main() {
	dag, err := concepts.LoadDir("data/concepts")
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK: %d concepts across %d domains\n", dag.Count(), len(dag.Domains()))
}

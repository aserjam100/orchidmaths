package main

// Merge the per-domain concept files into the copy the web app imports
// (`web/next-app/data/concepts.json`, typed by `web/next-app/lib/conceptData.ts`).
//
// `data/concepts/enrichment.json` is skipped on purpose. It is an overlay keyed by
// id -- encompasses, key_prerequisites, interference_group, variants -- and the
// loader applies it to the domain nodes rather than treating its entries as nodes,
// so including it would put ten id-only records into a file the browser reads as
// the graph. Before this skip the merge could never have been re-run without
// adding them; the copy in the tree was a snapshot taken before the file existed
// and it had quietly drifted 29 `grading_type` corrections behind the corpus.
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// overlayFiles are id-keyed overlays, not concept lists.
var overlayFiles = map[string]bool{"enrichment.json": true}

// creditToken names the licence family whose attribution has to reach the page. It is read from
// `data/lessons/mapping.json` (the ingest map) and kept only for concepts that have a lesson, because
// a token for a lesson nobody can open would ask a page to credit text it never shows.
//
// ORCCA is the reason the field exists. Its Licensing section requires a fixed sentence on "every
// digital format page view" of its content, which a browser cannot read out of the corpus, so the
// merge carries the token and `components/LessonCredits.tsx` holds the sentence. OpenStax and
// Algebrica state their terms in prose rather than a fixed sentence, and their credit is already the
// lesson body's first line, so they need no token.
var creditTokens = map[string]string{"orcca": "orcca"}

type lessonEntry struct {
	ConceptID string `json:"concept_id"`
	Source    string `json:"source"`
}

type mapEntry struct {
	Source string `json:"source"`
}

func readJSON(path string, v any) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v\n", path, err)
		os.Exit(1)
	}
	if err := json.Unmarshal(data, v); err != nil {
		fmt.Fprintf(os.Stderr, "parse %s: %v\n", path, err)
		os.Exit(1)
	}
}

// credits returns concept id -> credit token, for lessons that exist and whose publisher asks for a
// per-page-view attribution.
func credits() map[string]string {
	var lessons []lessonEntry
	readJSON("data/lessons/lessons.json", &lessons)
	taught := make(map[string]bool, len(lessons))
	for _, l := range lessons {
		taught[l.ConceptID] = true
	}
	var mapping map[string]mapEntry
	readJSON("data/lessons/mapping.json", &mapping)
	out := map[string]string{}
	for id, m := range mapping {
		if token, ok := creditTokens[m.Source]; ok && taught[id] {
			out[id] = token
		}
	}
	return out
}

func main() {
	entries, err := os.ReadDir("data/concepts")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read dir: %v\n", err)
		os.Exit(1)
	}
	byConcept := credits()
	var all []any
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") || overlayFiles[e.Name()] {
			continue
		}
		data, err := os.ReadFile(filepath.Join("data/concepts", e.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", e.Name(), err)
			os.Exit(1)
		}
		var batch []map[string]any
		if err := json.Unmarshal(data, &batch); err != nil {
			fmt.Fprintf(os.Stderr, "parse %s: %v\n", e.Name(), err)
			os.Exit(1)
		}
		for _, c := range batch {
			if id, ok := c["id"].(string); ok {
				if token, ok := byConcept[id]; ok {
					c["lesson_credit"] = token
				} else {
					delete(c, "lesson_credit")
				}
			}
			all = append(all, c)
		}
	}
	out, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile("web/next-app/data/concepts.json", out, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK: %d concepts merged into web/next-app/data/concepts.json, %d with a per-page credit\n",
		len(all), len(byConcept))
}

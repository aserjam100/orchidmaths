package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// TestDockerfileCopiesEveryDataFileTheServerLoads pins the deployment bug that
// made the goal-scope agent answer "no_candidates" in production while it
// worked locally: main.go loads data/destinations.json, the Dockerfile never
// copied it, and nothing noticed because a missing data file is a warning rather
// than a startup failure. The same omission also left the destination estimates
// degraded, which had nothing to do with AI.
//
// The list of files is read from main.go rather than written out here, so adding
// a dataPath call without adding the matching COPY fails in this test instead of
// on the deployed machine. A hand-maintained list would drift the same way the
// Dockerfile did.
func TestDockerfileCopiesEveryDataFileTheServerLoads(t *testing.T) {
	src, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	calls := regexp.MustCompile(`dataPath\("([^"]+)"\)`).FindAllStringSubmatch(string(src), -1)
	if len(calls) == 0 {
		t.Fatal("no dataPath calls found; this test is reading the wrong source")
	}

	dockerfile, err := os.ReadFile(filepath.Join("..", "..", "Dockerfile"))
	if err != nil {
		t.Fatalf("read Dockerfile: %v", err)
	}
	copied := make(map[string]bool)
	for _, m := range regexp.MustCompile(`(?m)^COPY data/(\S+)\s+/app/data/`).FindAllStringSubmatch(string(dockerfile), -1) {
		copied[strings.TrimSuffix(m[1], "/")] = true
	}

	for _, m := range calls {
		name := m[1]
		if name == "" {
			continue
		}
		if !copied[name] {
			t.Errorf("main.go loads data/%s but the Dockerfile never copies it into /app/data; "+
				"the deployed server will warn and carry on without it", name)
		}
	}
}

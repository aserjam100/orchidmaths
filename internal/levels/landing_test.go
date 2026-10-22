package levels

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// The landing page derived its level ranges from the corpus size, nine equal slices of 1250,
// while Compute awards titles at fixed steps of 32. The page put "Math Architect" at 1112
// mastered concepts; the server awarded it at 256. The check is on the generated artifact,
// because app/home/data.ts renders graphMeta.levels.
func TestLandingLevelsMatchServer(t *testing.T) {
	path := filepath.Join("..", "..", "web", "next-app", "data", "graph.meta.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v (run `npm run graph:build`)", path, err)
	}
	var meta struct {
		Levels []struct {
			Min   int    `json:"min"`
			Title string `json:"title"`
		} `json:"levels"`
	}
	if err := json.Unmarshal(raw, &meta); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	fix := "fix: set MASTERY_LEVELS in web/next-app/lib/graphPayload.ts to match levels.All and run `npm run graph:build`"
	if len(meta.Levels) != len(All) {
		t.Fatalf("landing page shows %d levels, the server has %d\n%s", len(meta.Levels), len(All), fix)
	}
	for i, l := range All {
		if meta.Levels[i].Min != l.Min || meta.Levels[i].Title != l.Title {
			t.Errorf("level %d: landing page shows %q from %d, the server awards %q from %d\n%s",
				i+1, meta.Levels[i].Title, meta.Levels[i].Min, l.Title, l.Min, fix)
		}
	}
}

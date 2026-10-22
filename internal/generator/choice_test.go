package generator

import (
	"regexp"
	"strings"
	"testing"
)

var optionLine = regexp.MustCompile(`\n\n([A-E])\) ([^\n]*)`)

func parseOptions(t *testing.T, q string) []string {
	t.Helper()
	var out []string
	for _, m := range optionLine.FindAllStringSubmatch(q, -1) {
		out = append(out, m[2])
	}
	return out
}

func TestMultipleChoiceBuildsConsecutiveOptions(t *testing.T) {
	p := MultipleChoice("Which form is it?", "modus ponens", "modus tollens", "affirming the consequent")
	opts := parseOptions(t, p.Question)
	if len(opts) != 3 {
		t.Fatalf("expected 3 options, got %d: %q", len(opts), p.Question)
	}
	letters := optionLine.FindAllStringSubmatch(p.Question, -1)
	for i, m := range letters {
		want := string(rune('A' + i))
		if m[1] != want {
			t.Errorf("option %d has letter %q, want %q", i, m[1], want)
		}
	}
	if p.Answer != "modus ponens" {
		t.Errorf("answer = %q, want the correct option's text", p.Answer)
	}
	count := 0
	for _, o := range opts {
		if o == p.Answer {
			count++
		}
	}
	if count != 1 {
		t.Errorf("the answer appears %d times among the options, want exactly once", count)
	}
}

func TestMultipleChoiceDropsDuplicateDistractors(t *testing.T) {
	p := MultipleChoice("Q", "yes", "yes", "no", " no ")
	opts := parseOptions(t, p.Question)
	if len(opts) != 2 {
		t.Fatalf("expected the duplicate distractors to collapse, got %v", opts)
	}
	if strings.Join(opts, "|") != "yes|no" && strings.Join(opts, "|") != "no|yes" {
		t.Errorf("unexpected options %v", opts)
	}
}

func TestMultipleChoiceKeepsTheAnswerWhenThereAreTooManyOptions(t *testing.T) {
	p := MultipleChoice("Q", "correct", "d1", "d2", "d3", "d4", "d5", "d6")
	opts := parseOptions(t, p.Question)
	if len(opts) != 5 {
		t.Fatalf("expected five options, got %d", len(opts))
	}
	found := false
	for _, o := range opts {
		if o == p.Answer {
			found = true
		}
	}
	if !found {
		t.Errorf("the correct answer was dropped from %v", opts)
	}
}

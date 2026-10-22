package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aserjam100/orchid/internal/concepts"
	"github.com/aserjam100/orchid/internal/engine"
	"github.com/aserjam100/orchid/internal/generator"
	"github.com/aserjam100/orchid/internal/storage"
)

// formattedServer is testServer with a generator that declares the syntax of
// its own answer, so the lab can be checked against a concept whose served
// format differs from its grading type. Re-registering "a" on a server that
// already knows it is refused, hence a fresh server rather than a swap.
func formattedServer(t *testing.T) *Server {
	t.Helper()
	d, _ := concepts.Build([]concepts.Concept{
		{ID: "a", Label: "A", Domain: "d", GradingType: "numeric", Prerequisites: []string{},
			MasteryThreshold: concepts.MasteryThreshold{Streak: 3, AvgTimeSeconds: 60}},
	})
	store, _ := storage.NewSQLiteStore(":memory:")
	t.Cleanup(func() { store.Close() })
	reg := generator.NewRegistry()
	reg.Register("a", &formattedTestGen{})
	return New(engine.New(store, d, reg, nil, nil), store, nil)
}

func labPost(t *testing.T, h http.HandlerFunc, path string, body interface{}) (int, map[string]interface{}) {
	t.Helper()
	raw, _ := json.Marshal(body)
	rec := httptest.NewRecorder()
	h(rec, httptest.NewRequest("POST", path, bytes.NewReader(raw)))
	var res map[string]interface{}
	if rec.Body.Len() > 0 {
		_ = json.Unmarshal(rec.Body.Bytes(), &res)
	}
	return rec.Code, res
}

func TestAdminLabGenerate_ServesQuestionsInTheDeclaredFormat(t *testing.T) {
	s := formattedServer(t)
	code, res := labPost(t, s.handleAdminLabGenerate, "/api/admin/lab/generate", map[string]interface{}{
		"concept_id": "a", "count": 2, "difficulty": 0.6,
	})
	if code != 200 {
		t.Fatalf("generate: %d", code)
	}
	if got := res["answer_input_mode"]; got != "text" {
		t.Errorf("answer_input_mode = %v, want the generator's text rather than the numeric default", got)
	}
	if got := res["answer_hint"]; got != "Quotient and remainder, e.g. 5 R 2" {
		t.Errorf("answer_hint = %v, want the generator's declaration", got)
	}
	if got := res["writes"]; got != false {
		t.Errorf("writes = %v, want false", got)
	}
	qs, _ := res["questions"].([]interface{})
	if len(qs) == 0 {
		t.Fatal("no questions generated")
	}
	q, _ := qs[0].(map[string]interface{})
	if q["answer"] != "4" {
		t.Errorf("question answer = %v, want 4 (the expected answer is what makes the lab judgeable)", q["answer"])
	}
	if q["answer_input_mode"] != "text" {
		t.Errorf("per-question answer_input_mode = %v, want text", q["answer_input_mode"])
	}
}

func TestAdminLabGenerate_RefusesAConceptWithNoGenerator(t *testing.T) {
	s := testServer(t)
	code, _ := labPost(t, s.handleAdminLabGenerate, "/api/admin/lab/generate", map[string]interface{}{
		"concept_id": "not.a.concept",
	})
	if code != 404 {
		t.Errorf("unknown concept = %d, want 404", code)
	}
}

func TestAdminLabGrade_ReturnsTheVerdictAndWritesNothing(t *testing.T) {
	s := testServer(t)
	grade := func(answer string) map[string]interface{} {
		t.Helper()
		code, res := labPost(t, s.handleAdminLabGrade, "/api/admin/lab/grade", map[string]interface{}{
			"concept_id": "a", "expected": "4", "answer": answer,
		})
		if code != 200 {
			t.Fatalf("grade %q: %d", answer, code)
		}
		return res
	}
	if got := grade("4")["correct"]; got != true {
		t.Errorf("correct answer graded %v, want true", got)
	}
	if got := grade("5")["correct"]; got != false {
		t.Errorf("wrong answer graded %v, want false", got)
	}
	if got := grade("4")["writes"]; got != false {
		t.Errorf("writes = %v, want false", got)
	}
	if code, _ := labPost(t, s.handleAdminLabGrade, "/api/admin/lab/grade", map[string]interface{}{
		"concept_id": "a", "expected": "", "answer": "4",
	}); code != 400 {
		t.Errorf("empty expected = %d, want 400", code)
	}
}

func TestAdminContent_FlagsOverriddenFormats(t *testing.T) {
	s := formattedServer(t)
	rec := httptest.NewRecorder()
	s.handleAdminContent(rec, httptest.NewRequest("GET", "/api/admin/content?overridden=1", nil))
	if rec.Code != 200 {
		t.Fatalf("content: %d", rec.Code)
	}
	var res struct {
		Concepts []struct {
			ID               string `json:"id"`
			AnswerHint       string `json:"answer_hint"`
			AnswerInputMode  string `json:"answer_input_mode"`
			AnswerOverridden bool   `json:"answer_overridden"`
		} `json:"concepts"`
		Total int `json:"total"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if res.Total != 1 || len(res.Concepts) != 1 {
		t.Fatalf("overridden filter returned %d concepts (total %d), want 1", len(res.Concepts), res.Total)
	}
	c := res.Concepts[0]
	if !c.AnswerOverridden || c.AnswerInputMode != "text" || c.AnswerHint == "" {
		t.Errorf("overridden row = %+v, want the served format and the flag", c)
	}

	// A concept whose generator declares nothing is not in the table.
	plain := testServer(t)
	rec = httptest.NewRecorder()
	plain.handleAdminContent(rec, httptest.NewRequest("GET", "/api/admin/content?overridden=1", nil))
	var plainRes struct {
		Total int `json:"total"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &plainRes); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if plainRes.Total != 0 {
		t.Errorf("plain generator listed %d overrides, want 0", plainRes.Total)
	}
}

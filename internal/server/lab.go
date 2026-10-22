package server

import (
	"net/http"
	"strings"

	"github.com/aserjam100/orchid/internal/generator"
)

// The lab is the admin's reproduction bench: search any concept, generate its
// questions at a chosen difficulty, try an answer with the same field the
// learner gets, and see the same verdict the engine would return — without
// writing a single attempt, XP point or mastery update.
//
// It exists because a moderator could previously see a report but not the
// question as served, and "I cannot reproduce it" is not a resolution. The
// replay in the report queue answers one report; the lab answers the next one
// before it is filed, and covers the concepts nobody has reported yet.

type labQuestion struct {
	Question        string `json:"question"`
	Answer          string `json:"answer"`
	Explanation     string `json:"explanation"`
	AnswerHint      string `json:"answer_hint"`
	AnswerInputMode string `json:"answer_input_mode"`
}

// POST /api/admin/lab/generate
// Body: { "concept_id": "...", "count": 3, "difficulty": 0.5 }
//
// The expected answer is included on purpose: this route requires content.read,
// and a moderator cannot judge a question without seeing what it grades to.
// None of it is learner-visible.
func (s *Server) handleAdminLabGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "method not allowed", 405)
		return
	}
	var req struct {
		ConceptID  string  `json:"concept_id"`
		Count      int     `json:"count"`
		Difficulty float64 `json:"difficulty"`
	}
	if err := decodeJSON(w, r, &req); err != nil {
		writeError(w, "invalid request", 400)
		return
	}
	conceptID := strings.TrimSpace(req.ConceptID)
	dag := s.eng.GetDAG()
	if dag == nil {
		writeError(w, "curriculum unavailable", 503)
		return
	}
	c := dag.Concept(conceptID)
	if c == nil {
		writeError(w, "unknown concept", 404)
		return
	}
	count := req.Count
	if count < 1 {
		count = 3
	}
	if count > 10 {
		count = 10
	}
	difficulty := req.Difficulty
	if difficulty < 0.05 || difficulty > 1 {
		difficulty = 0.5
	}
	reg := s.eng.GetGeneratorRegistry()
	if reg == nil {
		writeError(w, "generators unavailable", 503)
		return
	}
	problems, err := reg.BatchGenerate(conceptID, count, difficulty)
	if err != nil {
		// A curated-only concept has no generator; that is not a server fault.
		writeError(w, "no generator for this concept", 404)
		return
	}
	format := s.answerFormatOf(conceptID)
	questions := make([]labQuestion, len(problems))
	for i, p := range problems {
		questions[i] = labQuestion{
			Question:        p.Question,
			Answer:          p.Answer,
			Explanation:     p.Explanation,
			AnswerHint:      format.Hint,
			AnswerInputMode: format.InputMode,
		}
	}
	writeJSON(w, map[string]interface{}{
		"concept_id":        conceptID,
		"concept_name":      c.Label,
		"grading_type":      c.GradingType,
		"answer_hint":       format.Hint,
		"answer_input_mode": format.InputMode,
		"difficulty":        difficulty,
		"questions":         questions,
		"writes":            false,
	})
}

// POST /api/admin/lab/grade
// Body: { "concept_id": "...", "expected": "...", "answer": "..." }
//
// Pure. GradeAnswer reads the generator and the grader and returns a verdict;
// it never enters SubmitAnswer, so no attempt, XP, streak or mastery changes.
// The report replay and the lab share writeGradeResult so the two cannot drift.
func (s *Server) handleAdminLabGrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "method not allowed", 405)
		return
	}
	var req struct {
		ConceptID string `json:"concept_id"`
		Expected  string `json:"expected"`
		Answer    string `json:"answer"`
	}
	if err := decodeJSON(w, r, &req); err != nil {
		writeError(w, "invalid request", 400)
		return
	}
	conceptID := strings.TrimSpace(req.ConceptID)
	expected := strings.TrimSpace(req.Expected)
	answer := strings.TrimSpace(req.Answer)
	if conceptID == "" || expected == "" || answer == "" {
		writeError(w, "concept_id, expected and answer are required", 400)
		return
	}
	dag := s.eng.GetDAG()
	if dag == nil || dag.Concept(conceptID) == nil {
		writeError(w, "unknown concept", 404)
		return
	}
	s.writeGradeResult(w, conceptID, expected, answer)
}

// writeGradeResult is the one shape both the report replay and the lab return.
// writes:false is the server stating what it did, so a client can label the
// result "writes nothing" without asserting it on the server's behalf.
func (s *Server) writeGradeResult(w http.ResponseWriter, conceptID, expected, answer string) {
	res := s.eng.GradeAnswer(conceptID, expected, answer)
	format := s.answerFormatOf(conceptID)
	writeJSON(w, map[string]interface{}{
		"correct":           res.Correct,
		"score":             res.Score,
		"feedback":          res.Feedback,
		"diagnosis":         res.Diagnosis,
		"unavailable":       res.Unavailable,
		"graded_answer":     answer,
		"expected":          expected,
		"answer_hint":       format.Hint,
		"answer_input_mode": format.InputMode,
		"writes":            false,
	})
}

// answerFormatIsOverride reports whether a concept's served format differs from
// the default its grading type alone would give. The lab's override table lists
// exactly these, which is the same condition the format guard test checks: a
// concept where the grader and the input disagree is the class of bug that put a
// digits-only keypad in front of "5 R 2".
func (s *Server) answerFormatIsOverride(conceptID, gradingType string) (generator.AnswerFormat, bool) {
	served := s.answerFormatOf(conceptID)
	def := generator.DefaultAnswerFormat(gradingType)
	return served, served != def
}

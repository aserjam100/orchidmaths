package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/aserjam100/orchid/internal/storage"
)

func TestCreateReport_Guest(t *testing.T) {
	s := testServer(t)
	mux := http.NewServeMux()
	s.Register(mux)

	body, _ := json.Marshal(map[string]string{
		"concept_id":  "a",
		"kind":        "question",
		"question":    "2+2=?",
		"expected":    "4",
		"reason":      "wrong_answer",
		"detail":      "looks off",
		"reporter_id": "guest_123",
	})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/api/reports", bytes.NewReader(body)))
	if rec.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var res map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if res["status"] != "open" {
		t.Errorf("expected open status, got %v", res)
	}
}

func TestCreateReport_InvalidReason(t *testing.T) {
	s := testServer(t)
	mux := http.NewServeMux()
	s.Register(mux)

	body, _ := json.Marshal(map[string]string{
		"concept_id": "a",
		"kind":       "question",
		"reason":     "bogus",
	})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/api/reports", bytes.NewReader(body)))
	if rec.Code != 400 {
		t.Errorf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestCreateReport_DiagnosticSource(t *testing.T) {
	s := testServer(t)
	mux := http.NewServeMux()
	s.Register(mux)

	body, _ := json.Marshal(map[string]string{
		"concept_id": "arith.add.basic",
		"kind":       "question",
		"question":   "1+1=?",
		"reason":     "wrong_answer",
		"source":     "diagnostic",
		"session_id": "diag_123",
	})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/api/reports", bytes.NewReader(body)))
	if rec.Code != 200 {
		t.Fatalf("expected 200 for diagnostic source, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestCreateReport_InvalidSource(t *testing.T) {
	s := testServer(t)
	mux := http.NewServeMux()
	s.Register(mux)

	body, _ := json.Marshal(map[string]string{
		"concept_id": "a",
		"kind":       "question",
		"question":   "2+2=?",
		"reason":     "wrong_answer",
		"source":     "bogus",
	})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/api/reports", bytes.NewReader(body)))
	if rec.Code != 400 {
		t.Errorf("expected 400 for invalid source, got %d: %s", rec.Code, rec.Body.String())
	}
}

// A report has to store what the learner typed and the format the field
// offered. Without those a moderator can see the question and the expected
// answer but not the attempt, so "cannot type the answer" reports could not be
// reproduced — which is exactly how the digits-only keypad survived.
func TestCreateReport_CapturesTheAnswerAndFormat(t *testing.T) {
	s := testServer(t)
	mux := http.NewServeMux()
	s.Register(mux)

	body, _ := json.Marshal(map[string]string{
		"concept_id":        "a",
		"kind":              "question",
		"question":          "2+2=?",
		"expected":          "4",
		"reason":            "formatting",
		"reporter_id":       "guest_123",
		"answer_given":      "2 + 2",
		"grading_type":      "numeric",
		"answer_hint":       "Number only",
		"answer_input_mode": "text",
	})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/api/reports", bytes.NewReader(body)))
	if rec.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var res struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	report, err := s.repo.GetReport(res.ID)
	if err != nil || report == nil {
		t.Fatalf("read back report %d: %v", res.ID, err)
	}
	if report.AnswerGiven != "2 + 2" {
		t.Errorf("answer_given = %q, want the learner's typed answer", report.AnswerGiven)
	}
	if report.GradingType != "numeric" || report.AnswerHint != "Number only" || report.AnswerInputMode != "text" {
		t.Errorf("format not stored: type=%q hint=%q mode=%q",
			report.GradingType, report.AnswerHint, report.AnswerInputMode)
	}
}

func TestCreateReport_RejectsUnknownInputMode(t *testing.T) {
	s := testServer(t)
	mux := http.NewServeMux()
	s.Register(mux)

	body, _ := json.Marshal(map[string]string{
		"concept_id":        "a",
		"kind":              "question",
		"question":          "2+2=?",
		"reason":            "formatting",
		"answer_input_mode": "voice",
	})
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest("POST", "/api/reports", bytes.NewReader(body)))
	if rec.Code != 400 {
		t.Errorf("expected 400 for an unknown input mode, got %d: %s", rec.Code, rec.Body.String())
	}
}

// The replay must use the same grader the learner's submission used and must
// not write: a moderation tool that moved a learner's mastery would change the
// evidence while it was being examined.
func TestAdminReportGrade_ReplaysWithoutWriting(t *testing.T) {
	s := testServer(t)
	id, err := s.repo.CreateReport(storage.QuestionReport{
		ConceptID:       "a",
		Kind:            "question",
		Question:        "2+2=?",
		Expected:        "4",
		AnswerGiven:     "4",
		GradingType:     "numeric",
		AnswerHint:      "Number only",
		AnswerInputMode: "text",
		Reason:          "formatting",
		Status:          "open",
	})
	if err != nil {
		t.Fatalf("create report: %v", err)
	}

	post := func(answer string) map[string]interface{} {
		body, _ := json.Marshal(map[string]string{"answer": answer})
		req := httptest.NewRequest("POST", "/api/admin/reports/"+strconv.FormatInt(id, 10)+"/grade", bytes.NewReader(body))
		rec := httptest.NewRecorder()
		s.handleAdminReport(rec, req)
		if rec.Code != 200 {
			t.Fatalf("grade %q: %d %s", answer, rec.Code, rec.Body.String())
		}
		var res map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			t.Fatalf("decode: %v", err)
		}
		return res
	}

	if got := post("4")["correct"]; got != true {
		t.Errorf("correct answer graded %v, want true", got)
	}
	if got := post("5")["correct"]; got != false {
		t.Errorf("wrong answer graded %v, want false", got)
	}
	if got := post("4")["writes"]; got != false {
		t.Errorf("writes = %v, want false", got)
	}

	// The report itself is untouched: replaying is not moderating.
	after, err := s.repo.GetReport(id)
	if err != nil || after == nil {
		t.Fatalf("read back report: %v", err)
	}
	if after.Status != "open" {
		t.Errorf("status = %q after replay, want open", after.Status)
	}
}

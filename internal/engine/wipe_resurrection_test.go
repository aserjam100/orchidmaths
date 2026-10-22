package engine

import (
	"errors"
	"testing"
)

// Phase 0: a wipe must retire in-flight work on every answer path. The
// diagnostic and quiz legs are pinned at the HTTP layer
// (account_wipe_test.go); these pin the practice and study legs at the
// engine layer, including an answer graded across the wipe itself.

func assertWipedClean(t *testing.T, e *Engine, sid string) {
	t.Helper()
	if p, err := e.repo.GetAllProgress(sid); err != nil || len(p) != 0 {
		t.Fatalf("progress after wipe: %d rows, err=%v", len(p), err)
	}
	if a, err := e.repo.GetAttemptsForStudent(sid); err != nil || len(a) != 0 {
		t.Fatalf("attempts after wipe: %d rows, err=%v", len(a), err)
	}
	if st, err := e.repo.GetStudent(sid); err != nil || st == nil || st.XPTotal != 0 {
		t.Fatalf("xp after wipe: %+v, err=%v", st, err)
	}
}

// A question served before a reset must be refused after it: the in-memory
// session is evicted with the wipe, so the stale tab cannot resurrect rows.
func TestReset_InFlightPracticeCannotResurrectProgress(t *testing.T) {
	e := testEngine(t)
	st, _ := e.CreateStudent("wiped-practice")
	sess, _ := e.repo.CreateSession(st.ID)
	q, err := e.NextQuestion(sess.ID, st.ID)
	if err != nil {
		t.Fatalf("serve: %v", err)
	}
	if err := e.ResetAccountProgress(st.ID, ResetPhrase); err != nil {
		t.Fatalf("reset: %v", err)
	}
	if _, err := e.SubmitAnswer(sess.ID, st.ID, q.AttemptID, "42", 3.0); !errors.Is(err, ErrNoActiveQuestion) {
		t.Fatalf("stale practice submit = %v, want ErrNoActiveQuestion", err)
	}
	assertWipedClean(t, e, st.ID)

	// Starting over after the wipe still works: like the production
	// /reviews/session flow, a fresh session row is created before serving.
	sess2, _ := e.repo.CreateSession(st.ID)
	q2, err := e.NextQuestion(sess2.ID, st.ID)
	if err != nil {
		t.Fatalf("re-serve: %v", err)
	}
	if _, err := e.SubmitAnswer(sess2.ID, st.ID, q2.AttemptID, "42", 3.0); err != nil {
		t.Fatalf("fresh submit after wipe: %v", err)
	}
}

// The narrower race: an answer already being graded when the wipe lands.
// Reconstructed deterministically by restoring the pre-wipe question object
// the way a blocked grader would still hold it — no sleeps, no flakes.
func TestReset_MidGradePracticeCannotResurrectProgress(t *testing.T) {
	e := testEngine(t)
	st, _ := e.CreateStudent("wiped-midgrade")
	sess, _ := e.repo.CreateSession(st.ID)
	q, err := e.NextQuestion(sess.ID, st.ID)
	if err != nil {
		t.Fatalf("serve: %v", err)
	}
	e.mu.Lock()
	stale := e.sessions[sess.ID]
	e.mu.Unlock()
	if stale == nil {
		t.Fatal("no in-memory session to race with")
	}
	if err := e.ResetAccountProgress(st.ID, ResetPhrase); err != nil {
		t.Fatalf("reset: %v", err)
	}
	e.mu.Lock()
	e.sessions[sess.ID] = stale
	e.mu.Unlock()
	if _, err := e.SubmitAnswer(sess.ID, st.ID, q.AttemptID, "42", 3.0); !errors.Is(err, ErrNoActiveQuestion) {
		t.Fatalf("mid-grade submit = %v, want ErrNoActiveQuestion", err)
	}
	assertWipedClean(t, e, st.ID)
}

// The study path's positive control: work genuinely started after the wipe
// (anchor re-seeded by a re-serve, snapshotting the current generation)
// must still record. The generation guard refuses only questions anchored
// before the wipe — i.e. an answer graded across it, symmetric with the
// practice reconstruction above.
func TestReset_PostWipeStudyWorkStillRecords(t *testing.T) {
	e := testEngine(t)
	st, _ := e.CreateStudent("wiped-study-fresh")
	const question = "What is the answer?"
	if err := e.ResetAccountProgress(st.ID, ResetPhrase); err != nil {
		t.Fatalf("reset: %v", err)
	}
	// Same write a post-wipe re-serve would perform.
	e.SetStudyAnchorBatch(st.ID, "a", map[string]Anchor{question: {Answer: "42", Explanation: "Here is why."}})
	res, err := e.SubmitStudyAnswer(st.ID, "a", "42", 3.0, question)
	if err != nil {
		t.Fatalf("fresh study submit after wipe: %v", err)
	}
	if !res.Correct {
		t.Error("expected a correct grade for the anchored answer")
	}
	if p, _ := e.repo.GetAllProgress(st.ID); len(p) != 1 {
		t.Errorf("want exactly the fresh progress row, got %d", len(p))
	}
}

// Every live practice session carries its owner: eviction on wipe matches on
// studentID, so a creation site that forgets the stamp silently exempts that
// session (caught once already on the NextQuestion path).
func TestReset_SessionsCarryOwner(t *testing.T) {
	e := testEngine(t)
	st, _ := e.CreateStudent("owned")
	sess, _ := e.repo.CreateSession(st.ID)
	if _, err := e.NextQuestion(sess.ID, st.ID); err != nil {
		t.Fatalf("serve: %v", err)
	}
	e.mu.Lock()
	as := e.sessions[sess.ID]
	e.mu.Unlock()
	if as == nil || as.studentID != st.ID {
		t.Fatalf("practice session missing owner: %+v", as)
	}
}

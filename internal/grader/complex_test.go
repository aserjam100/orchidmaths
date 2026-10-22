package grader

import "testing"

// The complex grading path had no tests before this file, although eight shipped concepts depend on
// it and no other surface exercised the preprocessor. These are the formats those concepts use:
// rectangular answers with `i`, rational coefficients, decimal equivalents, polar form with whole
// degrees, and the `±` pair a two-root answer is written with.
//
// Every case runs through Router.Grade, which is the path a learner's answer takes. When the SymPy
// runtime is missing each test skips with that stated, because Result.Unavailable means "the grader
// could not run", which the engine deliberately does not record as a miss.

func complexRouter(t *testing.T) *Router {
	t.Helper()
	r := NewRouter()
	t.Cleanup(r.Close)
	return r
}

func TestComplexAcceptsEquivalentForms(t *testing.T) {
	r := complexRouter(t)
	cases := []struct{ name, expected, answer string }{
		{"imaginary unit", "-i", "-i"},
		{"square with a negative real part", "-5+12i", "-5+12i"},
		{"rational reciprocal", "3/25-4/25i", "3/25-4/25i"},
		{"decimal reciprocal", "3/25-4/25i", "0.12-0.16i"},
		{"linear solution", "2-3i", "2-3i"},
		{"geometric series", "4/5+2/5i", "4/5+2/5i"},
		{"decimal geometric series", "4/5+2/5i", "0.8+0.4i"},
		{"residue", "5+6i", "5+6i"},
		{"rational cross-ratio", "3/2", "3/2"},
		{"polar form with whole degrees", "5(cos 53° + i sin 53°)", "5(cos 53° + i sin 53°)"},
		{"plus-minus pair, positive root", "±2i", "2i"},
		{"plus-minus pair, negative root", "±2i", "-2i"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := r.Grade(GradingComplex, c.expected, c.answer)
			if res.Unavailable {
				t.Skip("SymPy runtime missing; the complex grader cannot run here")
			}
			if !res.Correct {
				t.Errorf("%q should be accepted for %q, feedback: %s", c.answer, c.expected, res.Feedback)
			}
		})
	}
}

func TestComplexRejectsWrongAnswers(t *testing.T) {
	r := complexRouter(t)
	cases := []struct{ name, expected, answer string }{
		{"sign of the imaginary part", "-i", "i"},
		{"sign of the real part", "2-3i", "2+3i"},
		{"reciprocal wrong sign", "3/25-4/25i", "3/25+4/25i"},
		{"cross-ratio inverted", "3/2", "2/3"},
		{"real instead of imaginary", "5+6i", "5+6"},
		{"wrong modulus in polar form", "5(cos 53° + i sin 53°)", "4(cos 53° + i sin 53°)"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := r.Grade(GradingComplex, c.expected, c.answer)
			if res.Unavailable {
				t.Skip("SymPy runtime missing; the complex grader cannot run here")
			}
			if res.Correct {
				t.Errorf("%q should be rejected for %q", c.answer, c.expected)
			}
		})
	}
}

// TestComplexReportsUnavailableWithoutSymPy is the other half of the contract: a grader that cannot
// run must say so rather than call the learner wrong. It is checked by the engine's own test
// (TestUngradedAnswerCarriesNoEvidence) once this flag is set, and the flag is what this asserts.
func TestComplexReportsUnavailableWithoutSymPy(t *testing.T) {
	r := complexRouter(t)
	res := r.Grade(GradingComplex, "2-3i", "2-3i")
	if res.Unavailable && res.Correct {
		t.Error("an unavailable grader must not also report the answer correct")
	}
	if !res.Unavailable && !res.Correct {
		t.Errorf("with the grading path present the answer must grade correct, got %+v", res)
	}
}

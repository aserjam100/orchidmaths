# ADR-005: Grading contract

*2026-09-27*

one canonical dispatch `Engine.GradeAnswer` = custom `GradedGenerator.Grade` → 9-type `Router` (6 pure-Go graders) → SymPy. Every server path (`/api/answer`, `/api/study/answer`, quiz, goal diagnostic) must use it; a bare router call mis-grades the 32 custom-grader concepts (`5 R 3`, `3 sqrt(2)`, matrices, ratios). SymPy ships in the image (pinned `grading/requirements.txt`) and the image build self-tests, so `symbolic`/`polynomial`/`expression`/`complex` grading is present in production; there is **no** pure-Go symbolic fallback (the old `symbolicGrader` is dead code). Go fast paths grade exact-equal and constant-value expressions without the subprocess. Consequence: `expected == answer` must always grade correct.

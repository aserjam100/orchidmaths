# ADR-010: Grading data repair

*2026-09-27*

`orchid -repair-grading[-dry-run]` re-grades persisted `attempts` with `Engine.GradeAnswer`, flips false negatives (`UpdateAttemptCorrect`), and rebuilds `concept_progress` by replaying corrected history with the mastery/SM-2 primitives. `Result.Unavailable` attempts are skipped; XP is not rewritten. Always dry-run first.

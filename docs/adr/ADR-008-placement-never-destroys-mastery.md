# ADR-008: Placement never destroys mastery

*2026-09-27*

`ApplyGoalResults` seeds only concepts still `UNSEEN`; a diagnostic/plan run must not overwrite streak, SM-2, attempts, or mastery of a concept the student has already started.

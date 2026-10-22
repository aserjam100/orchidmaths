# ADR-006: Grading never penalizes an infrastructure fault

*2026-09-27*

a grader that cannot run is "ungraded", never a student miss — no attempt, streak reset, weakness bump, or XP. `grader.Result.Unavailable` marks it. (The prod outage recorded ~32% of concepts as false misses because a missing SymPy runtime returned `correct:false`.)

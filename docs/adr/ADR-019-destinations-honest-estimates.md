# ADR-019: Destinations + honest estimates

*2026-09-29*

destinations are named course bundles resolved to union prerequisite closures (graph stays source of truth); plans persist evidence/snapshots in `server_sessions(kind=study_plan)`, frontier derives live and is never stored. `internal/xp` owns award math (effort-derived 1–5 bases × time × streak; miss 0; rush −1; diagnostic 0 XP) and the estimator imports it — estimates can never duplicate constants. `EstimateWorkload` prices XP + time tracks (don't-knows cost reduced time with full remaining work, floored against gaming; quizzes/diagnostics are time-only, never XP debt). Pace is `min(goal, trailing-14d adherence)`; dates render as months, never promises. Goal changes never move quiz eligibility or mastery. Deferred: enforced per-path budgets, exam entity, CAT diagnostic, scored shelf policy, full misconception taxonomy.

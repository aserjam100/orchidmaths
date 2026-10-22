# ADR-014: One diagnostic route

*2026-09-27*

the legacy `/api/diagnostic` + `/api/diagnostic/answer` (unauthenticated, trusted a client `correct` flag, no persistence, unused by the frontend) is removed, along with `Engine.SubmitDiagnosticAnswer`. `/api/goal/diagnostic/*` is the only diagnostic: it grades server-side via `Engine.GradeAnswer` against the stored problem. The legacy end-to-end test was migrated to the goal route.

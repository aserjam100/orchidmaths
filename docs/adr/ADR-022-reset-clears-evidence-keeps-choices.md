# ADR-022: Reset clears evidence, keeps choices

*2026-09-29*

`Repository.ResetProgress` wipes progress/SM-2/speeds/attempts/sessions/quiz completions, study-scoped KV, XP, diagnostic flag, and league fields in one transaction on both stores; account, credentials, avatar, settings, goal, course, and `study_plan_prefs` (destination/deadline/pace) survive. Legacy single-record plans migrate lazily and idempotently. `POST /api/account/reset` requires the typed phrase, clears engine memory, orders the KV clear last, and is retry-safe. Settings danger zone gates on the phrase, offers CSV export first, and names the keeps (incl. leaderboard zeroing). Post-reset plans show a neutral fresh-start state until a new baseline lands.

# ADR-023: Single Next head over a ranked shelf

*2026-09-30*

supersedes the navigation model of ADR-017/018/021 (reference/doing split stands). DAG + `NextSmart` stay the source of truth; `selectShelfHead` returns `{next, alternatives}` with item 0 as the committed head, `selectShelf` is a flat wrapper, `selectNextUp` is the empty-policy fallback. `Learn` owns the head (`NextUpCard`); `Profile` shows a compact summary linking to `/learn`; `LearnStepper` done re-fetches the head and advances via `Continue: {label}` click (no auto-advance, alternatives present). `/concept` redirects to `/study?concept=`, `/session` redirects to `/learn`; nav `Start` renamed `Learn`. Study lists lessons in Kahn topo order. `TaskShelf` deleted.

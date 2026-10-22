# ADR-017: Adaptive lesson loop

*2026-09-29*

`/learn` hosts the KP-interleaved loop; `Study` stays reference library, `Profile` task shelf deep-links to `/learn`. Pilot `Explanation`s (`arith.add`, `arith.div.remainder/long`, `frac.add.diff`, `frac.mixed.convert`, word variants) rewritten Given→Steps→Answer with substituted numbers. Prereqs are a soft dismissible banner via `GET /api/lessons/{id}/readiness` (no auto-redirect). Mastery is derive-only (`Engine.MasteryPct` + `MasteryBand`, no migration). Practice takes `?seed=&exclude=&difficulty=` so misses serve staged fresh variants. `mapping.json` kept (ingest scripts still read it).

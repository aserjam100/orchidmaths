# ADR-018: Attempt-first feed + task shelf

*2026-09-29*

`LearnStepper` is append-only — intro once, then each submit locks its card inline and the next question opens below; scrollback is the history, with a sticky tally header. No hint button and no gating buttons: post-mistake explanations render in the locked card, halt is a 3-miss safety net that appends an easier question plus prereq links. `masteryEstimate` (accuracy recent-weighted × difficulty × variation × time) drives advance/continue/intervene; 2-in-a-row is one input, not the rule. `selectShelf` proposes ≤5 eligible tasks behind a `SelectionPolicy` (v1 fixed weights: frontier-new, due review, weakness; XP shown per card); learner disposes. UI never names selection internals; critical path stays internal.

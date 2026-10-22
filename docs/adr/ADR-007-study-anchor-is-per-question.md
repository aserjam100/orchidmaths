# ADR-007: Study anchor is per-question

*2026-09-27*

practice stores the expected answer for **every** served question (keyed by question text, `SetStudyExpectedBatch`), not just the first. Anchoring only `questions[0]` guaranteed a mis-grade for any later question.

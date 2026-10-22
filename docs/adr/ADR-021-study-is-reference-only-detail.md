# ADR-021: Study is reference-only detail

*2026-09-29*

LessonDetail shows title, concept chips, slim prereqs (badge + not-started/in-progress/mastered words, no percentages — percentages stay in Profile), one collapsed "Worked examples (N)" group (first expanded, identical bodies rendered once with combined labels, no pointer links), per-block ⋯ report menus with auto-attached block ids, and Start-learning CTAs. Footer, quiz banner, embedded practice, and the scroll-back button are gone from `/study` + `/concept`. No hidden results, no doing: reference explains, Learn teaches. Diagnostic prompts show iff new (`concepts_mastered === 0`) and incomplete, everywhere the rule applies.

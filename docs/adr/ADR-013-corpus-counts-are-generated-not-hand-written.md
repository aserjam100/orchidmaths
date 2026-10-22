# ADR-013: Corpus counts are generated, not hand-written

*2026-09-27*

`scripts/counts.py` computes concepts/domains/KP files/KP entries/subgoals/courses/diagrams from `data/` and rewrites the README's badges, sentences, and domain table (`--write`). CI and `make validate` run `--check`, which fails on stale README counts and on a denylist of superseded numbers (`641`/`580`/`284`/`1923`) anywhere in the docs. Counts had drifted every content wave; they now cannot.

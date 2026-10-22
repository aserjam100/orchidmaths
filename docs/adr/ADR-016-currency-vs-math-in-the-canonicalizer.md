# ADR-016: Currency-vs-math in the canonicalizer

*2026-09-28*

a `$…$` span with a digit is prose (escaped to literal `$`) only when it has no math marker and is either purely numeric or spaced *with a letter* (wordy price pairings like `$5 … saves … $10`). Factorial `!` is a math marker, and letterless lists (`1, 2, 3`) stay math — the old space-only rule leaked literal `$7!$`/`$1, 2, 3$` text into 24 lessons. Bare numerals (`$1$`) stay escaped per the bare-numeral gate; fix those in content, not here.

# ADR-011: Structured choice input

*2026-09-27*

`multiple_choice` questions whose option list is embedded as `A) …` blocks (`lib/choices.ts` `parseChoices`) render tappable options (`components/ChoiceOptions.tsx`) in every answer host (diagnostic, quiz, review, study); picking one submits the option text, which `choiceGrader` grades by exact match. Non-parseable MC and all other types keep the free-text input. `SymbolPalette` gains a space key and the numeric hint names the fraction/mixed-number formats, so `4 1/10` is enterable on mobile. A server-side `Options` field on `generator.Problem` (rather than parsing the stem) is the intended follow-up.

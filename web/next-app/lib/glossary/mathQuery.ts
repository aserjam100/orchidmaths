import type { GlossaryEntry } from './types'

/**
 * Math-notation queries from the MathLive search field.
 *
 * MathLive reports LaTeX (`\forall x \in ...`), while entries store canonical
 * commands (`\forall`) and glyphs (`∀`). This bridges the two by extracting
 * matchable terms — backslash commands and non-ASCII symbols — and matching
 * them against entry latex/symbol fields.
 *
 * Matching stays case-sensitive on commands (`\Sigma` is not `\sigma`) for
 * the same reason text search is: case changes the meaning. Prose typed into
 * the math field yields no terms (MathLive splits words into spaced letters
 * rather than words, so there is nothing word-like to match) and therefore no
 * filtering — the name search box remains the path for words.
 */
export function mathLatexToTerms(latex: string): string[] {
  const terms: string[] = []
  const seen = new Set<string>()
  const push = (t: string): void => {
    if (t && !seen.has(t)) {
      seen.add(t)
      terms.push(t)
    }
  }
  for (const m of latex.matchAll(/\\[a-zA-Z]+/g)) push(m[0])
  for (const ch of latex) {
    const cp = ch.codePointAt(0) ?? 0
    if (cp > 127 && !/\s/.test(ch)) push(ch)
  }
  return terms
}

export function mathSearchEntries(entries: GlossaryEntry[], latex: string): GlossaryEntry[] {
  const terms = mathLatexToTerms(latex)
  if (terms.length === 0) return entries
  const out: GlossaryEntry[] = []
  const seen = new Set<string>()
  for (const t of terms) {
    for (const e of entries) {
      if (seen.has(e.id)) continue
      const hit = t.startsWith('\\')
        ? e.latex !== undefined && (e.latex === t || e.latex.includes(t))
        : e.symbol.includes(t)
      if (hit) {
        seen.add(e.id)
        out.push(e)
      }
    }
  }
  return out
}

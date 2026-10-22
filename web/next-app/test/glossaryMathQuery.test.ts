import { describe, it, expect } from 'vitest'
import { GLOSSARY_ENTRIES } from '../lib/glossary/data'
import { mathLatexToTerms, mathSearchEntries } from '../lib/glossary/mathQuery'

describe('math latex to search terms', () => {
  it('extracts backslash commands, skipping prose and structure', () => {
    expect(mathLatexToTerms('\\forall x \\in')).toEqual(['\\forall', '\\in'])
  })

  it('extracts non-ASCII symbols verbatim', () => {
    expect(mathLatexToTerms('∀')).toEqual(['∀'])
    expect(mathLatexToTerms('\\alpha + β')).toEqual(['\\alpha', 'β'])
  })

  it('dedupes repeated terms', () => {
    expect(mathLatexToTerms('\\pi + \\pi')).toEqual(['\\pi'])
  })

  it('keeps command casing distinct', () => {
    // Σ and σ mean different things; lowercasing here would merge them.
    expect(mathLatexToTerms('\\Sigma \\sigma')).toEqual(['\\Sigma', '\\sigma'])
  })

  it('yields no terms for empty or prose-only input', () => {
    expect(mathLatexToTerms('')).toEqual([])
    expect(mathLatexToTerms('   ')).toEqual([])
    expect(mathLatexToTerms('hello')).toEqual([])
  })
})

describe('math notation search', () => {
  const ids = (q: string) => mathSearchEntries(GLOSSARY_ENTRIES, q).map(e => e.id)

  it('finds entries by LaTeX command', () => {
    expect(ids('\\forall')).toContain('forall')
    expect(ids('\\int')).toContain('integral')
    expect(ids('\\lambda')).toContain('lambda')
  })

  it('finds entries by pasted symbol', () => {
    expect(ids('∀')).toContain('forall')
    expect(ids('∫')).toContain('integral')
    expect(ids('λ')).toContain('lambda')
  })

  it('shows everything when there is nothing matchable', () => {
    expect(mathSearchEntries(GLOSSARY_ENTRIES, '')).toHaveLength(GLOSSARY_ENTRIES.length)
    expect(mathSearchEntries(GLOSSARY_ENTRIES, 'hello')).toHaveLength(GLOSSARY_ENTRIES.length)
  })

  it('returns an empty list for unknown commands, not everything', () => {
    expect(ids('\\zzzqqq')).toEqual([])
  })
})

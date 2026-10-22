import { describe, it, expect } from 'vitest'
import { GLOSSARY_ENTRIES } from '../lib/glossary/data'
import {
  exactSymbolMatches,
  filterEntries,
  GLOSSARY_PAGE_SIZE,
  paginateEntries,
  parsePageParam,
  searchEntries,
  totalPagesFor,
} from '../lib/glossary/search'

describe('glossary search', () => {
  it('finds entries by exact symbol', () => {
    expect(searchEntries(GLOSSARY_ENTRIES, '∀').map(e => e.id)).toContain('forall')
    expect(searchEntries(GLOSSARY_ENTRIES, '∫').map(e => e.id)).toContain('integral')
    expect(searchEntries(GLOSSARY_ENTRIES, 'λ').map(e => e.id)).toContain('lambda')
  })

  it('finds entries by name', () => {
    expect(searchEntries(GLOSSARY_ENTRIES, 'integral').map(e => e.id)).toContain('integral')
    expect(searchEntries(GLOSSARY_ENTRIES, 'lambda').map(e => e.id)).toContain('lambda')
  })

  it('finds entries by alias and alternate spelling', () => {
    expect(searchEntries(GLOSSARY_ENTRIES, 'for all').map(e => e.id)).toContain('forall')
    expect(searchEntries(GLOSSARY_ENTRIES, 'n choose k').map(e => e.id)).toContain('binomial-coefficient')
  })

  it('finds entries by term and definition words', () => {
    expect(searchEntries(GLOSSARY_ENTRIES, 'eigenvalue').map(e => e.id)).toContain('eigenvalue')
    expect(searchEntries(GLOSSARY_ENTRIES, 'standard deviation').map(e => e.id)).toContain('standard-deviation')
  })

  it('matches LaTeX commands without knowing the name', () => {
    expect(exactSymbolMatches(GLOSSARY_ENTRIES, '\\forall').map(e => e.id)).toContain('forall')
  })

  it('keeps case-sensitive symbols distinct while text search stays forgiving', () => {
    // Uppercase Sigma (summation) and lowercase sigma (letter) are different entries.
    const upper = searchEntries(GLOSSARY_ENTRIES, 'Σ').map(e => e.id)
    expect(upper).toContain('summation-sigma')
    // Case-insensitive description search still works.
    expect(searchEntries(GLOSSARY_ENTRIES, 'INTEGRAL').map(e => e.id)).toContain('integral')
  })

  it('returns every entry on an empty query', () => {
    expect(searchEntries(GLOSSARY_ENTRIES, '').length).toBe(GLOSSARY_ENTRIES.length)
    expect(searchEntries(GLOSSARY_ENTRIES, '   ').length).toBe(GLOSSARY_ENTRIES.length)
  })

  it('filters by category and combines with a query', () => {
    const greek = filterEntries(GLOSSARY_ENTRIES, { query: '', category: 'greek' })
    expect(greek.length).toBeGreaterThanOrEqual(24)
    expect(greek.every(e => e.category === 'greek')).toBe(true)
    const scoped = filterEntries(GLOSSARY_ENTRIES, { query: 'angle', category: 'greek' })
    expect(scoped.map(e => e.id)).toContain('theta')
    expect(scoped.every(e => e.category === 'greek')).toBe(true)
  })

  it('handles ambiguous symbols with contextual entries', () => {
    // Pi: the letter, the constant and the product are separate entries that cross-link.
    const pi = searchEntries(GLOSSARY_ENTRIES, 'π')
    expect(pi.map(e => e.id)).toContain('pi-letter')
    expect(pi.map(e => e.id)).toContain('pi-constant')
    const letter = GLOSSARY_ENTRIES.find(e => e.id === 'pi-letter')
    expect(letter?.relatedSymbols).toContain('pi-constant')
    expect(letter?.relatedSymbols).toContain('product-notation')
    // Sigma: letter, summation and standard deviation link to each other.
    const sigma = GLOSSARY_ENTRIES.find(e => e.id === 'sigma-letter')
    expect(sigma?.relatedSymbols).toContain('summation-sigma')
    expect(sigma?.relatedSymbols).toContain('standard-deviation')
    // Absolute value documents its other bar-meanings.
    const abs = GLOSSARY_ENTRIES.find(e => e.id === 'absolute-value')
    expect(abs?.contexts?.length ?? 0).toBeGreaterThanOrEqual(3)
  })

  it('returns nothing matchable rather than everything on nonsense', () => {
    expect(searchEntries(GLOSSARY_ENTRIES, 'zzzqqq').length).toBe(0)
  })
})

describe('glossary pagination', () => {
  it('pages at twenty entries', () => {
    expect(GLOSSARY_PAGE_SIZE).toBe(20)
    expect(totalPagesFor(342)).toBe(18)
    expect(totalPagesFor(0)).toBe(1)
    expect(totalPagesFor(20)).toBe(1)
    expect(totalPagesFor(21)).toBe(2)
  })

  it('slices pages without overlap or gaps', () => {
    const p1 = paginateEntries(GLOSSARY_ENTRIES, 1)
    const p2 = paginateEntries(GLOSSARY_ENTRIES, 2)
    expect(p1).toHaveLength(20)
    expect(p2).toHaveLength(20)
    expect(new Set([...p1, ...p2].map(e => e.id)).size).toBe(40)
    const last = paginateEntries(GLOSSARY_ENTRIES, totalPagesFor(GLOSSARY_ENTRIES.length))
    expect(last.length).toBeGreaterThan(0)
    expect(last.length).toBeLessThanOrEqual(20)
  })

  it('parses and clamps ?page= values', () => {
    expect(parsePageParam(null, 18)).toBe(1)
    expect(parsePageParam('', 18)).toBe(1)
    expect(parsePageParam('abc', 18)).toBe(1)
    expect(parsePageParam('0', 18)).toBe(1)
    expect(parsePageParam('-3', 18)).toBe(1)
    expect(parsePageParam('2', 18)).toBe(2)
    expect(parsePageParam('99', 18)).toBe(18)
    expect(parsePageParam('5', 1)).toBe(1)
  })
})

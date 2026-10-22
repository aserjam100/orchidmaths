import Fuse, { type IFuseOptions } from 'fuse.js'
import type { GlossaryCategoryId, GlossaryEntry } from './types'

export interface GlossaryFilter {
  query: string
  category: GlossaryCategoryId | 'all'
}

// Fuse handles fuzzy name/alias/definition matching. Exact-symbol matching is
// done separately below because case and Unicode identity matter there
// (Sigma vs sigma mean different things) while text search stays forgiving.
const fuseOptions: IFuseOptions<GlossaryEntry> = {
  includeScore: true,
  ignoreLocation: true,
  threshold: 0.35,
  keys: [
    { name: 'name', weight: 3 },
    { name: 'aliases', weight: 3 },
    { name: 'symbol', weight: 2 },
    { name: 'latex', weight: 2 },
    { name: 'definition', weight: 1 },
    { name: 'examples.tex', weight: 1 },
  ],
}

function normalizeText(s: string): string {
  return s.trim().toLowerCase()
}

/** Exact-symbol lookup: raw symbol, LaTeX command or alias equality. */
export function exactSymbolMatches(entries: GlossaryEntry[], query: string): GlossaryEntry[] {
  const q = query.trim()
  if (!q) return []
  const qLower = q.toLowerCase()
  return entries.filter(e =>
    e.symbol === q ||
    e.symbol.includes(q) ||
    (e.latex && (e.latex === q || e.latex.toLowerCase() === qLower)) ||
    e.aliases.some(a => a.toLowerCase() === qLower),
  )
}

export function searchEntries(entries: GlossaryEntry[], query: string): GlossaryEntry[] {
  const q = query.trim()
  if (!q) return entries
  const exact = exactSymbolMatches(entries, q)
  const fuse = new Fuse(entries, fuseOptions)
  const fuzzy = fuse.search(normalizeText(q)).map(r => r.item)
  const seen = new Set(exact.map(e => e.id))
  const merged = [...exact]
  for (const e of fuzzy) {
    if (!seen.has(e.id)) {
      seen.add(e.id)
      merged.push(e)
    }
  }
  return merged
}

export function filterEntries(entries: GlossaryEntry[], filter: GlossaryFilter): GlossaryEntry[] {
  const inCategory =
    filter.category === 'all' ? entries : entries.filter(e => e.category === filter.category)
  return searchEntries(inCategory, filter.query)
}

/** Entries per glossary results page. */
export const GLOSSARY_PAGE_SIZE = 20

export function totalPagesFor(total: number, pageSize: number = GLOSSARY_PAGE_SIZE): number {
  return Math.max(1, Math.ceil(total / pageSize))
}

/** Parse a ?page= param into a valid 1-based page, clamping garbage and overflow. */
export function parsePageParam(raw: string | null, totalPages: number): number {
  const n = typeof raw === 'string' ? Number.parseInt(raw, 10) : NaN
  if (!Number.isFinite(n) || n < 1) return 1
  return Math.min(Math.floor(n), Math.max(1, totalPages))
}

export function paginateEntries(
  entries: GlossaryEntry[],
  page: number,
  pageSize: number = GLOSSARY_PAGE_SIZE,
): GlossaryEntry[] {
  const start = (Math.max(1, page) - 1) * pageSize
  return entries.slice(start, start + pageSize)
}

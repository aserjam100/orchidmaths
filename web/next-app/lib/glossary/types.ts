/**
 * Glossary data model.
 *
 * One canonical entry per notation concept. A symbol with several legitimate
 * meanings carries them in `contexts` rather than as duplicate entries; truly
 * different notations (e.g. summation Sigma vs the Greek letter sigma) are
 * distinct entries cross-linked through `relatedSymbols`.
 *
 * No entry links to a Orchid concept id (v1 decision): the glossary is a
 * standalone reference and must stay useful with no curriculum mapping.
 */
export type GlossaryCategoryId =
  | 'greek'
  | 'arithmetic'
  | 'algebra'
  | 'logic-sets'
  | 'geometry-trig'
  | 'calculus'
  | 'linear-discrete'
  | 'probability'
  | 'advanced'

export interface GlossarySource {
  /** Human-readable source name, e.g. "OpenStax Calculus Volume 1". */
  label: string
  /** Stable source URL. Must resolve; never invent one. */
  url: string
  /** Reuse terms as verified, e.g. "CC BY-NC-SA 4.0 (openstax.org footer); consulted as reference, original wording, no text reused". */
  license: string
}

export interface GlossaryExample {
  /** Inline LaTeX (no $ delimiters), KaTeX-compatible. */
  tex: string
  /** What the example shows, in plain English. */
  explanation: string
}

export interface GlossaryContext {
  /** The field or situation, e.g. "Geometry". */
  context: string
  /** What the symbol means there. */
  meaning: string
}

export interface GlossaryEntry {
  /** Stable unique id, e.g. "forall-quantifier". */
  id: string
  /** Canonical rendered symbol or notation, e.g. "∀". */
  symbol: string
  /** Optional LaTeX command(s), e.g. "\\forall". */
  latex?: string
  /** Standard mathematical name. */
  name: string
  /** Alternative names, spellings and search terms. */
  aliases: string[]
  /** Optional pronunciation hint. */
  pronunciation?: string
  /** Primary category. */
  category: GlossaryCategoryId
  /** Further applicable categories. */
  relatedCategories?: GlossaryCategoryId[]
  /** Clear, concise explanation in original wording. */
  definition: string
  /** How to read the notation aloud, where useful. */
  reading?: string
  /** At least one mathematically correct example. */
  examples: GlossaryExample[]
  /** Context-specific meanings for ambiguous symbols. */
  contexts?: GlossaryContext[]
  /** Common misunderstandings, where relevant. */
  commonMistakes?: string[]
  /** Ids of other glossary entries. Must resolve; checked by test/glossaryData.test.ts. */
  relatedSymbols?: string[]
  /** At least one source with verified reuse terms. */
  sources: GlossarySource[]
  /** Qualifications or conventions. */
  notes?: string
}

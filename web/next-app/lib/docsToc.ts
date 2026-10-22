// Shared between the docs generator (`scripts/build-docs.mjs`, node) and the renderer
// (`components/DocsMarkdown.tsx`, browser). It exists so a heading's id and the table-of-contents
// link to it are computed by the same rule and cannot disagree.
//
// That was a real bug before: the hand-written ToC in docs/system-design.md links to
// `#5-sm-2-algorithm--mastery` (GitHub's slugger keeps both hyphens around a removed "&"), while
// the renderer collapsed them to one, so two of the eleven links did not jump.

export interface TocHeading {
  /** The element id, and the ToC link target. */
  id: string
  /** The heading text, with markdown emphasis stripped. */
  text: string
  /** Heading level. Only h2 headings populate the table of contents. */
  level: number
}

// The minimum number of h2 headings a document needs before a table of contents is worth the
// sidebar. A three-section page does not need one.
export const TOC_MIN_HEADINGS = 4

/** Lowercase, non-alphanumerics to single hyphens, no leading or trailing hyphen. */
export function slugify(text: string): string {
  return text
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '')
}

// Strip the inline markdown that a heading may carry so the ToC label and the slug both come
// from the words, not the punctuation: `## \`A\` and **B**` should read "A and B".
function plainText(markdown: string): string {
  return markdown
    .replace(/`([^`]*)`/g, '$1')
    .replace(/\*\*([^*]*)\*\*/g, '$1')
    .replace(/\*([^*]*)\*/g, '$1')
    .replace(/\[([^\]]*)\]\([^)]*\)/g, '$1')
    .trim()
}

/**
 * Extract the h2 headings of a markdown document, in order, with unique ids.
 *
 * Duplicate headings get a numeric suffix, the same convention GitHub uses, so two sections that
 * happen to share a title still resolve to distinct anchors.
 */
export function extractHeadings(markdown: string): TocHeading[] {
  const out: TocHeading[] = []
  const seen = new Map<string, number>()
  for (const line of markdown.split('\n')) {
    const m = /^##\s+(.+?)\s*#*\s*$/.exec(line)
    if (!m) continue
    const text = plainText(m[1])
    if (!text) continue
    const base = slugify(text)
    const n = seen.get(base) ?? 0
    seen.set(base, n + 1)
    out.push({ id: n === 0 ? base : `${base}-${n}`, text, level: 2 })
  }
  return out
}

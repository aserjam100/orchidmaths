import { describe, it, expect } from 'vitest'
import { readFileSync } from 'node:fs'
import { join } from 'node:path'
import { extractHeadings, TOC_MIN_HEADINGS } from '../lib/docsToc'

// The docs pages build a table of contents from the headings in content/docs.generated.json, and
// each link points at the id the renderer gives that heading. Both come from the same slug
// function, and these tests hold that: a heading and its link cannot drift apart without a red
// build.

const docs = JSON.parse(
  readFileSync(join(process.cwd(), 'content', 'docs.generated.json'), 'utf8'),
) as Record<string, { body: string; headings: { id: string; text: string }[] }>

describe('docs table of contents', () => {
  it('stores headings that match a fresh extraction from the body', () => {
    // If a doc changed without regenerating, the stored headings are stale. `docs:check` catches
    // the same thing for the whole module; this states it for the headings specifically.
    for (const [slug, doc] of Object.entries(docs)) {
      expect(extractHeadings(doc.body), `${slug}: run npm run docs:build`).toEqual(doc.headings)
    }
  })

  it('gives every heading a unique id within its document', () => {
    for (const [slug, doc] of Object.entries(docs)) {
      const ids = doc.headings.map(h => h.id)
      expect(new Set(ids).size, `${slug} has duplicate heading ids`).toBe(ids.length)
    }
  })

  it('removes a hand-written table of contents from the rendered body', () => {
    // The site shows the sidebar ToC; the repo's own list would be a second one on the page.
    for (const [slug, doc] of Object.entries(docs)) {
      expect(doc.body, `${slug} still renders its own Table of Contents`).not.toMatch(/^##\s+Table of Contents/m)
    }
  })

  it('slugs an ampersand heading the same way the link expects', () => {
    // The bug this was written for: the repo ToC linked to #5-sm-2-algorithm--mastery (two
    // hyphens, GitHub's slugger) while the renderer produced one hyphen, so the link did nothing.
    expect(extractHeadings('## 5. SM-2 Algorithm & Mastery')[0].id).toBe('5-sm-2-algorithm-mastery')
    expect(extractHeadings('## 11. Key Design Decisions & Trade-offs')[0].id).toBe('11-key-design-decisions-trade-offs')
  })

  it('gives a long enough document a table of contents at all', () => {
    // A guard that finds nothing reports nothing: assert the threshold is met by real docs.
    const withToc = Object.values(docs).filter(d => d.headings.length >= TOC_MIN_HEADINGS)
    expect(withToc.length).toBeGreaterThan(0)
  })
})

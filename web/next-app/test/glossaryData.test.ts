import { describe, it, expect } from 'vitest'
import katex from 'katex'
import { GLOSSARY_ENTRIES } from '../lib/glossary/data'
import { GLOSSARY_CATEGORIES } from '../lib/glossary/categories'
import { lessonMacros, prepareLessonMath } from '../lib/lessonMath'

// Guards the bundled dataset: unique ids, resolvable cross-links, no
// same-category duplicates, and source/licensing metadata on every entry.
// Notably it does NOT check mathematical truth — that was verified by hand at
// authoring time against the cited sources. Unverifiable candidates were
// excluded instead (see scripts/glossary-validation-report.json).
describe('glossary dataset integrity', () => {
  it('publishes a substantial entry count', () => {
    // Canonical entries, not glyphs: Greek variants live inside one entry each.
    expect(GLOSSARY_ENTRIES.length).toBeGreaterThanOrEqual(150)
  })

  it('covers all 24 Greek letters with one entry each', () => {
    const greek = GLOSSARY_ENTRIES.filter(e => e.category === 'greek')
    for (const letter of [
      'alpha', 'beta', 'gamma-letter', 'delta-letter', 'epsilon', 'zeta',
      'eta', 'theta', 'iota', 'kappa', 'lambda', 'mu', 'nu', 'xi',
      'omicron', 'pi-letter', 'rho', 'sigma-letter', 'tau', 'upsilon',
      'phi', 'chi', 'psi', 'omega',
    ]) {
      expect(greek.map(e => e.id), `missing Greek entry ${letter}`).toContain(letter)
    }
  })

  it('covers every category', () => {
    const present = new Set(GLOSSARY_ENTRIES.map(e => e.category))
    for (const c of GLOSSARY_CATEGORIES) {
      expect(present.has(c.id), `empty category ${c.id}`).toBe(true)
    }
  })

  it('uses unique stable ids', () => {
    const ids = GLOSSARY_ENTRIES.map(e => e.id)
    expect(new Set(ids).size).toBe(ids.length)
    for (const id of ids) {
      expect(id).toMatch(/^[a-z0-9-]+$/)
    }
  })

  it('resolves every related-symbol reference', () => {
    const ids = new Set(GLOSSARY_ENTRIES.map(e => e.id))
    const dangling: string[] = []
    for (const e of GLOSSARY_ENTRIES) {
      for (const r of e.relatedSymbols ?? []) {
        if (!ids.has(r)) dangling.push(`${e.id} -> ${r}`)
      }
    }
    expect(dangling).toEqual([])
  })

  it('has no same-category duplicate symbols', () => {
    const seen = new Map<string, string>()
    const dupes: string[] = []
    for (const e of GLOSSARY_ENTRIES) {
      const key = `${e.category}::${e.symbol}`
      if (seen.has(key)) dupes.push(`${e.id} duplicates ${seen.get(key)} (${key})`)
      else seen.set(key, e.id)
    }
    expect(dupes).toEqual([])
  })

  it('gives every entry a definition, an example and a source', () => {
    for (const e of GLOSSARY_ENTRIES) {
      expect(e.definition.length, `${e.id} definition`).toBeGreaterThan(20)
      expect(e.examples.length, `${e.id} examples`).toBeGreaterThan(0)
      for (const ex of e.examples) {
        expect(ex.tex.length, `${e.id} example tex`).toBeGreaterThan(0)
        expect(ex.explanation.length, `${e.id} example explanation`).toBeGreaterThan(0)
      }
      expect(e.sources.length, `${e.id} sources`).toBeGreaterThan(0)
      for (const s of e.sources) {
        expect(s.label.length, `${e.id} source label`).toBeGreaterThan(0)
        expect(s.url, `${e.id} source url`).toMatch(/^https:\/\//)
        expect(s.license.length, `${e.id} source license`).toBeGreaterThan(0)
      }
    }
  })

  it('keeps every example inside the math pipeline (no throw, no raw leak)', () => {
    for (const e of GLOSSARY_ENTRIES) {
      for (const ex of e.examples) {
        const out = prepareLessonMath(`$${ex.tex}$`)
        expect(out, `${e.id} math span`).toMatch(/math-(inline|display)/)
      }
    }
  })

  it('renders every example through KaTeX with zero errors', () => {
    // Same strictness as the lesson corpus gate (test/latexCorpus.test.ts):
    // examples must render, not fall back to source text. No baseline — with
    // ~200 entries every span must be clean.
    const failures: string[] = []
    for (const e of GLOSSARY_ENTRIES) {
      for (const ex of e.examples) {
        const prepared = prepareLessonMath(`$${ex.tex}$`)
        const spanRe = /<span class="math-(display|inline)">([\s\S]*?)<\/span>/g
        let m: RegExpExecArray | null
        let spans = 0
        while ((m = spanRe.exec(prepared)) !== null) {
          spans++
          const inner = m[2]
            .replace(/&amp;/g, '&')
            .replace(/&lt;/g, '<')
            .replace(/&gt;/g, '>')
            .replace(/&#(\d+);/g, (_, n: string) => String.fromCharCode(Number(n)))
          try {
            katex.renderToString(inner, {
              displayMode: m[1] === 'display',
              throwOnError: true,
              trust: false,
              macros: lessonMacros,
            })
          } catch (err) {
            failures.push(`${e.id}: ${err instanceof Error ? err.message : String(err)}`)
          }
        }
        if (spans === 0) failures.push(`${e.id}: no math span produced`)
      }
    }
    expect(failures).toEqual([])
  })

  it('keeps presentation prefixes out of reading values', () => {
    // The card component renders `Read as {reading}` exactly once, so a
    // reading that already contains the prefix displays "Read as Read as …"
    // on all 201 cards that carry one. The prefix belongs to the component,
    // never to the data. Case-insensitive: "read As", "READ AS" and
    // lookalikes ("Read → as", "Read by naming") fail the same way.
    const offenders: string[] = []
    for (const e of GLOSSARY_ENTRIES) {
      if (!e.reading) continue
      if (/^\s*read\b/i.test(e.reading)) offenders.push(e.id)
    }
    expect(offenders).toEqual([])
  })

  it('carries no emoji in the dataset', () => {
    const emoji = /[\u{1F000}-\u{1FAFF}\u{2300}-\u{23FF}\u{2600}-\u{26FF}\u{2700}-\u{27BF}\u{2B00}-\u{2BFF}\uFE0F]/u
    // Same deliberate typography exclusions as test/noEmoji.test.ts: arrows,
    // check marks, floor/ceiling brackets and mathematical symbols are notation,
    // not pictographs. The ⌊⌋⌈⌉ in the floor-ceiling entry are the subject itself.
    const allowed = new Set('→←↑↓↔⇒⇔⟶⟹✓✗✔✕✖✦✧·•⌈⌉⌊⌋∈∑∏√∞≈×÷≠≥≤°²′…—–§†')
    for (const e of GLOSSARY_ENTRIES) {
      const text = [...JSON.stringify(e)].filter(ch => !allowed.has(ch)).join('')
      expect(emoji.test(text), `${e.id} contains emoji`).toBe(false)
    }
  })
})

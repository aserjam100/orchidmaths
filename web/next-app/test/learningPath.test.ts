import { describe, it, expect } from 'vitest'
import { readdirSync, readFileSync, statSync } from 'node:fs'
import { join, relative } from 'node:path'

// Reference access and learning eligibility are different things, and this file pins the
// boundary between them.
//
// ADR-047 closed `/study` as a learner workflow because it was a *catalogue*: a browsable
// library of whole lessons with its own search and domain index, linked from the header, the tab
// bar, the sidebar and the footer, so a learner could read their way to a concept instead of
// being taught it. ADR-054 restored `/study` as a **context-bound reference page** — one concept
// named in the URL, no catalogue, awarding nothing — because a prerequisite's mastery must not
// lock access to its material.
//
// So this file pins the two properties that keep the restoration honest, each stated as a
// property of the source rather than a snapshot, because the failure mode is a link or a control
// being *added*:
//
//   1. `/study` is not a second way into the curriculum. It is not in nav, not sitemapped, and
//      has no index/search/domain list. Discovery stays in `/domains`.
//   2. Reading is not evidence. No surface records exposure as progress, and whatever moves
//      mastery does it through the answer seam.
//
// What is deliberately still allowed: the `/api/study/answer` seam (Learn's answer endpoint — a
// name, not a destination), the server-side comments that explain the history, and `/study`
// links from the DAG neighbourhood (they are the reference, not a progression step).

const APP = join(__dirname, '..')

function sources(dir: string): string[] {
  const out: string[] = []
  for (const entry of readdirSync(dir)) {
    if (entry === 'node_modules' || entry.startsWith('.') || entry === 'out') continue
    const full = join(dir, entry)
    if (statSync(full).isDirectory()) out.push(...sources(full))
    else if (/\.(ts|tsx)$/.test(entry)) out.push(full)
  }
  return out
}

/** Every authored surface a learner can be shown: routes, components and the API client. */
const LEARNER_SURFACES = [...sources(join(APP, 'app')), ...sources(join(APP, 'components')), ...sources(join(APP, 'lib'))]

/**
 * Strip comments, so a file can explain the history without tripping the guard.
 *
 * This is the part that keeps the guard honest. Without it, every explanatory comment added
 * later would need a carve-out, carve-outs accumulate, and the check quietly stops meaning
 * anything — which is the exact shape of the guards this repo already had to fix twice.
 */
function code(src: string): string {
  return src
    .replace(/\/\*[\s\S]*?\*\//g, ' ')
    .replace(/(^|[^:])\/\/[^\n]*/g, '$1 ')
}

const studyPage = () => code(readFileSync(join(APP, 'app', 'study', 'page.tsx'), 'utf-8'))

describe('the reference surface is context-bound, not a catalogue', () => {
  it('scans the surfaces it claims to scan', () => {
    expect(LEARNER_SURFACES.length).toBeGreaterThan(100)
    expect(LEARNER_SURFACES.some((f) => f.endsWith(join('lib', 'nav.ts')))).toBe(true)
    expect(LEARNER_SURFACES.some((f) => f.endsWith(join('app', 'learn', 'page.tsx')))).toBe(true)
    expect(LEARNER_SURFACES.some((f) => f.endsWith(join('app', 'study', 'page.tsx')))).toBe(true)
  })

  it('is not a global entry point', () => {
    // A nav or sitemap entry is what made `/study` "a second place to approach the same
    // mathematics". A reference page reached from the flow is not a destination in the chrome.
    expect(code(readFileSync(join(APP, 'lib', 'nav.ts'), 'utf-8'))).not.toMatch(/['"`]\/study/)
    expect(code(readFileSync(join(APP, 'lib', 'site.ts'), 'utf-8'))).not.toMatch(/\/study/)
  })

  it('requires a concept, and forwards when there is none', () => {
    // No concept means no reference to show. It forwards to the loop rather than rendering an
    // index, which is the whole difference between a reference page and a library.
    const page = studyPage()
    expect(page).toMatch(/ConceptNotFound/)
    expect(page).toMatch(/router\.replace\('\/learn'\)/)
    expect(page).toMatch(/searchParams\.get\('concept'\)/)
  })

  it('carries no catalogue: no search, no domain index, no lesson list', () => {
    const page = studyPage()
    for (const banned of [/SearchBar/, /topoRank/, /conceptTarget/, /getCurriculumDomains?,/, /<input/]) {
      expect(page).not.toMatch(banned)
    }
  })

  it('is not the reference material\'s only home, and /learn still carries its own copy', () => {
    const stepper = code(readFileSync(join(APP, 'components', 'LearnStepper.tsx'), 'utf-8'))
    expect(stepper).toMatch(/ConceptReference/)
  })
})

describe('/learn is the one assessed surface', () => {
  it('offers no control that would record exposure as progress', () => {
    // Reading is not evidence. No "mark as learned", no "I studied this", no skip, on any of the
    // reference surfaces — and nothing awards mastery outside `submitStudyAnswer`.
    const stepper = code(readFileSync(join(APP, 'components', 'LearnStepper.tsx'), 'utf-8'))
    const reference = code(readFileSync(join(APP, 'components', 'ConceptReference.tsx'), 'utf-8'))
    const prereqs = code(readFileSync(join(APP, 'components', 'PrerequisitePanel.tsx'), 'utf-8'))
    for (const src of [stepper, reference, prereqs, studyPage()]) {
      expect(src).not.toMatch(/mark as (read|learned|done)/i)
      expect(src).not.toMatch(/i (studied|read) this/i)
      expect(src).not.toMatch(/skip practice/i)
    }
    // Only the answer seam may move progress.
    for (const src of [reference, prereqs, studyPage()]) {
      expect(src).not.toMatch(/submitStudyAnswer|setBand|setCorrectCount|AddXP/)
    }
  })

  it('keeps the orientation surfaces free of any link into the learning loop', () => {
    // Graph and the concept route orient; they do not teach.
    const graph = code(readFileSync(join(APP, 'app', 'graph', 'page.tsx'), 'utf-8'))
    expect(graph).not.toMatch(/\/study/)

    const flow = code(readFileSync(join(APP, 'components', 'ConceptGraphFlow.tsx'), 'utf-8'))
    expect(flow).toMatch(/countsAsMastered/)
    expect(flow).not.toMatch(/\/study/)
  })

  it('sends a concept that is already demonstrated to retrieval practice, not a re-teach', () => {
    // The graph node action branches on the concept's state. Asserting only that /learn is not
    // linked would pass with a single unconditional `/learn?concept=` link, which re-teaches a
    // mastered concept — the tutorial purgatory ADR-037 removed from the scheduler's own path.
    const flow = code(readFileSync(join(APP, 'components', 'ConceptGraphFlow.tsx'), 'utf-8'))
    expect(flow).toMatch(/selectedLearned \? '\/review' : `\/learn\?concept=/)
  })
})

describe('the deprecated stubs stay forwarders', () => {
  it('forwards without reaching their own route', () => {
    for (const route of ['concept', 'session'] as const) {
      const page = code(readFileSync(join(APP, 'app', route, 'page.tsx'), 'utf-8'))
      expect(page).toMatch(/router\.replace/)
      expect(page).toMatch(/\/learn/)
      expect(page).not.toMatch(new RegExp(`['"\`]\\/${route}`))
    }
  })
})

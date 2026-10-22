// Build the docs content module from the repository's markdown.
//
// The website's documentation pages render the same markdown a contributor reads in the repo,
// so there is one source and the pages cannot drift from it. Before this, the pages were
// hand-written paraphrases and had drifted on exactly the facts a reader would check first:
// the in-app architecture page called Postgres "a stub" and the diagnostic "binary search".
//
//   content/docs.generated.json   slug -> { title, description, source, body }
//
// Usage:
//   npm run docs:build    regenerate the module
//   npm run docs:check    fail if the module is out of date (CI)
//
//   node scripts/build-docs.mjs --overrides <file.json>
//     Read a doc source from the given map (`path -> text`) instead of the working tree. The
//     count sweep uses this: it stages the new text of a doc another writer has dirty edits in
//     without touching that file, so a build that read the working tree would embed the stale
//     copy and the counts one batch behind. An override key the DOCS list does not name is
//     ignored, so the caller can pass every staged path without knowing which are published.
//
// The output is committed, like web/next-app/data/concepts.json, because the type checker and
// the unit tests import it and they run before a build. Committing generated content is only
// safe with the check: `docs:check` regenerates in memory and diffs, so an edit to a repo doc
// that was not regenerated fails the build rather than shipping a stale page.
import { readFileSync, writeFileSync, mkdirSync, existsSync } from 'node:fs'
import { dirname, join } from 'node:path'
import { fileURLToPath } from 'node:url'
import { extractHeadings } from '../lib/docsToc.ts'

const here = dirname(fileURLToPath(import.meta.url))
const appRoot = join(here, '..')
const repoRoot = join(appRoot, '..', '..')
const outFile = join(appRoot, 'content', 'docs.generated.json')

// The docs the website publishes, and the metadata each page needs.
//
// This list is the boundary: a doc not named here does not reach the web. The docs that stay in
// the repo are the ones a learner has no use for: the defect register (code-defects.md), the
// corpus and content audits (content-expansion.md, lesson-media.md,
// glossary-provenance.md, wasc-uc-ag-research.md), the scheduling and assessment design notes
// (mastery-semantics.md, mastery-assessments.md), the figure and slide style notes
// (diagram-style.md, system-design-slides.md) and the device checklist (mobile-device-checklist.md).
// The list here was four names while eleven docs were outside it, so a reader could not tell which
// side of the boundary the other seven were on.
//
// `description` is written here rather than derived, because a meta description is editorial
// and the first sentence of an engineering doc is usually not one.
const DOCS = [
  {
    slug: 'architecture',
    source: 'docs/architecture.md',
    description: 'The five-layer stack — UI, API, core engine, grading and storage — and the computerised adaptive diagnostic.',
  },
  {
    slug: 'administration',
    source: 'docs/administration.md',
    description: 'Roles, the permission model, staff invitations, the last-owner invariant and the audit trail.',
  },
  {
    slug: 'system-design',
    source: 'docs/system-design.md',
    description: 'Full system design: schema, scheduler, SM-2, grading, the API surface, and the trade-offs behind each.',
  },
  {
    slug: 'math-input',
    source: 'docs/math-input.md',
    description: 'How a typed answer reaches a verdict: MathLive, MathJSON, the plain-text shim, and the mobile keypad.',
  },
  {
    slug: 'efficacy',
    source: 'docs/efficacy.md',
    description: 'How Orchid measures learning, and the honest limits of those measurements.',
  },
  {
    slug: 'curriculum',
    source: 'docs/curriculum/sg-moe-primary.md',
    description: 'Every MOE Primary Mathematics outcome, P1 to P6, and how fully Orchid practises it.',
  },
  {
    slug: 'licenses',
    source: 'THIRD_PARTY_NOTICES.md',
    description: 'Every third-party library, content source and font Orchid ships, with its license.',
  },
]

function titleOf(markdown, fallback) {
  const m = markdown.match(/^#\s+(.+)$/m)
  return m ? m[1].trim() : fallback
}

// Drop the leading H1. The page renders its own heading, and two titles on one page is worse
// than one.
function withoutTitle(markdown) {
  return markdown.replace(/^#\s+.*(\r?\n)+/, '')
}

// Remove a hand-written "Table of Contents" block. The website builds its own from the headings
// and shows it in the sidebar, so rendering the repo's list too would put two tables of contents
// on one page. The repo keeps its own for GitHub.
function withoutRepoToc(markdown) {
  return markdown.replace(/^##\s+Table of Contents[^\n]*\n[\s\S]*?(?=^##\s)/m, '')
}

export function build(overrides = {}) {
  const out = {}
  const missing = []
  for (const doc of DOCS) {
    const path = join(repoRoot, doc.source)
    const overridden = Object.prototype.hasOwnProperty.call(overrides, doc.source)
    if (!overridden && !existsSync(path)) {
      // A missing source is fatal, not skipped: a silently absent page is the failure this
      // whole arrangement exists to prevent.
      missing.push(doc.source)
      continue
    }
    const raw = overridden ? overrides[doc.source] : readFileSync(path, 'utf8')
    const body = withoutRepoToc(withoutTitle(raw)).trimEnd() + '\n'
    out[doc.slug] = {
      title: titleOf(raw, doc.slug),
      description: doc.description,
      source: doc.source,
      body,
      headings: extractHeadings(body),
    }
  }
  if (missing.length) {
    throw new Error(
      `docs source not found: ${missing.join(', ')}. ` +
      `Every name in build-docs.mjs must exist; there is no silent skip.`,
    )
  }
  return out
}

function serialise(data) {
  return JSON.stringify(data, null, 2) + '\n'
}

const check = process.argv.includes('--check')
const overridesArg = process.argv.indexOf('--overrides')
const overrides = overridesArg >= 0
  ? JSON.parse(readFileSync(process.argv[overridesArg + 1], 'utf8'))
  : {}
const next = serialise(build(overrides))

if (check) {
  const current = existsSync(outFile) ? readFileSync(outFile, 'utf8') : ''
  if (current !== next) {
    console.error(
      'FAIL: content/docs.generated.json is out of date with the repo docs.\n' +
      '      Run `npm run docs:build` and commit the result.',
    )
    process.exit(1)
  }
  console.log(`OK: docs content matches ${DOCS.length} repo source(s)`)
} else {
  mkdirSync(dirname(outFile), { recursive: true })
  writeFileSync(outFile, next)
  console.log(`OK: wrote ${outFile} from ${DOCS.length} source(s)`)
}

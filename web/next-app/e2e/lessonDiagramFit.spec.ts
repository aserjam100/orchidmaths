import { test, expect } from './fixtures'
import { readFileSync } from 'node:fs'
import { resolve } from 'node:path'

/**
 * A step diagram in `/learn` is scaled to its height bound, not clipped by it.
 *
 * `LessonDiagram` bounds a figure at 180px by default, and the bound belongs on the SVG. It used
 * to sit on the wrapping box, which also had `overflow: hidden`: a figure taller than the bound
 * was cut off rather than shrunk. That is not a corner case — 123 of the 127 mapped SVGs are
 * taller than 180px and the arithmetic ones are 360x200, so the bottom of nearly every diagram
 * was missing at any width that did not force a scale, which on a desktop panel is all of them.
 *
 * `lessonFigures.spec.ts` could not see it: that spec measures the reference panel, which renders
 * the article through `KatexContent` as a plain `<img>`, and an `<img>` scales under `max-height`
 * correctly. The clipping was only in the step renderer, which sets the SVG inline. This spec
 * covers that renderer, and its negative control restores the old rule to show the measurement
 * can report a clip — a "not clipped" assertion that cannot fail is not a guard.
 */

const ROOT = resolve(__dirname, '../../..')
const CONCEPT = 'arith.place.tens'
const FILE = 'teaching/arith.place.tens.md'
/** 360x200: taller than the 180px default bound, or this spec would prove nothing. */
const DIAGRAM = '/diagrams/arithmetic/place-tens.svg'

const BOX = '.lesson-diagram'

async function stubStep(page: import('@playwright/test').Page, diagram: string) {
  const body = readFileSync(resolve(ROOT, 'data/lessons', FILE), 'utf8')
  const title = body.split('\n')[0].replace(/^#\s*/, '')
  await page.route('**/api/lessons/body*', route =>
    route.fulfill({ json: { title, body, source: 'teaching', assets: [] } }),
  )
  await page.route('**/api/lessons*', route =>
    route.fulfill({ json: { lessons: { teaching: [{ title, body: '', concepts: [CONCEPT] }] } } }),
  )
  // The concept's diagram rides on its first step, which is the box under test.
  await page.route('**/api/lessons/*/kp*', route =>
    route.fulfill({
      json: {
        concept_id: CONCEPT,
        diagram,
        kps: [{ label: 'Counting and whole numbers', subgoals: [], worked_example: 'A worked example.' }],
      },
    }),
  )
  await page.route('**/api/lessons/*/readiness*', route =>
    route.fulfill({ json: { concept_id: CONCEPT, ready: true, weak: [], missing: [] } }),
  )
  await page.route('**/api/lessons/*/practice*', route =>
    route.fulfill({ json: { concept_id: CONCEPT, questions: [] } }),
  )
  await page.route('**/api/progress/**', route => route.fulfill({ json: {} }))
}

/**
 * Open `/learn` and wait until a diagram box is laid out with its SVG in it.
 *
 * The step is a `<details>` that starts closed, so a box that exists is not yet a box with a
 * height. Waiting on a height rather than on attachment is what keeps the measurement below from
 * reading 0 and passing for the wrong reason.
 */
async function openStepDiagram(page: import('@playwright/test').Page) {
  await page.goto(`/learn?concept=${encodeURIComponent(CONCEPT)}`)
  await page.waitForFunction(
    (sel) => {
      const boxes = Array.from(document.querySelectorAll(sel)) as HTMLElement[]
      return boxes.some((b) => b.clientHeight > 0 && b.querySelector('svg'))
    },
    BOX,
    { timeout: 30_000 },
  )
}

/** The first laid-out box: its height, its content height, and the drawn height of its SVG. */
async function stepDiagramSize(page: import('@playwright/test').Page) {
  return page.evaluate((sel) => {
    const boxes = Array.from(document.querySelectorAll(sel)) as HTMLElement[]
    const box = boxes.find((b) => b.clientHeight > 0 && b.querySelector('svg'))
    if (!box) throw new Error('no laid-out diagram box')
    const svg = box.querySelector('svg')!
    return {
      boxHeight: box.clientHeight,
      contentHeight: box.scrollHeight,
      svgHeight: Math.round(svg.getBoundingClientRect().height),
    }
  }, BOX)
}

test('a step diagram taller than the bound is scaled to fit, not clipped', async ({ page }) => {
  await stubStep(page, DIAGRAM)
  await openStepDiagram(page)

  const size = await stepDiagramSize(page)
  expect(size.svgHeight, 'the SVG drew at zero height, so nothing was measured').toBeGreaterThan(0)
  expect(
    size.contentHeight,
    `the figure is clipped: content is ${size.contentHeight}px in a ${size.boxHeight}px box`,
  ).toBeLessThanOrEqual(size.boxHeight + 1)
})

test('negative control: the old clamp reports a clip under the same measurement', async ({ page }) => {
  await stubStep(page, DIAGRAM)
  await openStepDiagram(page)

  // The rule this test exists to keep out: the bound on the box, the overflow hidden, and the
  // SVG free to draw at its natural height.
  await page.addStyleTag({
    content: `
      .lesson-diagram { overflow: hidden !important; max-height: 180px !important; }
      .lesson-diagram svg { max-height: none !important; }
    `,
  })

  const size = await stepDiagramSize(page)
  expect(
    size.contentHeight,
    'the check above cannot fail, so it is not a check',
  ).toBeGreaterThan(size.boxHeight + 1)
})

import { test, expect } from './fixtures'
import { readFileSync } from 'node:fs'
import { resolve } from 'node:path'

/**
 * A Orchid-authored figure keeps its text inside its own viewBox.
 *
 * An inline `<svg>` clips to its viewport, and every figure here is drawn to the pixel in its own
 * coordinate system: a label whose ink crosses `viewBox` loses its right-hand end on the page. Twelve
 * of the seventy local figures did, and each looked like a mistake rather than a crop — the place
 * value row read `1,247 = 1 thousand + 2 hundreds + 4 ten`, the resultant vector's `u+v` label sat at
 * x=410 in a 400-wide canvas, and a linear graph carried its `-4` tick 24px below a 300-tall one.
 *
 * The measurement is exact rather than estimated. `globals.css` forces every inlined figure's text to
 * JetBrains Mono, which advances 0.6 em per glyph, so a label's width is computable from its font
 * size; this spec measures the rendered `getBBox()` in a monospace face instead of trusting the
 * arithmetic, and then asserts the advance really is 0.6 em so that a hostile environment reports a
 * broken harness rather than a clean corpus.
 *
 * Only `source: local` assets are measured. The vendored set is Algebrica's, drawn on Algebrica's
 * canvases for a narrower face, and `docs/diagram-style.md` section 7 records the 43 of those whose
 * in-figure commentary is wider than the drawing — a separate decision, not ours to reflow here.
 */

const ROOT = resolve(__dirname, '../../..')
const PUBLIC = resolve(ROOT, 'web/next-app/public')

type Meta = Record<string, { source?: string }>

const meta: Meta = JSON.parse(readFileSync(resolve(PUBLIC, 'diagrams/meta.json'), 'utf8'))
const local = Object.entries(meta)
  .filter(([, v]) => v.source === 'local')
  .map(([asset]) => asset)
  .sort()

/** JetBrains Mono, and the monospace fallback beside it, advance 0.6 em per glyph. */
const ADVANCE = 0.6

const STYLE = `<style>.lesson-diagram text { font-family: monospace; }</style>`

async function overflow(
  page: import('@playwright/test').Page,
  markup: string,
): Promise<{ label: string; where: string }[]> {
  await page.setContent(
    `<body style="margin:0">${STYLE}<div class="lesson-diagram">${markup}</div></body>`,
  )
  await page.evaluate(() => document.fonts.ready)
  return page.evaluate(() => {
    const svg = document.querySelector('#host svg, .lesson-diagram svg') as SVGSVGElement | null
    if (!svg || !svg.viewBox?.baseVal) throw new Error('no viewBox on the figure')
    const vb = svg.viewBox.baseVal
    // Element space to viewBox space: the screen CTMs cancel the root's own scale.
    const inv = svg.getScreenCTM()!.inverse()
    const out: { label: string; where: string }[] = []
    for (const t of Array.from(svg.querySelectorAll('text')) as SVGTextElement[]) {
      const b = t.getBBox()
      if (!b.width) continue
      const m = inv.multiply(t.getScreenCTM()!)
      const corners = [
        [b.x, b.y],
        [b.x + b.width, b.y],
        [b.x, b.y + b.height],
        [b.x + b.width, b.y + b.height],
      ].map(([x, y]) => ({ x: m.a * x + m.c * y + m.e, y: m.b * x + m.d * y + m.f }))
      const xs = corners.map(p => p.x)
      const ys = corners.map(p => p.y)
      const over: [string, number][] = []
      if (Math.min(...xs) < -0.5) over.push(['left', -Math.min(...xs)])
      if (Math.max(...xs) > vb.width + 0.5) over.push(['right', Math.max(...xs) - vb.width])
      if (Math.min(...ys) < -0.5) over.push(['top', -Math.min(...ys)])
      if (Math.max(...ys) > vb.height + 0.5) over.push(['bottom', Math.max(...ys) - vb.height])
      if (over.length) {
        const [side, px] = over.reduce((a, c) => (c[1] > a[1] ? c : a))
        out.push({ label: t.textContent!.trim().slice(0, 40), where: `${side} by ${px.toFixed(1)}px` })
      }
    }
    return out
  })
}

test.describe('local diagram text', () => {
  test('the harness measures in the face the app renders with', async ({ page }) => {
    await page.setContent(
      `<body style="margin:0">${STYLE}<div class="lesson-diagram">` +
        `<svg viewBox="0 0 200 40"><text x="0" y="20" font-size="10">0123456789</text></svg>` +
        `</div></body>`,
    )
    const ratio = await page.evaluate(() => {
      const t = document.querySelector('text')!
      return t.getBBox().width / (10 * parseFloat(getComputedStyle(t).fontSize))
    })
    expect(Math.abs(ratio - 0.6)).toBeLessThan(0.01)
  })

  test('the measurement reports text that leaves the canvas', async ({ page }) => {
    // The negative control: 10 glyphs at 10px are 60px wide, starting at 80 in a 100-wide viewBox.
    const bad = await overflow(
      page,
      '<svg viewBox="0 0 100 40"><text x="80" y="20" font-size="10">0123456789</text></svg>',
    )
    expect(bad).toHaveLength(1)
    expect(bad[0].where).toMatch(/^right by \d+\.\d+px$/)
    expect(parseFloat(bad[0].where.replace(/[^\d.]/g, ''))).toBeGreaterThan(30)
  })

  test('every local figure keeps its text inside its viewBox', async ({ page }) => {
    expect(local.length).toBeGreaterThan(50)
    const failures: string[] = []
    for (const asset of local) {
      // meta.json keys are absolute URLs (`/diagrams/...`); resolve() would discard PUBLIC for them.
      const file = resolve(PUBLIC, asset.replace(/^\//, ''))
      const markup = readFileSync(file, 'utf8').replace(/<\?xml[^>]*\?>/, '')
      const bad = await overflow(page, markup)
      for (const b of bad) failures.push(`${asset}: ${JSON.stringify(b.label)} runs off the ${b.where}`)
    }
    expect(failures.join('\n')).toBe('')
  })
})

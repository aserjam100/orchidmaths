import { describe, it, expect } from 'vitest'
import { readFileSync } from 'node:fs'
import { join } from 'node:path'
import { DIAGRAMS } from '../lib/docsDiagrams'

// The rendered docs pages replace a diagram image with its interactive React Flow component.
// When a doc references a diagram with no component, the page silently falls back to the
// pre-rendered SVG -- which looks fine and hides the gap. This fails instead, so adding a
// diagram to a repo doc and forgetting the component is a red build, not a static picture.

const docs = JSON.parse(
  readFileSync(join(process.cwd(), 'content', 'docs.generated.json'), 'utf8'),
) as Record<string, { body: string }>

function referencedDiagrams(): Set<string> {
  const names = new Set<string>()
  for (const doc of Object.values(docs)) {
    for (const m of doc.body.matchAll(/diagrams\/([\w-]+?)(?:-(?:light|dark))?\.svg/g)) {
      names.add(m[1])
    }
  }
  return names
}

describe('rendered docs embed every diagram as a component', () => {
  const refs = referencedDiagrams()

  it('finds diagram references at all', () => {
    // A guard that reads nothing reports nothing.
    expect(refs.size).toBeGreaterThan(0)
  })

  it('has a React Flow component for each referenced diagram', () => {
    const missing = [...refs].filter(name => !(name in DIAGRAMS))
    expect(missing, `no interactive component for: ${missing.join(', ')}`).toEqual([])
  })
})

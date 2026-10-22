import { describe, it, expect, vi } from 'vitest'
import { findWarmUp } from '../lib/warmUp'
import type { ReadinessRes } from '../lib/api'

function res(id: string, eligible: boolean, unmet: string[] = [], met: string[] = []): ReadinessRes {
  return {
    concept_id: id,
    ready: eligible,
    weak: [],
    missing: [],
    eligible,
    prerequisites: [
      ...unmet.map(p => ({ conceptId: p, title: p.toUpperCase(), state: 'locked', met: false })),
      ...met.map(p => ({ conceptId: p, title: p.toUpperCase(), state: 'mastered', met: true })),
    ],
  }
}

// A locked topic offers the first prerequisite the server will serve. The server decides every
// step; the walk only chooses which unmet prerequisite to ask about next.
describe('findWarmUp', () => {
  it('offers a direct prerequisite the server will serve', async () => {
    const readiness = vi.fn(async (id: string) => res(id, true))
    expect(await findWarmUp(res('c', false, ['b']), readiness)).toEqual({ conceptId: 'b', title: 'B' })
  })

  it('walks down through a locked prerequisite', async () => {
    const graph: Record<string, ReadinessRes> = { b: res('b', false, ['a']), a: res('a', true) }
    const readiness = vi.fn(async (id: string) => graph[id])
    expect(await findWarmUp(res('c', false, ['b']), readiness)).toEqual({ conceptId: 'a', title: 'A' })
  })

  it('never asks about a prerequisite that is already met', async () => {
    const readiness = vi.fn(async (id: string) => res(id, true))
    await findWarmUp(res('c', false, ['b'], ['m']), readiness)
    expect(readiness).not.toHaveBeenCalledWith('m')
  })

  it('tries a sibling when the first branch dead-ends', async () => {
    const graph: Record<string, ReadinessRes> = {
      b: res('b', false, ['c']), // points back at the locked topic: a cycle, not a warm-up
      d: res('d', true),
    }
    const readiness = vi.fn(async (id: string) => graph[id])
    expect(await findWarmUp(res('c', false, ['b', 'd']), readiness)).toEqual({ conceptId: 'd', title: 'D' })
  })

  it('gives up after a bounded number of requests', async () => {
    let n = 0
    const readiness = vi.fn(async (id: string) => res(id, false, [`x${++n}`]))
    expect(await findWarmUp(res('c', false, ['x0']), readiness)).toBeNull()
    expect(readiness.mock.calls.length).toBeLessThanOrEqual(12)
  })
})

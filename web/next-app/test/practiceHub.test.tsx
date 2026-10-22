import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import LearnPage from '../app/learn/page'
import LearnStepper from '../components/LearnStepper'
import { getNext, getLessonKPs, getLessonPractice, getLessonReadiness } from '../lib/api'

vi.mock('next/navigation', () => ({
  useRouter: () => ({ push: vi.fn(), replace: vi.fn() }),
  usePathname: () => '/learn',
  useSearchParams: () => new URLSearchParams(),
}))

vi.mock('../lib/api', async (importOriginal) => {
  const mod = await importOriginal<typeof import('../lib/api')>()
  return {
    ...mod,
    getNext: vi.fn(),
    getLessonKPs: vi.fn(),
    getLessonPractice: vi.fn(),
    getLessonReadiness: vi.fn(),
    getSettings: vi.fn(async () => ({})),
  }
})

Element.prototype.scrollIntoView = vi.fn() as unknown as typeof Element.prototype.scrollIntoView

const rec = (id: string, href = `/learn?concept=${id}`) => ({
  id,
  conceptId: id,
  conceptTitle: id.toUpperCase(),
  kind: 'learn' as const,
  reason: 'new' as const,
  priority: 1,
  action: { type: 'learn' as const, href },
  badge: 'New',
  detail: 'Ready to learn',
  cta: 'Start →',
})

// The practice hub at /learn: the learner's grade, one button that starts the engine's pick, and
// Topics. A new learner's head is the placement test; the hub starts practice instead and offers
// the test in one line.
describe('practice hub', () => {
  beforeEach(() => {
    localStorage.clear()
    vi.mocked(getNext).mockReset()
  })

  it('asks for a grade first when none is stored', async () => {
    render(<LearnPage />)
    expect(await screen.findByRole('group', { name: 'Choose your grade' })).toBeInTheDocument()
    expect(getNext).not.toHaveBeenCalled()
  })

  it('starts the first concept rather than the placement test, and offers the test', async () => {
    localStorage.setItem('orchid_primary_level', 'P3')
    vi.mocked(getNext).mockResolvedValue({
      primary: { ...rec('diagnostic', '/onboard'), conceptId: '' },
      alternatives: [rec('a.one'), rec('a.two')],
      generatedAt: '',
    } as never)
    render(<LearnPage />)
    const start = await screen.findByRole('link', { name: /Start practising/ })
    expect(start.getAttribute('href')).toBe('/learn?concept=a.one')
    expect(screen.getByRole('heading', { name: 'Primary 3 practice' })).toBeInTheDocument()
    expect(screen.getByRole('link', { name: /skip ahead/ }).getAttribute('href')).toBe('/onboard')
    expect(screen.getByRole('link', { name: /Choose a topic/ }).getAttribute('href')).toBe('/levels')
  })

  it('follows the engine head when it is not the placement test', async () => {
    localStorage.setItem('orchid_primary_level', 'P3')
    vi.mocked(getNext).mockResolvedValue({ primary: rec('b.one'), alternatives: [], generatedAt: '' } as never)
    render(<LearnPage />)
    const start = await screen.findByRole('link', { name: /Start practising/ })
    expect(start.getAttribute('href')).toBe('/learn?concept=b.one')
    expect(screen.queryByRole('link', { name: /skip ahead/ })).toBeNull()
  })
})

// A locked topic offers a warm-up on the first prerequisite the server will serve, instead of a
// reading list.
describe('locked topic', () => {
  it('offers a one-tap warm-up on a servable prerequisite', async () => {
    vi.mocked(getLessonKPs).mockResolvedValue({ kps: [], diagram: null } as never)
    vi.mocked(getLessonPractice).mockResolvedValue({ questions: [], concept_id: 'arith.add.carry' } as never)
    vi.mocked(getLessonReadiness).mockImplementation(async (id: string) =>
      id === 'arith.add.carry'
        ? { concept_id: id, ready: false, weak: [], missing: [], eligible: false, prerequisites: [{ conceptId: 'arith.add.two', title: 'Two-digit addition', state: 'unlocked', met: false }] }
        : { concept_id: id, ready: true, weak: [], missing: [], eligible: true })
    render(<LearnStepper conceptId="arith.add.carry" />)
    const warm = await screen.findByRole('link', { name: /Start warm-up: Two-digit addition/ })
    expect(warm.getAttribute('href')).toMatch(/^\/learn\?concept=arith\.add\.two/)
    await waitFor(() => expect(getLessonPractice).not.toHaveBeenCalled())
  })
})

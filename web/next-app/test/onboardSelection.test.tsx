import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import OnboardPage from '../app/onboard/page'
import { scopeSize } from '../lib/diagnosticScope'
import { PRIMARY_LEVELS } from '../lib/primaryLevel'
import * as api from '../lib/api'
import * as primaryLevel from '../lib/primaryLevel'

const push = vi.fn()
const replace = vi.fn()
vi.mock('next/navigation', () => ({
  useRouter: () => ({ push, replace }),
}))
vi.mock('next/link', () => ({
  default: ({ children, href }: { children: React.ReactNode; href: string }) => (
    <a href={href}>{children}</a>
  ),
}))
vi.mock('../lib/api', async (importOriginal) => {
  const mod = await importOriginal<typeof import('../lib/api')>()
  return { ...mod, startGoalDiagnostic: vi.fn(), getGoalPlan: vi.fn() }
})
vi.mock('../lib/primaryLevel', async (importOriginal) => {
  const mod = await importOriginal<typeof import('../lib/primaryLevel')>()
  return { ...mod, setLevel: vi.fn(async () => {}) }
})
vi.mock('../hooks/useTheme', () => ({ useTheme: () => ({ mounted: true }) }))
vi.mock('../components/Header', () => ({ default: () => null }))
vi.mock('../components/Footer', () => ({ default: () => null }))
vi.mock('../components/BottomTabs', () => ({ default: () => null }))
vi.mock('../components/Loading', () => ({ default: () => <div>Loading</div> }))

/** The start button, whatever copy it currently carries. */
function startButton(): HTMLButtonElement {
  return screen.getByRole('button', { name: /Start diagnostic test/ }) as HTMLButtonElement
}

function scopeShown(): string | null {
  const m = startButton().textContent?.match(/([\d,]+) concepts in scope/)
  return m ? m[1].replace(/,/g, '') : null
}

function levelButton(name: string): HTMLElement {
  return screen.getByRole('radio', { name: new RegExp(`^${name}`) })
}

beforeEach(() => {
  push.mockReset()
  replace.mockReset()
  sessionStorage.clear()
  localStorage.clear()
  vi.mocked(api.startGoalDiagnostic).mockReset()
  vi.mocked(api.getGoalPlan).mockReset()
  vi.mocked(primaryLevel.setLevel).mockClear()
})

// The picker was a grid of all 21 topic domains, calculus and topology among them, on a product
// sold as Primary maths. It is now the child's level, and the level is the test's scope.
describe('/onboard level selection', () => {
  it('offers Primary 1 to 6 and no topic domains', async () => {
    render(<OnboardPage />)
    await waitFor(() => expect(levelButton('Primary 1')).toBeTruthy())
    for (const l of PRIMARY_LEVELS) expect(levelButton(l.name)).toBeTruthy()
    expect(screen.queryByText('Calculus')).toBeNull()
    expect(screen.queryByText('Select everything')).toBeNull()
    expect(startButton().disabled).toBe(true)
  })

  it('refuses to start with no level chosen', async () => {
    render(<OnboardPage />)
    await waitFor(() => expect(startButton()).toBeTruthy())
    fireEvent.click(startButton())
    expect(api.startGoalDiagnostic).not.toHaveBeenCalled()
  })

  it('reports the real scope of the chosen level, which grows with the level', async () => {
    render(<OnboardPage />)
    await waitFor(() => expect(levelButton('Primary 2')).toBeTruthy())
    fireEvent.click(levelButton('Primary 2'))
    const p2 = PRIMARY_LEVELS.find(l => l.level === 'P2')!
    expect(Number(scopeShown())).toBe(scopeSize(p2.targets))
    fireEvent.click(levelButton('Primary 5'))
    expect(Number(scopeShown())).toBeGreaterThan(scopeSize(p2.targets))
  })

  it('starts from the level already chosen', async () => {
    localStorage.setItem('orchid_primary_level', 'P4')
    render(<OnboardPage />)
    await waitFor(() => expect(levelButton('Primary 4').getAttribute('aria-checked')).toBe('true'))
  })

  it('saves the level and sends its course targets up', async () => {
    vi.mocked(api.startGoalDiagnostic).mockResolvedValue({
      session_id: 's1', concept_id: 'c1', concept_name: 'C', question: '1 + 1 = ?',
      progress: { answered: 0, estimated_total: 25, min_total: 25, max_total: 45, cover_done: 0, cover_size: 5, done: false },
    } as never)
    render(<OnboardPage />)
    await waitFor(() => expect(levelButton('Primary 3')).toBeTruthy())
    fireEvent.click(levelButton('Primary 3'))

    fireEvent.click(startButton())
    await screen.findByText('Before you begin')
    expect(screen.getByText(/Starting test on: Primary 3/)).toBeTruthy()
    fireEvent.click(screen.getByRole('button', { name: /Begin diagnostic/ }))

    await waitFor(() => expect(api.startGoalDiagnostic).toHaveBeenCalledTimes(1))
    expect(primaryLevel.setLevel).toHaveBeenCalledWith('P3')
    const p3 = PRIMARY_LEVELS.find(l => l.level === 'P3')!
    expect([...vi.mocked(api.startGoalDiagnostic).mock.calls[0][0]].sort()).toEqual([...p3.targets].sort())
  })

  it('offers a way into the topics without taking the test', async () => {
    render(<OnboardPage />)
    await waitFor(() => expect(screen.getByText(/Browse topics/)).toBeTruthy())
    const link = screen.getByText(/Browse topics/).closest('a')
    expect(link?.getAttribute('href')).toBe('/levels')
  })
})

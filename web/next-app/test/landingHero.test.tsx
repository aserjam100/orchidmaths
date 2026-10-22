import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import HomePage from '../app/page'
import ParentsPage from '../app/parents/page'
import { setLevel } from '../lib/primaryLevel'
import { ensureGuestId, ensureGuestToken } from '../lib/auth'

const pushMock = vi.fn()
const replaceMock = vi.fn()

vi.mock('next/navigation', () => ({
  useRouter: () => ({ push: pushMock, replace: replaceMock }),
  usePathname: () => '/',
  useSearchParams: () => new URLSearchParams(),
}))

vi.mock('../lib/primaryLevel', async (importOriginal) => {
  const mod = await importOriginal<typeof import('../lib/primaryLevel')>()
  return { ...mod, setLevel: vi.fn(async () => {}) }
})

vi.mock('../lib/auth', async (importOriginal) => {
  const mod = await importOriginal<typeof import('../lib/auth')>()
  return {
    ...mod,
    ensureGuestId: vi.fn(),
    ensureGuestToken: vi.fn(async () => 'guest-token'),
  }
})

// The landing is a grade picker. A student who came to practise picks a grade and the next
// screen asks a question: no placement test in between, no marketing to scroll past. The
// placement test and the explainers are one link away (/learn and /parents).
describe('landing grade picker', () => {
  beforeEach(() => {
    localStorage.clear()
    sessionStorage.clear()
    pushMock.mockClear()
    replaceMock.mockClear()
    vi.mocked(ensureGuestId).mockClear()
    vi.mocked(setLevel).mockClear()
  })

  it('offers all six Primary grades', async () => {
    render(<HomePage />)
    const group = await screen.findByRole('group', { name: 'Choose your grade' })
    expect(group.querySelectorAll('button')).toHaveLength(6)
  })

  it('saves the grade and opens practice, skipping the placement test', async () => {
    render(<HomePage />)
    ;(await screen.findByRole('button', { name: /Primary 3/ })).click()
    await waitFor(() => expect(pushMock).toHaveBeenCalledWith('/learn'))
    expect(setLevel).toHaveBeenCalledWith('P3')
    expect(pushMock).not.toHaveBeenCalledWith('/onboard')
    // The level is saved on the student row, so a guest needs an id before it is saved.
    expect(ensureGuestId).toHaveBeenCalled()
  })

  it('keeps Log in and the parents page one link away', async () => {
    render(<HomePage />)
    const main = await screen.findByRole('main')
    const hrefs = [...main.querySelectorAll('a')].map(a => a.getAttribute('href'))
    expect(hrefs).toContain('/login')
    expect(hrefs).toContain('/parents')
  })

  it('bounces a signed-in visitor to practice', async () => {
    localStorage.setItem('orchid_token', 't')
    render(<HomePage />)
    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith('/learn'))
  })
})

// The marketing page moved from / to /parents on 2026-10-22. Its hero's "Get started" is the most
// prominent action on it.
// It used to send a signed-out visitor to /profile — an empty dashboard with a
// zeroed heatmap and no question anywhere on it. It then went straight to /learn,
// which skipped the diagnostic entirely: the only route into placement was a small
// link on the profile page. It now opens the diagnostic, and "just practise" keeps
// the direct route into Learn.
describe('parents page hero', () => {
  beforeEach(() => {
    localStorage.clear()
    sessionStorage.clear()
    pushMock.mockClear()
    replaceMock.mockClear()
    vi.mocked(ensureGuestId).mockClear()
    vi.mocked(ensureGuestToken).mockClear()
  })

  it('sends "Get started" to the diagnostic, not to an empty Profile', async () => {
    render(<ParentsPage />)
    const cta = await screen.findByRole('button', { name: 'Get started' })
    cta.click()
    await waitFor(() => expect(pushMock).toHaveBeenCalledWith('/onboard'))
    expect(pushMock).not.toHaveBeenCalledWith('/profile')
  })

  it('offers a direct route into practice through the grade picker', async () => {
    render(<ParentsPage />)
    ;(await screen.findByRole('button', { name: /just practise/ })).click()
    await waitFor(() => expect(pushMock).toHaveBeenCalledWith('/'))
  })

  it('establishes a guest identity before leaving, so the first answer is gradeable', async () => {
    render(<ParentsPage />)
    ;(await screen.findByRole('button', { name: 'Get started' })).click()
    await waitFor(() => expect(pushMock).toHaveBeenCalledWith('/onboard'))
    // The anchor that grades the first answer is keyed by student, and the
    // practice endpoint only resolves one from a bearer token — so the guest
    // id and token have to exist before the next page fetches anything.
    expect(ensureGuestId).toHaveBeenCalled()
    expect(ensureGuestToken).toHaveBeenCalled()
  })

  it('does not bounce a signed-in visitor, who may be a parent on the child\'s account', async () => {
    localStorage.setItem('orchid_token', 't')
    render(<ParentsPage />)
    await screen.findByRole('button', { name: 'Get started' })
    expect(replaceMock).not.toHaveBeenCalled()
  })

  it('keeps the two hero actions on separate, non-conflicting backgrounds', async () => {
    render(<ParentsPage />)
    const primary = await screen.findByRole('button', { name: 'Get started' })
    const ghost = await screen.findByRole('link', { name: 'View on GitHub' })
    // The ghost button was once the primary button plus overrides, so both
    // `bg-orchid-blue` and `bg-transparent` (and both hover backgrounds) sat on one
    // element. Same property, same specificity: Tailwind's emitted order decided the
    // winner, and the GitHub action hovered to blue over the artwork. The ghost now
    // carries its own transparent pair below lg and themed `lg:` overrides for the
    // desktop panel — compared as exact class tokens, because the legitimate
    // `lg:hover:bg-orchid-blue-faint` contains `bg-orchid-blue` as a substring.
    const primaryTokens = primary.className.split(/\s+/)
    const ghostTokens = ghost.className.split(/\s+/)
    expect(primaryTokens).toContain('bg-orchid-blue')
    expect(primaryTokens).toContain('hover:bg-orchid-blue-hover')
    expect(ghostTokens).toContain('bg-transparent')
    expect(ghostTokens).not.toContain('bg-orchid-blue')
    expect(ghostTokens).not.toContain('hover:bg-orchid-blue-hover')
  })

  it('lists Practice in the header and does not advertise the closed library', async () => {
    render(<ParentsPage />)
    const header = await screen.findByRole('banner')
    const learn = header.querySelector('a[href="/"]')
    expect(learn).not.toBeNull()
    // The landing hardcodes its own header rather than inheriting lib/nav.ts, so it is the one
    // place the product's own navigation decision could quietly disappear — in either
    // direction. Learn leads, and the closed reference library is not offered at all.
    const hrefs = [...header.querySelectorAll('a')].map((a) => a.getAttribute('href'))
    // The brand mark is the first anchor, so the ordering claim is about the nav entries:
    // Learn leads them.
    expect(hrefs.lastIndexOf('/')).toBeGreaterThan(0)
    expect(hrefs.lastIndexOf('/')).toBeLessThan(hrefs.indexOf('/how-it-works'))
    expect(hrefs).not.toContain('/study')
    expect([...header.querySelectorAll('a')].some((a) => /study/i.test(a.textContent ?? ''))).toBe(false)
  })
})

import { describe, it, expect, vi } from 'vitest'
import { render, screen, fireEvent } from '@testing-library/react'
import PrerequisitePathDialog from '../components/PrerequisitePathDialog'

vi.mock('next/link', () => ({
  default: ({ children, href }: { children: React.ReactNode; href: string }) => (
    <a href={href}>{children}</a>
  ),
}))

const prereqs = [
  { conceptId: 'arith.add', title: 'Single-digit addition', state: 'locked', met: false },
  { conceptId: 'arith.zero', title: 'Zero', state: 'upcoming', met: false },
]

describe('PrerequisitePathDialog', () => {
  it('shows the target title and a loading state', () => {
    render(
      <PrerequisitePathDialog
        open
        targetTitle="Subtraction"
        prerequisites={[]}
        status="loading"
        onClose={() => {}}
      />,
    )
    expect(screen.getByRole('dialog').textContent).toContain('Subtraction')
    expect(screen.getByText(/Loading prerequisites/)).toBeTruthy()
  })

  it('renders every prerequisite with a working Study link carrying the return stack', () => {
    render(
      <PrerequisitePathDialog
        open
        targetTitle="Subtraction"
        prerequisites={prereqs}
        status="ready"
        returnStack={['/study?concept=arith.sub']}
        onClose={() => {}}
      />,
    )
    const links = screen.getAllByRole('link', { name: /Study →/ })
    expect(links).toHaveLength(2)
    expect(links[0].getAttribute('href')).toBe(
      '/study?concept=arith.add&return=%2Fstudy%3Fconcept%3Darith.sub',
    )
    expect(links[1].getAttribute('href')).toBe(
      '/study?concept=arith.zero&return=%2Fstudy%3Fconcept%3Darith.sub',
    )
  })

  it('shows the empty state without fabricating prerequisites', () => {
    render(
      <PrerequisitePathDialog
        open
        targetTitle="Single-digit addition"
        prerequisites={[]}
        status="empty"
        onClose={() => {}}
      />,
    )
    expect(screen.getByText('No prerequisites listed')).toBeTruthy()
    expect(screen.queryByRole('link', { name: /Study →/ })).toBeNull()
  })

  it('shows an actionable error with retry', () => {
    const onRetry = vi.fn()
    render(
      <PrerequisitePathDialog
        open
        targetTitle="Subtraction"
        prerequisites={[]}
        status="error"
        error="Couldn't load this topic's prerequisites."
        onClose={() => {}}
        onRetry={onRetry}
      />,
    )
    expect(screen.getByRole('alert')).toBeTruthy()
    fireEvent.click(screen.getByRole('button', { name: 'Retry' }))
    expect(onRetry).toHaveBeenCalledTimes(1)
  })

  it('Close dismisses without changing learning state', () => {
    const onClose = vi.fn()
    render(
      <PrerequisitePathDialog
        open
        targetTitle="Subtraction"
        prerequisites={prereqs}
        status="ready"
        onClose={onClose}
      />,
    )
    fireEvent.click(screen.getByRole('button', { name: 'Close' }))
    expect(onClose).toHaveBeenCalledTimes(1)
  })

  it('Escape dismisses the dialog', () => {
    const onClose = vi.fn()
    render(
      <PrerequisitePathDialog
        open
        targetTitle="Subtraction"
        prerequisites={prereqs}
        status="ready"
        onClose={onClose}
      />,
    )
    fireEvent.keyDown(screen.getByRole('dialog'), { key: 'Escape' })
    expect(onClose).toHaveBeenCalledTimes(1)
  })
})

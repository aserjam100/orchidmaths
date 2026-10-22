import { describe, it, expect, vi } from 'vitest'
import { render, screen, fireEvent } from '@testing-library/react'
import GlossarySearch from '../components/GlossarySearch'

function renderSearch(overrides: Partial<Parameters<typeof GlossarySearch>[0]> = {}) {
  const props: Parameters<typeof GlossarySearch>[0] = {
    query: '',
    category: 'all',
    resultCount: 342,
    rangeStart: 1,
    rangeEnd: 20,
    mode: 'text',
    mathLatex: '',
    onQueryChange: vi.fn(),
    onCategoryChange: vi.fn(),
    onModeChange: vi.fn(),
    onMathChange: vi.fn(),
    onClear: vi.fn(),
    ...overrides,
  }
  render(<GlossarySearch {...props} />)
  return props
}

describe('glossary search input modes', () => {
  it('shows the text box by default with a math toggle', () => {
    renderSearch()
    expect(screen.getByRole('searchbox', { name: 'Search symbols, names or meanings' })).toBeVisible()
    expect(screen.getByRole('button', { name: 'Text' })).toHaveAttribute('aria-pressed', 'true')
    expect(screen.getByRole('button', { name: 'Math: f(x)' })).toHaveAttribute('aria-pressed', 'false')
  })

  it('notifies when the math mode is chosen', () => {
    const props = renderSearch()
    fireEvent.click(screen.getByRole('button', { name: 'Math: f(x)' }))
    expect(props.onModeChange).toHaveBeenCalledWith('math')
  })

  it('renders the math field instead of the text box in math mode', async () => {
    renderSearch({ mode: 'math' })
    expect(screen.queryByRole('searchbox', { name: 'Search symbols, names or meanings' })).toBeNull()
    // Either the MathLive field (stubbed element) or its loading/fallback
    // plain input — both carry the same accessible name by contract.
    const field = await screen.findByLabelText('Search by math notation')
    expect(field).toBeVisible()
  })
})

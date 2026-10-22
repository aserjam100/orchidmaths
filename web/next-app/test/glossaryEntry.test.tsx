import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import GlossaryEntryCard from '../components/GlossaryEntry'
import { GLOSSARY_ENTRIES } from '../lib/glossary/data'

// The displayed "Read as" label must appear exactly once per card. The data
// holds the bare reading and the component owns the prefix; this test renders
// both sides of that contract (one Greek letter, one non-Greek entry) so a
// duplicated prefix fails here instead of on every card in production.
describe('glossary entry reading label', () => {
  for (const id of ['alpha', 'power']) {
    it(`renders a single Read-as label for ${id}`, () => {
      const entry = GLOSSARY_ENTRIES.find(e => e.id === id)
      expect(entry?.reading).toBeTruthy()
      render(<GlossaryEntryCard entry={entry!} onSelectRelated={() => {}} />)
      const labels = screen.getAllByText(/Read as/)
      expect(labels).toHaveLength(1)
      expect(labels[0].textContent).toBe(`Read as ${entry!.reading}`)
    })
  }
})

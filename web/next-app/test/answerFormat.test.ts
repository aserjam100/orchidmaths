import { describe, it, expect } from 'vitest'
import { formatForGradingType, formatForQuestion } from '../lib/answerFormat'

// The field used to be shaped by the concept's grading type alone, which
// describes the grader and not the answer: "5 R 2", "2 sqrt(5)", "4:7" and
// "[[1,2],[3,4]]" are all numeric concepts, and a digits-only keypad cannot
// produce any of them. The server now sends what the generator declared.
describe('formatForQuestion', () => {
  it('prefers the declaration the server sent over the grading type', () => {
    expect(
      formatForQuestion({
        grading_type: 'numeric',
        answer_hint: 'Quotient and remainder, e.g. 5 R 2',
        answer_input_mode: 'text',
      }),
    ).toEqual({ hint: 'Quotient and remainder, e.g. 5 R 2', inputMode: 'text' })
  })

  it('falls back to the type for a payload that predates the declaration', () => {
    expect(formatForQuestion({ grading_type: 'numeric' })).toEqual(formatForGradingType('numeric'))
  })

  it('ignores an input mode it does not know', () => {
    expect(formatForQuestion({ grading_type: 'symbolic', answer_input_mode: 'voice' }).inputMode).toBe('text')
  })

  it('keeps the type hint when the server sends an empty one', () => {
    expect(formatForQuestion({ grading_type: 'tuple', answer_hint: '   ' }).hint).toBe(
      formatForGradingType('tuple').hint,
    )
  })
})

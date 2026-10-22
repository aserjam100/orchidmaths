import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'

vi.mock('../lib/api', async (importOriginal) => {
  const actual = await importOriginal<typeof import('../lib/api')>()
  return { ...actual, gradeLabAnswer: vi.fn() }
})

vi.mock('../components/math/MathInput', () => ({
  MathAnswerInput: ({
    value,
    onChange,
    id,
    ariaLabel,
    inputMode,
  }: {
    value: string
    onChange: (v: string) => void
    id?: string
    ariaLabel?: string
    inputMode?: string
  }) => (
    <input
      id={id}
      aria-label={ariaLabel}
      data-input-mode={inputMode}
      value={value}
      onChange={(e) => onChange(e.target.value)}
    />
  ),
}))

import LabQuestionCard from '../components/LabQuestionCard'
import { gradeLabAnswer } from '../lib/api'

const question = {
  question: 'Divide 47 by 5. Give the quotient and remainder.',
  answer: '9 R 2',
  explanation: '47 = 9 x 5 + 2',
  answer_hint: 'Quotient and remainder, e.g. 5 R 2',
  answer_input_mode: 'text' as const,
}

function renderCard() {
  return render(
    <ul>
      <LabQuestionCard
        question={question}
        conceptId="arith.div.remainder"
        gradingType="numeric"
        index={0}
      />
    </ul>,
  )
}

describe('LabQuestionCard', () => {
  beforeEach(() => {
    vi.mocked(gradeLabAnswer).mockReset()
  })

  it('hides the expected answer and gives the field the declared mode and hint', () => {
    renderCard()
    expect(screen.getByText(/Divide 47 by 5/)).toBeTruthy()
    expect(screen.getByText('Quotient and remainder, e.g. 5 R 2')).toBeTruthy()
    expect(screen.queryByText('9 R 2')).toBeNull()
    expect(screen.getByLabelText('Answer to question 1').getAttribute('data-input-mode')).toBe('text')
    fireEvent.click(screen.getByRole('button', { name: 'Show expected answer' }))
    expect(screen.getByText('9 R 2')).toBeTruthy()
  })

  it('grades with the expected answer and says that nothing was written', async () => {
    vi.mocked(gradeLabAnswer).mockResolvedValue({
      correct: true,
      score: 1,
      feedback: 'Correct.',
      graded_answer: '9 R 2',
      expected: '9 R 2',
      answer_hint: question.answer_hint,
      answer_input_mode: 'text',
      writes: false,
    })
    renderCard()
    fireEvent.change(screen.getByLabelText('Answer to question 1'), { target: { value: '9 R 2' } })
    fireEvent.click(screen.getByRole('button', { name: 'Grade this answer' }))
    await waitFor(() => expect(screen.getByText('Correct')).toBeTruthy())
    expect(vi.mocked(gradeLabAnswer)).toHaveBeenCalledWith('arith.div.remainder', '9 R 2', '9 R 2')
    expect(screen.getByText(/Writes nothing/)).toBeTruthy()
  })

  it('shows the diagnosis on an incorrect answer', async () => {
    vi.mocked(gradeLabAnswer).mockResolvedValue({
      correct: false,
      score: 0,
      feedback: 'The remainder is 2.',
      diagnosis: 'wrong_remainder',
      graded_answer: '9',
      expected: '9 R 2',
      writes: false,
    })
    renderCard()
    fireEvent.change(screen.getByLabelText('Answer to question 1'), { target: { value: '9' } })
    fireEvent.click(screen.getByRole('button', { name: 'Grade this answer' }))
    await waitFor(() => expect(screen.getByText('Incorrect')).toBeTruthy())
    expect(screen.getByText('Diagnosis: wrong_remainder')).toBeTruthy()
  })

  it('warns when the grader could not evaluate the answer', async () => {
    vi.mocked(gradeLabAnswer).mockResolvedValue({
      correct: false,
      score: 0,
      unavailable: true,
      graded_answer: '',
      expected: '9 R 2',
      writes: false,
    })
    renderCard()
    fireEvent.change(screen.getByLabelText('Answer to question 1'), { target: { value: '9 R 2' } })
    fireEvent.click(screen.getByRole('button', { name: 'Grade this answer' }))
    await waitFor(() => expect(screen.getByText(/not evidence either way/)).toBeTruthy())
  })
})

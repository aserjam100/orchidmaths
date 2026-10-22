import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import ReportReplay from '../components/ReportReplay'
import { gradeAdminReport } from '../lib/api'
import type { QuestionReport } from '../lib/api'

vi.mock('../lib/api', async (importOriginal) => ({
  ...(await importOriginal<typeof import('../lib/api')>()),
  gradeAdminReport: vi.fn(),
}))

// The real input mounts the math editor, which is not what this test is about:
// it is about the stored words reaching a grader the learner never sees.
vi.mock('../components/math/MathInput', () => ({
  MathAnswerInput: ({ value, onChange, id, ariaLabel }: {
    value: string
    onChange: (v: string) => void
    id?: string
    ariaLabel?: string
  }) => <input id={id} aria-label={ariaLabel} value={value} onChange={(e) => onChange(e.target.value)} />,
}))

function report(overrides: Partial<QuestionReport> = {}): QuestionReport {
  return {
    id: 7,
    reporter_id: 'stu-1',
    concept_id: 'arith.div.remainder',
    kind: 'question',
    question: 'Divide 47 by 5. Give the quotient and remainder.',
    expected: '9 R 2',
    explanation: '47 = 9 x 5 + 2',
    lesson_id: '',
    source: 'practice',
    session_id: 's1',
    attempt_id: '',
    answer_given: '9.2',
    grading_type: 'numeric',
    answer_hint: 'Quotient and remainder, e.g. 5 R 2',
    answer_input_mode: 'text',
    reason: 'wrong_answer',
    detail: 'I typed the decimal and it was marked wrong.',
    status: 'open',
    resolution: '',
    resolved_by: '',
    resolved_at: '',
    created_at: '2026-10-20 10:00',
    ...overrides,
  }
}

beforeEach(() => {
  vi.mocked(gradeAdminReport).mockReset()
})

describe('ReportReplay', () => {
  it('shows the learner answer, the expected answer and the format they faced', () => {
    render(<ReportReplay report={report()} />)

    expect(screen.getByText('9 R 2')).toBeInTheDocument()
    expect(screen.getByText('9.2')).toBeInTheDocument()
    expect(screen.getByText(/Quotient and remainder, e.g. 5 R 2/)).toBeInTheDocument()
    expect(screen.getByText(/text \(grading type: numeric\)/)).toBeInTheDocument()
  })

  it('grades the answer through the report id and shows the verdict', async () => {
    vi.mocked(gradeAdminReport).mockResolvedValue({
      correct: true,
      score: 1,
      feedback: 'Right.',
      graded_answer: '9 R 2',
      expected: '9 R 2',
      writes: false,
    })

    render(<ReportReplay report={report()} />)
    fireEvent.click(screen.getByRole('button', { name: /grade this answer/i }))

    await waitFor(() => expect(screen.getByText('Correct')).toBeInTheDocument())
    expect(gradeAdminReport).toHaveBeenCalledWith(7, '9.2')
    expect(screen.getByText(/Writes nothing/)).toBeInTheDocument()
  })

  it('says plainly when the answer was never captured', () => {
    render(<ReportReplay report={report({ answer_given: '' })} />)

    expect(screen.getByText(/Not stored: this report was filed before answers were captured/)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /grade this answer/i })).toBeDisabled()
  })

  it('restores the learner answer after an edit', () => {
    render(<ReportReplay report={report()} />)

    const input = screen.getByLabelText('Answer to replay') as HTMLInputElement
    fireEvent.change(input, { target: { value: 'something else' } })
    fireEvent.click(screen.getByRole('button', { name: /restore the learner/i }))

    expect(input.value).toBe('9.2')
  })

  it('refuses to guess when the report has no concept to grade against', () => {
    render(<ReportReplay report={report({ concept_id: '', expected: '' })} />)

    expect(screen.getByText(/nothing to replay/)).toBeInTheDocument()
    expect(screen.queryByRole('button', { name: /grade this answer/i })).not.toBeInTheDocument()
  })
})

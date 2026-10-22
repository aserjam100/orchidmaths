'use client'

import { useState } from 'react'
import KatexContent from './KatexContent'
import { MathAnswerInput } from './math/MathInput'
import { gradeLabAnswer, getErrorMessage, type AdminGradeRes, type LabQuestion } from '../lib/api'

// One generated question, tryable in place with the same field the learner gets.
//
// The expected answer starts hidden. A moderator who can see the answer cannot
// tell whether the question is answerable, and "can a person type this" is the
// exact failure being investigated here: a digits-only keypad in front of
// "5 R 2" looks fine to anyone who has already read the answer.
export default function LabQuestionCard({
  question,
  conceptId,
  gradingType,
  index,
}: {
  question: LabQuestion
  conceptId: string
  gradingType: string
  index: number
}) {
  const [answer, setAnswer] = useState('')
  const [result, setResult] = useState<AdminGradeRes | null>(null)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState('')
  const [revealed, setRevealed] = useState(false)

  const fieldID = `lab-answer-${conceptId}-${index}`

  async function grade() {
    if (busy || !answer.trim()) return
    setBusy(true)
    setError('')
    try {
      setResult(await gradeLabAnswer(conceptId, question.answer, answer))
    } catch (e) {
      setError(getErrorMessage(e))
    } finally {
      setBusy(false)
    }
  }

  return (
    <li className="border border-orchid-border bg-orchid-surface">
      <div className="border-b border-orchid-border px-4 py-3">
        <div className="mb-2 flex items-baseline justify-between gap-3">
          <span className="font-mono text-[11px] uppercase tracking-wider text-orchid-muted">Question {index + 1}</span>
          <span className="font-mono text-[11px] text-orchid-muted">{question.answer_input_mode}</span>
        </div>
        <KatexContent className="text-orchid-primary text-sm">{question.question}</KatexContent>
      </div>

      <div className="p-4">
        <label htmlFor={fieldID} className="mb-1 block font-mono text-[11px] uppercase tracking-wider text-orchid-muted">
          Answer
        </label>
        <MathAnswerInput
          id={fieldID}
          value={answer}
          onChange={setAnswer}
          onSubmit={grade}
          gradingType={gradingType}
          conceptId={conceptId}
          inputMode={question.answer_input_mode}
          disabled={busy}
          ariaLabel={`Answer to question ${index + 1}`}
          placeholder="Your answer"
        />
        <p className="mt-2 font-mono text-[12px] text-orchid-secondary">{question.answer_hint}</p>

        {error && <p role="alert" className="mt-3 font-mono text-[12px] text-orchid-red">{error}</p>}

        <div className="mt-3 flex flex-wrap items-center gap-2">
          <button
            type="button"
            onClick={grade}
            disabled={busy || !answer.trim()}
            className="border border-orchid-primary px-3 h-9 font-mono text-xs text-orchid-primary hover:bg-orchid-hover disabled:opacity-50"
          >
            {busy ? 'Grading…' : 'Grade this answer'}
          </button>
          <button
            type="button"
            onClick={() => setRevealed((v) => !v)}
            className="border border-orchid-border px-3 h-9 font-mono text-xs text-orchid-secondary hover:text-orchid-primary"
          >
            {revealed ? 'Hide expected answer' : 'Show expected answer'}
          </button>
          <span className="font-mono text-[11px] text-orchid-muted">Writes nothing: no XP, streak, attempt or mastery.</span>
        </div>

        {revealed && (
          <dl className="mt-3 border border-orchid-border bg-orchid-code p-3 font-mono text-[12px]">
            <dt className="text-orchid-muted">Expected answer</dt>
            <dd className="break-words text-orchid-primary">{question.answer}</dd>
          </dl>
        )}

        {result && (
          <div className="mt-3 border border-orchid-border bg-orchid-code p-3">
            <div className="mb-2 flex items-baseline justify-between gap-3">
              <span className={result.correct ? 'font-mono text-xs text-orchid-green' : 'font-mono text-xs text-orchid-red'}>
                {result.correct ? 'Correct' : 'Incorrect'}
              </span>
              <span className="font-mono text-[11px] text-orchid-muted">score {result.score.toFixed(2)}</span>
            </div>
            {result.feedback && <KatexContent className="mt-2 text-sm text-orchid-secondary">{result.feedback}</KatexContent>}
            {result.diagnosis && (
              <p className="mt-2 font-mono text-[12px] text-orchid-secondary">Diagnosis: {result.diagnosis}</p>
            )}
            {result.unavailable && (
              <p className="mt-2 font-mono text-[12px] text-orchid-red">
                The grader could not evaluate this answer, so a wrong verdict is not evidence either way.
              </p>
            )}
            <p className="mt-2 font-mono text-[11px] text-orchid-muted">
              Graded {JSON.stringify(result.graded_answer)} against {JSON.stringify(result.expected)}
            </p>
          </div>
        )}
      </div>
    </li>
  )
}

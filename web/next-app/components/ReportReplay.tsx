'use client'

import { useState } from 'react'
import KatexContent from './KatexContent'
import { MathAnswerInput } from './math/MathInput'
import { gradeAdminReport, getErrorMessage, type AdminGradeRes, type QuestionReport } from '../lib/api'
import { formatForQuestion } from '../lib/answerFormat'

// Replaying a report is the only way to answer the two questions a moderator
// actually has: was the learner right, and is it still wrong? The stored words
// are what make that possible — the question, the expected answer and the
// concept — because a generator instance cannot be rebuilt from a seed.
//
// The grade endpoint is pure by construction: it calls GradeAnswer directly and
// never enters SubmitAnswer, so replaying writes no attempt, XP, streak or
// mastery. The panel says so, because an operator should not have to trust a
// comment to know whether pressing a button changed a learner's record.
export default function ReportReplay({ report }: { report: QuestionReport }) {
  const [answer, setAnswer] = useState(report.answer_given ?? '')
  const [result, setResult] = useState<AdminGradeRes | null>(null)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState('')

  const format = formatForQuestion({
    grading_type: report.grading_type,
    answer_hint: report.answer_hint,
    answer_input_mode: report.answer_input_mode,
  })
  const canReplay = Boolean(report.concept_id && report.expected)

  async function replay() {
    if (busy) return
    setBusy(true)
    setError('')
    try {
      setResult(await gradeAdminReport(report.id, answer))
    } catch (e) {
      setError(getErrorMessage(e))
    } finally {
      setBusy(false)
    }
  }

  return (
    <section className="mt-4 border border-orchid-border bg-orchid-surface">
      <header className="border-b border-orchid-border px-4 py-2">
        <h3 className="font-mono text-[11px] uppercase tracking-wider text-orchid-muted">Replay the attempt</h3>
      </header>

      <div className="p-4">
        {!canReplay ? (
          <p className="font-mono text-[12px] text-orchid-secondary">
            This report carries no concept or expected answer, so there is nothing to replay.
          </p>
        ) : (
          <>
            <dl className="grid grid-cols-1 gap-2 font-mono text-[12px] text-orchid-secondary sm:grid-cols-2">
              <div>
                <dt className="text-orchid-muted">Concept</dt>
                <dd className="break-all text-orchid-primary">{report.concept_id}</dd>
              </div>
              <div>
                <dt className="text-orchid-muted">Field the learner had</dt>
                <dd className="text-orchid-primary">
                  {format.inputMode}
                  {report.grading_type ? ` (grading type: ${report.grading_type})` : ''}
                </dd>
              </div>
              <div className="sm:col-span-2">
                <dt className="text-orchid-muted">Expected answer</dt>
                <dd className="break-words text-orchid-primary">{report.expected}</dd>
              </div>
              {report.answer_given ? (
                <div className="sm:col-span-2">
                  <dt className="text-orchid-muted">What the learner typed</dt>
                  <dd className="break-words text-orchid-primary">{report.answer_given}</dd>
                </div>
              ) : (
                <div className="sm:col-span-2">
                  <dt className="text-orchid-muted">What the learner typed</dt>
                  <dd className="text-orchid-secondary">
                    Not stored: this report was filed before answers were captured. Type the answer from the report
                    text, or ask the reporter.
                  </dd>
                </div>
              )}
            </dl>

            <div className="mt-4">
              <label htmlFor={`replay-${report.id}`} className="mb-1 block font-mono text-[11px] uppercase tracking-wider text-orchid-muted">
                Answer to grade
              </label>
              <MathAnswerInput
                id={`replay-${report.id}`}
                value={answer}
                onChange={setAnswer}
                onSubmit={replay}
                gradingType={report.grading_type}
                conceptId={report.concept_id}
                inputMode={format.inputMode}
                disabled={busy}
                ariaLabel="Answer to replay"
                placeholder="Your answer"
              />
              <p className="mt-2 font-mono text-[12px] text-orchid-secondary">{format.hint}</p>
            </div>

            {error && <p role="alert" className="mt-3 font-mono text-[12px] text-orchid-red">{error}</p>}

            <div className="mt-3 flex flex-wrap items-center gap-2">
              <button
                type="button"
                onClick={replay}
                disabled={busy || !answer.trim()}
                className="border border-orchid-primary px-3 h-9 font-mono text-xs text-orchid-primary hover:bg-orchid-hover disabled:opacity-50"
              >
                {busy ? 'Grading…' : 'Grade this answer'}
              </button>
              {report.answer_given && answer !== report.answer_given && (
                <button
                  type="button"
                  onClick={() => { setAnswer(report.answer_given ?? ''); setResult(null) }}
                  disabled={busy}
                  className="border border-orchid-border px-3 h-9 font-mono text-xs text-orchid-secondary disabled:opacity-50"
                >
                  Restore the learner&apos;s answer
                </button>
              )}
              <span className="font-mono text-[11px] text-orchid-muted">Writes nothing: no XP, streak, attempt or mastery.</span>
            </div>

            {result && (
              <div className="mt-4 border border-orchid-border bg-orchid-code p-3">
                <p className="font-mono text-[11px] uppercase tracking-wider text-orchid-muted">Replay verdict</p>
                <p className={`mt-1 font-mono text-sm ${result.correct ? 'text-orchid-green' : 'text-orchid-red'}`}>
                  {result.correct ? 'Correct' : 'Incorrect'}
                  <span className="ml-2 text-orchid-muted">score {result.score.toFixed(2)}</span>
                </p>
                {result.feedback && (
                  <KatexContent className="mt-2 text-sm text-orchid-secondary">{result.feedback}</KatexContent>
                )}
                {result.diagnosis && (
                  <p className="mt-2 font-mono text-[12px] text-orchid-secondary">Diagnosis: {result.diagnosis}</p>
                )}
                {result.unavailable && (
                  <p className="mt-2 font-mono text-[12px] text-orchid-red">
                    The grader could not evaluate this answer. The stored verdict is not evidence either way.
                  </p>
                )}
                <p className="mt-2 break-words font-mono text-[12px] text-orchid-muted">
                  Graded {JSON.stringify(result.graded_answer)} against {JSON.stringify(result.expected)}
                </p>
              </div>
            )}
          </>
        )}
      </div>
    </section>
  )
}

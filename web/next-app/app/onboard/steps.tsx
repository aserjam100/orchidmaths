'use client'

import type { RefObject } from 'react'
import Link from 'next/link'
import KatexContent from '../../components/KatexContent'
import SectionHeader from '../../components/SectionHeader'
import ProgressBar from '../../components/ProgressBar'
import ReportButton from '../../components/ReportButton'
import BriefingCard, { DIAGNOSTIC_BRIEFING } from '../../components/BriefingCard'
import SubmitErrorBlock, { type SubmitError } from '../../components/SubmitErrorBlock'
import DiagnosticResults from '../../components/DiagnosticResults'
import Loading from '../../components/Loading'
import type { DiagnosticProgress, GoalPlanRes } from '../../lib/api'
import type { AnswerFormat } from '../../lib/answerFormat'
import ChoiceOptions from '../../components/ChoiceOptions'
import type { PrimaryLevel } from '../../lib/primaryLevel'
import { MathAnswerInput, type MathFocusHandle } from '@/components/math/MathInput'

export function WelcomeStep({
  levels,
  selectedLevel,
  loading,
  hasPaused,
  selectedCount,
  confirming,
  onSelect,
  onStart,
  onResume,
  onBegin,
  onCancel,
}: {
  levels: PrimaryLevel[]
  selectedLevel: string | null
  loading: boolean
  hasPaused: boolean
  selectedCount: number
  confirming: boolean
  onSelect: (level: string) => void
  onStart: () => void
  onResume: () => void
  onBegin: () => void
  onCancel: () => void
}) {
  if (confirming) {
    const chosen = levels.find(l => l.level === selectedLevel)
    return (
      <div className="max-w-2xl mx-auto px-2">
        <div className="text-center mb-4 mt-6 sm:mt-8 min-w-0">
          <div className="font-mono text-[12px] uppercase text-orchid-muted mb-3">One quick check</div>
          <h1 className="font-serif text-2xl sm:text-3xl font-medium text-orchid-primary px-2">
            Before you begin
          </h1>
          <p className="text-orchid-secondary text-sm mt-3 px-2">
            Starting test on: {chosen?.name ?? 'your level'} · {selectedCount} concepts in scope
          </p>
          <p className="text-orchid-muted text-xs mt-2 px-2">
            The test adapts as it goes and usually takes about 30–45 minutes.
          </p>
        </div>
        <BriefingCard eyebrow="What to expect" items={DIAGNOSTIC_BRIEFING} />
        <div className="mt-6 flex flex-col sm:flex-row gap-3 justify-center px-2">
          <button
            type="button"
            onClick={onBegin}
            disabled={loading}
            className="border border-orchid-blue bg-orchid-blue text-white hover:opacity-90 rounded-none h-12 min-h-[44px] px-6 sm:px-10 font-medium text-sm disabled:opacity-50"
          >
            {loading ? (<><Loading inline size={13} /> Loading…</>) : 'Begin diagnostic →'}
          </button>
          <button
            type="button"
            onClick={onCancel}
            disabled={loading}
            className="font-mono text-xs text-orchid-muted hover:text-orchid-primary min-h-[44px] px-4 disabled:opacity-50"
          >
            ← Back
          </button>
        </div>
      </div>
    )
  }
  return (
    <div className="max-w-4xl mx-auto">
      <div className="text-center mb-4 mt-6 sm:mt-8 px-2 min-w-0">
        <div className="font-mono text-[12px] uppercase text-orchid-muted mb-3">Welcome to Orchid</div>
        <h1 className="font-serif text-2xl sm:text-4xl font-medium text-orchid-primary px-2">
          Let&apos;s find your starting point
        </h1>
        <p className="text-orchid-secondary text-sm mt-3 max-w-[500px] mx-auto px-2">
          Which Primary level is your child in? The test covers that level and the ones before it,
          then practice starts where your child is.
        </p>
      </div>

      <div role="radiogroup" aria-label="Primary level" className="grid grid-cols-2 sm:grid-cols-3 gap-2 sm:gap-3 mb-8 min-w-0">
        {levels.map(l => {
          const selected = l.level === selectedLevel
          return (
            <button
              type="button"
              role="radio"
              aria-checked={selected}
              key={l.level}
              onClick={() => onSelect(l.level)}
              className={`rounded-none p-3 sm:p-4 text-left transition-all text-sm min-h-[60px] min-w-0 overflow-hidden ${
                selected
                  ? 'bg-orchid-blue text-white'
                  : 'bg-orchid-surface border border-orchid-border text-orchid-secondary hover:border-orchid-blue hover:text-orchid-blue'
              }`}
            >
              <div className="font-medium text-[15px] sm:text-sm">{l.name}</div>
              <div className={`font-mono text-[12px] mt-1 line-clamp-2 ${selected ? 'text-white/70' : 'text-orchid-muted'}`}>
                {l.topics.slice(0, 3).map(t => t.topic).join(' · ')}
              </div>
            </button>
          )
        })}
      </div>

      <div className="text-center px-4">
        {hasPaused && (
          <button
            type="button"
            onClick={onResume}
            disabled={loading}
            className="border border-orchid-blue bg-orchid-blue text-white hover:opacity-90 rounded-none h-12 min-h-[44px] px-6 sm:px-10 font-medium text-sm disabled:opacity-50 max-w-full mb-3"
          >
            {loading ? (<><Loading inline size={13} /> Loading…</>) : 'Continue diagnostic →'}
          </button>
        )}
        <button
          type="button"
          onClick={onStart}
          disabled={selectedCount === 0 || loading}
          className="border border-orchid-blue text-orchid-blue hover:bg-orchid-blue-faint rounded-none h-12 min-h-[44px] px-6 sm:px-10 font-medium text-sm disabled:opacity-50 max-w-full"
        >
          {loading ? (<><Loading inline size={13} /> Loading…</>) : `Start diagnostic test (${selectedCount} concepts in scope)`}
        </button>
      </div>

      {/* The way out, and the reason it is a link rather than a flag.
          A learner can start any topic whose prerequisites are met without ever
          taking the diagnostic: /learn?concept=… is gated server-side, not by
          this flag, and /levels is already browsable. Nothing here records a
          result, seeds placement, or grants credit — it only makes an existing
          path findable from the screen a learner is actually looking at. */}
      <div className="text-center px-4 mt-6">
        <Link
          href="/levels"
          className="font-mono text-xs text-orchid-muted hover:text-orchid-blue underline underline-offset-2 min-h-[44px] inline-flex items-center"
        >
          Browse topics by level and start without the test →
        </Link>
      </div>
    </div>
  )
}

export function DiagnosticStep({
  question,
  conceptName,
  progress,
  questionCount,
  answerInput,
  lastResult,
  accuracy,
  loading,
  inputRef,
  conceptId,
  sessionId,
  onInputChange,
  onSubmit,
  onDontKnow,
  onNext,
  done,
  answerFormat,
  submitError,
  planError,
  onSkip,
  skipsLeft,
  onRestart,
  onRetryPlan,
  retryAvailable,
  onRetryQuestion,
}: {
  question: string
  conceptName: string
  progress: DiagnosticProgress | null
  questionCount: number
  answerInput: string
  lastResult: { correct: boolean; feedback: string } | null
  accuracy: { correct: number; total: number }
  loading: boolean
  inputRef: RefObject<MathFocusHandle | null>
  conceptId: string
  sessionId: string
  onInputChange: (value: string) => void
  onSubmit: () => void
  onDontKnow: () => void
  onNext: () => void
  done: boolean
  answerFormat: AnswerFormat
  submitError: SubmitError | null
  planError: string
  onSkip: () => void
  skipsLeft: number
  onRestart: () => void
  onRetryPlan: () => void
  retryAvailable: boolean
  onRetryQuestion: () => void
}) {
  return (
    <>
      <SectionHeader label={`Question ${progress ? progress.answered + 1 : questionCount}`} title={conceptName} />
      <div className="max-w-2xl mx-auto px-2 sm:px-0 min-w-0 overflow-hidden">
        <ProgressBar
          answered={progress ? progress.answered + 1 : questionCount}
          coverDone={progress ? progress.cover_done : 0}
          coverSize={progress ? progress.cover_size : 0}
        />
        <div className="bg-orchid-surface border border-orchid-border rounded-none p-4 sm:p-6 mb-6 w-full max-w-full min-w-0 overflow-hidden">
          <div className="bg-orchid-code border border-orchid-border rounded-none p-4 sm:p-6 text-center mb-4 w-full max-w-full min-w-0 overflow-hidden">
            <div className="w-full max-w-full min-w-0 overflow-hidden">
              <KatexContent className="text-orchid-primary text-lg font-mono font-light whitespace-pre-wrap break-words">{question}</KatexContent>
            </div>
          </div>
          <ChoiceOptions question={question} value={answerInput} onPick={onInputChange} disabled={loading || lastResult !== null} />
          <form
            onSubmit={e => { e.preventDefault(); onSubmit() }}
            className="flex flex-col sm:flex-row gap-3 min-w-0"
          >
            <label htmlFor="onboard-answer" className="sr-only">Your answer</label>
            <MathAnswerInput
              id="onboard-answer"
              value={answerInput}
              onChange={onInputChange}
              onSubmit={onSubmit}
              conceptId={conceptId}
              disabled={loading || lastResult !== null}
              status={lastResult ? (lastResult.correct ? 'correct' : 'incorrect') : 'default'}
              inputMode={answerFormat.inputMode}
              placeholder="Your answer..."
              className="sm:flex-1"
              focusRef={inputRef}
            />
            <button
              type="submit"
              disabled={!answerInput.trim() || loading || lastResult !== null}
              className="border border-orchid-blue text-orchid-blue hover:bg-orchid-blue-faint rounded-none h-12 min-h-[36px] px-6 font-medium text-sm disabled:opacity-50 shrink-0 w-auto self-end sm:self-auto"
            >
              Check Answer
            </button>
          </form>
          <div className="mt-2 flex items-center justify-between gap-2">
            <p className="font-mono text-[12px] text-orchid-muted">{answerFormat.hint}</p>
            {!lastResult && !loading && (
              <button
                type="button"
                onClick={onDontKnow}
                className="shrink-0 font-mono text-[12px] text-orchid-muted hover:text-orchid-red underline underline-offset-2"
              >
                I don&apos;t know
              </button>
            )}
          </div>
          <div className="mt-2 flex justify-end">
            <ReportButton
              key={question}
              conceptId={conceptId}
              kind="question"
              question={question}
              source="diagnostic"
              sessionId={sessionId}
              answerGiven={answerInput}
              answerHint={answerFormat.hint}
              answerInputMode={answerFormat.inputMode}
            />
          </div>
        </div>

        {lastResult && (
          <div className={`bg-orchid-surface border rounded-none p-4 mb-4 text-center ${lastResult.correct ? 'border-orchid-green' : 'border-orchid-red'}`}>
            <KatexContent className={lastResult.correct ? 'text-orchid-green' : 'text-orchid-red'}>{lastResult.feedback}</KatexContent>
            {!lastResult.correct && !done && retryAvailable && (
              <button
                type="button"
                onClick={onRetryQuestion}
                disabled={loading}
                className="mt-2 font-mono text-[12px] text-orchid-muted hover:text-orchid-blue underline underline-offset-2 disabled:opacity-50"
              >
                I made a silly mistake — retry
              </button>
            )}
          </div>
        )}

        {submitError ? (
          <div className="mb-4">
            <SubmitErrorBlock
              error={submitError}
              onRetry={onSubmit}
              onSkip={onSkip}
              skipsLeft={skipsLeft}
              onRestart={onRestart}
              restartLabel="Restart diagnostic"
              retrying={loading}
            />
          </div>
        ) : lastResult && !done ? (
          <div className="mb-4 text-center">
            <button
              type="button"
              autoFocus
              onClick={onNext}
              className="border border-orchid-blue text-orchid-blue hover:bg-orchid-blue-faint rounded-none h-12 px-8 font-medium text-sm shrink-0"
            >
              Next →
            </button>
          </div>
        ) : lastResult && done && planError ? (
          <div className="mb-4 border border-orchid-red bg-orchid-surface p-4 text-left" role="alert">
            <p className="font-mono text-xs text-orchid-red">Couldn&apos;t load your results.</p>
            <p className="mt-1 font-mono text-[12px] text-orchid-muted break-words [overflow-wrap:anywhere]">{planError}</p>
            <div className="mt-3">
              <button
                type="button"
                onClick={onRetryPlan}
                disabled={loading}
                className="border border-orchid-blue text-orchid-blue hover:bg-orchid-blue-faint px-4 h-10 font-mono text-xs disabled:opacity-50"
              >
                {loading ? 'Loading…' : 'Load results'}
              </button>
            </div>
          </div>
        ) : lastResult && done ? (
          <div className="mb-4 text-center text-orchid-muted text-xs font-mono">
            Preparing results…
          </div>
        ) : null}

        <div className="text-center text-orchid-muted text-xs font-mono">
          {accuracy.total > 0 && `${accuracy.correct}/${accuracy.total} correct`}
        </div>
      </div>
    </>
  )
}

export function ResultsStep({ plan, onStartPractice }: { plan: GoalPlanRes; onStartPractice: () => void }) {
  return (
    <div className="max-w-3xl mx-auto min-w-0 overflow-hidden px-2 sm:px-0">
      <SectionHeader label="Your results" title="Here's what we found" />
      <div className="mt-6 min-w-0 overflow-hidden">
        <DiagnosticResults plan={plan} onStartPractice={onStartPractice} />
      </div>
    </div>
  )
}

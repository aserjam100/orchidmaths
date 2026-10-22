'use client'

import { Suspense, useEffect, useState } from 'react'
import { useSearchParams } from 'next/navigation'
import Link from 'next/link'
import Header from '../../components/Header'
import BottomTabs from '../../components/BottomTabs'
import Footer from '../../components/Footer'
import LearnStepper from '../../components/LearnStepper'
import LessonCredits from '../../components/LessonCredits'
import GradePicker from '../../components/GradePicker'
import type { Shelf, ShelfItem } from '../../lib/nextUp'
import { usePrimaryLevel } from '../../hooks/usePrimaryLevel'
import { getLevel, levelInfo } from '../../lib/primaryLevel'
import { fetchShelf } from '../../lib/recommendations'
import { parseReturnStack } from '../../lib/learnReturn'

/**
 * The practice hub: the learner's grade, one big button that starts the engine's pick, and a
 * route to choosing a topic instead.
 *
 * The engine still chooses (one `/api/next` call); this page only decides how much of its answer
 * to show. A new learner's head is the placement test, and a student who came to practise should
 * not be sent to a test first, so the hub starts the first concept instead and offers the test in
 * one line underneath. The learner disposes: the alternatives and Topics are one tap away.
 */
function LearnEntry() {
  const [level] = usePrimaryLevel()
  const [checked, setChecked] = useState(false)
  const [picking, setPicking] = useState(false)
  // undefined while loading, null when the request failed.
  const [shelf, setShelf] = useState<Shelf | null | undefined>(undefined)

  // With no grade yet, the page opens on the picker, which is also what holds the shelf fetch
  // until the server has the grade.
  useEffect(() => {
    if (!getLevel()) setPicking(true)
    setChecked(true)
  }, [])

  // Not while picking: the new level reaches this browser before the server has saved it, and
  // a shelf fetched in between would be scoped to the old grade. Closing the picker refetches.
  useEffect(() => {
    if (!level || picking) return
    let cancelled = false
    setShelf(undefined)
    fetchShelf()
      .then(s => { if (!cancelled) setShelf(s) })
      .catch(() => { if (!cancelled) setShelf(null) })
    return () => { cancelled = true }
  }, [level, picking])

  if (!checked) {
    return <div className="p-6 font-mono text-xs text-orchid-muted">Loading…</div>
  }

  if (picking || !level) {
    return (
      <div className="mt-6 mb-16 px-4">
        <div className="max-w-xl mx-auto">
          <h1 className="font-serif text-2xl text-orchid-primary">Which grade are you in?</h1>
          <p className="mt-2 mb-6 text-sm text-orchid-secondary">Practice will stay at your grade. You can change it any time.</p>
          <GradePicker onChosen={() => setPicking(false)} />
          {level && (
            <button type="button" onClick={() => setPicking(false)} className="mt-4 font-mono text-xs text-orchid-secondary hover:text-orchid-blue">
              ← Keep {levelInfo(level)?.name}
            </button>
          )}
        </div>
      </div>
    )
  }

  const placement = shelf?.next.href === '/onboard'
  const start: ShelfItem | null = shelf ? (placement ? shelf.alternatives[0] ?? null : shelf.next) : null
  const others = shelf ? (placement ? shelf.alternatives.slice(1) : shelf.alternatives).slice(0, 3) : []

  return (
    <div className="mt-6 mb-16 px-4">
      <div className="max-w-xl mx-auto">
        <div className="flex items-baseline justify-between gap-3">
          <h1 className="font-serif text-2xl text-orchid-primary">{levelInfo(level)?.name} practice</h1>
          <button type="button" onClick={() => setPicking(true)} className="shrink-0 font-mono text-xs text-orchid-secondary hover:text-orchid-blue underline underline-offset-2">
            Change grade
          </button>
        </div>

        {shelf === undefined && <p className="mt-6 font-mono text-xs text-orchid-muted">Finding your next question…</p>}
        {shelf === null && <p className="mt-6 font-mono text-xs text-orchid-muted">Could not reach Orchid. Check your connection and reload.</p>}

        {start && (
          <Link
            href={start.href}
            className="mt-6 block border-2 border-orchid-blue bg-orchid-blue px-6 py-6 text-white transition-colors hover:bg-orchid-blue-hover min-w-0"
          >
            <span className="block font-serif text-2xl leading-tight">{start.kind === 'review' ? 'Start review' : 'Start practising'} →</span>
            <span className="mt-2 block font-mono text-[13px] text-white/85 break-words [overflow-wrap:anywhere]">{start.title}</span>
          </Link>
        )}
        {shelf && !start && (
          <p className="mt-6 text-sm text-orchid-secondary">Nothing is waiting right now. Pick a topic below to keep going.</p>
        )}

        {others.length > 0 && (
          <div className="mt-6">
            <h2 className="font-mono text-[12px] uppercase tracking-wider text-orchid-muted mb-2">Or try</h2>
            <ul className="border border-orchid-border divide-y divide-orchid-border bg-orchid-surface">
              {others.map((it, i) => (
                <li key={`${it.href}:${i}`}>
                  <Link href={it.href} className="flex items-center justify-between gap-3 px-4 py-3 min-h-[48px] font-mono text-xs text-orchid-primary hover:text-orchid-blue min-w-0">
                    <span className="min-w-0 break-words [overflow-wrap:anywhere]">{it.title}</span>
                    <span aria-hidden="true" className="shrink-0 text-orchid-blue">→</span>
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        )}

        <Link href="/levels" className="mt-4 flex items-center justify-between border border-orchid-border px-4 py-3 min-h-[48px] font-mono text-xs text-orchid-secondary hover:border-orchid-blue hover:text-orchid-blue">
          <span>Choose a topic</span>
          <span aria-hidden="true">→</span>
        </Link>

        {placement && (
          <p className="mt-6 font-mono text-[12px] text-orchid-muted">
            Already know the earlier grades?{' '}
            <Link href="/onboard" className="text-orchid-blue hover:text-orchid-blue-hover underline underline-offset-2">
              Take a short check to skip ahead
            </Link>
          </p>
        )}
      </div>
    </div>
  )
}

function LearnContent() {
  const searchParams = useSearchParams()
  const concept = searchParams.get('concept') || ''
  const returnStack = parseReturnStack(searchParams)

  if (!concept) {
    return <LearnEntry />
  }

  return (
    <div className="mt-6 mb-16 px-4">
      <div className="max-w-2xl mx-auto mb-4">
        <Link href="/learn" className="font-mono text-xs text-orchid-secondary hover:text-orchid-blue">← Practice</Link>
      </div>
      <LearnStepper conceptId={concept} returnStack={returnStack} />
      <div className="max-w-2xl mx-auto">
        <LessonCredits conceptId={concept} />
      </div>
    </div>
  )
}

export default function LearnPage() {
  return (
    <>
      <Header />
      <div className="pt-[var(--chrome-top)] lg:pt-0">
        <Suspense fallback={<div className="p-6 font-mono text-xs text-orchid-muted">Loading…</div>}>
          <LearnContent />
        </Suspense>
      </div>
      <BottomTabs />
      <Footer />
    </>
  )
}

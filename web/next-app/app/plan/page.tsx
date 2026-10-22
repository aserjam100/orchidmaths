'use client'

import { Suspense, useEffect, useState } from 'react'
import Link from 'next/link'
import { useSearchParams } from 'next/navigation'
import Header from '../../components/Header'
import BottomTabs from '../../components/BottomTabs'
import Footer from '../../components/Footer'
import PlanEditor from '../../components/PlanEditor'
import { isLoggedIn } from '../../lib/auth'

// Learning planner: destination + (daily effort XOR deadline) → estimate.
// Dates render as months (estimates, never promises); quiz eligibility and
// mastery are independent of every number on this page. The editor is shared
// with Settings (compact there); this page keeps the full estimate display.
//
// The page used to end at "← Back to profile", which was the only exit and was
// useless to the visitors most likely to arrive here: the assistant links
// straight to a destination, and a guest following that link has no profile to
// go back to. A plan states a pace; it does not say where to begin. The block at
// the bottom is the missing next step, and it is shown to everyone because
// /learn and /onboard both work for a guest.
function PlanContent() {
  const params = useSearchParams()
  // Set by an assistant link, so "start with calculus" lands on the planner
  // with calculus selected instead of on whatever sorts first.
  const destination = params.get('destination') ?? undefined
  // Read the session after mount rather than during render: isLoggedIn() reads
  // localStorage, and the static export renders this page on the server first.
  const [registered, setRegistered] = useState(false)
  useEffect(() => {
    setRegistered(isLoggedIn())
  }, [])
  const [justSaved, setJustSaved] = useState(false)

  return (
    <div className="max-w-2xl mx-auto mt-6 mb-16 px-4">
      <p className="font-mono text-[12px] uppercase tracking-wider text-orchid-muted">Learning planner</p>
      <h1 className="font-serif text-2xl text-orchid-primary mt-1">Plan your learning</h1>
      <PlanEditor initialDestination={destination} onSaved={() => setJustSaved(true)} />

      {justSaved && (
        <p role="status" className="mt-5 border border-orchid-border bg-orchid-surface px-3 py-2.5 font-mono text-xs text-orchid-secondary">
          Plan saved. A pace is not a starting point — use the diagnostic below to find out
          which concept you belong at.
        </p>
      )}

      <section className="mt-8 border-t border-orchid-border pt-4">
        <h2 className="font-mono text-[12px] uppercase tracking-wider text-orchid-muted">What to do next</h2>
        <p className="mt-1 text-sm text-orchid-secondary">
          A plan sets how much you study per day. It does not say where to begin: the
          diagnostic measures what you already know, then starts you at the right concept.
        </p>
        <div className="mt-3 flex flex-wrap gap-2">
          <Link
            href="/onboard"
            className="inline-flex min-h-[44px] items-center border border-orchid-blue px-4 py-2.5 font-mono text-xs text-orchid-blue hover:bg-orchid-blue-faint"
          >
            Take the diagnostic
          </Link>
          <Link
            href="/learn"
            className="inline-flex min-h-[44px] items-center border border-orchid-border px-4 py-2.5 font-mono text-xs text-orchid-secondary hover:border-orchid-blue hover:text-orchid-blue"
          >
            Start learning
          </Link>
          {registered && (
            <Link
              href="/profile"
              className="inline-flex min-h-[44px] items-center border border-orchid-border px-4 py-2.5 font-mono text-xs text-orchid-secondary hover:border-orchid-blue hover:text-orchid-blue"
            >
              Back to profile
            </Link>
          )}
        </div>
      </section>
    </div>
  )
}

export default function PlanPage() {
  return (
    <>
      <Header />
      <div className="pt-[var(--chrome-top)] lg:pt-0">
        <Suspense fallback={<div className="p-6 font-mono text-xs text-orchid-muted">Loading…</div>}>
          <PlanContent />
        </Suspense>
      </div>
      <BottomTabs />
      <Footer />
    </>
  )
}

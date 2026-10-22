'use client'

import { Suspense } from 'react'
import Link from 'next/link'
import { useSearchParams } from 'next/navigation'
import Header from '../../../../components/Header'
import BottomTabs from '../../../../components/BottomTabs'
import Footer from '../../../../components/Footer'
import SectionHeader from '../../../../components/SectionHeader'
import AdminNav from '../../../../components/AdminNav'
import LearnStepper from '../../../../components/LearnStepper'

// The learner route for one concept, mounted from the lab.
//
// This is the same component /learn?concept=... renders, not a mock-up of it, so
// an operator can see what a learner sees: the knowledge points, the served
// question, the field it gets and the feedback block. It is a real learning
// session — answering here writes evidence — which is why the header says so.
//
// A concept whose prerequisites are unmet still shows the lock the learner would
// see. That is not a limitation of the preview; it is the behaviour under test.
// For a locked concept the lab's generator is the place to see the questions.
function Playground() {
  const params = useSearchParams()
  const concept = params.get('concept') ?? ''

  return (
    <>
      {!concept ? (
        <p className="border border-orchid-border bg-orchid-surface p-4 font-mono text-xs text-orchid-secondary">
          No concept named. Pick one in the{' '}
          <Link href="/admin/lab" className="text-orchid-blue">
            lab
          </Link>
          .
        </p>
      ) : (
        <>
          <div className="mb-4 border border-orchid-border bg-orchid-surface p-4">
            <p className="font-mono text-[12px] leading-relaxed text-orchid-secondary">
              This is the live learner route for <span className="text-orchid-primary">{concept}</span>. Answers
              submitted here count as a real attempt for your account. To try an answer without writing evidence, use
              the{' '}
              <Link href="/admin/lab" className="text-orchid-blue">
                lab
              </Link>
              .
            </p>
          </div>
          <LearnStepper conceptId={concept} />
        </>
      )}
    </>
  )
}

export default function AdminLabPlayPage() {
  return (
    <>
      <Header />
      <div className="pt-[var(--chrome-top)] lg:pt-0">
        <main className="mx-auto w-full max-w-[900px] min-w-0 px-4 sm:px-6 py-8 sm:py-12 overflow-x-hidden">
          <SectionHeader label="Administration" title="Learner view" />
          <AdminNav />
          <Suspense fallback={<p className="font-mono text-xs text-orchid-muted">Loading…</p>}>
            <Playground />
          </Suspense>
        </main>
      </div>
      <Footer />
      <BottomTabs />
    </>
  )
}

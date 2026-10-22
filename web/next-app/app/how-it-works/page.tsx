'use client'

import Header from '../../components/Header'
import Footer from '../../components/Footer'
import BottomTabs from '../../components/BottomTabs'
import TableOfContents from '../../components/TableOfContents'
import { bodyStyle, h1Style, sections } from './styles'
import {
  ConceptGraphSection,
  DiagnosticSection,
  GeneratorsSection,
  MeasuringLearningSection,
  SchedulerSection,
  ScoringSection,
  SpacedRepetitionSection,
  StudentModelSection,
  SymbolicGradingSection,
} from './sections'

export default function HowItWorksPage() {
  return (
    <>
      <Header links={[{ label: 'Learn', href: '/learn' }, { label: 'Docs', href: '/docs' }, { label: 'About', href: '/note' }]} />
      <div className="min-h-screen w-full max-w-full min-w-0 overflow-x-hidden">
      <div className="flex max-w-[960px] mx-auto p-4 sm:p-[32px_24px] gap-4 sm:gap-10 max-lg:flex-col max-lg:pb-[calc(80px+env(safe-area-inset-bottom))] min-w-0">
        {/* The table of contents is the shared component the docs pages use. It replaced an
            inline sidebar whose mobile form was a fixed horizontal bar under the header that
            overlapped it and scrolled sideways on a phone. The replacement is a sticky sidebar
            on desktop and a bottom sheet behind a floating button on mobile, in one place. */}
        <TableOfContents headings={sections.map(s => ({ id: s.id, text: s.label }))} />

        <main className="max-w-[720px] flex-1 min-w-0 overflow-hidden">
          <section id="intro" className="mb-12 pb-8 scroll-mt-28" style={{ borderBottom: '0.5px solid var(--border)' }}>
            <h1 style={h1Style}>
              How Orchid Works
            </h1>
            <p style={bodyStyle}>
              The engine behind the learning: a technical explanation of the concept graph, student
              model, diagnostic algorithm, task selection and scoring system.
            </p>
          </section>

          <ConceptGraphSection />
          <StudentModelSection />
          <SpacedRepetitionSection />
          <DiagnosticSection />
          <SchedulerSection />
          <ScoringSection />
          <GeneratorsSection />
          <SymbolicGradingSection />
          <MeasuringLearningSection />

          <section className="mt-12 pt-6 text-center" style={{ borderTop: '0.5px solid var(--border)' }}>
            <p style={{ ...bodyStyle, textAlign: 'center' }}>
              Ready to find your starting point?
            </p>
            <div className="flex gap-3 justify-center mt-4 max-sm:flex-col max-sm:items-center">
              <a href="/onboard" style={{ fontFamily: "var(--font-jetbrains-mono), 'JetBrains Mono', monospace", fontSize: '14px', color: 'var(--bg)', background: 'var(--accent-blue)', padding: '10px 22px', textDecoration: 'none' }}>
                Start diagnostic test →
              </a>
              <a href="/learn" style={{ fontFamily: "var(--font-jetbrains-mono), 'JetBrains Mono', monospace", fontSize: '14px', color: 'var(--text-secondary)', border: '0.5px solid var(--border-strong)', padding: '10px 22px', textDecoration: 'none' }}>
                Start learning →
              </a>
            </div>
          </section>
        </main>
      </div>
    </div>
      <Footer />
      <BottomTabs />
    </>
  )
}

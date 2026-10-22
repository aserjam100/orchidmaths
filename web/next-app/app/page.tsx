'use client'

import Link from 'next/link'
import { useTheme } from '../hooks/useTheme'
import { useHomeRedirect } from '../hooks/useHomeRedirect'
import Header from '../components/Header'
import Footer from '../components/Footer'
import GradePicker from '../components/GradePicker'
import OrchidBackdrop from '../components/OrchidBackdrop'

// The landing is for a student who came to practise: pick a grade, and the next screen asks a
// question. What Orchid is and how it decides moved to /parents, one link away. Logged-in visits
// bounce straight to practice.
export default function HomePage() {
  const { mounted } = useTheme()
  const homeChecked = useHomeRedirect(mounted)

  if (!mounted || !homeChecked) {
    return <div style={{ background: 'var(--bg)', minHeight: '100vh' }} />
  }

  return (
    <>
      <Header links={[{ label: 'For parents', href: '/parents' }, { label: 'Log in', href: '/login' }]} />
      {/* The artwork fills this section behind the picker. The bottom padding is the room the
          flowers need below the last line of text, so they frame the content rather than sit
          under the buttons. */}
      <div className="relative isolate overflow-hidden pb-[clamp(180px,50vw,240px)] sm:pb-[clamp(120px,14vw,220px)]">
        <OrchidBackdrop />
        <main className="mx-auto flex min-h-[60vh] max-w-xl flex-col justify-center px-4 pt-12 sm:pt-16 min-w-0">
          <h1 className="text-balance font-serif text-[clamp(2rem,6vw,3rem)] leading-[1.05] tracking-tight text-orchid-primary">
            Maths practice for Primary 1 to 6
          </h1>
          <p className="mt-3 text-[15px] text-orchid-secondary">
            Pick your grade to start. No sign-up needed.
          </p>
          <div className="mt-8">
            <GradePicker />
          </div>
          <p className="mt-8 font-mono text-[12px] text-orchid-muted">
            Have an account?{' '}
            <Link href="/login" className="text-orchid-blue hover:text-orchid-blue-hover underline underline-offset-2">
              Log in
            </Link>{' '}
            to keep your progress.
          </p>
          <p className="mt-2 font-mono text-[12px] text-orchid-muted">
            <Link href="/parents" className="hover:text-orchid-blue underline underline-offset-2">
              For parents: how Orchid works →
            </Link>
          </p>
        </main>
      </div>
      <Footer />
    </>
  )
}

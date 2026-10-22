'use client'

import { useRouter } from 'next/navigation'
import { useTheme } from '../../hooks/useTheme'
import { ensureGuestId, ensureGuestToken } from '../../lib/auth'
import { useScrollReveal } from '../../hooks/useScrollReveal'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import {
  CoverageSection,
  FaqSection,
  FeaturesSection,
  HeroSection,
  StatesSection,
  TrustSection,
} from '../home/sections'

// The page that used to be the landing. A student lands on the grade picker at `/` and starts
// practising; a parent who wants to know what Orchid is and how it decides comes here. The
// ranking section is left out while the leaderboard is hidden. It does not bounce a logged-in
// visitor, because a parent may well be signed in on the child's account.
export default function ParentsPage() {
  const { mounted } = useTheme()
  const { push } = useRouter()
  useScrollReveal(mounted)

  if (!mounted) {
    return <div style={{ background: 'var(--bg)', minHeight: '100vh' }} />
  }

  // The placement test grades answers, so it needs the guest token first. "Just practise" goes
  // back to the grade picker, which is the practice route now.
  function startTest() {
    ensureGuestId()
    ensureGuestToken().finally(() => push('/onboard'))
  }

  return (
    <>
      <Header links={[{ label: 'Practice', href: '/' }, { label: 'How it works', href: '/how-it-works' }, { label: 'Docs', href: '/docs' }, { label: 'About', href: '/note' }, { label: 'Log in', href: '/login' }]} />
      <main className="max-w-container mx-auto px-4 sm:px-6 overflow-x-clip min-w-0">
        <HeroSection onGetStarted={startTest} onPractise={() => push('/')} />
        <TrustSection />
        <FeaturesSection />
        <StatesSection />
        <CoverageSection />
        <FaqSection />
      </main>
      <Footer />
    </>
  )
}

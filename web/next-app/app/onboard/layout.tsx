import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Onboarding \u00b7 Orchid',
  description: 'Select what you want to learn and let Orchid build a personalized study plan.',
}

export default function OnboardLayout({ children }: { children: React.ReactNode }) {
  return children
}

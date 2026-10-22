import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Efficacy · Orchid Docs',
  description: 'How Orchid measures learning: first/second-pass instrumentation, retention, and the evidence figure.',
}

export default function EfficacyLayout({ children }: { children: React.ReactNode }) {
  return children
}

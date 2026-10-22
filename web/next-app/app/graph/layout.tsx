import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Knowledge Graph \u00b7 Orchid',
  description: 'Explore the Orchid concept map · view your progress across the full curriculum.',
}

export default function GraphLayout({ children }: { children: React.ReactNode }) {
  return children
}

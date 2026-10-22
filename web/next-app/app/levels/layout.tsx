import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Topics | Orchid',
  description:
    'Primary 1 to 6, mapped to the MOE 2021 Primary Mathematics syllabus: each level’s topics and the concepts Orchid practises for them.',
}

export default function LevelsLayout({ children }: { children: React.ReactNode }) {
  return children
}

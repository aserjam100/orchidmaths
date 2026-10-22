import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'For parents | Orchid',
  description:
    'How Orchid works: adaptive practice for Primary 1 to 6, mapped to the Singapore MOE Primary Mathematics syllabus.',
}

export default function ParentsLayout({ children }: { children: React.ReactNode }) {
  return children
}

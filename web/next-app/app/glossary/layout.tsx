import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Mathematical glossary | Orchid',
  description:
    'A searchable reference for mathematical symbols, Greek letters, notation and terminology — from arithmetic through university mathematics.',
}

export default function GlossaryLayout({ children }: { children: React.ReactNode }) {
  return children
}

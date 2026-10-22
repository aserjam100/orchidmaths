'use client'

import Header from '../../components/Header'
import Footer from '../../components/Footer'
import BottomTabs from '../../components/BottomTabs'

export default function DocsLayout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <Header links={[
        { label: 'Levels', href: '/levels' },
        { label: 'Docs', href: '/docs' },
        { label: 'Architecture', href: '/docs/architecture' },
        { label: 'Contributing', href: '/docs/contributing' },
        { label: 'About', href: '/note' },
      ]} />
      <main>{children}</main>
      <Footer />
      <BottomTabs />
    </>
  )
}

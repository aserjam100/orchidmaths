import type { Metadata } from 'next'
import Link from 'next/link'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import BottomTabs from '../../components/BottomTabs'

export const metadata: Metadata = {
  title: 'About — Orchid',
  description: 'What Orchid is, how it follows the MOE Primary Mathematics syllabus, and where it comes from.',
}

const monoFont = "var(--font-jetbrains-mono), 'JetBrains Mono', monospace"
const bodyFont = "var(--font-space-grotesk), 'Space Grotesk', serif"

const textStyle: React.CSSProperties = {
  fontFamily: bodyFont,
  fontSize: '0.95rem',
  color: 'var(--text-secondary)',
  lineHeight: 1.85,
  marginBottom: '1rem',
}

const h2Style: React.CSSProperties = {
  fontFamily: bodyFont,
  fontWeight: 400,
  fontSize: '1.3rem',
  color: 'var(--text-primary)',
  marginTop: '2.5rem',
  marginBottom: '1rem',
  borderBottom: '0.5px solid var(--border)',
  paddingBottom: '0.5rem',
}

const quoteStyle: React.CSSProperties = {
  fontFamily: bodyFont,
  fontSize: '1.05rem',
  fontStyle: 'italic',
  color: 'var(--text-primary)',
  lineHeight: 1.7,
  borderLeft: '2px solid var(--accent-blue)',
  paddingLeft: '1rem',
  margin: '1.5rem 0',
}

const citeStyle: React.CSSProperties = {
  display: 'block',
  fontFamily: monoFont,
  fontSize: '13px',
  fontStyle: 'normal',
  color: 'var(--text-muted)',
  marginTop: '0.5rem',
}

const ipaStyle: React.CSSProperties = {
  fontFamily: monoFont,
  fontSize: '14px',
  color: 'var(--text-muted)',
  marginBottom: '0.75rem',
}

// About Orchid. The upstream project kept its creator's personal note here; that note is
// Mathua's author's story, not Orchid's, so it was replaced at the fork with plain facts. The
// last section is a placeholder for Orchid's own founder note.
type Block =
  | { text: string }
  | { parts: Array<string | { strong: string }> }
  | { quote: string; cite: string }
  | { ipa: string }

type Section = { heading: string; blocks: Block[] }

const SECTIONS: Section[] = [
  {
    heading: 'What Orchid is',
    blocks: [
      {
        parts: [
          'Orchid is an adaptive maths practice app for primary school students in Singapore. Every topic sits in a graph of prerequisites, and a child moves on only when they are ',
          { strong: 'both accurate and quick' },
          ' on what comes before. Mastered topics come back for review before they fade.',
        ],
      },
      {
        text: 'Questions are generated, not stored, so a child cannot memorise answers: the same skill gives a new question every time.',
      },
    ],
  },
  {
    heading: 'How Orchid follows the MOE syllabus',
    blocks: [
      {
        text: 'Each learning outcome in the MOE 2021 Primary Mathematics syllabus (Primary 1 to 6, Standard) is mapped to the Orchid topics that practise it, and marked full, partial or not yet covered. The Primary 1 to 6 courses are built from that mapping, and the coverage report shows every outcome, including the ones Orchid does not cover yet.',
      },
      {
        text: 'Orchid is independent and is not affiliated with or endorsed by the Ministry of Education, Singapore.',
      },
    ],
  },
  {
    heading: 'Where Orchid comes from',
    blocks: [
      {
        text: 'Orchid is built on Mathua, an open-source adaptive maths engine by Wisdom Chukwuemeka, released under the MIT licence. The orchid is the national flower of Singapore.',
      },
    ],
  },
  // TODO(founder): replace this section with your own note: who you are and why you are building Orchid.
  {
    heading: 'A note from the founder',
    blocks: [{ text: 'Coming soon.' }],
  },
]

export default function NotePage() {
  return (
    <>
      <Header />
      <div className="max-w-[680px] mx-auto px-4 sm:px-6 py-12 sm:py-20 pb-[calc(80px+env(safe-area-inset-bottom))] lg:pb-20 overflow-x-hidden min-w-0">
        <Link href="/" style={{ fontFamily: monoFont, fontSize: '14px', color: 'var(--text-muted)' }}>
          ← Back
        </Link>
        <div style={{ fontFamily: monoFont, fontSize: '13px', color: 'var(--text-muted)', marginTop: '1.5rem', marginBottom: '0.5rem' }}>
          ABOUT
        </div>
        <h1 style={{
          fontFamily: "var(--font-space-grotesk), 'Space Grotesk', serif",
          fontWeight: 400,
          fontSize: '1.9rem',
          color: 'var(--text-primary)',
          marginBottom: '2rem',
          borderBottom: '0.5px solid var(--border)',
          paddingBottom: '0.5rem',
        }}>
          About Orchid
        </h1>

        {SECTIONS.map(section => (
          <section key={section.heading}>
            <h2 style={h2Style}>{section.heading}</h2>
            {section.blocks.map((block, i) => {
              if ('quote' in block) {
                return (
                  <blockquote key={i} style={quoteStyle}>
                    {`\u201c${block.quote}\u201d`}
                    <cite style={citeStyle}>{block.cite}</cite>
                  </blockquote>
                )
              }
              if ('ipa' in block) {
                return <p key={i} style={ipaStyle}>{block.ipa}</p>
              }
              if ('parts' in block) {
                return (
                  <p key={i} style={textStyle}>
                    {block.parts.map((part, j) =>
                      typeof part === 'string' ? part : <strong key={j}>{part.strong}</strong>,
                    )}
                  </p>
                )
              }
              return <p key={i} style={textStyle}>{block.text}</p>
            })}
          </section>
        ))}

        <div className="flex gap-3 justify-center mt-10 max-sm:flex-col max-sm:items-center">
          <Link href="/onboard" style={{ fontFamily: monoFont, fontSize: '14px', color: 'var(--bg)', background: 'var(--accent-blue)', padding: '10px 22px', textDecoration: 'none' }}>
            Start diagnostic test →
          </Link>
          <Link href="/learn" style={{ fontFamily: monoFont, fontSize: '14px', color: 'var(--text-secondary)', border: '0.5px solid var(--border-strong)', padding: '10px 22px', textDecoration: 'none' }}>
            Start learning →
          </Link>
        </div>
      </div>
      <Footer />
      <BottomTabs />
    </>
  )
}

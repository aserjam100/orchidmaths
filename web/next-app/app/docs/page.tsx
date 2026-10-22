import type { Metadata } from 'next'
import Link from 'next/link'

export const metadata: Metadata = {
  title: 'Documentation \u2014 Orchid',
  description: 'Architecture, system design, administration, math input and the contributing guide for Orchid, adaptive maths practice for Singapore primary schools.',
}

const monoFont = "var(--font-jetbrains-mono), 'JetBrains Mono', monospace"
const bodyFont = "var(--font-space-grotesk), 'Space Grotesk', serif"

// The list is data, not six copied blocks. Each entry's title and description match the
// generated docs module, so the index and the page it links to cannot disagree. Adding a doc is
// one line here and a route beside the others.
const DOCS = [
  { href: '/docs/architecture', title: 'Architecture', blurb: 'The five-layer stack — UI, API, core engine, grading and storage — and the computerised adaptive diagnostic.' },
  { href: '/docs/system-design', title: 'System design', blurb: 'Schema, scheduler, SM-2, grading, the API surface, and the trade-offs behind each.' },
  { href: '/docs/administration', title: 'Administration', blurb: 'Roles, the permission model, staff invitations, the last-owner invariant and the audit trail.' },
  { href: '/docs/math-input', title: 'Math input', blurb: 'How a typed answer reaches a verdict: MathLive, MathJSON, the plain-text shim, and the mobile keypad.' },
  { href: '/docs/contributing', title: 'Contributing', blurb: 'Add a concept, write a generator, pass the validator.' },
  { href: '/docs/curriculum', title: 'MOE syllabus coverage', blurb: 'Every Primary 1 to 6 learning outcome in the MOE syllabus, and how fully Orchid practises it.' },
  { href: '/docs/efficacy', title: 'Efficacy', blurb: 'How Orchid measures learning, and the honest limits of those measurements.' },
]

export default function DocsIndexPage() {
  const textStyle: React.CSSProperties = {
    fontFamily: bodyFont,
    fontSize: '0.95rem',
    color: 'var(--text-secondary)',
    lineHeight: 1.7,
  }

  const rowStyle: React.CSSProperties = {
    display: 'flex',
    gap: '1rem',
    padding: '0.75rem 0',
    borderBottom: '0.5px solid var(--border)',
    color: 'inherit',
    textDecoration: 'none',
  }

  return (
    <div className="max-w-[720px] mx-auto px-4 sm:px-6 py-12 sm:py-20 pb-[calc(80px+env(safe-area-inset-bottom))] lg:pb-20 overflow-x-hidden min-w-0">
      <Link href="/" style={{ fontFamily: monoFont, fontSize: '14px', color: 'var(--text-muted)' }}>
        ← Back
      </Link>
      <h1 style={{
        fontFamily: "var(--font-space-grotesk), 'Space Grotesk', serif",
        fontWeight: 400,
        fontSize: '1.9rem',
        color: 'var(--text-primary)',
        marginBottom: '0.5rem',
        borderBottom: '0.5px solid var(--border)',
        paddingBottom: '0.5rem',
      }}>
        Documentation
      </h1>

      <p style={{ ...textStyle, marginBottom: '2.5rem' }}>
        Architecture, design and the contributing guide. Every page is generated from the same
        markdown the repository holds, so the two cannot drift apart.
      </p>

      <div style={{ marginBottom: '3rem' }}>
        {DOCS.map((doc, i) => (
          <Link key={doc.href} href={doc.href} className="link-underline" style={rowStyle}>
            <span style={{ fontFamily: monoFont, fontSize: '14px', color: 'var(--accent-blue)', minWidth: '1.5rem', paddingTop: '1px' }}>
              {i + 1}.
            </span>
            <div>
              <div style={{ fontFamily: "var(--font-space-grotesk), 'Space Grotesk', serif", fontWeight: 400, fontSize: '1.1rem', color: 'var(--text-primary)', marginBottom: '2px' }}>
                {doc.title}
              </div>
              <p style={{ ...textStyle, margin: 0, fontSize: '0.9rem', color: 'var(--text-muted)' }}>
                {doc.blurb}
              </p>
            </div>
          </Link>
        ))}
      </div>

      <div style={{ borderTop: '0.5px solid var(--border)', paddingTop: '3rem' }}>
        <h2 style={{
          fontFamily: "var(--font-space-grotesk), 'Space Grotesk', serif",
          fontWeight: 400,
          fontSize: '1.3rem',
          color: 'var(--text-primary)',
          marginBottom: '1rem',
        }}>
          Quickstart <span style={{ fontFamily: monoFont, fontSize: '13px', color: 'var(--text-muted)' }}>(for contributors)</span>
        </h2>

        <pre style={{
          fontFamily: monoFont,
          fontSize: '14px',
          color: 'var(--text-secondary)',
          whiteSpace: 'pre',
          overflowX: 'auto',
          lineHeight: 1.6,
          margin: 0,
          width: '100%',
          maxWidth: '100%',
          minWidth: 0,
        }}>
{`git clone https://github.com/aserjam100/orchid.git
cd orchid
go build -o bin/orchid ./cmd/orchid

# Run the web server (SQLite by default)
./bin/orchid --serve --port 8080`}
        </pre>

        <p style={{
          fontFamily: bodyFont,
          fontStyle: 'italic',
          fontSize: '0.9rem',
          color: 'var(--text-muted)',
          marginTop: '1rem',
          lineHeight: 1.7,
        }}>
          Requirements: Go 1.26+ · New to Orchid? Start with <Link href="/how-it-works" className="link-underline" style={{ color: 'var(--accent-blue)' }}>How it works</Link>. Licenses are at <Link href="/licenses" className="link-underline" style={{ color: 'var(--accent-blue)' }}>Licenses</Link>.
        </p>
      </div>
    </div>
  )
}

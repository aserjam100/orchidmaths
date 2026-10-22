'use client'

import Link from 'next/link'
import { DICEBEAR_VERSION } from '../lib/dicebear'
import OrchidMark from './OrchidMark'

interface FooterProps {
  className?: string
}

const FOOTER_SHAPES = ['alpha', 'beta', 'gamma']
const shapeUrl = (seed: string) =>
  `https://api.dicebear.com/${DICEBEAR_VERSION}/shapes/svg?seed=${seed}&backgroundColor=1c3a5e`

const linkClass = 'block opacity-80 transition-opacity hover:opacity-100'
const headingClass = 'font-mono text-[11px] uppercase tracking-[0.2em] text-orchid-muted'

// The credits strip names each license the shipped content carries, not only the software.
// Orchid is a fork of Orchid (MIT). Some lessons are adapted from ORCCA (CC BY 4.0), whose
// licence requires attribution in the manner it specifies, so it gets its own line. A short
// line cannot carry title, author, source and license for every source, so "Licenses" links
// to /licenses, which renders THIRD_PARTY_NOTICES.md in full.
const CREDITS = [
  'MIT License',
  'Built on Mathua (MIT)',
  'Avatars by DiceBear',
  'Some lessons adapted from ORCCA (CC BY 4.0)',
  'Glossary: OpenStax, DLMF and others',
]

export default function Footer({ className = '' }: FooterProps) {
  return (
    // A plain band on the page's own colours. The orchid artwork that used to sit behind it
    // set its height to the picture's, about 550px at desktop width; it now lives behind the
    // home page's body (OrchidBackdrop).
    <footer className={`footer relative mt-[10px] border-t border-orchid-border text-orchid-primary ${className}`}>
      <div className="footer-content">
        <div className="footer-content-row">
          <div className="md:max-w-[34ch]">
            <div>
              <div className="font-sans text-lg"><OrchidMark /> Orchid</div>
              <div className="mt-2 flex gap-2 md:mt-3">
                {FOOTER_SHAPES.map(seed => (
                  <img
                    key={seed}
                    src={shapeUrl(seed)}
                    alt=""
                    width={24}
                    height={24}
                    className="size-6 ring-1 ring-white/10 md:size-8"
                    loading="lazy"
                    referrerPolicy="no-referrer"
                  />
                ))}
              </div>
            </div>
            <p className="mt-3 text-sm leading-relaxed text-orchid-muted">
              Adaptive maths practice for Singapore primary schools.
            </p>
          </div>

          <div className="grid grid-cols-2 gap-x-10 gap-y-8 sm:gap-x-16">
            <nav aria-labelledby="footer-learn">
              <h2 id="footer-learn" className={headingClass}>Learn</h2>
              <ul className="mt-3 space-y-2 font-mono text-xs">
                <li><Link href="/levels" className={linkClass}>Topics</Link></li>
                <li><Link href="/graph" className={linkClass}>Concept graph</Link></li>
              </ul>
            </nav>

            <nav aria-labelledby="footer-resources">
              <h2 id="footer-resources" className={headingClass}>Resources</h2>
              <ul className="mt-3 space-y-2 font-mono text-xs">
                <li><Link href="/how-it-works" className={linkClass}>How it works</Link></li>
                <li><Link href="/docs" className={linkClass}>Docs</Link></li>
                <li><Link href="/docs/contributing" className={linkClass}>Contributing</Link></li>
                <li><Link href="/note" className={linkClass}>About</Link></li>
                <li>
                  <a
                    href="https://github.com/aserjam100/orchid"
                    target="_blank"
                    rel="noreferrer"
                    className={linkClass}
                  >
                    GitHub
                  </a>
                </li>
              </ul>
            </nav>
          </div>
        </div>

        <div className="footer-credits border-t pt-5 font-mono text-[11px] opacity-80">
          {CREDITS.map(credit => (
            <span key={credit} className="after:ml-3 after:content-['·']">
              {credit}
            </span>
          ))}
          <Link href="/licenses" className="underline hover:no-underline">
            Licenses
          </Link>
          {' · '}
          <Link href="/privacy" className="underline hover:no-underline">
            Privacy
          </Link>
        </div>
      </div>
    </footer>
  )
}

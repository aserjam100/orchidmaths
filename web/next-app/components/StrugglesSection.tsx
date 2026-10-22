'use client'

import Link from 'next/link'
import type { WeaknessRes } from '../lib/api'

interface Props {
  weaknesses: WeaknessRes | null
}

// A one-line "needs attention" summary: how many concepts are weak, and the way to review them.
// It was a boxed panel with a per-concept bar list and a per-domain expander — the heaviest block
// on the page — and every number in it was decided by the same count this now states outright.
export default function StrugglesSection({ weaknesses }: Props) {
  const entries = weaknesses?.by_domain ?? {}
  const domains = Object.keys(entries)
  if (domains.length === 0) return null

  let struggling = 0
  let total = 0
  for (const items of Object.values(entries)) {
    for (const it of items) {
      total++
      if (it.weakness >= 0.6) struggling++
    }
  }
  const needsPractice = total - struggling

  return (
    <div className="min-w-0">
      <h3 className="font-mono text-[12px] uppercase tracking-wider text-orchid-muted mb-2">
        Needs attention
      </h3>
      <div className="flex flex-wrap items-center justify-between gap-x-4 gap-y-2 border-[0.5px] border-orchid-border px-3 py-2 min-w-0">
        <span className="font-mono text-[12px] min-w-0">
          {struggling > 0 && <span className="text-orchid-red">{struggling} struggling</span>}
          {struggling > 0 && needsPractice > 0 && <span className="text-orchid-muted"> · </span>}
          {needsPractice > 0 && <span className="text-orchid-muted">{needsPractice} need practice</span>}
          <span className="text-orchid-muted">
            {total > 0 ? ' · ' : ''}{total} concept{total !== 1 ? 's' : ''} · {domains.length} domain{domains.length !== 1 ? 's' : ''}
          </span>
        </span>
        <Link
          href="/review"
          className="shrink-0 font-mono text-[12px] text-orchid-blue hover:text-orchid-blue-hover whitespace-nowrap"
        >
          Review now →
        </Link>
      </div>
    </div>
  )
}

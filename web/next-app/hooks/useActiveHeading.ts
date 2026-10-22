'use client'

import { useEffect, useState } from 'react'

// Which section is currently in view, for a table of contents to highlight.
//
// Lifted out of app/how-it-works/page.tsx, which had this inline and was the only page with a
// table of contents. The docs pages need the same behaviour, and a second copy would drift.
//
// The rootMargin is the part that matters: it treats the top fifth and bottom 70% of the
// viewport as out of play, so a heading is "active" when it reaches the upper third of the
// screen rather than the moment its first pixel appears at the bottom.
export function useActiveHeading(ids: string[]): string {
  const [active, setActive] = useState(ids[0] ?? '')

  useEffect(() => {
    if (ids.length === 0) return
    const observer = new IntersectionObserver(
      entries => {
        // The entries array is not ordered by document position, so pick the highest one that is
        // on screen rather than whichever fires last. Otherwise scrolling up jumps the highlight
        // ahead of the reader.
        const onScreen = entries
          .filter(e => e.isIntersecting)
          .sort((a, b) => a.boundingClientRect.top - b.boundingClientRect.top)
        if (onScreen.length > 0) setActive(onScreen[0].target.id)
      },
      { rootMargin: '-20% 0px -70% 0px' },
    )
    for (const id of ids) {
      const el = document.getElementById(id)
      if (el) observer.observe(el)
    }
    return () => observer.disconnect()
    // ids is a fresh array each render; depend on its contents, not its identity.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ids.join('|')])

  return active
}

'use client'

import { useEffect, useLayoutEffect, useState } from 'react'

// The orchid artwork, drawn behind the body of a page.
//
// It used to be the footer's backdrop, where the picture set the footer's height: at desktop
// width a 2.6:1 print is about 550px tall, which left a wide empty band under a few lines of
// links. It now sits behind the main content instead, filling the box it is given with
// `object-fit: cover` anchored to the bottom, so the flowers frame the content and the box's
// own content decides its height.
//
// Each entry carries two pictures, because a landscape print across a phone is a short slot
// and a portrait print across a desktop is a thin strip. `<picture>` with a `max-width: 639px`
// source swaps between them. The desktop prints leave the middle empty and put the flowers at
// the sides, which is what lets centred content sit on top of them.
//
// PROVENANCE: every picture is Orchid's own, drawn by `art/orchid-art.py` and rendered to PNG
// (`make art`). The files keep their `public/footer/` path because `make art` writes there.
// `test/orchidArt.test.tsx` checks that the files exist and declare their real dimensions.
export interface OrchidArtVariant {
  file: string
  width: number
  height: number
}

export interface OrchidArt {
  desktop: OrchidArtVariant
  mobile: OrchidArtVariant
}

export const ORCHID_ART: OrchidArt[] = [
  {
    desktop: { file: '/footer/footer-orchid-wide.png', width: 1672, height: 643 },
    mobile: { file: '/footer/footer-orchid-tall.png', width: 1080, height: 1920 },
  },
  {
    desktop: { file: '/footer/footer-orchid-wide.png', width: 1672, height: 643 },
    mobile: { file: '/footer/footer-orchid-tall-2.png', width: 1080, height: 1920 },
  },
  {
    desktop: { file: '/footer/footer-orchid-wide-2.png', width: 1672, height: 643 },
    mobile: { file: '/footer/footer-orchid-tall.png', width: 1080, height: 1920 },
  },
  {
    desktop: { file: '/footer/footer-orchid-wide-2.png', width: 1672, height: 643 },
    mobile: { file: '/footer/footer-orchid-tall-2.png', width: 1080, height: 1920 },
  },
]

const useIsomorphicLayoutEffect = typeof window === 'undefined' ? useEffect : useLayoutEffect

// Picks one entry once per mount. The pick happens in a layout effect and never during render:
// Math.random() in render would disagree with the server's markup and fail hydration, and a
// layout effect lands the swap before the first paint.
function useOrchidArt(): OrchidArt {
  const [art, setArt] = useState<OrchidArt>(ORCHID_ART[0])

  useIsomorphicLayoutEffect(() => {
    setArt(ORCHID_ART[Math.floor(Math.random() * ORCHID_ART.length)])
  }, [])

  return art
}

/**
 * Fills its nearest positioned ancestor, behind that ancestor's content. The ancestor needs
 * `relative isolate` so the negative z-index stays inside it.
 */
export default function OrchidBackdrop() {
  const art = useOrchidArt()

  return (
    <picture className="orchid-backdrop pointer-events-none">
      <source
        media="(max-width: 639px)"
        srcSet={art.mobile.file}
        width={art.mobile.width}
        height={art.mobile.height}
      />
      <img
        src={art.desktop.file}
        alt=""
        aria-hidden="true"
        width={art.desktop.width}
        height={art.desktop.height}
        decoding="async"
        className="orchid-backdrop-img"
      />
    </picture>
  )
}

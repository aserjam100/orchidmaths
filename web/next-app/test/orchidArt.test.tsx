import { describe, it, expect } from 'vitest'
import fs from 'node:fs'
import path from 'node:path'
import { ORCHID_ART } from '../components/OrchidBackdrop'

import { APP_ROOT } from './helpers/roots'

// The home page draws one of ORCHID_ART behind its body on every load, and each
// entry names two pictures: a landscape for desktop and a portrait for phones.
//
// A wrong `width`/`height` is worth guarding because nothing else would notice:
// the image still renders and still looks plausible.
//
// The second failure is a file that is listed but absent, which 404s on exactly
// the visits that pick it and leaves the backdrop empty -- intermittent by
// construction, because the pick is random.

type Dimensions = { width: number; height: number }

/** Real pixel dimensions, read from the file's own header. */
function dimensionsOf(abs: string): Dimensions {
  const head = fs.readFileSync(abs).subarray(0, 64 * 1024)

  // JPEG: walk the marker segments to the first SOF, which carries the frame size.
  if (head[0] === 0xff && head[1] === 0xd8) {
    for (let i = 2; i < head.length - 9; i++) {
      if (head[i] !== 0xff) continue
      const marker = head[i + 1]
      const len = head.readUInt16BE(i + 2)
      if (marker >= 0xc0 && marker <= 0xcf && ![0xc4, 0xc8, 0xcc].includes(marker)) {
        return { height: head.readUInt16BE(i + 5), width: head.readUInt16BE(i + 7) }
      }
      i += len
    }
  }

  // PNG: IHDR is always the first chunk, at a fixed offset.
  if (head.readUInt32BE(0) === 0x89504e47) {
    return { width: head.readUInt32BE(16), height: head.readUInt32BE(20) }
  }

  throw new Error(`unrecognised image header in ${abs}`)
}

const variants = ORCHID_ART.flatMap(art => [
  { ...art.desktop, slot: 'desktop' as const },
  { ...art.mobile, slot: 'mobile' as const },
])

describe('ORCHID_ART', () => {
  it('is not empty', () => {
    // The sentinel. Every assertion below passes vacuously against an empty pool,
    // which is the shape of a guard that reads nothing.
    expect(ORCHID_ART.length).toBeGreaterThan(0)
    expect(variants.length).toBe(ORCHID_ART.length * 2)
  })

  it('names only files that exist', () => {
    for (const art of variants) {
      const abs = path.join(APP_ROOT, 'public', art.file.replace(/^\//, ''))
      expect(fs.existsSync(abs), `${art.slot} ${art.file} is listed but not on disk`).toBe(true)
    }
  })

  it('declares the real dimensions of every file', () => {
    for (const art of variants) {
      const abs = path.join(APP_ROOT, 'public', art.file.replace(/^\//, ''))
      const { width, height } = dimensionsOf(abs)
      expect(art.width, `${art.file} declares width ${art.width}, file is ${width}`).toBe(width)
      expect(art.height, `${art.file} declares height ${art.height}, file is ${height}`).toBe(height)
    }
  })

  it('pairs a landscape desktop with a portrait mobile', () => {
    // The whole premise of the art direction. A landscape shown across a phone is
    // a letterbox slot; a portrait shown across a desktop is a thin strip with
    // dead space either side. Swapping them would still render, and would still
    // look like a deliberate choice -- which is what makes it worth asserting.
    for (const art of ORCHID_ART) {
      const desktop = art.desktop.width / art.desktop.height
      const mobile = art.mobile.width / art.mobile.height
      expect(desktop, `${art.desktop.file} is not a landscape`).toBeGreaterThanOrEqual(1.6)
      expect(mobile, `${art.mobile.file} is not a portrait`).toBeLessThanOrEqual(0.8)
    }
  })
})

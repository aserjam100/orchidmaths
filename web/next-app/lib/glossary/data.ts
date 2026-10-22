import type { GlossaryEntry } from './types'
import greek from '../../data/glossary/greek.json'
import arithmetic from '../../data/glossary/arithmetic.json'
import algebra from '../../data/glossary/algebra.json'
import logic from '../../data/glossary/logic.json'
import geometry from '../../data/glossary/geometry.json'
import calculus from '../../data/glossary/calculus.json'
import linear from '../../data/glossary/linear.json'
import probability from '../../data/glossary/probability.json'
import advanced from '../../data/glossary/advanced.json'

// Static bundle: every entry ships with the app. No network, no service,
// no concept mapping — the glossary never touches learner state.
export const GLOSSARY_ENTRIES: GlossaryEntry[] = [
  ...(greek as GlossaryEntry[]),
  ...(arithmetic as GlossaryEntry[]),
  ...(algebra as GlossaryEntry[]),
  ...(logic as GlossaryEntry[]),
  ...(geometry as GlossaryEntry[]),
  ...(calculus as GlossaryEntry[]),
  ...(linear as GlossaryEntry[]),
  ...(probability as GlossaryEntry[]),
  ...(advanced as GlossaryEntry[]),
]

export const GLOSSARY_VERSION = 1

export function glossaryById(id: string): GlossaryEntry | undefined {
  return GLOSSARY_ENTRIES.find(e => e.id === id)
}

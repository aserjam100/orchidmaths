import { getLessonReadiness, type EligibilityPrereq, type ReadinessRes } from './api'

/**
 * The concept to practise first when a topic is locked.
 *
 * A locked topic used to end in a list of its prerequisites with a "Study" link on each, and a
 * prerequisite is often locked itself, so a student who came to practise found a reading list
 * instead of a question. This walks down the prerequisites to the first one the server will
 * serve questions for, and the page offers that one as a warm-up.
 *
 * It decides nothing about mastery. Every step is the server's own eligibility verdict from
 * `/readiness`; the client only chooses which unmet prerequisite to ask about next, in the order
 * the server listed them.
 */
export interface WarmUp {
  conceptId: string
  title: string
}

// A bound on requests, not on depth. The prerequisite closure of a Primary 6 topic is well over
// a hundred concepts, and a learner waiting on a spinner is worse than the panel it replaces.
const MAX_REQUESTS = 12

function unmet(r: ReadinessRes): EligibilityPrereq[] {
  return (r.prerequisites ?? []).filter(p => !p.met)
}

export async function findWarmUp(
  locked: ReadinessRes,
  readiness: (conceptId: string) => Promise<ReadinessRes> = getLessonReadiness,
): Promise<WarmUp | null> {
  const seen = new Set<string>([locked.concept_id])
  // Depth first, in the server's order: the first prerequisite listed is followed all the way
  // down before its siblings are tried.
  const stack = unmet(locked).reverse()
  let requests = 0
  while (stack.length > 0 && requests < MAX_REQUESTS) {
    const p = stack.pop()!
    if (seen.has(p.conceptId)) continue
    seen.add(p.conceptId)
    requests++
    const r = await readiness(p.conceptId)
    if (r.eligible !== false) return { conceptId: p.conceptId, title: p.title }
    stack.push(...unmet(r).reverse())
  }
  return null
}

import graphMetaJson from '../../data/graph.meta.json'
import moeCoverageJson from '../../data/moe.coverage.json'
// A type-only import, so it costs the landing bundle nothing: graphPayload.ts is otherwise
// pulled in by the idle-loaded graph chunk and the build script, and zod with it.
// This file redeclared the interface locally, which meant the landing page's view of the
// artifact was a second schema — adding quizGateXP to the real one type-checked here as a
// missing property rather than as a change to the page, which is a quieter way to drift.
import type { GraphMeta } from '../../lib/graphPayload'

const graphMeta = graphMetaJson as GraphMeta

export const PIPELINE_STATES = [
  { label: 'UNSEEN', status: 'unseen' as const },
  { label: 'LEARNING', status: 'learning' as const },
  { label: 'PRACTICING', status: 'practicing' as const },
  { label: 'MASTERED', status: 'mastered' as const },
  { label: 'DECAYING', status: 'decaying' as const },
]

// The landing page used to hardcode 150 while the engine gates the quiz at 50 (xp.QuizGateXP,
// rescaled by ADR-020). One number now, guarded from Go by TestLandingQuizGateMatchesEngine.
export const quizGateXP = graphMeta.quizGateXP

// The page is about MOE Primary, so its coverage comes from the MOE mapping rather than the
// whole corpus. The corpus is 1250 concepts across 21 domains, most of them beyond Primary
// (calculus alone is 205), and a table of those under a "Primary 1 to 6" headline promised a
// child topology. scripts/curriculum_moe.py writes this file; `--check` fails on drift.
export interface MoeLevelCoverage {
  level: string
  name: string
  outcomes: number
  full: number
  partial: number
  none: number
}

export const moeCoverage = moeCoverageJson as {
  outcomeCount: number
  conceptCount: number
  levels: MoeLevelCoverage[]
}

// Thresholds come from Go's levels.All through the generated artifact, guarded by
// TestLandingLevelsMatchServer. They were once derived from the corpus size, which put the top
// title at 1112 mastered concepts while the server awarded it at 256.
export const levels = graphMeta.levels.map((l, i, all) => {
  const next = all[i + 1]
  return {
    num: String(i + 1).padStart(2, '0'),
    name: l.title,
    range: next ? `${l.min}–${next.min - 1}` : `${l.min}+`,
    ...(next ? {} : { elite: true }),
  }
})

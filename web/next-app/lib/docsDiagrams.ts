'use client'

import dynamic from 'next/dynamic'
import type { ComponentType } from 'react'

// The interactive version of each diagram the repository docs embed. Keyed by the diagram name
// with its -light/-dark suffix removed, so one entry covers both variants a doc may reference.
//
// A diagram absent here still renders, as its committed pre-rendered SVG. That is the fallback,
// not the plan: `test/docsDiagrams.test.ts` fails when a doc references a diagram with no
// component, so the fallback cannot become the silent answer.
export const DIAGRAMS: Record<string, ComponentType> = {
  architecture: dynamic(() => import('../components/ArchitectureFlow'), { ssr: false }),
  'system-design-architecture': dynamic(() => import('../components/ArchitectureFlow'), { ssr: false }),
  'system-design-data-model': dynamic(() => import('../components/SystemDesignFlow/DataModelFlow'), { ssr: false }),
  'system-design-request-flow': dynamic(() => import('../components/SystemDesignFlow/RequestFlow'), { ssr: false }),
  'system-design-grading-pipeline': dynamic(() => import('../components/GradingFlow'), { ssr: false }),
  'system-design-cat-diagnostic': dynamic(() => import('../components/DiagnosticFlow'), { ssr: false }),
  'cat-diagnostic': dynamic(() => import('../components/DiagnosticFlow'), { ssr: false }),
  'system-design-sm2-mastery': dynamic(() => import('../components/SM2MasteryFlow'), { ssr: false }),
  'student-model': dynamic(() => import('../components/StudentModelFlow'), { ssr: false }),
}

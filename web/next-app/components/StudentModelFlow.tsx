'use client'

import FlowDiagram, { type FlowNodeDef, type FlowEdgeDef } from './FlowDiagram'

export default function StudentModelFlow() {
  const nodes: FlowNodeDef[] = [
    { id: 'unseen', label: 'UNSEEN' },
    { id: 'learning', label: 'LEARNING' },
    { id: 'practicing', label: 'PRACTICING' },
    { id: 'mastered', label: 'MASTERED', variant: 'terminal' },
    { id: 'decaying', label: 'DECAYING', variant: 'error' },
  ]

  // Labels follow internal/mastery/machine.go. Advancement is decided by evidence, so the
  // rungs are named for what the evidence shows, not for a streak. The previous labels were
  // also off by one rung: "first correct" sat on LEARNING -> PRACTICING and "streak reached"
  // on PRACTICING -> MASTERED, which is not the ladder the state diagram draws.
  const edges: FlowEdgeDef[] = [
    { id: 'e-unseen-learning', source: 'unseen', target: 'learning', label: 'first\ncorrect', color: 'green', animated: true },
    { id: 'e-learning-practicing', source: 'learning', target: 'practicing', label: 'consistent\nattempts', color: 'green', animated: true },
    { id: 'e-practicing-mastered', source: 'practicing', target: 'mastered', label: 'evidence\nreached', color: 'green', animated: true },
    { id: 'e-mastered-decaying', source: 'mastered', target: 'decaying', label: '14 days\nidle', color: 'blue', animated: true },
    { id: 'e-decaying-practicing', source: 'decaying', target: 'practicing', label: 'review\ncorrect', dashed: true, color: 'green' },
  ]

  return <FlowDiagram nodes={nodes} edges={edges} direction="TB" height={420} rankSep={150} nodeSep={80} allowZoom />
}

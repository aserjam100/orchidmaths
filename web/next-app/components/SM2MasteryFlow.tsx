'use client'

import FlowDiagram, { type FlowNodeDef, type FlowEdgeDef } from './FlowDiagram'

// The diagram named "SM-2 & Mastery" had only a pre-rendered SVG. This replaces it with two
// React Flow panels so the section's most-consulted figure is interactive like the rest.
//
// Both panels follow the code: the ladder is internal/mastery/machine.go (Advances one rung on
// evidence, never on a streak), and the interval branch is internal/scheduler/sm2.go.
const stateNodes: FlowNodeDef[] = [
  { id: 'unseen', label: 'UNSEEN' },
  { id: 'learning', label: 'LEARNING' },
  { id: 'practicing', label: 'PRACTICING' },
  { id: 'mastered', label: 'MASTERED', variant: 'terminal' },
  { id: 'decaying', label: 'DECAYING', variant: 'error' },
]

const stateEdges: FlowEdgeDef[] = [
  { id: 'e1', source: 'unseen', target: 'learning', label: 'first\ncorrect', color: 'green', animated: true },
  { id: 'e2', source: 'learning', target: 'practicing', label: 'consistent\nattempts', color: 'green', animated: true },
  { id: 'e3', source: 'practicing', target: 'mastered', label: 'evidence\nreached', color: 'green', animated: true },
  { id: 'e4', source: 'mastered', target: 'decaying', label: '14 days\nidle', color: 'blue', animated: true },
  { id: 'e5', source: 'decaying', target: 'practicing', label: 'review\ncorrect', color: 'green', dashed: true },
]

const sm2Nodes: FlowNodeDef[] = [
  { id: 'quality', label: 'Grade review\nquality 0-5' },
  { id: 'grow', label: 'Ease factor up\nrepetitions + 1' },
  { id: 'reset', label: 'Reset\nrepetitions to 0', variant: 'error' },
  { id: 'interval', label: 'Interval: 1 day, then 6 days,\nthen previous × ease\n× learning speed' },
  { id: 'schedule', label: 'Schedule next review', variant: 'terminal' },
]

const sm2Edges: FlowEdgeDef[] = [
  { id: 's1', source: 'quality', target: 'grow', label: '3 or more', color: 'green' },
  { id: 's2', source: 'quality', target: 'reset', label: 'below 3', color: 'red' },
  { id: 's3', source: 'grow', target: 'interval', color: 'green', animated: true },
  { id: 's4', source: 'reset', target: 'interval', label: 'interval 1', color: 'red' },
  { id: 's5', source: 'interval', target: 'schedule', color: 'blue', animated: true },
]

export default function SM2MasteryFlow() {
  return (
    <div style={{ display: 'grid', gap: '1rem' }}>
      <FlowDiagram nodes={stateNodes} edges={stateEdges} direction="TB" height={360} rankSep={110} nodeSep={70} />
      <FlowDiagram nodes={sm2Nodes} edges={sm2Edges} direction="LR" height={300} rankSep={110} nodeSep={50} />
    </div>
  )
}

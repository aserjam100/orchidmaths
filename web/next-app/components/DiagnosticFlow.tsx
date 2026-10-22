'use client'

import { useMemo, useEffect } from 'react'
import {
  ReactFlow,
  Background,
  Controls,
  Handle,
  Position,
  useNodesState,
  useEdgesState,
  type Node,
} from '@xyflow/react'
import '@xyflow/react/dist/base.css'
import { useTheme } from '../hooks/useTheme'
import { themeColors } from './FlowDiagram'

const monoFont = "var(--font-jetbrains-mono), 'JetBrains Mono', monospace"

const NODE_W = 150
const CX = 200

function BaseNode({ data, variant }: { data: { label: string }; variant: 'start' | 'decision' | 'process' | 'terminal' }) {
  const { theme } = useTheme()
  const isDark = theme === 'dark'
  const c = themeColors[isDark ? 'dark' : 'light']

  if (variant === 'start') {
    return (
      <div style={{
        width: 28, height: 28, borderRadius: '50%', position: 'relative',
        background: isDark ? c.accentBlue : c.accentBlue,
        border: 'none',
        display: 'flex', alignItems: 'center', justifyContent: 'center',
        boxShadow: '0 2px 6px rgba(0,0,0,0.15)',
      }}>
        <Handle type="source" position={Position.Bottom} style={{ opacity: 0, pointerEvents: 'none' }} />
        <span style={{ color: '#fff', fontSize: '14px', lineHeight: 1 }}>●</span>
      </div>
    )
  }

  const style: React.CSSProperties = {
    fontFamily: monoFont,
    fontSize: '12px',
    padding: '10px 14px',
    textAlign: 'center',
    lineHeight: 1.5,
    borderRadius: 0,
    minWidth: 120,
    maxWidth: 180,
    boxShadow: '0 1px 3px rgba(0,0,0,0.06)',
  }

  if (variant === 'decision') {
    Object.assign(style, {
      background: isDark ? '#1e1e24' : '#fafafa',
      border: '0.5px solid ' + (isDark ? '#3f3f46' : '#d4d4d8'),
      borderLeft: '2px solid ' + c.accentBlue,
      color: c.textPrimary,
      fontWeight: 500,
    })
  } else if (variant === 'terminal') {
    Object.assign(style, {
      background: isDark ? '#1a1a2e' : '#f0f4ff',
      border: '0.5px solid ' + c.accentBlue,
      color: c.accentBlue,
      fontWeight: 500,
      fontSize: '11px',
      letterSpacing: '0.05em',
    })
  } else {
    Object.assign(style, {
      background: c.surface,
      border: '0.5px solid ' + c.border,
      color: c.textSecondary,
    })
  }

  return (
    <div style={style}>
      <Handle type="target" position={Position.Top} style={{ opacity: 0, pointerEvents: 'none' }} />
      <span style={{ whiteSpace: 'pre-wrap' }}>{data.label}</span>
      {variant !== 'terminal' && (
        <Handle type="source" position={Position.Bottom} style={{ opacity: 0, pointerEvents: 'none' }} />
      )}
    </div>
  )
}

const nodeTypes = {
  startNode: ({ data, ...props }: { data: { label: string } }) => <BaseNode data={data} {...props} variant="start" />,
  decisionNode: ({ data, ...props }: { data: { label: string } }) => <BaseNode data={data} {...props} variant="decision" />,
  processNode: ({ data, ...props }: { data: { label: string } }) => <BaseNode data={data} {...props} variant="process" />,
  terminalNode: ({ data, ...props }: { data: { label: string } }) => <BaseNode data={data} {...props} variant="terminal" />,
}

export default function DiagnosticFlow() {
  const { theme, mounted } = useTheme()
  const isDark = theme === 'dark'
  const c = themeColors[isDark ? 'dark' : 'light']

  // A single column: the algorithm is a loop, not a binary search. These labels follow
  // internal/diagnostic/cat.go -- compressedCover, infoGain, and shouldStop -- and the same
  // drawing is in diagrams/system-design-cat-diagnostic.d2. This component said "test midpoint"
  // and "move forward / backward", which described the binary search the code replaced and was
  // shown to learners on /docs/architecture and /diagnose.
  const gap = 56
  const yStart = 4
  const yCover = yStart + 28 + 20
  const yPick = yCover + 44 + gap
  const yQuestion = yPick + 44 + gap
  const yPropagate = yQuestion + 44 + gap
  const yStop = yPropagate + 44 + gap
  const yFrontier = yStop + 44 + gap
  const midX = CX - NODE_W / 2

  const initialNodes: Node[] = [
    { id: 'start', type: 'startNode', position: { x: CX - 14, y: yStart }, data: { label: '' } },
    { id: 'cover', type: 'decisionNode', position: { x: midX, y: yCover }, data: { label: 'Build covering set\nroots + evenly spaced\ntarget 30 concepts' } },
    { id: 'pick', type: 'processNode', position: { x: midX, y: yPick }, data: { label: 'Pick highest\ninformation gain' } },
    { id: 'question', type: 'processNode', position: { x: midX, y: yQuestion }, data: { label: 'Ask one question\nrecord evidence' } },
    { id: 'propagate', type: 'processNode', position: { x: midX, y: yPropagate }, data: { label: 'Propagate evidence\n±0.3 to neighbours' } },
    { id: 'stop', type: 'decisionNode', position: { x: midX, y: yStop }, data: { label: 'Cover done AND\nasked ≥ 25,\nor asked ≥ 45' } },
    { id: 'frontier', type: 'terminalNode', position: { x: midX, y: yFrontier }, data: { label: 'FRONTIER AT\nLARGEST BELIEF DROP' } },
  ]

  const [nodes, , onNodesChange] = useNodesState(initialNodes)

  const styledEdges = useMemo(() => {
    const edgeStyle = (color: string) => ({ stroke: color, strokeWidth: 1.5 })
    const labelStyle = (color: string) => ({ fontFamily: monoFont, fontSize: '11px', fill: color, fontWeight: 500 })

    return [
      { id: 'e-start-cover', source: 'start', target: 'cover', style: edgeStyle(c.borderStrong) },
      { id: 'e-cover-pick', source: 'cover', target: 'pick', style: edgeStyle(c.borderStrong) },
      { id: 'e-pick-question', source: 'pick', target: 'question', style: edgeStyle(c.borderStrong) },
      { id: 'e-question-propagate', source: 'question', target: 'propagate', style: edgeStyle(c.borderStrong) },
      { id: 'e-propagate-stop', source: 'propagate', target: 'stop', style: edgeStyle(c.borderStrong) },
      {
        id: 'e-stop-frontier', source: 'stop', target: 'frontier',
        label: 'Yes', animated: true,
        labelStyle: labelStyle(c.accentGreen),
        style: edgeStyle(c.accentGreen),
      },
      {
        id: 'e-stop-pick', source: 'stop', target: 'pick',
        label: 'No',
        labelStyle: labelStyle(c.accentBlue),
        style: { ...edgeStyle(c.accentBlue), strokeDasharray: '5 5' },
        type: 'smoothstep',
      },
    ]
  }, [c])

  const [edges, setEdges, onEdgesChange] = useEdgesState(styledEdges)

  useEffect(() => {
    setEdges(styledEdges)
  }, [styledEdges, setEdges])

  if (!mounted) return <div style={{ height: 580, width: '100%' }} />

  return (
    <div style={{
      height: 580, width: '100%',
      border: '0.5px solid ' + c.border,
      background: isDark ? 'rgba(0,0,0,0.12)' : 'rgba(0,0,0,0.02)',
    }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        nodeTypes={nodeTypes}
        fitView
        fitViewOptions={{ padding: 0.25 }}
        nodesConnectable={false}
        preventScrolling={false}
        proOptions={{ hideAttribution: true }}
      >
        <Background color={c.border} gap={20} size={0.5} />
        <Controls />
      </ReactFlow>
    </div>
  )
}

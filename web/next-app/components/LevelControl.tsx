'use client'

import { useState } from 'react'
import { toast } from 'sonner'
import { usePrimaryLevel } from '../hooks/usePrimaryLevel'
import { PRIMARY_LEVELS } from '../lib/primaryLevel'

/** The learner's Primary level, which scopes Learn and Review (lib/primaryLevel.ts). */
export default function LevelControl() {
  const [level, setLevel] = usePrimaryLevel()
  const [saving, setSaving] = useState(false)

  async function choose(next: string) {
    setSaving(true)
    try {
      await setLevel(next)
      toast.success('Level saved.')
    } catch {
      toast.error('Could not save your level. Try again.')
    } finally {
      setSaving(false)
    }
  }

  return (
    <div role="radiogroup" aria-label="Primary level" className="flex flex-wrap gap-2">
      {PRIMARY_LEVELS.map(l => {
        const active = l.level === level
        return (
          <button
            key={l.level}
            type="button"
            role="radio"
            aria-checked={active}
            disabled={saving}
            onClick={() => void choose(l.level)}
            className={`min-h-[44px] px-4 font-mono text-xs border disabled:opacity-60 ${
              active
                ? 'border-orchid-blue bg-orchid-blue text-white'
                : 'border-orchid-border text-orchid-secondary hover:border-orchid-blue hover:text-orchid-blue'
            }`}
          >
            {l.name}
          </button>
        )
      })}
    </div>
  )
}

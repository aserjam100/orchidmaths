'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'
import { ensureGuestId, isLoggedIn } from '../lib/auth'
import { PRIMARY_LEVELS } from '../lib/primaryLevel'
import { usePrimaryLevel } from '../hooks/usePrimaryLevel'

/**
 * Six buttons, one per Primary level. Choosing one saves it as the learner's level and opens
 * practice, with no placement test in between: a student who came to practise should be
 * answering a question one tap after they land. The test stays available from `/learn`.
 *
 * `onChosen` replaces the navigation for a caller that is already the practice page.
 *
 * A guest gets an id before the level is saved, because the level is stored on the student row
 * and the first answer is graded against it.
 */
export default function GradePicker({ next = '/learn', onChosen }: { next?: string; onChosen?: () => void }) {
  const { push } = useRouter()
  const [level, setLevel] = usePrimaryLevel()
  const [saving, setSaving] = useState<string | null>(null)

  async function choose(l: string) {
    setSaving(l)
    if (!isLoggedIn()) ensureGuestId()
    try {
      await setLevel(l)
    } catch {
      // The level is already in this browser, so practice still opens; the server falls back
      // to all of Primary until the save succeeds from Settings or Topics.
      toast.error('Could not save your grade. Practice will cover all of Primary for now.')
    }
    setSaving(null)
    if (onChosen) onChosen()
    else push(next)
  }

  return (
    <div role="group" aria-label="Choose your grade" className="grid grid-cols-2 sm:grid-cols-3 gap-3">
      {PRIMARY_LEVELS.map(l => {
        const mine = l.level === level
        return (
          <button
            key={l.level}
            type="button"
            onClick={() => void choose(l.level)}
            disabled={saving !== null}
            aria-pressed={mine}
            className={`flex min-h-[88px] flex-col items-center justify-center gap-1 border-2 px-4 py-4 transition-colors disabled:opacity-60 ${
              mine
                ? 'border-orchid-blue bg-orchid-blue text-white'
                : 'border-orchid-border bg-orchid-surface text-orchid-primary hover:border-orchid-blue hover:text-orchid-blue'
            }`}
          >
            <span className="font-serif text-3xl leading-none">{l.level}</span>
            <span className={`font-mono text-[12px] ${mine ? 'text-white/80' : 'text-orchid-muted'}`}>
              {saving === l.level ? 'Opening…' : l.name}
            </span>
          </button>
        )
      })}
    </div>
  )
}

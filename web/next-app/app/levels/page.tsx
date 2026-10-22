'use client'

/**
 * The curriculum by Primary level: level, then its MOE syllabus topics, then the concepts that
 * practise each topic.
 *
 * It replaces `/domains` as a student's way to browse. `/domains` lists all 21 topic domains of
 * the corpus, calculus and topology among them, which a Primary learner has no use for. Every
 * concept here is one `scripts/curriculum_moe.py` mapped to an outcome, and topics the syllabus
 * has but Orchid cannot practise yet are listed as such rather than left out.
 *
 * A learner may open a level above their own. Practice on its own is scoped to their level by
 * the server; choosing a concept here asks for it explicitly, which is the "if they ask" half.
 */

import { useEffect, useMemo, useState } from 'react'
import Link from 'next/link'
import { toast } from 'sonner'
import Header from '../../components/Header'
import BottomTabs from '../../components/BottomTabs'
import SectionHeader from '../../components/SectionHeader'
import Footer from '../../components/Footer'
import { useTheme } from '../../hooks/useTheme'
import { usePrimaryLevel } from '../../hooks/usePrimaryLevel'
import { getProgress, type ConceptProgress } from '../../lib/api'
import { getGuestId, getUserInfo } from '../../lib/auth'
import { PRIMARY_LEVELS, levelInfo, type LevelTopic } from '../../lib/primaryLevel'

function statusLabel(p: ConceptProgress | undefined): string {
  switch (p?.status?.toUpperCase()) {
    case 'MASTERED':
      return 'Mastered'
    case 'DECAYING':
      return 'Due for review'
    case 'LEARNING':
    case 'PRACTICING':
      return 'In progress'
    default:
      return 'Not started'
  }
}

function coverageNote(t: LevelTopic): string | null {
  if (t.concepts.length === 0) return 'Not practised in Orchid yet'
  if (t.none > 0 || t.partial > 0) return 'Partly practised in Orchid'
  return null
}

export default function LevelsPage() {
  const { mounted } = useTheme()
  const [myLevel, setMyLevel] = usePrimaryLevel()
  const [viewed, setViewed] = useState<string | null>(null)
  const [progress, setProgress] = useState<Record<string, ConceptProgress>>({})
  const [saving, setSaving] = useState(false)

  // Open on the learner's own level, or Primary 1 before one is chosen.
  const shown = viewed ?? myLevel ?? PRIMARY_LEVELS[0].level
  const level = levelInfo(shown) ?? PRIMARY_LEVELS[0]
  const aboveMine = myLevel !== null && PRIMARY_LEVELS.findIndex(l => l.level === shown) > PRIMARY_LEVELS.findIndex(l => l.level === myLevel)

  useEffect(() => {
    if (!mounted) return
    const id = getUserInfo()?.student_id ?? getGuestId()
    if (!id) return
    getProgress(id).then(setProgress).catch(() => setProgress({}))
  }, [mounted])

  const strands = useMemo(() => {
    const out: { strand: string; topics: LevelTopic[] }[] = []
    for (const t of level.topics) {
      const last = out[out.length - 1]
      if (last && last.strand === t.strand) last.topics.push(t)
      else out.push({ strand: t.strand, topics: [t] })
    }
    return out
  }, [level])

  async function chooseLevel() {
    setSaving(true)
    try {
      await setMyLevel(level.level)
      toast.success(`Your grade is now ${level.name}.`)
    } catch {
      toast.error('Could not save your grade. Try again.')
    } finally {
      setSaving(false)
    }
  }

  if (!mounted) return <div style={{ background: 'var(--bg)', minHeight: '100vh' }} />

  return (
    <>
      <Header />
      <main className="max-w-container mx-auto px-4 sm:px-6 py-8 sm:py-12 pb-[calc(80px+env(safe-area-inset-bottom))] lg:pb-12 min-w-0 overflow-x-hidden">
        <SectionHeader label="Curriculum" title="Topics" />
        <p className="text-orchid-secondary text-sm max-w-[60ch] mb-6">
          Pick a topic to practise it. Topics follow the MOE Primary Mathematics syllabus, and you can
          open any grade, including ones above yours.
        </p>

        <div role="tablist" aria-label="Primary level" className="flex flex-wrap gap-2 mb-6">
          {PRIMARY_LEVELS.map(l => {
            const active = l.level === shown
            return (
              <button
                key={l.level}
                type="button"
                role="tab"
                aria-selected={active}
                onClick={() => setViewed(l.level)}
                className={`min-h-[44px] px-4 font-mono text-xs border ${
                  active
                    ? 'border-orchid-blue bg-orchid-blue text-white'
                    : 'border-orchid-border text-orchid-secondary hover:border-orchid-blue hover:text-orchid-blue'
                }`}
              >
                {l.name}
                {l.level === myLevel && <span className={active ? 'text-white/70' : 'text-orchid-muted'}> · yours</span>}
              </button>
            )
          })}
        </div>

        <div role="tabpanel" aria-label={level.name}>
          <div className="flex flex-col sm:flex-row sm:items-center gap-3 mb-6 border border-orchid-border bg-orchid-surface p-4">
            <p className="flex-1 font-mono text-[12px] text-orchid-secondary">
              {level.level === myLevel
                ? `${level.name} is your grade. Practice stays within it.`
                : aboveMine
                  ? `${level.name} is above your grade. You can still practise its topics.`
                  : myLevel
                    ? `Your grade is ${levelInfo(myLevel)?.name}.`
                    : 'You have not chosen a grade yet, so practice covers all of Primary.'}
            </p>
            {level.level !== myLevel && (
              <button
                type="button"
                onClick={() => void chooseLevel()}
                disabled={saving}
                className="shrink-0 border border-orchid-blue text-orchid-blue hover:bg-orchid-blue-faint px-4 py-2 font-mono text-[12px] min-h-[44px] disabled:opacity-60"
              >
                {saving ? 'Saving…' : `Make ${level.name} my grade`}
              </button>
            )}
          </div>

          {strands.map(s => (
            <section key={s.strand} className="mb-8">
              <h2 className="font-mono text-[12px] uppercase tracking-wider text-orchid-muted mb-3">{s.strand}</h2>
              <ul className="border border-orchid-border divide-y divide-orchid-border bg-orchid-surface">
                {s.topics.map(t => {
                  const note = coverageNote(t)
                  return (
                    <li key={t.topic} className="px-4 py-3 min-w-0">
                      <div className="font-serif text-[15px] text-orchid-primary">{t.topic}</div>
                      {note && <div className="font-mono text-[11px] uppercase tracking-wider text-orchid-muted mt-0.5">{note}</div>}
                      {t.concepts.length > 0 && (
                        <ul className="mt-2 space-y-1">
                          {t.concepts.map(c => (
                            <li key={c.id} className="flex flex-col sm:flex-row sm:items-center gap-1 sm:gap-4 min-w-0">
                              <span className="flex-1 min-w-0 font-mono text-xs text-orchid-secondary break-words">{c.label}</span>
                              <span className="font-mono text-[11px] uppercase tracking-wider text-orchid-muted sm:w-32">
                                {statusLabel(progress[c.id])}
                              </span>
                              <Link
                                href={`/learn?concept=${encodeURIComponent(c.id)}`}
                                className="self-start sm:self-auto border border-orchid-blue text-orchid-blue hover:bg-orchid-blue-faint px-3 py-1 font-mono text-[12px] min-h-[36px] inline-flex items-center"
                              >
                                Practise
                              </Link>
                            </li>
                          ))}
                        </ul>
                      )}
                    </li>
                  )
                })}
              </ul>
            </section>
          ))}

          <p className="font-mono text-[12px] text-orchid-muted">
            Coverage of every outcome is in the{' '}
            <Link href="/docs/curriculum" className="text-orchid-blue hover:text-orchid-blue-hover">
              coverage report
            </Link>
            .
          </p>
        </div>
      </main>
      <Footer />
      <BottomTabs />
    </>
  )
}

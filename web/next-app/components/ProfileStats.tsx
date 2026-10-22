'use client'

import type { Scores } from '../lib/api'
import Avatar from './Avatar'

interface Props {
  name: string
  scores: Scores
  avatarSeed?: string
  avatarUrl?: string
  avatarPreset?: number | null
  /** Drop the outer border and padding, when a parent already draws the card. */
  bare?: boolean
  /** Skip the MASTERED tile, when the parent's position block already reports it. */
  hideMastered?: boolean
}

export const LEVEL_NAMES = {
  'Novice': '01',
  'Apprentice': '02',
  'Student': '03',
  'Scholar': '04',
  'Adept': '05',
  'Expert': '06',
  'Master': '07',
  'Grandmaster': '08',
  'Math Architect': '09',
} satisfies Record<string, string>

export default function ProfileStats({ name, scores, avatarSeed, avatarUrl, avatarPreset, bare, hideMastered }: Props) {
  const xpPct = scores.daily_xp_goal > 0
    ? Math.min(100, Math.round((scores.xp_today / scores.daily_xp_goal) * 100))
    : 0
  const levelNum = LEVEL_NAMES[scores.level] ?? '·'

  return (
    <div className={`flex flex-col sm:flex-row flex-wrap gap-4 sm:gap-5 items-start w-full max-w-full min-w-0 overflow-hidden ${bare ? '' : 'border-[0.5px] border-orchid-border p-4 sm:p-5 sm:px-6'}`}>
      <div className="flex flex-row gap-4 sm:contents w-full sm:w-auto">
        <div className="shrink-0 self-start">
          <Avatar seed={avatarSeed ?? name} name={name} size={72} url={avatarUrl} preset={avatarPreset ?? undefined} />
        </div>
        <div className="border-[0.5px] border-orchid-border-strong px-4 py-3.5 text-center min-w-[72px] sm:min-w-[80px] shrink-0 self-start">
          <div className="font-mono text-[10px] text-orchid-muted mb-1">
            LV
          </div>
          <div className="font-mono text-[36px] font-normal text-orchid-blue leading-none">
            {levelNum}
          </div>
          <div className="font-mono text-[11px] text-orchid-primary mt-1">
            {scores.level}
          </div>
        </div>
      </div>

      <div className="flex-1 min-w-0 w-full sm:min-w-0">
        <h2 className="font-serif text-[1.2rem] font-normal text-orchid-primary mb-0.5 truncate">
          {name}
        </h2>

        <div className="flex flex-wrap gap-x-6 gap-y-3 mt-2.5">
          <div className="min-w-0">
            <div className="font-mono text-[12px] text-orchid-muted">XP TODAY</div>
            <div className="font-mono text-[16px] text-orchid-primary">
              {scores.xp_today}
              <span className="font-mono text-[12px] text-orchid-muted"> / {scores.daily_xp_goal}</span>
            </div>
            <div className="mt-1 h-1 w-full max-w-[200px] bg-orchid-border">
              <div
                className="h-full bg-orchid-blue transition-all duration-300"
                style={{ width: `${xpPct}%` }}
              />
            </div>
          </div>

          <div className="min-w-0">
            <div className="font-mono text-[12px] text-orchid-muted">TOTAL XP</div>
            <div className="font-mono text-[16px] text-orchid-primary">
              {scores.lifetime_points.toLocaleString()}
            </div>
          </div>

          <div className="min-w-0">
            <div className="font-mono text-[12px] text-orchid-muted">STREAK</div>
            <div className={`font-mono text-[16px] ${scores.current_streak > 0 ? 'text-orchid-blue' : 'text-orchid-muted'}`}>
              {scores.current_streak} day{scores.current_streak !== 1 ? 's' : ''}
            </div>
          </div>

          {!hideMastered && (
            <div className="min-w-0">
              <div className="font-mono text-[12px] text-orchid-muted">MASTERED</div>
              <div className="font-mono text-[16px] text-orchid-primary">
                {scores.concepts_mastered}
              </div>
            </div>
          )}

          {/* The weekly score left with the leaderboard on 2026-10-22: it is the leaderboard's
              number, and its only link went there. */}
        </div>
      </div>
    </div>
  )
}

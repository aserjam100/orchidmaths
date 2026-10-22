import levelsJson from '../data/moe.levels.json'
import { authedFetch, ensureGuestToken, getAuthHeaders, isLoggedIn } from './auth'
import { prerequisiteClosure } from './diagnosticScope'

const API_BASE = process.env.NEXT_PUBLIC_API_URL || ''

/**
 * The learner's Primary level, which is how Orchid is organised for a student.
 *
 * Everything a student sees used to be organised by topic domain across the whole corpus:
 * 1250 concepts in 21 domains, most of them beyond Primary. A level is one of the MOE courses
 * `sg-p1` .. `sg-p6` that `scripts/curriculum_moe.py` writes, and each is cumulative: Primary 4
 * targets everything up to and including P4. The server reads the level from the student's
 * course and scopes practice to it (`internal/engine/level.go`); this module is the client's
 * half, and keeps a copy in localStorage so the level is known before `/api/auth/me` answers.
 */

export interface LevelTopic {
  topic: string
  strand: string
  full: number
  partial: number
  none: number
  concepts: { id: string; label: string }[]
}

export interface PrimaryLevel {
  level: string
  name: string
  courseId: string
  targets: string[]
  topics: LevelTopic[]
}

export const PRIMARY_LEVELS = levelsJson as PrimaryLevel[]

const LEVEL_KEY = 'orchid_primary_level'
export const LEVEL_EVENT = 'primary-level-changed'

export function levelInfo(level: string | null | undefined): PrimaryLevel | undefined {
  return PRIMARY_LEVELS.find(l => l.level === level)
}

/** The level a course ID names, or null for a course that is not a Primary level. */
export function levelForCourse(courseId: string | null | undefined): string | null {
  return PRIMARY_LEVELS.find(l => l.courseId === courseId)?.level ?? null
}

export function getLevel(): string | null {
  try {
    const v = localStorage.getItem(LEVEL_KEY)
    return levelInfo(v) ? v : null
  } catch {
    return null
  }
}

function storeLevel(level: string) {
  try {
    localStorage.setItem(LEVEL_KEY, level)
  } catch { /* storage unavailable — the level still reaches the server */ }
  window.dispatchEvent(new Event(LEVEL_EVENT))
}

/**
 * Takes the server's record when it names a level. Called whenever user info arrives, so a
 * level chosen on another device wins over this browser's copy.
 */
export function adoptServerLevel(courseId: string | null | undefined) {
  const level = levelForCourse(courseId)
  if (level && level !== getLevel()) storeLevel(level)
}

/**
 * Chooses a level: stored locally at once, then saved as the student's course. A guest gets a
 * guest token first, because the course is saved on the student row and the guest row is
 * created with the token.
 */
export async function setLevel(level: string): Promise<void> {
  const info = levelInfo(level)
  if (!info) return
  storeLevel(level)
  if (!isLoggedIn()) await ensureGuestToken()
  const res = await authedFetch(`${API_BASE}/api/courses/${info.courseId}/diagnostic`, {
    method: 'POST',
    headers: { ...getAuthHeaders() },
  })
  if (!res.ok) throw new Error(`Could not save level: ${res.status}`)
}

const scopeCache = new Map<string, Set<string>>()

/**
 * Every concept a learner at this level may be offered: the course targets and their
 * prerequisites, the same closure the server walks in `planning.PathForCourse`. With no level
 * chosen, it is all of Primary, which is also the server's default.
 */
export function levelScope(level: string | null | undefined): Set<string> {
  const info = levelInfo(level) ?? PRIMARY_LEVELS[PRIMARY_LEVELS.length - 1]
  let scope = scopeCache.get(info.level)
  if (!scope) {
    scope = new Set(prerequisiteClosure(info.targets))
    scopeCache.set(info.level, scope)
  }
  return scope
}

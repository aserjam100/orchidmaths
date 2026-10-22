'use client'

import { useCallback, useEffect, useState } from 'react'
import { getUserInfo } from '../lib/auth'
import { LEVEL_EVENT, adoptServerLevel, getLevel, setLevel as saveLevel } from '../lib/primaryLevel'

/**
 * The learner's Primary level, kept in step with the account. User info carries the server's
 * course, and it wins over this browser's copy whenever it names a level, so a level chosen on
 * another device arrives here on the next `auth-changed`.
 */
export function usePrimaryLevel(): [string | null, (level: string) => Promise<void>] {
  const [level, setLevelState] = useState<string | null>(null)

  useEffect(() => {
    const sync = () => {
      adoptServerLevel(getUserInfo()?.course_id)
      setLevelState(getLevel())
    }
    sync()
    window.addEventListener(LEVEL_EVENT, sync)
    window.addEventListener('auth-changed', sync)
    window.addEventListener('storage', sync)
    return () => {
      window.removeEventListener(LEVEL_EVENT, sync)
      window.removeEventListener('auth-changed', sync)
      window.removeEventListener('storage', sync)
    }
  }, [])

  const setLevel = useCallback((next: string) => saveLevel(next), [])
  return [level, setLevel]
}

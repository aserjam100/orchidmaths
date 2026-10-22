'use client'

import { Skeleton } from '../ui/skeleton'

// Mirrors the authed /profile layout (max-w-[820px]): one ProfileHeader card (identity + XP +
// where you are), the recommended-task rows, then the Activity heatmap.
export default function ProfileSkeleton() {
  return (
    <div
      role="status"
      aria-live="polite"
      aria-busy="true"
      aria-label="Loading profile"
      className="mx-auto w-full max-w-[820px] min-w-0 px-4 sm:px-6 py-8 sm:py-12 overflow-x-hidden"
    >
      {/* ProfileHeader: short sentences with the square avatar box on the left. */}
      <div className="border-[0.5px] border-orchid-border p-4 sm:px-5 w-full min-w-0">
        <div className="flex items-start gap-4 min-w-0">
          <Skeleton className="size-[80px] shrink-0 rounded-none" />
          <div className="min-w-0 flex-1">
            <Skeleton className="h-4 w-48" />
            <Skeleton className="mt-2 h-3.5 w-3/4" />
            <Skeleton className="mt-2 h-3.5 w-2/3" />
          </div>
        </div>
      </div>

      {/* Recommended task rows */}
      <div className="mt-6 border-[0.5px] border-orchid-border">
        {Array.from({ length: 3 }).map((_, i) => (
          <div
            key={i}
            className={i < 2 ? 'border-b-[0.5px] border-orchid-border px-3 py-2.5' : 'px-3 py-2.5'}
          >
            <Skeleton className="h-4 w-2/3 mb-1.5" />
            <Skeleton className="h-3 w-1/2" />
          </div>
        ))}
      </div>

      {/* Activity heatmap */}
      <section className="mt-8 flex min-w-0 flex-col items-stretch">
        <Skeleton className="h-5 w-24 mb-4" />
        <div className="w-full border border-orchid-border bg-orchid-surface p-4">
          <div className="grid grid-cols-[repeat(20,minmax(0,1fr))] gap-1">
            {Array.from({ length: 80 }).map((_, i) => (
              <Skeleton key={i} className="h-3 w-full rounded-[2px]" />
            ))}
          </div>
        </div>
      </section>
    </div>
  )
}

import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Leaderboard \u00b7 Orchid',
  description: 'Weekly leaderboard for Orchid · compete, improve and rise through the ranks.',
}

export default function LeaderboardLayout({ children }: { children: React.ReactNode }) {
  return children
}

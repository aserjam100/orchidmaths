export interface NavLink {
  label: string
  href: string
}

// HeaderLink is the Header's prop name for the same shape.
export type HeaderLink = NavLink

// Single source of truth for site navigation (Header + BottomTabs).
//
// Order carries meaning: Learn leads, and the orientation surfaces follow it.
//
// Learn is the only surface that asks a question, grades an answer and explains it. It used to
// lead into Study — a full reference library — because reading was treated as a way to learn.
// It is not: only demonstrated answers move a concept, so a learner who could read their way
// through the corpus had found a path around the only thing that counts. Reference material
// stayed; the route to it from navigation did not. It now lives inside `/learn`, scoped to the
// concept being learned (see components/ConceptReference).
//
// Bottom tabs are the primary mobile nav and stay at three: Practice, Topics, Me.
// Home is not a tab, because a student's home is practice and the logo leads there. The mobile compass menu shows ONLY
// overflow links not already in the tabs, so no destination appears twice on one screen.

// LEARN_LINKS is the doing group, and it is shown to everyone — a guest can
// answer questions, they just cannot have them remembered.
export const LEARN_LINKS: NavLink[] = [{ label: 'Practice', href: '/learn' }]

// REVIEW_LINKS is a Learn sub-mode: it grades, awards XP and reschedules
// memory, so it belongs beside Learn rather than looking like its own product.
// Auth-only, because /api/reviews/* sits behind authMiddleware — offering it to
// a guest would be a link to a 401.
export const REVIEW_LINKS: NavLink[] = [{ label: 'Review', href: '/review' }]

// DISPLAY_LINKS is the orientation group: where things sit in the curriculum and where the
// learner stands in it. None of these grade an answer.
export const DISPLAY_LINKS: NavLink[] = [
  // `/levels` is labelled Topics: a student looks for "what can I practise", and the page lists
  // their grade's MOE topics. It replaced Domains, which listed the corpus's 21 domains.
  { label: 'Topics', href: '/levels' },
  { label: 'How it works', href: '/how-it-works' },
]

// The leaderboard left navigation on 2026-10-22, at the owner's request, until there are enough
// learners for a ranking to mean something. The route still works.

// The concept graph and the glossary left student navigation on 2026-10-22. A student who came
// to practise has no use for either, and every extra destination is one more thing to decide
// before the first question. Both routes still work and the footer links them.

export const LOGGED_IN_LINKS: NavLink[] = [
  { label: 'Profile', href: '/profile' },
  { label: 'Settings', href: '/settings' },
]

export const LOGGED_OUT_LINKS: NavLink[] = [{ label: 'Login', href: '/login' }]

// ADMIN_LINKS is the administrative entry, offered only to staff (moderator, admin, owner).
//
// It belongs in the mobile overflow specifically: the desktop rail on `/profile` has carried the
// Admin entry since Admin V2, but on a phone `/profile` is reachable from the tab bar only, so an
// administrator who used their phone had no route to administration at all except typing a URL.
// Offering it grants nothing — every `/api/admin` route re-reads the role server-side
// (ADR-048, ADR-049) — so this is the invitation the server was always the authority for.
export const ADMIN_LINKS: NavLink[] = [{ label: 'Admin', href: '/admin' }]

// Hrefs covered by the bottom tabs — never repeated in the compass menu.
// Review is deliberately absent: the tab bar stays at three, and Review reaches
// the header row and the compass overflow instead.
export const TAB_HREFS: ReadonlySet<string> = new Set([
  '/learn',
  '/levels',
  '/profile',
])

// Full desktop header row: Learn first, then the orientation surfaces, then the account links.
export function desktopLinks(loggedIn: boolean): NavLink[] {
  return [
    ...LEARN_LINKS,
    ...(loggedIn ? REVIEW_LINKS : []),
    ...DISPLAY_LINKS,
    ...(loggedIn ? LOGGED_IN_LINKS : LOGGED_OUT_LINKS),
  ]
}

// Mobile compass overflow: destinations missing from the bottom tabs.
//
// `isStaff` appends the Admin entry. It defaults to false, so a caller that has not asked the
// question — a guest, a learner, or a component with no auth state — gets exactly the learner
// menu it got before, and only an account the server has already identified as staff sees it.
export function overflowLinks(loggedIn: boolean, isStaff = false): NavLink[] {
  const base = desktopLinks(loggedIn).filter(l => !TAB_HREFS.has(l.href))
  return isStaff ? [...base, ...ADMIN_LINKS] : base
}

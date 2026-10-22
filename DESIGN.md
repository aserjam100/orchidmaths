# Orchid Design System

A mathematics learning platform design system built on a geometric sans (Space Grotesk) paired with a technical monospace (JetBrains Mono). Dark-first, warm-neutral, blue-accented. Precision above decoration.

---

## Design Philosophy

**One family, two voices.** Space Grotesk carries all text — headings and body use weight alone for hierarchy. JetBrains Mono carries code, labels, nav, and structure. A grotesque for reading, a coding mono for structure: the pairing reads modern and technical without clashing.

The palette is defined as CSS custom properties in `web/next-app/app/globals.css`, and the colors are authored in **OKLCH**, not hex. OKLCH is perceptually uniform, so a lightness step is an even step to the eye, and the light and dark sets can share hue and chroma while differing only in lightness. A warm **gold** (`#c8a96e`) survives as the brand mark in two places that predate the token system — the 3D concept graph (`components/MathConceptGraph3D.tsx`) and the global error page (`app/global-error.tsx`) — but it is not a theme token and does not follow light and dark mode. The theme accent is blue.

This document describes the tokens; `globals.css` is the source. If the two disagree, the CSS is right.

---

## Color Palette

Every value below is declared in `:root` for light mode and `.dark` for dark mode. They are semantic tokens: components reference `var(--token)`, never a literal color, so a theme change is one file.

### Core surfaces and text

| Token | Light | Dark | Usage |
|---|---|---|---|
| `--bg` | `oklch(1 0 0)` | `oklch(0.16 0.004 264)` | Page background |
| `--surface` | `oklch(0.95 0 0)` | `oklch(0.21 0 0 / 0.8)` | Cards, raised panels |
| `--surface-elevated` | `oklch(0.92 0 0)` | `oklch(0.24 0 0)` | Hover states, code blocks |
| `--surface-highlight` | `oklch(0.89 0 0)` | `oklch(0.27 0 0)` | Active elements |
| `--text-primary` | `oklch(0.16 0.004 264)` | `oklch(1 0 0)` | Headings, body text |
| `--text-secondary` | `oklch(0.54 0 0)` | `oklch(0.66 0 0)` | Secondary text, descriptions |
| `--text-muted` | `oklch(0.54 0 0)` | `oklch(0.66 0 0)` | Labels, metadata, captions |

### Borders and code

| Token | Light | Dark | Usage |
|---|---|---|---|
| `--border` | `oklch(0.85 0.02 264)` | `oklch(0.35 0.02 264)` | Card borders, dividers |
| `--border-strong` | `oklch(0.75 0.03 264)` | `oklch(0.45 0.03 264)` | Strong separators |
| `--code-bg` | `oklch(0.95 0 0)` | `oklch(0.21 0 0 / 0.8)` | Code block backgrounds |
| `--division` | `oklch(0.85 0.02 264 / 0.45)` | `oklch(0.35 0.02 264 / 0.45)` | Section dividers |
| `--graph-surface` | `oklch(0.95 0 0 / 0.7)` | `oklch(0.21 0 0 / 0.65)` | Concept-graph panel |

### Semantic accents

| Token | Light | Dark | Usage |
|---|---|---|---|
| `--accent-blue` | `oklch(0.6 0.16 255)` | `oklch(0.72 0.145 230)` | Links, interactive states, focus rings |
| `--accent-blue-hover` | `oklch(0.53 0.17 255)` | `oklch(0.79 0.14 230)` | Blue hover state |
| `--accent-teal` | `oklch(0.68 0.15 250)` | `oklch(0.68 0.15 250)` | Secondary emphasis |
| `--accent-green` | `#059669` | `#10b981` | Correct answers, mastered |
| `--accent-red` | `#dc2626` | `#ef4444` | Incorrect answers |

Green and red stay as hex because their values are tuned to match the grading semantics rather than the OKLCH ramp.

### Heatmap ramp

The activity heatmap uses a five-step ramp, `--heat-0` through `--heat-4`, with matching foreground colors `--heat-fg-1` through `--heat-fg-4`. Light mode runs pale blue to deep blue; dark mode runs near-black to bright cyan. The foreground tokens flip between near-black and white so a label stays legible on every step.

---

## Typography

Fonts are loaded with `next/font/local` in `app/layout.tsx` and exposed as `--font-space-grotesk` and `--font-jetbrains-mono`.

| Role | Font Family | Weight | Size / Line | Responsive |
|---|---|---|---|---|
| Page Title | `'Space Grotesk', sans-serif` | 400 | `clamp(1.5rem, 5vw, 1.9rem)` / 1.2 | `max-sm:text-[1.5rem]` (SectionHeader) |
| Section Heading | `'Space Grotesk', sans-serif` | 400 | `1.9rem` / 1.2 | `max-sm:text-[1.5rem]` |
| Card Heading | `'Space Grotesk', sans-serif` | 400 | `1.1rem` / 1.3 | |
| Body | `'Space Grotesk', sans-serif` | 400 | 16px / 1.75 | `max-width: 640px` on mobile, bumps to 17px |
| Body Small | `'Space Grotesk', sans-serif` | 400 | 15px / 1.6 | |
| Inline Code | `'JetBrains Mono', monospace` | 400 | 0.9em | |
| Nav / Labels | `'JetBrains Mono', monospace` | 400 | 12–13px / 1.5 | |
| Code Blocks | `'JetBrains Mono', monospace` | 400 | 14px / 1.6 | |
| Metric Numbers | `'JetBrains Mono', monospace` | 400 | 2rem / 1 | `font-variant-numeric: tabular-nums` |
| ASCII Art | `'JetBrains Mono', monospace` | 400 | 13px / 1.6 | `overflow-x: auto` |

The scale was raised by one pixel on 2026-10-22 because the interface read too small. Tailwind's
`text-xs` and `text-sm` are overridden to 13px and 15px in `tailwind.config.js`, and every
`text-[9px]` to `text-[13px]` moved up one step. Graph node labels, SVG chart text and the math
keyboard kept their sizes, because their boxes are drawn around the text.

### Fluid typography pattern

Where font sizes need to adapt to screen width, use `clamp()`:

```css
font-size: clamp(min, preferred, max);
```

Examples:
- Hero heading: `clamp(2.2rem, 5vw, 3.8rem)`
- Body subtext: `clamp(0.95rem, 2.5vw, 1.1rem)`
- Small labels: `clamp(11px, 3vw, 12px)`

---

## Spacing & Layout

| Token | Value | Note |
|---|---|---|
| Container max-width | `1100px` (`max-w-container`) | Applied via `mx-auto px-6 max-sm:px-4` on all pages |
| Section padding (y) | `py-20 max-sm:py-12` | 80px desktop, 48px mobile |
| Code block padding | `16px 20px` | Unified across docs pages |
| Flex gap (default) | `24px` | Between section columns |
| Flex gap (tight) | `12px` | Between inline elements |
| Column gap (stacked mobile) | `16px` (`max-sm:gap-4`) | After stacking to `flex-col` |

---

## Borders

| Element | Style |
|---|---|
| Cards, panels | `0.5px solid var(--border)` |
| Strong separators | `0.5px solid var(--border-strong)` |
| Section dividers | `0.5px solid var(--border)` |
| Blue accents | `2px solid var(--accent-blue)` (left-border on formula blocks) |
| Buttons | No border-radius (`borderRadius: 0`) |
| Code blocks | `border none` with `background: var(--surface)` |

Prefer 0.5px borders throughout — they render as 1px on most displays and feel lighter than full 1px lines. No border-radius on buttons. Code blocks get surface backgrounds instead of borders.

---

## Link Hover Effect

All clickable text links use the `.link-underline` class — a line that animates from left to right on hover via a `::after` pseudo-element:

```css
.link-underline {
  position: relative;
  text-decoration: none;
}

.link-underline::after {
  content: '';
  position: absolute;
  left: 0;
  bottom: -2px;
  width: 100%;
  height: 1px;
  background: var(--accent-blue);
  transform: scaleX(0);
  transform-origin: left;
  transition: transform 0.2s ease;
}

.link-underline:hover::after {
  transform: scaleX(1);
}
```

Applied to: Header nav links, Footer links, docs page links, in-page reference links.

---

## Components

### Header

Sticky blurred nav bar at the top of every page. In `components/Header.tsx`.

- Position: `sticky top-0 z-100`, backdrop-filter: `blur(8px)`
- Background: semitransparent surface, driven by `--surface`
- Font: JetBrains Mono, 13px links (`text-xs`), 15px brand name (`text-sm`)
- Mobile: horizontal scroll (`overflow-x: auto`) with `white-space: nowrap` — no hamburger, no JS state
- Theme toggle: inline button in bar (not floating), hidden until `mounted`

### Footer

In `components/Footer.tsx`. Links use `display: flex; flex-wrap: wrap` with `gap: 4px 12px` — wraps gracefully on any screen width. Tagline in Space Grotesk.

### FormulaBlock

In `components/FormulaBlock.tsx`. Renders formulas with an accent left-border. Uses `white-space: pre-wrap` + `max-width: 100%` to prevent horizontal overflow on mobile.

### Section Header

In `components/SectionHeader.tsx`. Renders a centered page-section heading with an accent left-border block. Title font: Space Grotesk, `1.9rem max-sm:1.5rem`.

### DomainTable and ProgressionLevels

Both use `width: 100%` and JetBrains Mono for labels/structure. DomainTable uses `overflow-x: auto` for safe table scrolling.

---

## Diagrams (D2)

All architecture and workflow diagrams are pre-rendered D2 SVGs. No browser-side rendering. No Mermaid runtime.

Source files live in `web/next-app/diagrams/*.d2`, and the rendered SVGs are committed under `docs/diagrams/` and mirrored to `web/next-app/public/diagrams/`.

| Diagram | File | Used in |
|---|---|---|
| Architecture (5-layer stack) | `architecture.d2` | `docs/architecture.md`, `docs/system-design.md` |
| System design (architecture) | `system-design-architecture.d2` | `docs/system-design.md` |
| Data model | `system-design-data-model.d2` | `docs/system-design.md` |
| Request flow | `system-design-request-flow.d2` | `docs/system-design.md` |
| Grading pipeline | `system-design-grading-pipeline.d2` | `docs/system-design.md` |
| SM-2 & mastery | `system-design-sm2-mastery.d2` | `docs/system-design.md` |
| CAT diagnostic | `cat-diagnostic.d2`, `system-design-cat-diagnostic.d2` | `docs/system-design.md`, in-app docs |

Generate with `npm run diagrams`.

---

## Responsive Patterns

| Pattern | Implementation |
|---|---|
| Container padding | `px-6 max-sm:px-4` — 24px → 16px |
| Vertical spacing | `py-20 max-sm:py-12` — 80px → 48px |
| Column stacking | `flex flex-wrap gap-8 max-sm:flex-col max-sm:gap-4` |
| Nav overflow | Horizontal scroll on mobile (`overflow-x: auto`) |
| Footer wrapping | `display: flex; flex-wrap: wrap; gap: 4px 12px` |
| SVG scaling | `max-width: 100%; height: auto` |
| Code blocks | `overflow-x: auto` + `white-space: pre-wrap` for formulas |
| Fluid typography | `clamp()` on font sizes that need to scale |
| Hidden on mobile | `max-sm:hidden` for dividers/elements that break stacking |
| Vertical divider fix | `max-sm:hidden` on `borderLeft` dividers between stacked columns |

---

## Tailwind Configuration

Key customizations in `tailwind.config.js`:

```js
maxWidth: {
  container: '1100px',
},
fontFamily: {
  serif: ['var(--font-space-grotesk)', 'Space Grotesk', 'sans-serif'],
  sans: ['var(--font-space-grotesk)', 'Space Grotesk', 'sans-serif'],
  mono: ['var(--font-jetbrains-mono)', 'JetBrains Mono', 'monospace'],
}
```

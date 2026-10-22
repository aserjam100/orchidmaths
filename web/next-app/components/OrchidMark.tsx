// Orchid's mark: five petals and a lip, the same drawing as public/icons/icon.svg, inlined so it
// takes the surrounding text colour and size. It replaced the upstream λ glyph; a flower glyph
// would sit in the dingbat range test/noEmoji.test.ts forbids.
export default function OrchidMark({ className = '' }: { className?: string }) {
  return (
    <svg
      viewBox="-120 -130 240 240"
      width="1em"
      height="1em"
      aria-hidden="true"
      data-testid="orchid-mark"
      className={`inline-block shrink-0 align-[-0.125em] ${className}`}
      fill="currentColor"
    >
      <g>
        <ellipse cx="0" cy="-60" rx="27" ry="60" />
        <ellipse cx="0" cy="-60" rx="27" ry="60" transform="rotate(72)" />
        <ellipse cx="0" cy="-60" rx="27" ry="60" transform="rotate(144)" />
        <ellipse cx="0" cy="-60" rx="27" ry="60" transform="rotate(216)" />
        <ellipse cx="0" cy="-60" rx="27" ry="60" transform="rotate(288)" />
      </g>
      <circle r="14" fill="var(--bg, #fff)" />
    </svg>
  )
}

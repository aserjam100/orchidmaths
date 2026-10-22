# Domains in the complex plane

A **domain** is a set that is both open and connected. The two requirements are independent, and a set that fails either one is not a domain, however natural it looks.

## Openness

Openness says the set has no boundary points of its own: every point is surrounded by a small disk that stays inside. The open unit disk is open; the closed disk is not, and the real axis is not open in the complex plane, because a point on the line has no disk that stays on it.

Openness is about the interior, not about shape, so a set with a hole can still be open as long as the hole's boundary is excluded.

## Connectedness

Connectedness says the set cannot be split into two disjoint, non-empty open pieces. Two disjoint open disks are open but not connected, so they are not a domain even though each piece is.

The punctured plane \(\mathbb{C} \setminus \{0\}\) is connected despite the missing point: two points can be joined by a path that goes around the origin. Connectedness is about whether paths can travel inside the set, not about whether the set has holes.

## Both at once

A domain needs the two properties together: an open set in which any two points can be joined by a path that stays inside. The punctured plane, an open disk, the upper half-plane and an annulus \(\{z : 1 < |z| < 2\}\) all qualify.

The closed disk fails openness, and the union of two disjoint open disks fails connectedness — the two failure modes are different, and an explanation that separates them is the one that teaches the definition.

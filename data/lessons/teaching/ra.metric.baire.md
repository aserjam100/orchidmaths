# Baire's Category Theorem

**Baire category theorem:** a complete metric space cannot be written as a countable union of nowhere dense sets, so it is of the second category in itself.

## Nowhere dense and meagre

A set is nowhere dense when the closure of its interior is empty, so it is small in the sense that it cannot fill up any ball. A set is meagre, or of the first category, when it is a countable union of nowhere dense sets. The Baire category theorem says that a complete metric space is not meagre in itself: it cannot be covered by countably many nowhere dense sets. The theorem is an existence statement in disguise, because if a space cannot be meagre, then any countable union of closed sets with empty interior must fail to be the whole space.

## The nested-ball proof

The proof assumes a countable cover by nowhere dense sets and builds nested closed balls, each chosen to avoid the previous set entirely and with radius halved. The balls are nested with radii shrinking to zero, so in a complete space their intersection is a single point by the nested ball property. That point lies in no set of the cover, contradicting the assumption that the cover is everything. The trick is to use completeness exactly where shrinking closed balls guarantee a common point.

## Consequences

The theorem shows that the reals cannot be a countable union of nowhere dense sets, and it proves the existence of functions that are continuous nowhere, since the functions continuous at some point form a meagre set while the whole space of functions is complete. It is also the standard tool for proving that the pointwise limit of continuous functions has a point of continuity, and for showing that various sets defined by generic behaviour are dense. Each use converts the impossibility of a meagre cover into a positive existence result.

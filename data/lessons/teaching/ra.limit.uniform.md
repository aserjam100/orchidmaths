# Uniform Continuity

**Uniform continuity:** a function is uniformly continuous on a set when one distance delta controls the output for every pair of points, so the same delta works everywhere on the set.

## The definition

Uniform continuity strengthens ordinary continuity by making the choice of delta independent of the point. For every positive epsilon there is a single positive delta such that any two points of the set within delta of each other have outputs within epsilon, wherever those points lie. Ordinary continuity allows delta to depend on the point as well as on epsilon, which is weaker. The distinction matters because uniform continuity is what lets a limiting argument pass from local control to a statement holding across the whole set.

## Continuous but not uniformly

Ordinary continuity does not imply uniform continuity. The function which sends x to one over x is continuous on the interval from nought to one but not uniformly so, because near nought the graph becomes so steep that the required delta shrinks to nothing as the point approaches nought. The squaring function is continuous on the whole line but not uniformly continuous there, because points far out need a much smaller delta than points near the origin for the same output tolerance.

## Heine-Cantor and its uses

Uniform continuity is recovered on compact sets. The Heine-Cantor theorem states that a function continuous on a closed bounded interval is uniformly continuous there, so the compactness of the domain repairs the failure of the examples. Uniform continuity is the hypothesis that lets sequences of functions of a variable interchange limits with continuity and integration. It is also what guarantees that a Cauchy sequence of inputs produces a Cauchy sequence of outputs.

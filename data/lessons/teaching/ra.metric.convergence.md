# Convergence in Metric Spaces

**Convergence:** a sequence tends to a point when the distance to it can be made arbitrarily small beyond some index; the limit of a convergent sequence is unique in a metric space.

## The definition and uniqueness

A sequence in a metric space converges to a point when, for every positive epsilon, there is an index beyond which the distance from the term to that point is less than epsilon. The definition is the epsilon condition of the line with the absolute value replaced by the metric. A sequence has at most one limit, because two distinct limits would have disjoint balls around them, and a term far out would have to lie in both. Uniqueness uses the fact that distinct points in a metric space are separated by a positive distance.

## Cauchy sequences

A sequence is Cauchy when its own terms eventually get and stay arbitrarily close to one another, with no limit mentioned. Every convergent sequence is Cauchy, because two later terms are each close to the limit and hence close to each other. The converse requires completeness: a Cauchy sequence need not converge in a general metric space, and the rationals, with a Cauchy sequence approaching the square root of two, are the standard witness. The distinction between Cauchy and convergent is exactly the distinction between a space and its completion.

## Subsequences and boundedness

A convergent sequence is bounded, since its tail lies in one ball and only finitely many terms lie outside it. Every subsequence of a convergent sequence converges to the same limit, because the indices of a subsequence eventually exceed any threshold. These transfer properties make the metric-space theory coincide with the theory on the line wherever only the metric is used, and they are the reasons the abstract definitions were chosen to look the way they do.

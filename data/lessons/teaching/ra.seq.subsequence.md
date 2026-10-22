# Subsequences

**Subsequence:** a sequence obtained by keeping infinitely many terms of a given sequence in their original order, so the indices increase strictly.

## Definition and inheritance

A subsequence picks terms from a sequence at strictly increasing positions. Because the positions increase, a subsequence keeps the order of the original and simply discards terms. If the original sequence converges to L, every subsequence converges to L as well: the threshold that worked for the original also works for the subsequence, since its positions are all at least as large. Dropping terms cannot change where the remaining terms settle, only how quickly they arrive.

## Convergence through subsequences

A sequence converges to L exactly when every subsequence converges to L. One direction is inheritance. For the other, suppose every subsequence converged to L but the original did not; then some epsilon would leave infinitely many terms outside the epsilon neighbourhood of L, and those terms form a subsequence staying away from L, a contradiction. This equivalence is useful because a convenient subsequence can prove convergence, while two subsequences with different limits prove divergence.

## Limits and cluster points

The limits of convergent subsequences of a sequence are its limit points, also called cluster points. A convergent sequence has exactly one limit point, its limit. A bounded sequence that is not convergent has at least two, which is how the alternating sequence of minus one and one has both as limit points. The set of limit points measures how much a bounded sequence wanders, and the largest and smallest of them are the limit superior and limit inferior.

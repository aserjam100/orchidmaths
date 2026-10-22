# The Riemann Integral

**Riemann integral:** the limit of sums of the form value times subinterval length as the partition is refined, defined when the upper and lower sums meet at a single number.

## Partitions and sums

A partition cuts an interval into finitely many subintervals. On each subinterval a sample point is chosen and the function value there is multiplied by the length of the subinterval; the sum of those products is a Riemann sum. The integral is the limit of these sums as the partition is refined and the longest subinterval shrinks to zero, provided the same limit is reached no matter how the sample points are chosen. The condition is about the whole family of sums, not a single sequence of them.

## Upper and lower sums

The upper sum uses the supremum of the function on each subinterval and the lower sum uses its infimum, so the lower sum never exceeds the upper sum for the same partition. Refining a partition raises the lower sums and lowers the upper sums, so both sequences become tighter. The function is Riemann integrable exactly when the upper and lower sums can be brought arbitrarily close, which happens exactly when the set of discontinuities has measure zero. Continuous functions and monotone functions both pass this test.

## Worked: a right-endpoint sum

For the function that reports its input, the right-endpoint sum over equal subintervals sums the sample values times the common length. The arithmetic collapses to the square of the interval length times the triangular sum of the first n integers divided by n, which reduces to the interval length squared times n plus one over two n. As n grows this tends to half the square of the interval length, which is the integral, and the error is exactly the interval length squared over two n.

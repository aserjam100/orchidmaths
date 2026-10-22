# Comparison and Ratio Tests

**Comparison and ratio tests:** criteria that decide convergence of a series by bounding its terms by those of a known series, or by the limiting ratio of successive terms.

## The comparison test

If the terms of a series are nonnegative and no larger than the terms of a series already known to converge, then the partial sums of the first are bounded by those of the second, so the first converges as well. Reversing the inequality the other way, if the terms dominate those of a known divergent series, the larger series diverges too. The test works because a series of nonnegative terms is controlled entirely by the boundedness of its partial sums, so a termwise bound transfers that boundedness in the right direction.

## The ratio test

The ratio test looks at the limit of the absolute ratio of successive terms. If that limit is less than one, the terms eventually shrink at least geometrically, so the series converges absolutely. If the limit is greater than one, the terms eventually grow and cannot tend to zero, so the series diverges. The value one is the inconclusive boundary: the harmonic series and the convergent series of reciprocal squares both have ratio limit one while behaving oppositely, so the test can say nothing there.

## The root test and the boundary

The root test takes the limit of the n-th root of the absolute term and applies the same thresholds: below one converges, above one diverges, exactly one is inconclusive. It is most useful when the terms contain n-th powers, where the root removes the exponent. Both tests share the boundary at one, and the p-series supplies the answer there: the reciprocal n-th powers converge exactly when the exponent exceeds one, so the boundary case must be settled separately by comparison.

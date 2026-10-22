# The p-Series

**p-series:** the series of reciprocal powers; it converges exactly when the exponent exceeds one, and it is the reference for the comparison tests.

## The threshold

The p-series has terms equal to the reciprocal of the exponent-th power of the index. It converges exactly when the exponent exceeds one and diverges otherwise, including at the exponent one, where it is the harmonic series and diverges even though its terms tend to zero. The proof uses the integral test, comparing the series with the integral of the reciprocal power, which converges to a finite number only when the exponent exceeds one. The threshold is strict, so the exponent one is the boundary case that diverges.

## Why the boundary differs

At the exponent one the partial sums grow like the logarithm of the number of terms, which is unbounded, so the series diverges slowly. Just above one the partial sums stay bounded, because the integral of a steeper reciprocal power settles at a finite limit. The distinction between these two behaviours is the reason the comparison tests need a reference series and a strict inequality. Terms shrinking to zero is necessary for convergence but nowhere near sufficient.

## Comparison and the tests

The p-series is the yardstick for the comparison tests: a series with smaller positive terms than a convergent p-series converges, and one with larger terms than a divergent p-series diverges. If the terms behave asymptotically like a reciprocal power, the limit comparison test gives the same verdict, with the more slowly decaying power deciding the divergence. Recognising the decay rate of the terms as a power is often the fastest route to a verdict.

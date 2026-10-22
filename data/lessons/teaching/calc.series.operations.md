# Operations on Series

**Operations on series:** term-by-term addition, scalar multiplication and grouping preserve convergence, so convergent series can be combined like finite sums under simple rules.

## Adding and scaling

Two convergent series can be added term by term, and the sum of the result is the sum of the two sums, because the partial sums add and the limit of a sum is the sum of the limits. Multiplying every term of a convergent series by a constant scales every partial sum, so the series stays convergent and its sum is scaled by the same constant. These two operations make the convergent series over a fixed index set into a kind of vector space, closed under finite linear combinations with the expected sums.

## Grouping

Inserting parentheses into a convergent series does not change its sum, because the grouped partial sums are a subsequence of the original partial sums and a subsequence of a convergent sequence has the same limit. Removing parentheses is not always allowed: the alternating harmonic series converges, but regrouped it can be made to converge to a different number or to diverge. The asymmetry is the reason grouping is permitted only from a convergent series to its grouped form, and not the other way.

## Products and rearrangement

Rearranging the terms of an absolutely convergent series never changes its sum, but a conditionally convergent series can be rearranged to produce any sum, which is the substance of the Riemann rearrangement theorem. Multiplying two series needs the Cauchy product, whose convergence is guaranteed when at least one factor is absolutely convergent. These cautions matter because the ordinary algebra of finite sums does not transfer without the extra hypothesis of absolute convergence.

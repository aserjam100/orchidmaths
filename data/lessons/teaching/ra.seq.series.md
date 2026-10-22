# Series and Partial Sums

**Series:** the formal sum of the terms of a sequence; it converges when its sequence of partial sums converges, and its sum is the limit of those partial sums.

## Partial sums

A series is not added all at once; it is studied through its partial sums. The N-th partial sum adds the first N terms, and the series is said to converge when the sequence of partial sums converges. The sum of the series is that limit, and it can differ from any finite partial sum. A series with nonnegative terms always has an increasing sequence of partial sums, so it converges exactly when those partial sums are bounded above, by the monotone convergence theorem.

## Worked: the geometric series

For the geometric series with first term one and ratio r, the partial sum differs from one over one minus r by a remainder that is r to the N over one minus r. When the size of r is less than one, that remainder tends to zero, so the series converges to one over one minus r. When the size of r is at least one the remainder does not tend to zero and the series diverges. The single condition on the size of the ratio decides the whole family.

## The divergence test

If a series converges, its terms must tend to zero, because the term is the difference of two consecutive partial sums and both partial sums approach the same limit. The contrapositive is the divergence test: if the terms do not tend to zero, the series diverges. The test is one-way, and the harmonic series shows why: its terms tend to zero yet the series diverges, since its partial sums can be grouped to exceed any fixed bound. A passed divergence test therefore proves nothing.

# The Cauchy Criterion for Series

**Cauchy criterion for series:** a series converges exactly when, beyond some position, the sum of any block of consecutive later terms can be made smaller than any prescribed positive number.

## The statement

A series converges exactly when for every positive epsilon there is a position N such that the sum of the terms from any position after N up to any later position has absolute value less than epsilon. Equivalently, every block of consecutive tail terms is small. This is the Cauchy criterion applied to the sequence of partial sums, since the sum of a block of terms is the difference of two partial sums. It decides convergence from the terms alone, without naming the sum.

## Why it matters

The criterion is the bridge between convergence and the size of the terms. It never mentions a candidate limit, so it can be applied when the sum is unknown, and it is the tool behind the comparison and ratio tests. In the reals it is also a completeness statement: the partial sums converge precisely because a Cauchy sequence in the reals converges. The same criterion characterises convergence in any complete space, which is why it reappears throughout analysis.

## Worked: a geometric tail

For a geometric series with ratio r of size less than one, any block of terms beginning at position m has absolute value at most the size of the first term of the block divided by one minus the size of r. As m grows, the first term of the block shrinks to zero, so the block bound can be made smaller than any epsilon. This is the Cauchy criterion in action for a series already known to converge, and it shows how the remainder of a geometric series is controlled by its first omitted term.

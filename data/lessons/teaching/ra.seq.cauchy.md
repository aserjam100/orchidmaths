# Cauchy Sequences

**Cauchy sequence:** a sequence whose terms eventually get and stay arbitrarily close to one another, so that any two terms beyond some position are within any prescribed positive distance.

## The definition

A sequence is Cauchy when for every positive epsilon there is a position N such that any two terms from N onward are within epsilon of each other. The definition compares the terms directly and never names a limit, which is what makes it useful when the limit is unknown. Every convergent sequence is Cauchy, because two terms both close to the limit are close to each other by the triangle inequality. The interest is in the converse, which needs completeness.

## Cauchy and convergent

In the real numbers a sequence converges exactly when it is Cauchy, a property called completeness of the reals. In a general ordered field the two can differ. Every Cauchy sequence is bounded, since beyond the threshold all terms lie in a fixed interval and only finitely many come before. A Cauchy sequence with a convergent subsequence converges, to the subsequence limit, because the Cauchy property pulls the whole tail to that limit.

## Worked: the rationals are incomplete

Take the rational numbers obtained by truncating the decimal expansion of the square root of two after more and more digits. Any two of these truncations are eventually within any prescribed epsilon of each other, so the sequence is Cauchy in the rationals. If it converged to a rational limit, that limit would have to square to two, and no rational does. So the sequence is Cauchy but has no limit among the rationals, which is exactly the failure of completeness that the reals repair.

# The Archimedean Property

**Archimedean property:** for every real number x there is a positive integer n with n greater than x, so the natural numbers are unbounded in the reals and there is no positive number smaller than every reciprocal.

## No infinitely small positives

The Archimedean property says the integers grow past every real. Equivalently, for every positive epsilon there is a positive integer N with one over N less than epsilon. The property rules out an infinitesimal, a positive number smaller than every reciprocal of an integer. It follows from completeness, and it is what lets a finite grid approximate any interval however finely, which is the engine behind the density of the rationals.

## Worked: finding N

Suppose epsilon is one hundredth. The inequality one over N less than one hundredth is the same as N greater than one hundred, so the smallest integer that works is one hundred and one. In general the condition one over N less than a positive epsilon is N greater than the reciprocal of epsilon, and the smallest integer satisfying it is the floor of that reciprocal plus one. Existence is all the definition of a limit needs; the exact value of N is only a convenience.

## Consequences

From the Archimedean property several familiar facts follow. The sequence one over n tends to zero, because for any epsilon the tail lies below it. The floor function is well defined, since the integers are unbounded below as well as above. And the rationals can approximate any real, because a sufficiently fine grid of step one over n has a point inside any interval. Each of these is a place where a proof would otherwise need to assume that the integers do not run out.

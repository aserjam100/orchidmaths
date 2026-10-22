# The Intermediate Value Theorem

**Intermediate value theorem:** a continuous function on a closed bounded interval takes every value between its endpoint values, so if the endpoints straddle a target there is a point where the function equals it.

## The statement

If a function is continuous on the closed interval from a to b, and y is any number between the values of the function at a and at b, then there is a point of the interval where the function equals y. In particular, if the two endpoint values have opposite signs, the function has a root in the interval. Continuity is essential: a function that jumps over the target value at a single point can have endpoint values straddling the target with no point where it is attained.

## Proof by bisection

Suppose the function is below the target at the left end and above it at the right. Halve the interval; whichever half still has the function below at its left end and above at its right end is kept, and the other discarded. This produces nested closed bounded intervals of halving length. Their common point is the only candidate, and continuity there forces the value to equal the target, since the left endpoints have values below the target and the right endpoints have values above it and both kinds of endpoint converge to the common point.

## Consequences

The theorem is a root-finding and fixed-point theorem in disguise. It shows that every odd-degree polynomial has a real root, that a continuous function mapping an interval into itself has a fixed point when the endpoints are ordered suitably, and that the continuous image of an interval is again an interval. Its limitation is that it proves existence without locating the point, which is why the bisection method is run afterwards to approximate the root to a desired tolerance.

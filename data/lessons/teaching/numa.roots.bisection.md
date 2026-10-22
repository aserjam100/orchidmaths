# The Bisection Method

**Bisection:** repeatedly halve an interval that contains a sign change of the function, keeping the half whose endpoints still bracket a root.

## Halving a bracket

If a continuous function has opposite signs at two points, it has a root between them. Evaluating the midpoint decides which half still brackets a root, and the interval is replaced by that half. Each step halves the uncertainty in the location of the root, so the width falls by a factor of two every time.

## Worked: counting the steps

Starting from an interval of width eight, each bisection halves the width. To bring the width down to one half, the ratio is sixteen, and the base-two logarithm of sixteen is four, so four steps suffice. In general the number of steps is the smallest integer at least as large as the logarithm base two of the starting width divided by the tolerance.

## Strengths and limits

Bisection always converges when the initial bracket is valid, and its error bound is known in advance. It uses only the sign of the function, so it needs no derivative and cannot diverge. Its price is speed: it discards the size of the function values, so it converges linearly and slowly compared with methods that use them.

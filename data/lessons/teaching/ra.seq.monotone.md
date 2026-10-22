# Monotone Convergence

**Monotone convergence theorem:** an increasing sequence that is bounded above converges to its supremum, and a decreasing sequence bounded below converges to its infimum.

## Bounded monotone sequences converge

Let a sequence increase and stay below some upper bound. The set of its terms is nonempty and bounded above, so by completeness it has a least upper bound. Given a positive epsilon, some term is above the supremum minus epsilon, because otherwise the supremum minus epsilon would be a smaller upper bound. Monotonicity then keeps every later term between that one and the supremum, so the tail is within epsilon of the limit. The same argument runs downward for a decreasing sequence.

## Worked: a square-root recurrence

Take the sequence starting at zero with each term the square root of a plus the previous term, where a is positive. The terms increase because the square root of a plus x exceeds x whenever x is below the positive root of the equation L squared equals a plus L. They are bounded above by that root, since the square root stays below it. So the sequence converges, and passing to the limit in the recurrence gives L squared equals a plus L, whose positive solution is one plus the square root of one plus four a, all over two.

## Why it is useful

Monotone convergence gives existence without an explicit formula for the limit. It is the standard way to define constants such as e as the limit of increasing bounded sums, and to define the sum of a series of nonnegative terms as the supremum of its partial sums. It also underlies the comparison and integral tests for series, where boundedness and monotonicity of the partial sums are exactly what is checked. Existence first, value afterwards, is the pattern.

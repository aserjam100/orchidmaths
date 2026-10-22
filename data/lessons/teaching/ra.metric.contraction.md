# The Contraction Mapping Principle

**Contraction mapping principle:** a contraction of a complete metric space has exactly one fixed point, and the iterates from any starting point converge to it.

## The statement

A map of a metric space to itself is a contraction when it shrinks every distance by a fixed factor strictly less than one. The contraction mapping principle, also called the Banach fixed point theorem, states that a contraction of a complete metric space has exactly one fixed point, and that the sequence of iterates from any starting point converges to it. The proof shows the iterates form a Cauchy sequence, using completeness for the limit and the contraction to show the iterates get close, then uses the contraction once more to show the limit is the only fixed point.

## Error and rate

The principle is constructive: each iterate is closer to the fixed point, and the distance to the fixed point is bounded by the contraction factor over one minus the factor times the distance between the first two iterates. The bound is a geometric estimate, so the error falls at a fixed rate per step, which is exactly first-order convergence. This is why the principle is not only an existence theorem but also the justification for the fixed-point iteration, a numerical method for solving equations written in the form x equals the map of x.

## Worked: an affine contraction

For the affine map that multiplies by a fraction and adds a constant, the fixed point solves the linear equation obtained by setting the value equal to the map of it. Solving gives the constant divided by one minus the fraction, which is the limit of the iterates from any starting point. The fraction has absolute value below one, so the map contracts, and the explicit formula is the closed form of the geometric error estimate. This is the simplest full instance of the principle.

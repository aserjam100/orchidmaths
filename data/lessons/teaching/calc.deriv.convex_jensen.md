# Convexity and Jensen's Inequality

**Convex function:** a function whose graph lies below every chord between two of its points; equivalently, one with a nonnegative second derivative, and it satisfies Jensen's inequality.

## Two faces of convexity

A function is convex when the straight chord joining any two points of its graph lies above the graph between them. For a twice-differentiable function this is equivalent to a nonnegative second derivative everywhere on the interval, so the slope never decreases. Convexity is preserved by addition and by multiplication by a positive constant, and the sum of a convex and a linear function is convex. The two descriptions, geometric and analytic, are used interchangeably, and each is easier in different problems.

## Jensen's inequality

Jensen's inequality generalises the chord condition to more than two points: for a convex function and weights that sum to one, the value at the weighted average is at most the weighted average of the values. For two points with weights one half each this is exactly the midpoint of the chord lying above the graph. The inequality is reversed for concave functions, and equality holds only where the function is linear. It is the standard tool for proving inequalities about means and for bounding expectations of convex functions.

## Subgradients and uses

A convex function on an open interval is continuous but need not be differentiable everywhere, as the absolute value shows at zero; where it is differentiable the derivative is increasing. This monotone derivative is what makes convex optimisation tractable, because a zero derivative, where one exists, is a global minimum. Jensen's inequality and convexity together explain why averaging a convex cost is cheaper than costing the average, which is the intuition behind many risk and optimisation results.

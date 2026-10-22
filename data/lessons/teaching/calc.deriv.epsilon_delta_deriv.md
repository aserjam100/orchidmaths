# The Derivative as a Limit

**Epsilon-delta derivative:** the precise statement that a difference quotient tends to the derivative, which controls the difference quotient rather than the function values.

## The precise statement

The derivative of a function at a point is defined as the limit of the difference quotient, and writing that limit in epsilon-delta form controls the quotient. For every positive tolerance there must be a positive distance such that any input within that distance of the point, but not equal to it, makes the difference quotient within the tolerance of the derivative. The exclusion of the point itself is essential, because the quotient is undefined there. This is the same quantifier structure as the limit of a function, applied to the quotient rather than to the function.

## What differentiability gives

Differentiability at a point implies continuity there. The difference quotient stays near a finite number, so the change in the function is that number times the change in the input plus a small error, and as the input approaches the point the whole change tends to zero. The converse fails, and the absolute value at zero is the standard counterexample. The derivative can also fail to be continuous even where it exists, so continuity of the function is a strictly weaker property than differentiability.

## One-sided and two-sided limits

The derivative is defined by a two-sided limit, so a corner, where the one-sided slopes disagree, has no derivative even though both one-sided derivatives exist. A vertical tangent has infinite one-sided slopes and is likewise not differentiable. The epsilon-delta definition makes these failures precise: at a corner no single number can be within tolerance of the quotient from both sides at once. Recognising the two-sided requirement is what separates a derivative from the one-sided rates that appear in piecewise problems.

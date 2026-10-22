# Multiple Roots and Modified Newton

**Multiple root:** a root at which the function and some of its derivatives vanish, at which ordinary Newton iteration slows to linear convergence.

## Counting the multiplicity

A root has multiplicity m when the function and its first m minus one derivatives all vanish there while the m-th derivative does not. A simple root has multiplicity one, a double root two. Near a multiple root the graph is tangent to the axis, and the flatness changes how every method behaves.

## Why Newton slows down

At a simple root the derivative is nonzero and Newton converges quadratically. At a root of multiplicity m greater than one the derivative is also zero, so the quadratic term that drives the fast convergence vanishes and the method drops to linear convergence. Each step still reduces the error, but only by the fixed factor one minus one over m.

## The modified iteration

Multiplying the Newton step by the multiplicity m restores quadratic convergence when m is known. When m is unknown it can be estimated from the ratio of successive function values. The alternative is to apply Newton to the function divided by its derivative, which has a simple root wherever the original had a multiple one.

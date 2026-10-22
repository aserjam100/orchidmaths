# Uniform Convergence

**Uniform convergence:** a mode of convergence of functions in which one index controls the whole domain at once, so the approximation error is bounded simultaneously everywhere.

## One index for the whole domain

A sequence of functions converges uniformly on a set when, for every positive epsilon, there is an index beyond which every function in the sequence is within epsilon of the limit at every point of the set at once. The one index serves the whole domain, which is what the word uniform names. Pointwise convergence allows the index to depend on the point as well as on epsilon, and the two notions can differ: the sequence of powers of the input converges to zero at every point strictly inside the unit interval, yet not uniformly there, because the index needed near one grows without bound.

## The sup norm criterion

Uniform convergence is equivalent to the supremum of the distance between the n-th function and the limit tending to zero as n grows. Reading the definition through the supremum turns a statement about every point into a statement about a single number, and it makes the difference from pointwise convergence visible: pointwise convergence only bounds the error at each point separately. The criterion is the practical test, because the supremum of the error can often be computed directly.

## What uniformity buys

Uniform convergence is strong enough to preserve continuity, boundedness and the value of an integral. The uniform limit of continuous functions is continuous, and the limit of the integrals is the integral of the limit on a bounded interval. Pointwise convergence preserves none of these, which is why the twentieth century recast large parts of analysis in terms of uniform and stronger modes of convergence. Uniform convergence of a series is the hypothesis under which a series of functions can be integrated term by term.

# Convergence in Probability

**Convergence in probability:** a sequence $X_n$ converges to $X$ when, for every positive $\varepsilon$, the probability that the two differ by more than $\varepsilon$ tends to zero.

## The epsilon condition

The requirement is that for each fixed tolerance the chance of exceeding it can be made as small as we like by taking the index large enough. The tolerance is fixed first and the probability then shrinks, so the statement controls the size of a typical departure rather than the behaviour of any one sequence. This is the weakest of the modes of convergence.

## Relation to distributional convergence

Convergence in probability implies convergence in distribution, because controlling the distance between the variables controls their laws. The reverse implication fails in general, since convergence in distribution constrains only the marginal laws and says nothing about how the variables are related to one another. The two notions coincide only when the limit is a constant.

## The weak law as its home

The weak law of large numbers is exactly the statement that the sample mean converges in probability to the population mean. Its Chebyshev proof needs a finite variance, and the epsilon condition is what the resulting bound feeds into. Convergence in probability is the natural language for laws of large numbers.

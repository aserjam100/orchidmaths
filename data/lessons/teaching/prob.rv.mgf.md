# Moment Generating Functions

**Moment generating function:** $M(t) = E[e^{tX}]$, defined for the values of $t$ where the expectation is finite. Its derivatives at zero are the moments of $X$.

## Generating the moments

Expanding $e^{tX}$ as a power series and taking expectations term by term shows that the coefficient of $t^n/n!$ is the $n$-th moment. Differentiating $n$ times at zero therefore extracts that moment. The function is called generating because a single expression encodes the entire moment sequence, when the series converges around zero.

## Worked: reading the normal

The normal distribution with mean $\mu$ and variance $\sigma^2$ has $M(t) = e^{\mu t + \sigma^2 t^2/2}$. Reading the exponent, the coefficient of $t$ is the mean and twice the coefficient of $t^2$ is the variance. For $M(t) = e^{2t + 2t^2}$ the mean is two and the variance is four, which can be confirmed by differentiating twice at zero.

## Sums of independent variables

If two independent variables have moment generating functions $M_X$ and $M_Y$, then their sum has $M_X M_Y$, because the exponential of a sum factors and independence splits the expectation. This turns the distribution of a sum into a product of functions, which is how the binomial, Poisson and normal families are shown to be closed under the appropriate sums.

# Independent Random Variables

**Independence:** two random variables are independent when their joint distribution factors into the product of their marginals everywhere, and only then.

## The factorisation

Independence means $p(x, y) = p_X(x) p_Y(y)$ at every pair, with the same statement for densities. Equivalently, the events that $X$ lies in a set and $Y$ lies in another are independent for every choice of the two sets. The joint law is then fully reconstructed from the two margins.

## Worked: expectations factor

For independent variables the expectation of a product splits, $E[XY] = E[X]E[Y]$, because the double sum factorises. Two independent Bernoulli variables with success probability one half each have $E[XY] = 0.25$ and hence zero covariance. Independence is the hypothesis that makes many expectations collapse into products.

## Stronger than uncorrelatedness

Zero covariance does not imply independence. If $X$ is equally likely to be plus one or minus one and $Y$ is the square of $X$, the covariance is zero because the product averages out, yet $Y$ is a deterministic function of $X$, so the two are as dependent as two variables can be. Independence is a statement about the whole joint law, not just its second moments.

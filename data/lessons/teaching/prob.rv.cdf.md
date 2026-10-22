# The Distribution Function

**Distribution function:** $F(x) = P(X \le x)$. It is non-decreasing and right-continuous, with limit zero at minus infinity and limit one at plus infinity.

## Accumulating probability

The distribution function gathers all the probability up to the point $x$. Monotonicity follows because the events $\{X \le x\}$ grow as $x$ increases. Right-continuity is a convention that matches the use of a closed inequality, and the two endpoint limits record that the variable is finite with probability one. Together these properties characterise distribution functions completely.

## Worked: an interval

The probability that $X$ falls in the interval $(a, b]$ is $F(b) - F(a)$, because subtracting the accumulated probability below $a$ leaves the probability of the interval ending at $b$. If the distribution function jumps at a point, the size of the jump is the probability that $X$ equals that point exactly. Continuous random variables have no jumps, so every single point has probability zero.

## Jumps and densities

A distribution function is a step function exactly when the random variable is discrete. It is an integral of a density, and hence continuous, when the random variable is continuous. The general distribution function may do both: a staircase with a smooth part between the steps. This is why one object, the distribution function, describes every random variable without a case split.

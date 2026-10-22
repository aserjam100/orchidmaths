# Chebyshev's Inequality

**Chebyshev's inequality:** For a variable with mean $\mu$ and standard deviation $\sigma$, $P(|X - \mu| \ge k\sigma) \le 1/k^2$ for every positive $k$.

## A bound from the variance

Apply Markov's inequality to the nonnegative variable $(X - \mu)^2$. Its mean is the variance $\sigma^2$, so the probability that it reaches $k^2\sigma^2$ is at most $\sigma^2/(k^2\sigma^2) = 1/k^2$. That event is exactly the event $|X - \mu| \ge k\sigma$, which gives the stated bound. Adding the second moment to the toolkit sharpens what the mean alone could say.

## Worked: two standard deviations

With $k = 2$ the bound is $1/4$, so at least three quarters of the probability lies within two standard deviations of the mean, whatever the distribution. With $k = 4$ the bound is $1/16$, and with $k = 5$ it is $1/25$. In the form $P(|X - \mu| \ge d) \le \sigma^2/d^2$, the same inequality answers questions posed in the original units.

## Toward the law of large numbers

Apply the inequality to the sample mean of $n$ independent variables. The sample mean has the same expectation and variance $\sigma^2/n$, so the probability that it departs from the mean by more than a fixed amount is bounded by $\sigma^2/(n\varepsilon^2)$, which tends to zero. That is the weak law of large numbers, and it falls out of Chebyshev's inequality as its first application.

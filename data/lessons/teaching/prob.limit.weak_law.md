# The Weak Law of Large Numbers

**Weak law of large numbers:** for independent observations with mean $\mu$ and finite variance, the sample mean converges in probability to $\mu$.

## Convergence in probability

The sample mean converges in probability to the population mean when, for every positive $\varepsilon$, the probability that the two differ by more than $\varepsilon$ tends to zero. This is a statement about the shrinking chance of a departure, not about any single sequence of averages settling down. The word weak marks it against the stronger almost-sure convergence of the strong law.

## Worked: Chebyshev's bound

The sample mean of $n$ independent observations has the population mean and variance $\sigma^2/n$. Chebyshev's inequality bounds the chance of a departure of at least $\varepsilon$ by $\sigma^2/(n\varepsilon^2)$. With standard deviation two, one hundred observations and a threshold of one, the bound is $4/100 = 0.04$, and it continues to fall as the sample grows.

## Weak versus strong

Almost-sure convergence is strictly stronger than convergence in probability: the strong law implies the weak law, and there are sequences that converge in probability without converging almost surely. The weak law needs only a finite variance for the Chebyshev proof, while the strong law holds under a finite mean. Knowing which mode a result uses is what tells you how strong the hypothesis has to be.

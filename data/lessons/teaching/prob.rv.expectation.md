# Expectation of a Function

**Law of the unconscious statistician:** For a function $g$, the expectation of $g(X)$ is $\sum g(x)p(x)$ in the discrete case and $\int g(x)f(x)\,dx$ in the continuous case, without first finding the distribution of $g(X)$.

## The law of the unconscious statistician

To compute the expectation of a function of $X$ one never needs the distribution of the new variable. Each value of $X$ is weighted by its own probability and transformed by $g$ at the same time. The identity is the reason the second moment, the variance and the moment generating function are all routine sums or integrals of the original variable.

## Worked: the second moment

A discrete variable takes the values zero, one and two with probabilities each proportional to weights two, three and five, so the probabilities are 0.2, 0.3 and 0.5. The second moment is $0 \cdot 0.2 + 1 \cdot 0.3 + 4 \cdot 0.5 = 2.3$. The squares are weighted, not the values, which is exactly what the law prescribes for $g(x) = x^2$.

## Linearity and moments

Expectation is linear: $E[aX + b] = aE[X] + b$ whatever the distribution, and the same holds for sums of random variables even when they are dependent. Linearity is what makes the sample mean an unbiased estimate of the population mean. Higher moments are computed the same way, and the collection of moments, when it exists, pins down many distributions.

# The Memoryless Property

**Memoryless:** A lifetime is memoryless when $P(X > s + t \mid X > s) = P(X > t)$ for all positive $s$ and $t$. The geometric and exponential distributions have this property.

## No memory of the past

The conditional probability that the lifetime exceeds $s + t$, given that it already exceeded $s$, equals the unconditional probability of exceeding $t$. Age carries no information: a used item is as good as a new one, in the model. This is a strong assumption that fits waiting times between independent events and discrete success counts, and it is clearly false for a machine that wears out.

## Worked: the geometric distribution

For a geometric variable counting trials until the first success, $P(X > n) = (1-p)^n$. Then $P(X > s + t \mid X > s) = (1-p)^{s+t}/(1-p)^s = (1-p)^t = P(X > t)$. The calculation works because the survival probabilities are powers, so the ratio depends only on the remaining time. The exponential distribution has the same power structure in continuous time.

## Why only two distributions

The memoryless equation forces the survival function to satisfy $S(s+t) = S(s)S(t)$, and the only well-behaved nonnegative solutions are exponentials. Discrete trials give the geometric distribution; continuous time gives the exponential. This uniqueness is why the property is a strong modelling claim rather than a general convenience: adopting it commits the model to one shape.

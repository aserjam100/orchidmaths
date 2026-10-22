# The Probability Density Function

**Density:** A continuous random variable has a density $f$ with $P(a \le X \le b) = \int_a^b f(x)\,dx$ and $\int f(x)\,dx = 1$ over the whole line.

## Probability as area

For a continuous variable, probability is the area under the density over the interval of interest. The density itself is not a probability: it can exceed one on a short interval, and it carries units of probability per unit of the variable. Only its integral has the units of probability, which is why the two ideas must not be conflated.

## Worked: integrating a density

Let $X$ have density $f(x) = 2x$ on the unit interval and zero outside. Then $P(X \le a)$ is the integral of $2x$ from zero to $a$, which equals $a$ squared. At $a = 0.5$ this is 0.25, and at $a = 0.8$ it is 0.64. The density integrates to one over the whole interval, which confirms it is normalised.

## Points have probability zero

If $X$ is continuous, the probability that it equals any single value is zero, because an integral over a single point has zero width. That is why a density can be changed at finitely many points without affecting any probability, and why continuous models state intervals rather than exact values. A discontinuity in the density is harmless; a positive mass at a point is not allowed.

# Slutsky's Theorem

**Slutsky's theorem:** if $X_n$ converges in distribution to $X$ and $Y_n$ converges in probability to a constant $c$, then sums and products converge in distribution to the corresponding combinations with $c$.

## Mixing the two modes

The theorem lets one factor converge weakly while the other collapses to a constant. The constant-limit factor is strong enough to be treated as a fixed number in the limit, so the weak factor keeps its distributional limit. This mix of modes is exactly what appears when a statistic is divided by a random but consistent scale estimate.

## Replacing a random scale

In practice the denominator of a test statistic is rarely constant, but it is usually a consistent estimate of a fixed scale. Slutsky's theorem justifies replacing it by its probability limit, so the standardised statistic keeps its normal limit. This is the step that makes t-type statistics asymptotically normal.

## From continuous mapping

The theorem follows from the continuous mapping theorem applied to the pair, together with the fact that a probability limit may be inserted as a constant. The product case needs the weak factor to be bounded in probability, which tightness supplies. Every step is an application of weak convergence to a continuous function of the pair.

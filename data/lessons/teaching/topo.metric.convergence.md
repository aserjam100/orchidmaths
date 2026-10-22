# Convergence in Metric Spaces

**Convergence:** A sequence $x_n$ converges to $x$ when for every $\varepsilon > 0$ there is an $N$ such that $d(x_n, x) < \varepsilon$ for all $n > N$. Equivalently, every ball around $x$ contains all but finitely many terms.

## Eventually inside every ball

The definition is a promise about the tail: the first finitely many terms do not matter, and no matter how small the tolerance, the sequence eventually stays within it. In $\mathbb{R}$ this is the familiar epsilon-N definition, and in any metric space it uses the metric in the same way.

## Worked: three limits

1. $x_n = 1/n$: given $\varepsilon$, choose $N > 1/\varepsilon$; then $1/n < \varepsilon$ for all $n > N$, so the limit is 0.
2. $x_n = 3 + 1/n$: the tail is 3 plus a vanishing correction, so the limit is 3.
3. $x_n = (7n)/(2n+1)$: divide by $n$ to get $7/(2 + 1/n)$, which tends to $7/2 = 3.5$.

## Limits are unique

If $x_n$ converges to $x$ and to $y$ with $x$ different from $y$, let $r = d(x, y)/2$. The tails eventually lie in both balls of radius $r$, which are disjoint, a contradiction. So a metric-space limit is unique, and the notation $\lim x_n = x$ is unambiguous.

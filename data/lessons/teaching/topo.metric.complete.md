# Completeness

**Complete space:** A metric space is complete when every Cauchy sequence converges to a point of the space. Cauchyness is about the terms getting close to each other; completeness is the promise that this is enough to guarantee a limit inside the space.

## Cauchy sequences do not name their limit

A Cauchy sequence satisfies the tolerance condition without reference to a target: for every $\varepsilon$ there is an $N$ beyond which any two terms are within $\varepsilon$. Every convergent sequence is Cauchy, but the converse can fail when the space has gaps.

## Worked: the rationals are not complete

1. Approximate $\sqrt{2}$ by rationals: 1, 1.4, 1.41, 1.414, and so on.
2. The terms get arbitrarily close to each other, so the sequence is Cauchy in $\mathbb{Q}$.
3. Its limit is $\sqrt{2}$, which is not rational, so the sequence has no limit in $\mathbb{Q}$.

The completion of $\mathbb{Q}$ fills exactly these gaps and produces $\mathbb{R}$.

## Compact implies complete

A Cauchy sequence in a compact metric space has a convergent subsequence, because every compact metric space is sequentially compact. The whole Cauchy sequence then converges to the same limit: once the tail is close to itself and one term is close to the limit, every later term is close to the limit. So compact spaces are complete, while complete spaces need not be compact.

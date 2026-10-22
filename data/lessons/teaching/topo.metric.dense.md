# Dense Sets and Separability

**Dense set:** A subset $D$ of a space $X$ is dense when its closure is all of $X$, equivalently when every nonempty open set meets $D$. A space is separable when it has a countable dense subset.

## Meeting every open set

Density is a covering condition: no open set can avoid $D$ entirely. It does not require $D$ to contain its limit points, so a dense set can be open, closed, or neither. The rationals are dense in the reals, and so are the irrationals: each meets every interval.

## Worked: the rationals are dense

1. Take any open interval $(a, b)$ with $a < b$.
2. Between any two reals there is a rational, so choose $q$ in $(a, b)$.
3. Thus every nonempty open interval meets $\mathbb{Q}$, which is exactly density.

The same argument shows that $\mathbb{Q}$ is countable, so $\mathbb{R}$ is separable. A countable dense subset is a small skeleton that still touches every open region.

## Separable spaces

Separability is a countability axiom: it rules out spaces too large to be approximated by a countable set. $\mathbb{R}^n$ is separable because the points with rational coordinates are countable and dense. An uncountable discrete space is not separable, since the only dense subset is the whole space.

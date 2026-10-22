# Orthogonal Complements

**Orthogonal complement:** For a subspace $U$ of an inner-product space $V$, the orthogonal complement $U^\perp$ is the set of vectors orthogonal to every vector of $U$. It is a subspace, it meets $U$ only at the zero vector, and together they span $V$.

## Vectors orthogonal to a whole subspace

Checking orthogonality to a spanning set is enough: if $v$ is orthogonal to each basis vector of $U$, it is orthogonal to every linear combination. In $\mathbb{R}^n$ this turns into a homogeneous linear system whose solutions form $U^\perp$.

## Worked: complements in R^3

1. Let $U$ be the line spanned by $(1,0,0)$.
2. A vector $(x,y,z)$ is orthogonal to it exactly when $x = 0$.
3. Those vectors form the $yz$-plane, of dimension 2.

For the plane $x + y + z = 0$, the orthogonal complement is the normal line spanned by $(1,1,1)$, of dimension 1.

## Dimension and direct sums

Every vector splits uniquely as $u + u^\perp$ with $u \in U$ and $u^\perp \in U^\perp$, so $V = U \oplus U^\perp$. Taking dimensions gives $\dim U + \dim U^\perp = \dim V$, the identity behind projection and least squares. In $\mathbb{R}^n$ the double complement returns the original subspace: $(U^\perp)^\perp = U$.

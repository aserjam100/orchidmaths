# Injective, Surjective and Invertible Maps

**Three properties:** A linear map $T: V \to W$ is injective when $T(u) = T(v)$ forces $u = v$, surjective when every vector of $W$ is hit, and invertible when a linear inverse map exists. For finite-dimensional spaces, the three are linked by dimension.

## Three properties of a linear map

Injectivity is a statement about the kernel: $T$ is injective exactly when $\ker T = \{0\}$. Surjectivity is a statement about the image: $T$ is surjective exactly when $\operatorname{im} T = W$. An invertible map must be both, and its inverse is automatically linear.

## Worked: a map that cannot be injective

1. Take a linear map $T: \mathbb{R}^3 \to \mathbb{R}^2$.
2. Its rank is at most the dimension of the codomain, so the rank is at most 2.
3. Rank-nullity gives nullity at least $3 - 2 = 1$, so the kernel contains a nonzero vector and $T$ is not injective.

The same counting shows that a map $\mathbb{R}^2 \to \mathbb{R}^3$ cannot be surjective.

## The square-matrix shortcut

When the domain and codomain have the same finite dimension $n$, injective, surjective and invertible are equivalent. A trivial kernel gives rank $n$, and a map of rank $n$ from an $n$-dimensional space hits everything. For a square matrix this is also the statement that the determinant is nonzero, and the inverse is represented by the inverse matrix.

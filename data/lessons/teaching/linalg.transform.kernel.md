# Kernel and Image

**Kernel and image:** For a linear map $T: V \to W$, the kernel is $\{v \in V : T(v) = 0\}$ and the image is $\{T(v) : v \in V\}$. The kernel is a subspace of the domain, the image is a subspace of the codomain, and rank-nullity ties their dimensions together.

## What maps to zero

The kernel collects the directions the map collapses. A nonzero kernel vector means two different inputs share an output, so the map is not injective. The image collects the outputs that are actually reached, so its dimension is the rank of the map.

## Worked: kernel dimension by rank

1. Take $A = [[1, 2], [2, 4]]$. The second row is twice the first, so the rank is 1.
2. Rank-nullity gives nullity $2 - 1 = 1$: the vector $(2, -1)$ maps to zero.
3. Take $A = [[1, 0, 0], [0, 1, 0]]$ from $\mathbb{R}^3$ to $\mathbb{R}^2$. The rank is 2, so the nullity is $3 - 2 = 1$: the third coordinate is free.

So the kernel dimension is the number of free variables after elimination, and the rank is the number of pivots.

## Image, rank and injectivity

The image of a matrix map is its column space, whose dimension is the rank. A linear map is injective exactly when its kernel is the zero subspace, because $T(u) = T(v)$ is the same as $T(u - v) = 0$. For a map from an $n$-dimensional space, rank-nullity says $\dim\ker T + \operatorname{rank} T = n$, so the kernel measures exactly how much of the domain is lost.

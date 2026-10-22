# Betti Numbers

**Betti number:** The $n$-th Betti number $b_n$ is the rank of the $n$-th homology group, the number of independent $n$-dimensional holes. $b_0$ counts connected components, $b_1$ counts independent loops, and $b_2$ counts enclosed voids.

## Counting independent holes

Rank counts generators, not raw cycles: a loop that bounds a disk contributes nothing to $b_1$ because it is a boundary. For a graph, $b_1 = E - V + 1$, the number of independent cycles. For a closed orientable surface of genus $g$, $b_1 = 2g$ and $b_2 = 1$.

## Worked: the torus and the sphere

1. A circle has $b_1 = 1$: one independent loop.
2. A torus has $b_0 = 1$, $b_1 = 2$ and $b_2 = 1$: two loops and one void.
3. A sphere has $b_0 = 1$, $b_1 = 0$ and $b_2 = 1$: no loops, one void.

## Betti numbers and Euler characteristic

The Euler characteristic is the alternating sum of the Betti numbers: $\chi = b_0 - b_1 + b_2 - \cdots$. For the torus, $1 - 2 + 1 = 0$; for the sphere, $1 - 0 + 1 = 2$. Betti numbers carry more information than $\chi$, because they separate the dimensions that the alternating sum mixes together.

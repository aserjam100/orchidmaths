# The Chromatic Polynomial

**Chromatic polynomial:** For a graph $G$, the function $P(G,k)$ counts the proper colorings of $G$ with $k$ colors. It is a polynomial in $k$, and its value at a positive integer is the number of colorings with that many colors.

## Counting colorings as a polynomial

Color the vertices one at a time. The number of choices at each step depends on how many already-colored neighbours the vertex has, so the count is a product of terms like $k - j$. The result is a polynomial, and the chromatic number is the smallest positive $k$ for which it is nonzero.

## Worked: paths, cycles and complete graphs

1. A path on $n$ vertices has $P(P_n,k) = k(k-1)^{n-1}$: pick the first color freely and each later vertex avoids its predecessor's color.
2. A cycle has $P(C_n,k) = (k-1)^n + (-1)^n(k-1)$.
3. The complete graph has $P(K_n,k) = k(k-1)\cdots(k-n+1)$, the falling factorial.
4. A tree on $n$ vertices has the same polynomial as $P_n$.

For example, $P(P_4,3) = 3 \cdot 2^3 = 24$, and $P(C_4,3) = 2^4 + 2 = 18$.

## Evaluating and factoring

The chromatic polynomial satisfies deletion-contraction: $P(G,k) = P(G-e,k) - P(G/e,k)$ for a non-loop edge $e$, where $G/e$ contracts $e$. It also factors over the blocks of the graph, because colorings of separate blocks combine independently. Its degree is the number of vertices and its leading coefficient is 1.

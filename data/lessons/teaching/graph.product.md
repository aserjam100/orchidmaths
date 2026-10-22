# Graph Products

**Cartesian product:** The Cartesian product $G \square H$ has vertex set $V(G) \times V(H)$, with $(u,v)$ adjacent to $(u',v')$ when $u = u'$ and $vv'$ is an edge of $H$, or $v = v'$ and $uu'$ is an edge of $G$. Its edge count is $|E(G)|\,|V(H)| + |V(G)|\,|E(H)|$.

## Two graphs at once

A vertex of the product is a pair, one coordinate from each factor. An edge changes exactly one coordinate along an edge of that factor. The product therefore combines the two graphs while keeping the coordinates independent, and its edge count is the sum of the edges contributed by changing the first coordinate and by changing the second.

## Worked: the Cartesian edge count

1. Take $G = P_3$ with 3 vertices and 2 edges, and $H = P_3$ with the same.
2. Edges from changing the first coordinate: $|E(G)|\,|V(H)| = 2 \cdot 3 = 6$.
3. Edges from changing the second: $|V(G)|\,|E(H)| = 3 \cdot 2 = 6$.

The product has 12 edges and 9 vertices, which is the $3 \times 3$ grid graph.

## Grids and hypercubes

The product of two paths is a grid; the product of $n$ copies of $K_2$ is the $n$-dimensional hypercube $Q_n$, with $2^n$ vertices and $n2^{n-1}$ edges. Products are a standard way to build larger graphs with controlled structure, and their distances add: $d_{G \square H}((u,v),(u',v')) = d_G(u,u') + d_H(v,v')$.

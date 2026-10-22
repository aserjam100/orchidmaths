# Menger's Theorem

**Menger's theorem:** For non-adjacent vertices $u$ and $v$ in a graph, the maximum number of internally vertex-disjoint $u$-$v$ paths equals the minimum number of vertices whose deletion separates $u$ from $v$. The edge version replaces paths by edge-disjoint paths and vertex cuts by edge cuts.

## Disjoint paths and cuts

The theorem is a min-max duality. A family of disjoint paths is a packing of routes, and a separating set is a cover of all routes. No packing can be larger than any cover, and Menger says the two optima are equal. This is the same shape as König's theorem for matchings and the max-flow min-cut theorem.

## Worked: a small separation

1. Take two triangles sharing a single vertex $x$.
2. Any path from a vertex in the first triangle to a vertex in the second passes through $x$.
3. So the minimum separating set has size 1, and there is exactly one internally disjoint path between such a pair.

The equality holds: the packing and the cut both have size 1.

## Packing and covering

Menger implies the connectivity chain $\kappa(G) \le \lambda(G) \le \delta(G)$ and characterizes $k$-connected graphs as those with $k$ internally disjoint paths between every pair of vertices. The theorem is the combinatorial prototype of the max-flow min-cut theorem, which is its weighted, directed generalization.

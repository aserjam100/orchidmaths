# Matchings

**Matching:** A set of edges no two of which share a vertex. A maximum matching has the largest possible size, and a maximal matching is one that cannot be enlarged by adding another edge. Every maximum matching is maximal, but the converse fails.

## Disjoint edges

A matching pairs vertices without reusing any of them. In a path or a cycle the best is to take alternate edges; in a star every edge shares the centre, so only one edge can be chosen. The matching number $\nu(G)$ is the size of a maximum matching.

## Worked: paths, cycles and stars

1. $P_4$ has a matching of size 2, namely the first and third edges.
2. $P_5$ also has maximum matching size 2; one vertex must be left unmatched.
3. $C_6$ has size 3, alternating around the cycle.
4. $K_{1,5}$ has size 1, because all five edges meet at the centre.

## Maximum and perfect matchings

A perfect matching is a matching that covers every vertex, so it exists only when the vertex count is even. Augmenting paths characterize maximum matchings: a matching is maximum exactly when there is no augmenting path between two unmatched vertices. In bipartite graphs, maximum matching is computed by augmenting paths or by reducing to flow.

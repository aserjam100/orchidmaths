# Vertex Covers

**Vertex cover:** A set of vertices that meets every edge. A minimum vertex cover has the smallest possible size, written $\tau(G)$. The complement of a vertex cover is an independent set, so $\tau(G) + \alpha(G) = |V(G)|$.

## Vertices that meet every edge

A cover chooses vertices and pays for them; every edge must have at least one chosen endpoint. On a path, choosing alternate internal vertices covers all edges. On a star, the centre alone covers everything. Cover and matching are dual: every edge of a matching needs a distinct cover vertex, so $\nu(G) \le \tau(G)$.

## Worked: paths and stars

1. $P_4$ has minimum vertex cover 2, the two internal vertices.
2. $K_{1,5}$ has minimum vertex cover 1, the centre.
3. $K_{2,3}$ has minimum vertex cover 2, the smaller side.
4. $C_6$ has minimum vertex cover 3, alternate vertices.

## König's theorem

In a bipartite graph, the maximum matching size equals the minimum vertex cover size: $\nu(G) = \tau(G)$. This is König's theorem, and it gives a polynomial algorithm for both problems on bipartite graphs. For general graphs equality fails; the odd cycle $C_5$ has matching number 2 and cover number 3.

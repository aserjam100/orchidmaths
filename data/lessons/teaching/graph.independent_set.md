# Independent Sets

**Independent set:** A set of vertices no two of which are adjacent. A maximum independent set has the largest possible size, written $\alpha(G)$. Its complement is a vertex cover, so $\alpha(G) + \tau(G) = |V(G)|$.

## Vertices with no edges between them

An independent set is a set of vertices that can share a color, which is why independent sets and colorings are linked. In a bipartite graph, each side is independent. In a complete graph, only a single vertex is independent. Finding a maximum independent set is NP-hard in general, but polynomial on bipartite graphs through König's theorem.

## Worked: paths and complete bipartite

1. $P_4$ has maximum independent set 2, for example vertices 1 and 3.
2. $C_6$ has size 3, the alternate vertices.
3. The star $K_{1,5}$ has size 5, all the leaves.
4. $K_{2,3}$ has size 3, the larger side.

## Complement of a cover

If $C$ is a vertex cover, then $V \setminus C$ is independent: an edge with both ends outside $C$ would not be covered. Conversely, the complement of an independent set covers every edge. So maximizing an independent set is the same as minimizing a cover, and in bipartite graphs König's theorem turns both into maximum matching.

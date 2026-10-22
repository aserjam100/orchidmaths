# Subgraphs and Induced Subgraphs

**Subgraph:** A graph $H$ is a subgraph of $G$ when $V(H) \subseteq V(G)$ and $E(H) \subseteq E(G)$. A spanning subgraph keeps every vertex. An induced subgraph $G[S]$ keeps a vertex set $S$ and every edge of $G$ whose two endpoints lie in $S$.

## Keeping vertices or edges

There are two different ways to shrink a graph. Deleting edges keeps all the vertices and produces a spanning subgraph; deleting vertices removes the vertices and all their incident edges. An induced subgraph is determined entirely by its vertex set, because no edge inside $S$ may be dropped.

## Worked: induced versus spanning

1. Start with the square $C_4$ on vertices 1, 2, 3 and 4.
2. Spanning subgraph: keep all four vertices and only the edges 12 and 34. The result is two disjoint edges.
3. Induced subgraph on $S = \{1,2,3\}$: keep the edges 12 and 23, because 13 is not an edge of $C_4$. The result is the path $P_3$.

So the same vertex set can support different subgraphs, and "induced" fixes which edges survive.

## Deletion and containment

Deleting a vertex $v$ is written $G - v$, and it removes $v$ and every edge incident to it. Deleting an edge $e$ is written $G - e$ and keeps every vertex. A subgraph is proper when it is not the whole graph. These operations are the building blocks of connectivity: a cut vertex is a vertex whose deletion increases the number of components, and a bridge is an edge whose deletion does the same.

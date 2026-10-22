# Tutte's Theorem

**Tutte's theorem:** A graph $G$ has a perfect matching exactly when, for every vertex set $S$, the graph $G - S$ has at most $|S|$ odd components. The condition is necessary and sufficient, and it applies to every graph, not only bipartite ones.

## The odd-component condition

Deleting $S$ leaves components. Each odd component must send at least one matched edge to $S$, since an odd number of vertices cannot be paired internally. With at most $|S|$ odd components and only $|S|$ vertices in $S$ to absorb them, the count balances exactly.

## Worked: an odd vertex count

1. Take a graph with an odd number of vertices and let $S$ be empty.
2. The whole graph is one component, and it is odd.
3. The condition requires at most 0 odd components, which fails.

So no graph with an odd number of vertices has a perfect matching, as expected.

## Generalizing Hall

For bipartite graphs, Tutte's condition reduces to Hall's marriage theorem. The extra power appears in non-bipartite graphs, where odd cycles can be matched internally and the condition is genuinely about odd components. A bridge in a graph with a perfect matching must itself be in the matching, which forces each side of the bridge to contain an odd number of vertices.

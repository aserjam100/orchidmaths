# Strong Connectivity

**Strongly connected:** A directed graph is strongly connected when every vertex can reach every other vertex by a directed path.

## Reachability both ways

Strong connectivity requires reachability in both directions for every ordered pair of vertices. It is stronger than connectivity of the underlying undirected graph: a directed path whose edges all point one way is connected but not strongly connected. A strongly connected component is a maximal set of vertices that are mutually reachable, and the components partition the vertices of the digraph.

## Worked: the condensation

Contract each strongly connected component to a single vertex and keep the edges that run between components. The result is the condensation of the digraph, and it is always acyclic: a directed cycle among components would make them mutually reachable and merge them. For example, two directed cycles joined by a single one-way edge give a condensation that is a two-vertex path. The condensation is therefore a DAG whose vertices summarise the reachability structure.

## Testing strong connectivity

Linear-time algorithms find the components: Tarjan's algorithm uses one depth-first search and a stack, and Kosaraju's algorithm runs a depth-first search, then another on the reversed graph. Reversing every edge preserves strong connectivity, so a single search from any vertex, plus a search from that vertex in the reversed graph, decides strong connectivity in two passes. The condensation also gives a necessary condition for a directed graph to have a topological order over its components.

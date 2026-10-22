# Topological Sort

**Topological order:** A topological order of a directed graph is a linear ordering of its vertices such that every edge points from an earlier vertex to a later one. Such an order exists exactly when the graph is a DAG.

## Ordering a DAG

A directed cycle has no first vertex, because every vertex on the cycle has a predecessor on the cycle, so a cyclic digraph cannot be ordered. Conversely, every acyclic digraph has a source, a vertex of in-degree zero, and removing it leaves an acyclic digraph, so a topological order exists by induction. Topological orders are the formal version of a dependency order, such as the order in which a build system must compile files.

## Worked: Kahn's algorithm

Kahn's algorithm makes the induction constructive. Start with the set of vertices of in-degree zero and output one of them, then delete it and its outgoing edges, updating the in-degrees of its neighbours. Any neighbour whose in-degree drops to zero joins the ready set, and the process repeats until the graph is empty. If vertices remain but none has in-degree zero, those vertices contain a directed cycle, which certifies that no topological order exists. The algorithm runs in time linear in the number of vertices and edges.

## Uniqueness and choice

A topological order is not unique in general, and any member of the ready set may be taken next, which is why different runs of Kahn's algorithm can produce different valid orders. The order is unique exactly when the DAG has a Hamiltonian directed path, so that at every step exactly one vertex has in-degree zero. In an application this freedom is useful: a scheduler can break ties by priority without ever violating the dependency edges.

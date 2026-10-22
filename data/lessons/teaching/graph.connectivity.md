# Vertex and Edge Connectivity

**Connectivity numbers:** The vertex connectivity $\kappa(G)$ is the minimum number of vertices whose deletion disconnects $G$ or leaves one vertex. The edge connectivity $\lambda(G)$ is the minimum number of edges whose deletion disconnects $G$. A disconnected graph has both numbers equal to 0.

## How hard a graph is to disconnect

These numbers measure redundancy. A large $\kappa$ means many vertex-disjoint routes between any two vertices; a large $\lambda$ means many edge-disjoint routes. For a complete graph $K_n$, $\kappa = \lambda = n-1$, because deleting all but one vertex is necessary to disconnect it.

## Worked: complete, cycle and bipartite

1. $\kappa(K_5) = 4$, since four deletions leave one vertex.
2. $\kappa(C_6) = 2$, since one deletion leaves a connected path and two can split it.
3. $\kappa(P_5) = 1$, since any internal vertex is a cut vertex.
4. $\kappa(K_{3,3}) = 3$, because deleting one side disconnects the graph and no smaller set does.

## The inequality chain

Menger's theorem gives $\kappa(G) \le \lambda(G) \le \delta(G)$, where $\delta$ is the minimum degree. A vertex cut can be built from the edges at a vertex, and an edge cut from the edges at a minimum-degree vertex. The chain is often strict: a graph can have minimum degree 4 but vertex connectivity 2.

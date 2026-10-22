# Planar graphs and Euler's formula

A graph is **planar** when it can be drawn in the plane with no edges crossing. For a connected planar graph, the number of vertices, edges and faces satisfies **Euler's formula** \(V - E + F = 2\), where the outer region counts as one face.

## Euler's formula

The formula is proved by building the graph one edge at a time. A tree has \(F = 1\), and \(V - E = 1\) for a tree, so the formula holds; adding an edge that closes a cycle adds exactly one face and changes \(V - E\) by minus one, and the two changes cancel.

Reading it as a formula for faces is the usual use: \(F = E - V + 2\). A connected planar graph with eight vertices and twelve edges has \(12 - 8 + 2 = 6\) faces, counting the outer one.

## The edge bound

A simple planar graph with \(V \ge 3\) has at most \(3V - 6\) edges. The reason is counting: every face is bounded by at least three edges, and every edge borders two faces, so \(2E \ge 3F\); substituting Euler's formula gives the bound.

The bound is sharp: the tetrahedron has \(V = 4\) and \(6 = 3 \cdot 4 - 6\) edges. Being over the bound rules planarity out immediately, which is the quickest disproof for a dense graph.

## When planarity fails

\(K_5\) has five vertices and ten edges, but \(3 \cdot 5 - 6 = 9\), so it is not planar. \(K_{3,3}\) has six vertices and nine edges, which is under the bound, so the edge count does not settle it — it fails instead because every face would need at least four edges, giving \(2E \ge 4F\) and hence \(E \le 2V - 4 = 8 < 9\).

**Kuratowski's theorem** completes the picture: a graph is planar exactly when it contains no subdivision of \(K_5\) or \(K_{3,3}\). The two forbidden graphs are the whole obstruction, and the edge bounds are the cheap tests that catch most cases first.

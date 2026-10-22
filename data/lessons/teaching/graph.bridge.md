# Bridges

**Bridge:** An edge $e$ is a bridge of a connected graph when deleting $e$ disconnects the graph. Equivalently, $e$ lies on no cycle, because an edge on a cycle has an alternate route between its endpoints.

## Edges whose removal disconnects

Bridges are the edge analogue of cut vertices. Every edge of a tree is a bridge; a cycle has none. A connected graph with no bridge is 2-edge-connected, and its maximal 2-edge-connected pieces are its 2-edge-connected components, joined by bridges.

## Worked: trees and cycles

1. The path $P_4$ has three edges and every one is a bridge.
2. The cycle $C_5$ has five edges and none is a bridge.
3. The star $K_{1,4}$ has four bridges, one at each leaf.
4. A tree on six vertices has five edges and all five are bridges.

A square with one diagonal has no bridge: every edge lies on a triangle or on the square.

## Bridges and cut vertices

A bridge's endpoints need not be cut vertices: in a path of two edges, the middle vertex is a cut vertex and both edges are bridges, but in a single edge both endpoints have degree 1 and the edge is a bridge without either endpoint being a cut vertex of a larger graph. A bridge always produces a separation of the graph into two sides, which is why bridges appear in matching conditions such as Tutte's theorem.

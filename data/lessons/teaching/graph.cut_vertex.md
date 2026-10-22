# Cut Vertices

**Cut vertex:** A vertex $v$ is a cut vertex of a connected graph when deleting $v$ disconnects the remaining graph. Equivalently, some pair of vertices other than $v$ has every path between them passing through $v$.

## Vertices whose removal disconnects

Cut vertices are the bottlenecks of a connected graph. A graph with no cut vertex is 2-connected, and every connected graph decomposes into maximal 2-connected pieces, its blocks, joined at cut vertices. Finding cut vertices is a depth-first search that tracks discovery and low values.

## Worked: paths and stars

1. The path $P_4$ has two internal vertices, and removing either splits the path, so it has 2 cut vertices.
2. The path $P_5$ has three internal vertices, hence 3 cut vertices.
3. The star $K_{1,4}$ has exactly one cut vertex, its centre; removing a leaf leaves the rest connected.

A cycle has no cut vertex, because removing one vertex leaves a path, which is still connected.

## Blocks and 2-connectivity

A block is a maximal subgraph with no cut vertex; every edge lies in exactly one block. A connected graph is 2-connected exactly when it has no cut vertex, which by Menger is the same as every pair of vertices having two internally disjoint paths. The block-cut tree records how the blocks meet at cut vertices.

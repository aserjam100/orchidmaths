# Graph Minors

**Minor:** A graph $H$ is a minor of $G$ when $H$ can be obtained from $G$ by deleting vertices, deleting edges and contracting edges. Contraction identifies the two ends of an edge and removes the loop it creates.

## The three operations

A minor is built with only three moves: delete a vertex, delete an edge, or contract an edge. Each move can only shrink the graph, so a minor of a minor is a minor, and the relation is a partial order up to isomorphism. Subdivision is the opposite operation and is not allowed: a graph is generally not a minor of its own subdivision, because contraction can undo a subdivision but deletion cannot create the missing vertices.

## Worked: a K_5 minor from a larger clique

The complete graph $K_6$ contains $K_5$ as a subgraph: delete one vertex and every edge incident with it, and five mutually adjacent vertices remain. Deleting a vertex is a minor operation, so $K_5$ is a minor of $K_6$. In the other direction $K_5$ can appear without appearing as a subgraph: any graph that contains a subdivision of $K_5$ also has $K_5$ as a minor, because contracting the edges of each subdivided path identifies its endpoints. Minors detect nonplanarity that no subgraph may witness.

## Minor-closed families

A class of graphs is minor-closed when every minor of a member is again a member. Planar graphs are minor-closed: deleting and contracting edges inside a planar drawing cannot create a crossing. The Robertson-Seymour theorem says that every minor-closed class is characterised by a finite set of forbidden minors, which is why planarity has the two forbidden minors $K_5$ and $K_{3,3}$. Hadwiger's conjecture ties the chromatic number to clique minors: a graph with chromatic number $k$ is conjectured to have a $K_k$ minor.

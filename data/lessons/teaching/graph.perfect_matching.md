# Perfect Matchings

**Perfect matching:** A matching that covers every vertex exactly once. It pairs all the vertices into disjoint edges, so the graph must have an even number of vertices. A graph with a perfect matching has one when the vertices can be paired along existing edges.

## Covering every vertex

Perfect matchings are the strongest form of pairing. Every regular bipartite graph has one, every complete graph with an even number of vertices has one, and every cycle with an even number of vertices has one. An odd path or an odd cycle cannot, because one vertex would remain uncovered.

## Worked: even versus odd

1. $P_4$ has the perfect matching consisting of edges 1-2 and 3-4.
2. $P_5$ has an odd number of vertices, so no perfect matching exists; its maximum matching leaves one vertex out.
3. $C_6$ has the alternate-edge perfect matching.

## Hall and regular bipartite graphs

Hall's marriage theorem characterizes perfect matchings in bipartite graphs: every set of vertices on one side must have at least as many neighbours on the other. Regular bipartite graphs satisfy Hall's condition, so they have perfect matchings. Tutte's theorem removes the bipartite restriction and characterizes perfect matchings in all graphs.

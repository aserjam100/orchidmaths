# Perfect Graphs

**Perfect graph:** A graph $G$ is perfect when $\omega(H) = \chi(H)$ for every induced subgraph $H$, where $\omega$ is the clique number and $\chi$ the chromatic number. The clique is an obvious lower bound on the chromatic number; perfection says the bound is always tight.

## Clique number equals chromatic number

In a perfect graph, the largest clique explains exactly how many colors are needed, on the graph and on every induced subgraph. Complete graphs, bipartite graphs, interval graphs and chordal graphs are perfect. Odd cycles of length at least 5 are the smallest imperfections: their clique number is 2 but they need 3 colors.

## Worked: bipartite and odd cycles

1. A bipartite graph has clique number at most 2 and chromatic number at most 2, and every induced subgraph is bipartite.
2. The cycle $C_5$ has $\omega = 2$ and $\chi = 3$, so it is not perfect.
3. The complement of $C_5$ is also $C_5$, so it fails the same way.

## The strong perfect graph theorem

The strong perfect graph theorem (Chudnovsky, Robertson, Seymour and Thomas) characterizes perfect graphs: a graph is perfect exactly when neither it nor its complement contains an induced odd cycle of length at least 5. Such a cycle is an odd hole, and its complement is an odd antihole. The theorem settled a conjecture from 1960.

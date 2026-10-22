# Bipartite graphs and matchings

A graph is **bipartite** when its vertices split into two parts with every edge running between them. Equivalently, it is bipartite exactly when it is 2-colourable: colour one part black and the other white, and every edge joins different colours.

## Two-colourable graphs

The equivalence is operational. Colour a vertex, colour its neighbours the other colour, and continue; a conflict means an odd cycle, and a graph is bipartite precisely when it has no odd cycle. So bipartiteness is testable by breadth-first colouring, and a single odd cycle is the obstruction.

Complete bipartite graphs \(K_{m,n}\) have all \(mn\) edges between the parts and none inside. Their degrees are read off the parts: each of the \(m\) vertices has degree \(n\), and each of the \(n\) has degree \(m\).

## Matchings

A **matching** is a set of edges with no shared endpoint. In a bipartite graph the question is how many pairs can be formed at once, and the maximum matching size is the answer to assignment problems: workers to jobs, machines to tasks.

For a small graph the maximum is found by trying assignments: left vertices \(a, b, c\) with edges \(a1, a2, b2, c3\) admit the matching \(\{a1, b2, c3\}\), of size three, and no matching can be larger because there are only three left vertices.

## Hall's condition

**Hall's theorem** says a matching saturating the left part exists exactly when every set \(S\) of left vertices has at least \(\lvert S \rvert\) neighbours. The condition is necessary because those neighbours must supply the partners, and it is sufficient — which is the hard direction and the content of the theorem.

Hall's condition is also a way to *prove* a maximum is not larger than it is: exhibit a set \(S\) whose neighbourhood is smaller than \(S\), and no matching can saturate \(S\). That is the certificate a learner can check without searching.

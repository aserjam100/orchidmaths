# The Complement of a Graph

**Complement:** For a simple graph $G$ on $n$ vertices, the complement $\overline{G}$ has the same vertex set and exactly the edges that $G$ does not have. Its edge count is $\binom{n}{2} - |E(G)|$.

## Edges that are not there

The complement records absence rather than presence. A pair of vertices is adjacent in $\overline{G}$ exactly when it is non-adjacent in $G$. Taking the complement twice returns the original graph. The construction turns questions about independent sets into questions about cliques, because an independent set in $G$ is a clique in $\overline{G}$.

## Worked: the complement of P_4

1. The path $P_4$ has 4 vertices and 3 edges.
2. A graph on 4 vertices has $\binom{4}{2} = 6$ possible edges.
3. So $\overline{P_4}$ has $6 - 3 = 3$ edges: the two diagonals and the edge joining the endpoints.

In general the complement of $P_4$ is again a path, so $P_4$ is self-complementary.

## Self-complementary graphs

A graph is self-complementary when it is isomorphic to its complement, which forces it to have exactly half of the possible edges. $P_4$ is the smallest example. Self-complementary graphs exist only when $n \equiv 0$ or $1 \pmod 4$, because $\binom{n}{2}$ must be even. The complement is a cheap source of invariants: a graph and its complement share the same automorphism group.

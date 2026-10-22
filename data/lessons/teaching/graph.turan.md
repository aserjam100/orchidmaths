# Turán's Theorem

**Turán's theorem:** Among all graphs on $n$ vertices that contain no $K_{r+1}$, the maximum number of edges is attained by the complete balanced $r$-partite graph $T_r(n)$.

## The Turán graph

The Turán graph $T_r(n)$ has its $n$ vertices split into $r$ parts that are as equal as possible, with every pair in different parts joined. It contains no $K_{r+1}$, because a clique can use at most one vertex from each part. Its edge count is $\lfloor (1 - 1/r) n^2 / 2 \rfloor$. When $r = 2$ this is Mantel's theorem: the triangle-free extremal graph is the balanced complete bipartite graph.

## Worked: K_4-free on six vertices

Take $r = 3$ and $n = 6$. The three parts each have size two, and every pair in different parts is joined, giving $3 \cdot (2 \cdot 2) = 12$ edges. The graph has no $K_4$, because four mutually adjacent vertices would need four different parts and only three exist. The bound formula agrees: $\lfloor (1 - 1/3) \cdot 36 / 2 \rfloor = \lfloor 12 \rfloor = 12$.

## The general bound and beyond

Turán's theorem is the first extremal result: it says that forbidding one small subgraph caps the edge count and identifies the extremal graph. The formula $\lfloor (1 - 1/r) n^2 / 2 \rfloor$ grows quadratically in $n$ for fixed $r$. Forbidding a non-bipartite graph gives a similar quadratic cap, while forbidding a bipartite graph allows far more edges, a contrast made precise by the Erdős-Stone theorem.

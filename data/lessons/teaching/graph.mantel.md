# Mantel's Theorem

**Mantel's theorem:** A triangle-free simple graph on $n$ vertices has at most $\lfloor n^2/4 \rfloor$ edges, and the bound is attained by a complete bipartite graph with parts as equal as possible.

## Triangle-free graphs

A graph is triangle-free when it contains no $K_3$. Every bipartite graph is triangle-free, because a triangle would need an odd cycle, but not every triangle-free graph is bipartite: the cycle $C_5$ is the smallest counterexample. Mantel's theorem says that among all triangle-free graphs on $n$ vertices, the bipartite ones already achieve the maximum possible edge count.

## Worked: five vertices

With $n = 5$ the bound is $\lfloor 25/4 \rfloor = 6$. The complete bipartite graph $K_{2,3}$ has parts of sizes two and three, is triangle-free, and has $2 \cdot 3 = 6$ edges. So the bound is tight here. Any graph on five vertices with seven edges must contain a triangle, because the average degree would be $14/5 = 2.8$, and a standard argument shows that a triangle-free graph with this density cannot exist.

## The extremal example and a proof sketch

The extremal graph is the complete bipartite graph whose two parts differ in size by at most one, so both parts have size about $n/2$ and the edge count is about $n^2/4$. One proof fixes a vertex $v$ and observes that its neighbours form an independent set, since any edge among them with $v$ would make a triangle. Summing over vertices and bounding the degrees gives the $n^2/4$ bound, which is exactly the edge count of the balanced complete bipartite graph.

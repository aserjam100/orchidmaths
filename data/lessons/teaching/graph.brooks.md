# Brooks' Theorem

**Brooks' theorem:** If a connected graph $G$ has maximum degree $\Delta$, then its chromatic number is at most $\Delta$, unless $G$ is a complete graph or an odd cycle, where it is $\Delta + 1$.

## One color below greedy

The greedy algorithm colors a graph with at most $\Delta + 1$ colors: order the vertices and give each the smallest color not used by its earlier neighbours. Brooks says the extra color is almost never needed. The two exceptional families are exactly the graphs in which every vertex has $\Delta$ neighbours that already use all $\Delta$ colors.

## Worked: the exceptions

1. The complete graph $K_4$ has $\Delta = 3$ but needs 4 colors, so it is an exception.
2. The odd cycle $C_5$ has $\Delta = 2$ but needs 3 colors, so it is an exception.
3. A connected graph with $\Delta = 3$ that is not $K_4$ is 3-colorable.

Even cycles are not exceptions: $C_4$ has $\Delta = 2$ and is 2-colorable.

## Why complete graphs and odd cycles

In a complete graph every vertex sees all the others, so no color can be shared. In an odd cycle the alternation of two colors closes with a conflict. Brooks' proof removes a vertex, colors the rest greedily, and uses a Kempe-chain argument to free a color for the removed vertex unless the graph has one of those two structures.

# Edge Coloring

**Edge coloring:** An assignment of colors to edges so that edges sharing a vertex get different colors. The edge chromatic number $\chi'(G)$ is the smallest number of colors needed. It satisfies $\Delta \le \chi' \le \Delta + 1$, where $\Delta$ is the maximum degree.

## Coloring edges, not vertices

Edge coloring is vertex coloring of the line graph, but the structure gives a much tighter bound. Every edge at a vertex needs a different color, so $\chi' \ge \Delta$. Vizing's theorem says one extra color always suffices, which splits graphs into class 1 ($\chi' = \Delta$) and class 2 ($\chi' = \Delta + 1$).

## Worked: complete graphs and cycles

1. $K_4$ has $\Delta = 3$ and $\chi' = 3$: a round-robin schedule of four players in three rounds.
2. $K_5$ has $\Delta = 4$ but $\chi' = 5$, the classic class-2 example.
3. An even cycle has $\chi' = 2$; an odd cycle has $\chi' = 3$.
4. The star $K_{1,4}$ has $\chi' = 4$, because all four edges meet at the centre.

## Vizing's theorem

Vizing proved $\chi' \le \Delta + 1$ for every simple graph, and both values occur. Deciding whether a graph is class 1 is NP-hard in general. For bipartite graphs König's line-coloring theorem gives the exact value $\chi' = \Delta$, which is why timetabling and scheduling problems on bipartite graphs are tractable.

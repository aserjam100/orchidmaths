# List Coloring

**List coloring:** Each vertex receives a list of allowed colors, and a list coloring chooses one color from each list so that adjacent vertices differ. A graph is $k$-choosable when every assignment of lists of size $k$ admits a list coloring. The list chromatic number $\chi_\ell(G)$ is the smallest such $k$.

## Colors chosen from lists

List coloring is a generalization of ordinary coloring: giving every vertex the same list of $k$ colors recovers a $k$-coloring. The lists can be different, which makes the problem harder. A graph that is $k$-colorable need not be $k$-choosable, so $\chi_\ell(G) \ge \chi(G)$.

## Worked: even and odd cycles

1. An even cycle is 2-choosable: color from one vertex around, avoiding the previous color, and the even length closes consistently.
2. An odd cycle is not 2-choosable; three colors are needed in the worst case.
3. A path is 2-choosable: color from one end, and each next vertex has at most one forbidden color.
4. Complete bipartite graphs can require more colors than their chromatic number, so choosability is genuinely stronger.

## Choosability versus chromatic number

The gap between $\chi$ and $\chi_\ell$ can be arbitrarily large: bipartite graphs can have arbitrarily high list chromatic number. The list version of Brooks' theorem says $\chi_\ell \le \Delta + 1$ for connected graphs that are not complete or odd cycles, matching the ordinary bound. List coloring is the right model when each item has its own set of allowed slots, as in channel assignment.

# The Moore Bound

**Moore bound:** A graph with maximum degree $d$ and diameter two has at most $d^2 + 1$ vertices.

## Degree and diameter

Fix a vertex $v$. It contributes one vertex at distance zero, at most $d$ neighbours at distance one, and at most $d(d-1)$ vertices at distance two, because each neighbour has at most $d-1$ further neighbours. A graph of diameter two contains every vertex within that reach, so the total is at most $1 + d + d(d-1) = d^2 + 1$. The argument extends to any diameter $k$ and gives the Moore bound for the degree-diameter problem.

## Worked: the Petersen graph

With $d = 3$ the bound is $3^2 + 1 = 10$, and the Petersen graph attains it: it is 3-regular, has diameter two, and has exactly ten vertices. The cycle $C_5$ attains the bound for $d = 2$, since $2^2 + 1 = 5$. When the bound is met every vertex has $d$ neighbours and no two neighbours of $v$ share another neighbour, so the graph is as efficient as a network can be.

## Moore graphs and the degree-diameter problem

A Moore graph meets the bound exactly. The known Moore graphs are the odd cycles for $d = 2$, the Petersen graph for $d = 3$, the Hoffman-Singleton graph for $d = 7$, and the complete graph for $d = 2$ in diameter one. The degree-diameter problem asks for the largest graph with given degree and diameter, and the Moore bound is the first upper bound against which every construction is measured.

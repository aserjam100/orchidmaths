# Graph Genus

**Genus:** The genus of a graph is the smallest number of handles a surface must have for the graph to embed in it without crossings. The sphere has genus zero, the torus genus one, and each handle adds one.

## Embedding surfaces

An embedding draws the vertices as points and the edges as curves that meet only at shared endpoints. A planar graph embeds in the sphere, so its genus is zero. A graph that cannot be drawn in the plane without crossings may embed in the torus, which is the sphere with one handle, and the genus counts the smallest number of handles that suffices. The genus is a measure of how far the graph is from planar.

## Worked: K_5 on the torus

$K_5$ is nonplanar, so its genus is at least one. It embeds in the torus, which has Euler characteristic $2 - 2g = 0$. With $n = 5$ vertices and $m = 10$ edges, Euler's formula on a surface, $n - m + f = 2 - 2g$, gives the number of faces as $f = 2 - 2g - n + m = 0 - 5 + 10 = 5$. Five faces is consistent with a torus embedding, so the genus of $K_5$ is one. The same surface carries $K_6$ and $K_7$.

## Genus formulas

The complete graph $K_n$ has genus $\lceil (n-3)(n-4)/12 \rceil$ for $n \ge 3$, which is one for $n = 5, 6, 7$ and two for $n = 8$. The complete bipartite graph $K_{a,b}$ has genus $\lceil (a-2)(b-2)/4 \rceil$, which is one for $K_{3,3}$. Planar graphs and cycles have genus zero. On a surface of genus $g$ the bound on the edges of a simple graph is $m \le 3(n + 2g - 2)$.

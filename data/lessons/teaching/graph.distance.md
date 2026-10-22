# Distance and Diameter

**Distance:** The distance $d(u,v)$ between two vertices is the length of a shortest path joining them, or infinity when they lie in different components. The diameter of a connected graph is the largest distance over all pairs, and the radius is the smallest eccentricity.

## Shortest paths between vertices

Distance counts edges, not weighted lengths. In an unweighted graph it is computed by breadth-first search from a source, which visits vertices in nondecreasing distance. The eccentricity of a vertex is its largest distance to any other vertex; the diameter is the largest eccentricity, and the radius is the smallest.

## Worked: paths, cycles and grids

1. In the path $P_5$ the endpoints are four edges apart, so the diameter is 4.
2. In the cycle $C_6$ the farthest pair is three steps apart, so the diameter is 3.
3. In the complete graph $K_5$ every pair is adjacent, so the diameter is 1.

For an $r \times c$ grid, opposite corners are $(r-1) + (c-1)$ steps apart, which is the diameter.

## Diameter and radius

Every connected graph satisfies $\mathrm{rad}(G) \le \mathrm{diam}(G) \le 2\,\mathrm{rad}(G)$. A graph with diameter 1 is complete; a graph with diameter 2 is common in networks, where a small diameter means short routing. Removing a vertex can increase the diameter, which is one reason connectivity and distance are studied together.

# Outerplanar Graphs

**Outerplanar:** A graph is outerplanar when it can be drawn in the plane so that every vertex lies on the outer face. Every outerplanar graph is planar, but the drawing is more restrictive.

## Drawing on the outer face

Put the vertices on a circle in order and draw the edges inside without crossings. Maximal outerplanar graphs are exactly the triangulations of a polygon: every bounded face is a triangle and the outer face contains all the vertices. Trees are outerplanar with the trivial drawing, and so is every cycle. A graph that needs an interior vertex, such as $K_4$, is not outerplanar.

## Worked: the edge bound

A maximal outerplanar graph on $n$ vertices is a triangulation of an $n$-gon. The polygon has $n$ sides, and a triangulation of an $n$-gon uses $n - 3$ non-crossing diagonals, for $n + (n-3) = 2n - 3$ edges. No outerplanar graph has more, because adding an edge inside a maximal one would cross an existing edge or leave the outer face. For $n = 5$ the bound is $2 \cdot 5 - 3 = 7$ edges.

## Forbidden minors and treewidth

Outerplanar graphs are characterised by forbidden minors: a graph is outerplanar exactly when it has no $K_4$ minor and no $K_{2,3}$ minor. Since outerplanar graphs are a subclass of planar graphs with a single outer face, they have treewidth at most two. The maximal outerplanar graphs are the graphs of treewidth two that are planar and have a drawing with all vertices outside, which is why they are also called triangulated polygons.

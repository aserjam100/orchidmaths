# Kuratowski's Theorem

**Kuratowski's theorem:** A finite graph is planar if and only if it contains no subdivision of $K_5$ or $K_{3,3}$. Those two graphs are the complete graph on five vertices and the complete bipartite graph with three vertices on each side.

## Subdivisions

A subdivision of a graph replaces an edge by a path whose internal vertices are new and have degree two. Subdivision neither creates nor destroys a planar embedding: stretching an edge cannot make a drawing cross, and shortening a path back to an edge cannot remove a crossing. Kuratowski's theorem therefore says that the two graphs are the only obstructions up to subdivision, and that every nonplanar graph hides one of them.

## Worked: why K_{3,3} is not planar

Suppose $K_{3,3}$ were planar. It has $n = 6$ vertices and $m = 9$ edges. A planar embedding with $n$ vertices and $m$ edges satisfies $m \le 3n - 6$ when $n \ge 3$, which gives $9 \le 12$, so the crude bound does not settle it. The sharper bipartite bound does: a triangle-free planar graph has every face of length at least four, so $2m \ge 4f$. With Euler's formula $n - m + f = 2$ this gives $m \le 2n - 4 = 8$, and $9 > 8$ is the contradiction. So $K_{3,3}$ is nonplanar.

## Forbidden subdivisions and minors

The theorem is often stated in terms of minors instead of subdivisions. Wagner's theorem says that a graph is planar exactly when it has neither $K_5$ nor $K_{3,3}$ as a minor. The two statements are equivalent because a graph contains a subdivision of $H$ exactly when it has an $H$ minor whose branch vertices have degree at least three, and for $K_5$ and $K_{3,3}$ the two conditions coincide.

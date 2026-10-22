# Bases and Subbases

**Basis:** A basis for a topology is a family of open sets such that every open set is a union of members of the family, and for any point in the intersection of two basis elements there is a basis element containing the point and staying inside the intersection.

## Generating a topology from smaller pieces

A basis is a set of building blocks: the topology is all unions of those blocks. The second condition makes intersections behave, which is what ensures the unions are closed under finite intersection. In a metric space the open balls form a basis, and the topology they generate is the metric topology.

## Worked: open intervals generate the line

1. Take the family of all open intervals $(a, b)$ in $\mathbb{R}$.
2. Any open set is a union of open intervals around its points, so it is a union of basis elements.
3. Given a point in $(a,b) \cap (c,d)$, the intersection is itself an open interval containing the point, so the basis condition holds.

A subbasis is weaker: it is any family whose finite intersections form a basis. The open rays $(-\infty, b)$ and $(a, \infty)$ form a subbasis for the usual topology on $\mathbb{R}$, because their intersections are the open intervals.

## Basis versus subbasis

A basis is closed under the operation of taking a smaller basis element inside an intersection; a subbasis need not be. Starting from a subbasis, take all finite intersections to get a basis, then all unions to get the topology. Starting from a basis, only the unions are needed.

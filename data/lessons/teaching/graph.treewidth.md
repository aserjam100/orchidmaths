# Treewidth

**Treewidth:** The treewidth of a graph $G$ is the smallest width of a tree decomposition of $G$, where the width is one less than the largest bag. It measures how close $G$ is to a tree.

## Tree decompositions

A tree decomposition is a tree whose nodes are bags, each a set of vertices of $G$. Every vertex of $G$ lies in a connected subtree of bags, and every edge of $G$ lies in some bag. The width is the size of the largest bag minus one, and the treewidth is the smallest width over all decompositions. A tree has a decomposition whose bags are its edges, of width one, and every graph with at least one edge has treewidth at least one.

## Worked: the cycle C_6

Take the six edges of $C_6$ as the six bags and join consecutive bags in a path. Each bag holds two adjacent vertices, so the largest bag has size two and the width is one. This is a valid decomposition, but the width is not yet optimal: no decomposition of a cycle has width one, because a graph of treewidth one is a forest and $C_6$ is not a forest. So the treewidth of $C_6$ is two. In general $C_n$ has treewidth two for $n \ge 3$.

## Bounded treewidth and algorithms

Many problems that are hard on an arbitrary graph become easy on a graph of bounded treewidth by dynamic programming over the decomposition: graph coloring, independent set and Hamiltonian path are all decidable in time linear in $n$ for fixed width. Complete graphs are the other extreme, since $K_n$ needs one bag holding all $n$ vertices and has treewidth $n - 1$. An $r$-by-$c$ grid has treewidth $\min(r,c)$, so large grids have large treewidth even though every vertex has degree at most four.

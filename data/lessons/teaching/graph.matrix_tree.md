# The Matrix-Tree Theorem

**Matrix-tree theorem:** The number of spanning trees of a connected graph is the determinant of any cofactor of its Laplacian: delete one row and the same column of $L$, and the determinant is the tree count.

## Cofactors and trees

The determinant is independent of which row and column are deleted, because every cofactor of $L$ equals the same product of nonzero Laplacian eigenvalues divided by $n$. For a connected graph the result is a positive integer, the number of distinct spanning trees. This turns a counting problem into a determinant and is the standard tool for tree counts in structured graphs.

## Worked: K_4

For $K_4$ the Laplacian is $4I - J$, where $J$ is the all-ones matrix. Deleting one row and column leaves a $3 \times 3$ matrix with diagonal 4 and off-diagonal $-1$, whose determinant is 16. Cayley's formula confirms it: $n^{n-2} = 4^2 = 16$. Each spanning tree of $K_4$ is a labelled tree on four vertices, and there are sixteen of them.

## Cayley's formula and examples

Cayley's formula is the special case of the complete graph: $K_n$ has $n^{n-2}$ spanning trees, which is 3 for $K_3$ and 125 for $K_5$. A cycle $C_n$ has $n$ spanning trees, one for each edge removed, and a path has exactly one because it is already a tree. The complete bipartite graph $K_{a,b}$ has $a^{b-1} b^{a-1}$ spanning trees, which is $2^2 \cdot 3 = 12$ for $K_{2,3}$.

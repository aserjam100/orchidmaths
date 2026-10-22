# The Laplacian Matrix

**Laplacian:** For a graph with degree matrix $D$ and adjacency matrix $A$, the Laplacian is $L = D - A$. Its diagonal entry $L_{ii}$ is the degree of vertex $i$.

## Degrees on the diagonal

The Laplacian has the vertex degrees on the diagonal, $-1$ in entry $(i,j)$ when $i$ and $j$ are adjacent, and 0 otherwise. Every row and column sums to zero, so the all-ones vector is an eigenvector with eigenvalue zero. The matrix is symmetric and positive semidefinite, because the quadratic form $x^{\top} L x$ equals the sum of $(x_i - x_j)^2$ over all edges. The name comes from the discrete analogue of the continuous Laplacian.

## Worked: zero eigenvalues

The multiplicity of the eigenvalue zero equals the number of connected components of the graph. If the graph is connected, only the all-ones vector lies in the kernel, so zero has multiplicity one. Two disjoint edges have two components, and the kernel is spanned by one indicator vector per component, so zero has multiplicity two and the Laplacian eigenvalues are 0, 0, 1, 1. Counting zero eigenvalues is therefore a matrix test for connectivity.

## Using the Laplacian

Writing the quadratic form as a sum over edges makes the Laplacian a bridge between linear algebra and graph structure. The matrix-tree theorem reads the number of spanning trees from a cofactor of $L$. Its second-smallest eigenvalue measures how well connected the graph is and controls the speed of random walks and the quality of spectral partitions, which is why the Laplacian appears in clustering and in the analysis of electrical networks.

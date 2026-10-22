# The Spectral Radius

**Spectral radius:** The spectral radius of a graph is the largest absolute value of an eigenvalue of its adjacency matrix. For a connected graph this is the largest eigenvalue itself.

## The largest eigenvalue

The spectral radius $\rho$ is at least the average degree and at most the maximum degree, by the Perron-Frobenius theorem for nonnegative matrices. When the graph is connected, the eigenvector for $\rho$ can be chosen with all entries positive, and $\rho$ is a simple eigenvalue. The spectral radius therefore measures the long-run growth of walks: the number of walks of length $k$ grows like $\rho^k$.

## Worked: regular graphs

A $d$-regular graph has the all-ones vector as an eigenvector with eigenvalue $d$, and every other eigenvalue is at most $d$ in absolute value, so the spectral radius is exactly $d$. This settles many families at once: $K_n$ is $(n-1)$-regular with spectral radius $n - 1$, every cycle $C_n$ is 2-regular with spectral radius two, and the Petersen graph is 3-regular with spectral radius three. The complete bipartite graph $K_{4,4}$ is 4-regular, so its spectral radius is four.

## Perron-Frobenius and bounds

Equality in $\rho \ge$ average degree holds precisely for a regular component, and $\rho = $ maximum degree holds precisely when the component with the maximum degree is regular. The spectral radius controls mixing and expansion: a small gap between $\rho$ and the second eigenvalue means a well-mixed graph, and this is the adjacency-matrix counterpart of the Laplacian's algebraic connectivity.

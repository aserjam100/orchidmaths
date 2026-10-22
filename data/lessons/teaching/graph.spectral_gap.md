# Algebraic Connectivity

**Algebraic connectivity:** The algebraic connectivity of a connected graph is the second-smallest eigenvalue $\lambda_2$ of its Laplacian. The smallest eigenvalue is always zero.

## The second-smallest eigenvalue

Because the Laplacian is positive semidefinite, its eigenvalues can be ordered $0 = \lambda_1 \le \lambda_2 \le \cdots \le \lambda_n$. The graph is connected exactly when $\lambda_2 > 0$, and the size of $\lambda_2$ measures how hard it is to cut the graph into two large pieces. A small $\lambda_2$ means a bottleneck: two dense regions joined by few edges.

## Worked: the cycle C_6

For the cycle $C_n$ the Laplacian eigenvalues are $2 - 2\cos(2\pi k/n)$ for $k = 0, \dots, n-1$. The second one, at $k = 1$, is $2 - 2\cos(2\pi/n)$. At $n = 6$ this is $2 - 2\cos(\pi/3) = 2 - 1 = 1$, and at $n = 4$ it is $2 - 2\cos(\pi/2) = 2$. A long cycle has small algebraic connectivity, because it can be cut in half by removing two edges.

## Connectivity and expansion

Several families have known values: the complete graph $K_n$ has $\lambda_2 = n$, the complete bipartite graph $K_{a,b}$ has $\lambda_2 = \min(a,b)$, and a star on at least two vertices has $\lambda_2 = 1$. The vector for $\lambda_2$, called the Fiedler vector, gives a near-optimal partition of the vertices into two parts with few crossing edges, which is the basis of spectral clustering.

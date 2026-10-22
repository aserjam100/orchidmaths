# The Adjacency Matrix

**Adjacency matrix:** For a simple graph on $n$ vertices, the adjacency matrix $A$ is the $n \times n$ matrix with $A_{ij} = 1$ when $i$ and $j$ are adjacent and $A_{ij} = 0$ otherwise.

## From graph to matrix

The matrix is symmetric with zero diagonal, and the sum of row $i$ is the degree of vertex $i$. A loopless undirected graph and its adjacency matrix carry the same information, so every graph property can be read from the matrix. For a multigraph the entries count edges, and for a weighted graph they hold the weights. Powers of $A$ count walks.

## Worked: walks of length two

The $(i,j)$ entry of $A^2$ counts the walks of length two from $i$ to $j$. In the cycle $C_4$ with vertices 1, 2, 3, 4 in order, $(A^2)_{11} = 2$, from the walks $1-2-1$ and $1-4-1$, and $(A^2)_{13} = 2$, from $1-2-3$ and $1-4-3$. The diagonal entry $(A^2)_{ii}$ equals the degree of $i$, because each neighbour gives one closed walk of length two.

## What the spectrum sees

The eigenvalues of $A$ are the spectrum of the graph, and they encode structure. The number of closed walks of length $k$ is the sum of the $k$-th powers of the eigenvalues. A graph is bipartite exactly when its spectrum is symmetric about zero, so a connected graph is bipartite when its smallest eigenvalue is $-\rho$. Cospectral graphs share the spectrum but may not be isomorphic, which is why the spectrum alone does not determine a graph.

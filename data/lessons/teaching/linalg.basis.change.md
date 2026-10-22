# Coordinates and Change of Basis

**Coordinates:** Once a basis $B = (b_1, \dots, b_n)$ of $V$ is fixed, every vector has a unique list of coefficients: $v = c_1 b_1 + \cdots + c_n b_n$. That list is the coordinate vector $[v]_B$, and it depends on the basis even though $v$ does not.

## Coordinates depend on the basis

In the standard basis of $\mathbb{R}^2$, the vector $(3,5)$ is simply $(3,5)$. In the basis $(1,1)$ and $(1,-1)$ the same vector is $4(1,1) - 1(1,-1)$, so its coordinates are $(4,-1)$. The vector has not moved; the description has.

## Worked: solving for coordinates

1. Write $(3,5) = a(1,1) + b(1,-1)$.
2. This gives the system $a + b = 3$ and $a - b = 5$.
3. Solving gives $a = 4$ and $b = -1$, so the coordinates are $(4,-1)$.

Each coordinate problem is a linear system whose matrix has the basis vectors as columns.

## The change-of-basis matrix

Put the new basis vectors as the columns of a matrix $P$. Then $P[v]_B = v$, so $[v]_B = P^{-1} v$: the change-of-basis matrix is $P^{-1}$. Changing from one nonstandard basis to another composes two such maps. The determinant of $P$ must be nonzero, which is exactly the statement that the new vectors form a basis.

# Projection onto a Subspace

**Projection:** The projection of $v$ onto a subspace $U$ is the unique vector $p \in U$ such that $v - p$ is orthogonal to $U$. It is the closest point of $U$ to $v$, and for a line spanned by $u$ it is the scalar projection formula $\frac{v \cdot u}{u \cdot u} u$.

## The closest point in a subspace

The residual $v - p$ is perpendicular to the subspace, so Pythagoras gives $\|v - q\|^2 = \|v - p\|^2 + \|p - q\|^2$ for any other $q$ in $U$. The projection therefore minimises distance, which is why least squares is a projection problem.

## Worked: projecting onto a line

1. Project $(3,4)$ onto the line spanned by $(1,1)$.
2. The coefficient is $(3 + 4)/(1 + 1) = 3.5$.
3. So the projection is $3.5(1,1) = (3.5, 3.5)$, and the residual $(3,4) - (3.5,3.5) = (-0.5,0.5)$ is orthogonal to $(1,1)$.

Projecting $(1,3)$ onto the same line gives $2(1,1) = (2,2)$, because $(1+3)/2 = 2$.

## Projection matrices

For a subspace with an orthonormal basis collected as the columns of $Q$, the projection is $p = Q Q^T v$, and the projection matrix is $P = Q Q^T$. More generally $P = A(A^T A)^{-1} A^T$ for any basis matrix $A$. A projection matrix is symmetric and idempotent, $P^2 = P$, which is the algebraic statement that projecting twice changes nothing.

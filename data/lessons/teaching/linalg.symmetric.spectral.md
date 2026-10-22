# The Spectral Theorem

**Spectral theorem:** Every real symmetric matrix is orthogonally diagonalizable. There is an orthogonal matrix $Q$ and a diagonal matrix $\Lambda$ with real entries such that $A = Q\Lambda Q^T$, so the columns of $Q$ are an orthonormal basis of eigenvectors.

## Symmetric matrices have real eigenvalues

If $Ax = \lambda x$ with $x$ nonzero, then $\lambda\|x\|^2 = x^T A x$. For a real symmetric $A$ the scalar $x^T A x$ is real, so $\lambda$ is real. Eigenvectors for distinct eigenvalues are orthogonal: symmetry turns $x^T A y$ into both $\mu x^T y$ and $\lambda x^T y$, forcing $(\lambda - \mu)x^T y = 0$.

## Worked: orthogonally diagonalizing a symmetric matrix

1. Take $A = [[2, 1], [1, 2]]$.
2. The characteristic polynomial is $(2-\lambda)^2 - 1 = \lambda^2 - 4\lambda + 3$, so the eigenvalues are 3 and 1.
3. Eigenvectors are $(1,1)$ for 3 and $(1,-1)$ for 1; they are orthogonal, and normalizing gives $Q = \frac{1}{\sqrt2}[[1,1],[1,-1]]$.

## Why orthogonality matters

An orthonormal eigenbasis means $Q^{-1} = Q^T$, so no matrix inverse is needed. The decomposition exposes a quadratic form as a weighted sum of squares, which is the basis for definiteness tests and for principal component analysis. It also makes powers easy: $A^k = Q\Lambda^k Q^T$, because the inner $Q^TQ$ factors collapse to the identity.

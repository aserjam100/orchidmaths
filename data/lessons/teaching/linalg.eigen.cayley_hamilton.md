# The Cayley-Hamilton Theorem

**Cayley-Hamilton theorem:** Every square matrix satisfies its own characteristic polynomial. If $p(\lambda) = \det(A - \lambda I)$, then $p(A) = 0$, where the constant term is multiplied by the identity matrix.

## Every matrix satisfies its characteristic polynomial

The theorem holds even when the matrix is not diagonalizable, which is what makes it a structural result rather than a consequence of having enough eigenvectors. For a 2x2 matrix with characteristic polynomial $\lambda^2 - \operatorname{tr}(A)\lambda + \det(A)$, it gives $A^2 - \operatorname{tr}(A)A + \det(A)I = 0$.

## Worked: a 2x2 example

1. Take $A = [[2,1],[1,2]]$.
2. The characteristic polynomial is $\lambda^2 - 4\lambda + 3$, so Cayley-Hamilton gives $A^2 - 4A + 3I = 0$.
3. Check: $A^2 = [[5,4],[4,5]]$, $-4A = [[-8,-4],[-4,-8]]$, and $+3I = [[3,0],[0,3]]$; the sum is the zero matrix.

## Powers and the minimal polynomial

The relation rewrites $A^2$ as $4A - 3I$, and multiplying by $A$ rewrites every higher power as a combination of $I$ and $A$. The minimal polynomial is the monic polynomial of least degree that annihilates $A$; it divides the characteristic polynomial, and equality holds when the eigenvalues are distinct. Cayley-Hamilton is therefore the reason a matrix's powers live in a finite-dimensional space spanned by its lower powers.

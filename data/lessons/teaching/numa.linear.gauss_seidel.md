# The Gauss-Seidel Method

**Gauss-Seidel method:** like Jacobi, but each updated value is used immediately in the equations that follow, which usually converges faster.

## Using fresh values at once

Gauss-Seidel visits the unknowns in order and, as soon as a new value is computed, uses it in the equations that come after in the same sweep. The information therefore travels across the mesh within one sweep rather than waiting for the next. This small change of bookkeeping often roughly halves the number of iterations.

## Worked: the second unknown

When computing the second unknown, Gauss-Seidel uses the first unknown's new value from the same sweep rather than its old value. That makes the second equation's estimate a little closer to the solution, and the effect compounds down the sweep. The result is the same converged solution as Jacobi, reached sooner.

## Its guarantee and its matrix form

Like Jacobi, Gauss-Seidel is guaranteed to converge when the matrix is strictly diagonally dominant. It can be written as a splitting in which the iteration matrix involves the lower triangle of the matrix, because the already-updated entries are the lower-triangular part. Successive over-relaxation generalises the method by extrapolating with an over-relaxation parameter.

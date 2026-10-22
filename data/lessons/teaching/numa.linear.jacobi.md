# The Jacobi Method

**Jacobi method:** solve each equation of a linear system for its own unknown using values from the previous sweep, and repeat until the values settle down.

## Iterating instead of eliminating

Instead of eliminating unknowns directly, the Jacobi method solves the first equation for the first unknown, the second for the second, and so on, using only values from the previous sweep. Each sweep produces a new vector from the old one. When the values stop changing the system is solved.

## Worked: one sweep step

For the equation four x plus y equals nine, with the previous value of y equal to one, the new x is nine minus one, all over four, which is two. The other equations are treated the same way to complete the sweep. The new values are not used until the next sweep begins, which is what distinguishes Jacobi from Gauss-Seidel.

## Convergence and its condition

The Jacobi iteration converges when the matrix is strictly diagonally dominant, meaning each diagonal entry is larger in absolute value than the sum of the others in its row. A diagonally dominant matrix also has a unique solution, so the method is safe in that case. For general matrices the iteration can diverge, which is why diagonal dominance is the usual test.

# Gaussian Elimination

**Gaussian elimination:** solve a linear system by using each equation to eliminate one unknown from the equations below it, then back-substitute.

## Clearing below the pivot

Elimination takes the first unknown and uses the first equation as a pivot to remove it from every equation below. The same is then done for the second unknown with the second equation, and so on, until the system is upper triangular. Each step subtracts a multiple of the pivot row, chosen so that the target entry becomes zero.

## Worked: a two-by-two system

For the system with coefficients two and one in the first row and one and three in the second, the pivot is two. Subtracting half the first row from the second removes the first unknown and leaves a single equation for the second, giving it the value three. Back-substitution then gives the first unknown the value one.

## Cost and the back substitution

Elimination on n equations costs about one third of n cubed operations, and back-substitution costs about one half of n squared, so the elimination dominates for large systems. The method is direct: it finishes in a fixed number of steps rather than iterating. Its accuracy depends on the size of the pivots, which is why pivoting matters.

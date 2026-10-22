# LU Factorization

**LU factorization:** write the matrix as the product of a lower-triangular matrix and an upper-triangular matrix, so the same elimination can be reused for many right-hand sides.

## Separating the elimination from the data

Gaussian elimination is a sequence of row operations that turns the matrix into an upper-triangular one. Recording the multipliers in a lower-triangular matrix turns that process into a factorization of the original matrix. Once the factors are known, a new right-hand side is handled by two cheap triangular solves instead of a fresh elimination.

## Worked: the multiplier

The multiplier that clears the entry below the pivot is that entry divided by the pivot. For a matrix whose first entry is two and whose lower-left entry is four, the multiplier is two. Each multiplier is exactly one entry of the lower-triangular factor, so the factorization is read off from the elimination rather than computed separately.

## Solving and factoring again

With the factorization in hand, solving is a forward substitution with the lower factor followed by a back-substitution with the upper one. Pivoting interleaves row interchanges with the factors, giving a permuted factorization in which the lower factor is still triangular. The factorization costs the same as one elimination and pays off when many right-hand sides share the same matrix.

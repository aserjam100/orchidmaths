# The rank-nullity theorem

## Rank and nullity

The rank of a matrix is the dimension of its column space, which is the number of independent columns. The nullity is the dimension of its null space, which is the set of vectors the matrix sends to zero.

A matrix that collapses a dimension has a null space of dimension at least one, so the two numbers trade against each other.

## The theorem

Rank plus nullity equals the number of columns:

\[
\operatorname{rank}(A) + \operatorname{nullity}(A) = n.
\]

The number of rows does not appear, which is the part worth noticing. A matrix with five columns and rank two has a null space of dimension three, whatever its number of rows.

## What it says about solutions

The theorem is the counting behind the shape of a solution set. A system written as \(Ax = b\) has a solution only when \(b\) lies in the column space, and when it does, the solutions form a copy of the null space shifted by one particular solution.

So the nullity is how much freedom the solutions have, and the rank is how much of the space the matrix can reach.

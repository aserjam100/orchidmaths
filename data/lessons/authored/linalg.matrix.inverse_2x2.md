# The inverse of a two-by-two matrix

## The two-by-two formula

For a matrix with entries \(a\), \(b\), \(c\) and \(d\), the inverse is

\[
A^{-1} = \frac{1}{ad - bc}\begin{pmatrix} d & -b \\ -c & a \end{pmatrix}.
\]

The diagonal entries swap, the other two change sign, and the whole thing is divided by the determinant.

## Why the determinant divides

The determinant measures how much the matrix scales area. Its inverse has to undo that scaling, so it carries the reciprocal, and when the determinant is zero there is no scaling to undo because the matrix has collapsed the plane onto a line.

That is why a zero determinant means no inverse exists, and it is not a special case but the arithmetic of the formula failing to make sense.

## Inverses of larger matrices

For larger matrices there is no comparable formula to memorise: the inverse is found by row reduction, and the determinant is one of the numbers the process produces rather than something to divide by at the start.

The two-by-two case is worth knowing by hand because it appears constantly, and because the general method's behaviour is visible in it.

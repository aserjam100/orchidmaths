# The Power Method

**Power method:** repeatedly multiply a vector by the matrix and normalise, so the vector grows toward the dominant eigenvector and the growth factor approaches the dominant eigenvalue.

## Letting the dominant direction win

Most vectors can be written as a combination of the matrix's eigenvectors. Each multiplication by the matrix stretches each component by its eigenvalue, so after many multiplications the component with the largest eigenvalue dominates the rest. Normalising after each step keeps the numbers manageable while the direction settles down.

## Worked: a symmetric matrix

For the symmetric matrix with diagonal entries two and two and off-diagonal entry one, the eigenvalues are three and one. After enough steps the iterate lines up with the eigenvector of the eigenvalue three, and the ratio of successive normalised values approaches three. The power method therefore returns the dominant eigenvalue, three.

## Rate, limits and variants

The method converges at a rate set by the ratio of the second-largest eigenvalue to the largest, so it is slow when they are close. It needs the dominant eigenvalue to be unique in absolute value, or the iteration will oscillate between competing directions. The inverse power method uses the inverse of the matrix to find the eigenvalue nearest a chosen shift instead.

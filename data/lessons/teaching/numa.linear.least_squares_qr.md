# Least Squares by QR

**Least squares by QR:** factor the design matrix into an orthogonal and an upper-triangular factor, then solve a triangular system, avoiding the squaring of the condition number.

## An orthogonal route to the fit

The normal equations solve least squares but square the condition number of the design matrix, so they lose accuracy on ill-conditioned problems. Factoring the matrix into an orthogonal factor and an upper-triangular factor lets the orthogonal part be dropped, leaving a triangular system of the same size. Orthogonal transformations preserve lengths, so they do not amplify rounding error.

## Worked: the triangular system

After the factorization the least-squares problem reduces to an upper-triangular system in the coefficients. Solving it by back-substitution gives the same fit the normal equations describe, but computed without ever forming the product of the matrix with its transpose. The residual is read from the discarded part of the transformed right-hand side.

## Why it is preferred

QR is the recommended method whenever the design matrix might be ill conditioned, because it avoids the factor-of-two loss in the condition number that the normal equations introduce. Householder reflections do the zeroing stably, and pivoting can be added to handle rank deficiency by dropping the insignificant columns. The extra cost over the normal equations is modest.

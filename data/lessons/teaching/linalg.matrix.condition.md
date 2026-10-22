# The Condition Number

**Condition number:** For an invertible matrix, the condition number is $\kappa(A) = \sigma_{\max}/\sigma_{\min}$, the ratio of the largest singular value to the smallest. It measures how much the matrix can stretch one direction relative to another.

## How much a matrix can stretch

Singular values are the lengths of the semi-axes of the image of the unit ball. A condition number near 1 means the matrix distorts all directions about equally; a large condition number means some directions are stretched far more than others, so the matrix is close to singular.

## Worked: condition numbers from singular values

1. Singular values 10 and 2 give $\kappa = 10/2 = 5$.
2. Singular values 5 and 0.5 give $\kappa = 5/0.5 = 10$.
3. Equal singular values 9 and 9 give $\kappa = 1$, the best possible conditioning.

## Why conditioning matters

Solving $Ax = b$ can amplify errors by roughly the condition number: a relative error in $b$ may produce a relative error up to $\kappa$ times as large in $x$. A large condition number means digits are lost, so the answer is sensitive to rounding and measurement noise. This is why the SVD is used to diagnose a nearly singular system before solving it.

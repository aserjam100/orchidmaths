# Quadratic Forms and Definiteness

**Quadratic form:** For a symmetric matrix $A$, the quadratic form is $q(x) = x^T A x = \sum_i a_{ii} x_i^2 + 2\sum_{i<j} a_{ij} x_i x_j$. Definiteness describes its sign: positive definite if $q(x) > 0$ for every nonzero $x$, negative definite if $q(x) < 0$, indefinite if it takes both signs, and semidefinite if it can be zero away from the origin.

## From a symmetric matrix to a quadratic form

The symmetric part of a matrix determines the quadratic form, because the skew part cancels in $x^T A x$. In an orthonormal eigenbasis the form becomes $\lambda_1 y_1^2 + \dots + \lambda_n y_n^2$, so the signs of the eigenvalues decide the definiteness.

## Worked: classifying 2x2 forms

1. $q(x,y) = x^2 + y^2$ is positive definite: it is a sum of squares and vanishes only at the origin.
2. $q(x,y) = x^2 - y^2$ is indefinite: positive at $(1,0)$ and negative at $(0,1)$.
3. The form with matrix $[[2,1],[1,2]]$ has eigenvalues 3 and 1, both positive, so it is positive definite.

## Sylvester's criterion

A symmetric matrix is positive definite exactly when all its leading principal minors are positive. For $[[a,b],[b,c]]$ that means $a > 0$ and $ac - b^2 > 0$. The determinant test is the 2x2 case of the same rule, and it is why a positive determinant alone is not enough: the leading entry must also be positive.

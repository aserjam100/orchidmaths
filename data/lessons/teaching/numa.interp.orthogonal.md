# Orthogonal Polynomials

**Orthogonal polynomials:** a sequence of polynomials whose inner products with one another vanish, which makes the least-squares system diagonal and the coefficients independent.

## Why orthogonality helps

If the basis functions are orthogonal, every off-diagonal inner product in the normal equations is zero, so the matrix is diagonal and each coefficient can be computed on its own. Building the fit then becomes a list of independent quotients rather than a coupled system. The inner product is an integral with a weight for continuous data and a sum for discrete data.

## Worked: a coefficient

With an orthogonal basis, the coefficient of each basis function is the inner product of the data with that function divided by the squared norm of the function. There is no coupling to the other coefficients, so each can be found immediately. This is the same formula as the projection of a vector onto an orthogonal basis.

## The classical families

Legendre polynomials are orthogonal with respect to the constant weight on the interval from minus one to one, and Chebyshev polynomials with respect to the reciprocal square-root weight. Using the plain powers of x instead is a poor choice, because they are highly correlated and the normal equations become ill conditioned. Orthogonal families are the stable bases for polynomial fitting.

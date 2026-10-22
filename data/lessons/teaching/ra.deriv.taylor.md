# Taylor's Theorem

**Taylor's theorem:** a sufficiently differentiable function equals its Taylor polynomial plus a remainder, so the polynomial approximates the function with a quantified error.

## The polynomial and the remainder

Taylor's theorem writes a function as its Taylor polynomial of degree n at a point plus a remainder. The polynomial matches the value and the first n derivatives at the point, so it hugs the function there. The remainder measures the error away from the point, and the Lagrange form of the remainder is the next derivative evaluated at some intermediate point, divided by the factorial of n plus one, times the distance raised to n plus one. The theorem is therefore an exact identity, not merely an expansion.

## Taylor series and its convergence

When the remainder tends to zero as the degree grows, the Taylor polynomials converge to the function and the Taylor series is valid. For the exponential, the sine and the cosine the remainder tends to zero everywhere, so the series represent the functions on the whole line. For other functions the remainder may fail to tend to zero, and then the Taylor series converges to something other than the function, or fails to converge at all away from the point. Convergence of the series must always be justified by the remainder.

## Worked: a coefficient

The coefficient of the n-th power in the Taylor series of a function at a point is the n-th derivative there divided by n factorial. For the exponential with a scale factor, the n-th derivative at zero is the scale raised to n, so the coefficient of the cube is the cube of the scale divided by six. Reading a coefficient is therefore a matter of evaluating a single derivative and dividing by a factorial, which is exactly what the definition of the Taylor polynomial prescribes.

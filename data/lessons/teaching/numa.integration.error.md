# Error Bounds for Quadrature

**Quadrature error:** the difference between a rule's estimate and the true integral, governed by the width of the pieces and the size of the relevant derivative.

## Where the error comes from

A quadrature rule is exact for polynomials up to some degree, and its error is the first term it cannot capture. For the trapezoid rule that is the second derivative times the cube of the width, and for Simpson's rule the fourth derivative times the fourth power of the width. The derivative tells you how curved the integrand is, and the width tells you how well the rule follows it.

## Worked: halving the width

The composite trapezoid error is proportional to the square of the width, so halving every subinterval reduces the error by a factor of four. The composite Simpson error is proportional to the fourth power of the width, so halving it divides the error by sixteen. These ratios are the quickest way to check that a computation is behaving as its theory predicts.

## Using the bounds in practice

The bounds need a bound on a derivative, which is often the hard part. When the derivative is unknown one compares two estimates at different widths and uses their difference as a practical error estimate. This is the idea behind Romberg integration and behind every adaptive routine.

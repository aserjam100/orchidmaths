# Error Bounds for Numerical Integration

**Error bound:** the guaranteed maximum error of the trapezoid and Simpson rules, read from a derivative of the integrand and the step size.

## Where the bound comes from

The trapezoid rule replaces the curve on each subinterval by a chord. Repeating this on every subinterval and summing the local errors is what leaves the global error, and the local error on one subinterval is governed by the second derivative of the integrand. This gives a bound equal to the maximum size of the second derivative, times the cube of the interval length, divided by twelve times the square of the number of subintervals. The bound is therefore smaller when the curve is straighter and when the subintervals are finer.

## Doubling the subintervals

Because the trapezoid error falls like the square of the number of subintervals, halving the step reduces it by about a factor of four. This is the practical rule for choosing how many subintervals to use: compute the bound, compare it with the tolerance, and increase the number of subintervals until the bound is below the tolerance. The fourth power of the step governs Simpson's rule, so it converges faster, at the cost of needing an extra function value per subinterval.

## Reading the estimate honestly

The bound uses the largest size of the second derivative on the interval, which is often not known exactly and must itself be bounded. A bound that is larger than necessary is still correct, just pessimistic, and refinements compare successive estimates instead of bounding the derivative. The same structure recurs for other quadrature rules, each with its own derivative order and power of the step, so understanding one bound gives the pattern for the rest.

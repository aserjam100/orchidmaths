# Gaussian Quadrature

**Gaussian quadrature:** choose both the nodes and the weights to make a rule with as few points as possible that is exact for polynomials of the highest possible degree.

## Choosing the nodes too

Newton-Cotes rules fix the nodes evenly and choose the weights, but Gaussian quadrature chooses the nodes as well. With n nodes there are two n free quantities, the positions and the weights, so the rule can be made exact for polynomials up to degree two n minus one. Two points are therefore exact for cubics, and the nodes are the roots of the Legendre polynomial of that degree.

## Worked: two points on the standard interval

On the interval from minus one to one, two-point Gauss quadrature uses the nodes at plus and minus one over the square root of three, with equal weights. Because the nodes are symmetric, the estimate for an even power is twice the value at the positive node. For the fourth power that is twice one ninth, which is 0.222222, while the exact integral is two fifths.

## Why it is the most efficient

Gaussian quadrature achieves the highest possible degree of exactness for a given number of function evaluations, which makes it the rule of choice for smooth integrands. Its price is that the nodes and weights must be tabulated or computed, and the nodes lie strictly inside the interval, which is awkward when the integrand is only defined in part of it. Variants with a weight function handle integrals over infinite intervals.

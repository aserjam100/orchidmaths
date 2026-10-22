# Integrating to Infinity

**Improper integral:** an integral whose interval is infinite or whose integrand becomes unbounded, which must be handled by a change of variable or a special rule rather than a naive truncation.

## Infinite intervals

A finite quadrature rule cannot sample an infinite interval. One standard approach maps the infinite interval to a finite one by a change of variable, such as the tangent substitution, and then applies an ordinary rule on the finite range. Alternatively, a Gauss-type rule with Laguerre or Hermite weights absorbs the decay of the integrand into the weights.

## Worked: the danger of truncation

Truncating an infinite interval and ignoring the tail is tempting because it is easy, but a slowly decaying integrand can leave a large, silent error beyond the cut. Doubling the cut then barely improves the answer, and there is no easy way to know how much mass was missed. A substitution removes the guesswork and gives a rule with a real error bound.

## Unbounded integrands

An improper integral of the second kind has an integrand that blows up at a point inside or at the edge of the interval, which is a different difficulty from an infinite range. A substitution that removes the singularity turns it into a bounded integrand that an ordinary rule can handle. Failing that, a rule with nodes clustered toward the singular point can capture the behaviour directly.

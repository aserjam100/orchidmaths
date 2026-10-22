# Area in polar coordinates

## Half the square of the radius

The area swept by a polar curve between two angles is

\[
A = \frac{1}{2}\int_{\alpha}^{\beta} r^2\,d\theta.
\]

The half is the same half that appears in the area of a sector, and it is there because the area element in polar coordinates is a triangle with sides \(r\) and \(r\,d\theta\).

## Why the half

A thin wedge of angle \(d\theta\) has area approximately a triangle with base \(r\,d\theta\) and height \(r\), which is half the product. Summing the wedges gives the integral, and the squared radius is what makes the answer grow fast with distance.

For a constant radius the integral is just half the square of the radius times the angle, which is the fraction of the circle's area that the angle covers.

## Sweeping a region

The formula finds the area swept from the origin, so it applies to a region described by a radius as a function of the angle and not to a region bounded by straight lines. A circle, a cardioid or a rose fits the description; a square does not.

The most common mistake is to integrate the radius rather than its square, which gives a length instead of an area. Checking the units is enough to catch it.

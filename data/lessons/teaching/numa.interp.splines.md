# Cubic Splines

**Cubic spline:** a piecewise cubic polynomial that matches the data at the knots and matches its first and second derivatives across every knot, giving a smooth interpolant without high-degree oscillation.

## Pieces joined smoothly

Instead of one polynomial through all the points, a spline uses a cubic on each interval between knots. At each knot the two neighbouring cubics are made to agree in value and in their first and second derivatives, so the whole curve is twice differentiable. This smoothness is what makes splines look and behave well.

## Worked: the boundary conditions

The derivative matching gives enough equations for all the unknown coefficients except two, which the boundary conditions supply. The natural spline sets the second derivative to zero at the two endpoints, as if the curve were straight there. Other choices prescribe an endpoint slope, and all give a unique spline.

## Why splines beat one polynomial

A single high-degree interpolating polynomial can oscillate between the data, as Runge's phenomenon shows. Low-degree pieces cannot oscillate that way, and a cubic spline's error on each interval is proportional to the fourth power of the knot spacing times the fourth derivative. Splines are therefore the default tool for smooth interpolation and computer graphics.

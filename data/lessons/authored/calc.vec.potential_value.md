# Evaluating a potential

## Recovering the potential

When a field passes the cross-partial test it is the gradient of a function, and that function is found by integrating one component and matching the other. For \((2x, 2y)\), integrating \(2x\) in \(x\) gives \(x^2\) plus a function of \(y\), and comparing the \(y\)-derivative of that with \(2y\) shows the function is \(y^2\).

The result is \(x^2 + y^2\), with an arbitrary constant still to be fixed.

## Using the initial condition

A single known value removes the constant. If the potential is required to be zero at the origin, then the constant is zero and the potential at \((3, 4)\) is \(9 + 16 = 25\).

That value is also what a line integral of the field from the origin to that point returns, which is the practical use of a potential: it replaces an integral along a path with two evaluations.

## Why the constant matters

The constant changes the potential but not any difference of potentials, so it never changes a line integral between two points. Fixing it is a convention that makes the function single-valued rather than a piece of physics.

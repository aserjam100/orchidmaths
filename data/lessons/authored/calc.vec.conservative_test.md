# Testing a field for a potential

## The cross-partial test

A field is conservative exactly when it is the gradient of a function, and on a simply connected region that happens exactly when

\[\frac{\partial Q}{\partial x} = \frac{\partial P}{\partial y}.\]

The test is one computation on the two components, and for \(\mathbf{F} = (2x + 3y,\ 4x + 5y)\) it gives 4 against 3, so the field is not conservative and no potential exists.

## Finding the potential

When the test passes, integrate either component to recover the potential. For \((2x,\ 2y)\) the cross partials are both zero, so integrating \(P\) in \(x\) gives \(x^2 + g(y)\), and differentiating that in \(y\) and comparing with \(Q\) gives \(g'(y) = 2y\). The potential is \(x^2 + y^2 + C\).

The constant matters for the integral but not for a difference of potentials, which is all a line integral asks for.

## When the test is not enough

The test is about a simply connected region. The field \((-y, x)/(x^2 + y^2)\) passes it everywhere except the origin, and its integral around a circle centred there is not zero, so no single potential exists on the punctured plane. The failure is in the domain, not in the computation.

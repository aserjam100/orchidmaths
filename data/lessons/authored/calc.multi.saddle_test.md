# The second-derivative test in two variables

## The discriminant

At a critical point of \(f(x,y)\) in two variables, the test uses

\[D = f_{xx}f_{yy} - \left(f_{xy}\right)^2.\]

The mixed derivative is squared and subtracted, which is what makes the test different from the one-variable case: a large \(f_{xy}\) can pull a point from a minimum into a saddle.

## Minima, maxima and saddles

With \(D > 0\) the point is a minimum when \(f_{xx} > 0\) and a maximum when \(f_{xx} < 0\). With \(D < 0\) it is a saddle, and with \(D = 0\) the test says nothing.

For \(f = 3x^2 + 4y^2\) at the origin, \(f_{xx} = 6\), \(f_{yy} = 8\) and \(f_{xy} = 0\), so \(D = 48\) and the point is a minimum. Changing the second coefficient to \(-4y^2\) makes \(D = -48\) and the same origin a saddle.

## Why the test fails when D is zero

The formula is the determinant of the matrix of second derivatives, so \(D = 0\) says that matrix is singular and the quadratic approximation has a flat direction. Higher-order terms then decide the shape, and no test built only from second derivatives can see them.

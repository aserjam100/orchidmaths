# The Jacobian as a change of measure

## The determinant of partials

Changing variables in several variables stretches area and volume, and the local stretch factor is the determinant of the matrix of partial derivatives, written \(\partial(x,y)/\partial(u,v)\).

Its absolute value converts an area element in one coordinate system into the area element in the other: \(dx\,dy = \left|\partial(x,y)/\partial(u,v)\right| du\,dv\).

## Polar coordinates

For \(x = r\cos\theta\) and \(y = r\sin\theta\) the partials are \(\cos\theta\), \(\sin\theta\), \(-r\sin\theta\) and \(r\cos\theta\), and the determinant is

\[r\left(\sin^2\theta + \cos^2\theta\right) = r.\]

So the factor is the radius itself, and evaluating it at \(r = 5\) gives 5. Away from the origin the factor is never zero, and at the origin the whole change of variables collapses the plane to a point.

## Why the factor appears

A small rectangle in \((r, \theta)\) space has sides \(dr\) and \(r\,d\theta\): the arc traced by a small angle grows with the radius, which is the geometry behind the algebraic factor \(r\).

The Jacobian is the general version of that observation. It is also what makes an integral computable in a coordinate system suited to the region, since the region's boundary can be constant in the new variables while the factor carries the distortion.

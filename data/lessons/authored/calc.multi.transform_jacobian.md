# The Jacobian of a linear change of variables

## The determinant

A linear change of variables \(x = au + bv\), \(y = cu + dv\) has a constant matrix of partials, and its determinant

\[\frac{\partial(x,y)}{\partial(u,v)} = ad - bc\]

is the factor by which the transformation multiplies every area. For \(x = 2u + 3v\), \(y = u + 4v\) the determinant is \(8 - 3 = 5\).

## Area scaling

A unit square in the \((u,v)\) plane maps to a parallelogram whose area is the absolute value of the determinant. That is why an integral changes by exactly this factor: \(dx\,dy = \left|ad - bc\right| du\,dv\).

A determinant of zero collapses the plane onto a line, and the change of variables is then not invertible, which is the algebraic statement that a whole dimension of information has been lost.

## Composition

Composing two changes of variables multiplies their Jacobians, because the matrices multiply and the determinant is multiplicative. That property is what lets a change be broken into two simpler ones, which is how polar coordinates are often applied after a translation.

# The curl in three dimensions

## Three cross-partial differences

For a field \((P, Q, R)\) in space, the curl is the vector

\[\left(\frac{\partial R}{\partial y} - \frac{\partial Q}{\partial z},\ \frac{\partial P}{\partial z} - \frac{\partial R}{\partial x},\ \frac{\partial Q}{\partial x} - \frac{\partial P}{\partial y}\right).\]

Each component is the planar curl of the field's behaviour in the plane perpendicular to that axis, which is why the formula is cyclic: the \(z\) component is the one from the plane.

## Reading one component

For \((ay, bz, cx)\) the components are \(-b\), \(-c\) and \(-a\). Checking the \(x\) component: the \(y\)-derivative of the third component is zero and the \(z\)-derivative of the second is \(b\), so the component is \(-b\).

The pattern is worth noticing rather than memorising: each component pairs one derivative of a later component against one of an earlier one.

## What the curl measures

The curl at a point is the axis and rate of the local rotation, which is why its direction matters in three dimensions where the planar curl was a single number. A field with zero curl is a gradient, on a simply connected region, and a field with zero divergence is incompressible; a field can have either property without the other.

# Lagrange multipliers

## Gradients in line

At an extremum of \(f\) subject to a constraint \(g = c\), the two gradients are parallel:

\[\nabla f = \lambda \nabla g.\]

If they were not, a small move along the constraint curve would still change \(f\), so the point would not be an extremum. The constant \(\lambda\) is the multiplier, and it measures how much the optimum would improve if the constraint were relaxed.

## The method

Write the two partial equations together with the constraint and solve. Maximising \(xy\) subject to \(x + y = 10\) gives \(y = \lambda\), \(x = \lambda\) and so \(x = y = 5\), with a maximum value of 25.

Three equations in three unknowns is the usual shape, and the constraint is one of them: it is not solved and substituted, it is carried along.

## Why the tangency condition

The constraint curve is a level set of \(g\), so its tangent is perpendicular to \(\nabla g\). The level sets of \(f\) are perpendicular to \(\nabla f\). An extremum along the curve is a point where the curve touches a level set of \(f\) rather than crossing it, so the two normals point the same way, which is the parallel condition again.

That geometric picture also explains the failure modes: a corner of the constraint, or a point where \(\nabla g\) vanishes, needs to be checked separately because the tangent is not defined there.

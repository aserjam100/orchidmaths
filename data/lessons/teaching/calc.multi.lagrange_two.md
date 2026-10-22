# Lagrange Multipliers with Two Constraints

**Two constraints:** one multiplier is introduced per constraint, so the gradient of the objective lies in the plane spanned by the two constraint gradients at an extremum.

## The system of equations

To extremise a function subject to two constraints, the gradient of the objective is set equal to one multiple of the first constraint gradient plus another multiple of the second. Along with the two constraint equations themselves, this gives as many equations as unknowns: the coordinates and the two multipliers. The geometric meaning is that the objective's gradient lies in the plane spanned by the two constraint gradients, so no small step that keeps both constraints satisfied can increase the objective.

## The geometry of two constraints

Two smooth constraints cut out a curve, and a small motion along that curve is tangent to both constraint surfaces, so it is perpendicular to both constraint gradients. At an extremum the directional derivative of the objective along every such tangent must vanish, which holds exactly when the objective's gradient is a combination of the two constraint gradients. If the two constraint gradients are parallel, the constraints are degenerate at that point and the method must be applied on a smaller set; this is the analogue of a vanishing derivative in the one-constraint case.

## Solving and checking

Solving the system usually means eliminating the multipliers first and then the constraints, or using the multipliers to express coordinates in terms of them. Every solution of the system is a candidate, and the extrema are found among the candidates, including any points where the constraint gradients are degenerate. When the constrained set is compact the objective attains its extrema, so the candidate list is finite and the comparison is direct, exactly as in unconstrained optimisation.

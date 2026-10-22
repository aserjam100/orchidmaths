# Green's theorem

## Boundary and interior

Green's theorem equates an integral around a closed curve with an integral over the region it encloses:

\[\oint_C P\,dx + Q\,dy = \iint_R \left(\frac{\partial Q}{\partial x} - \frac{\partial P}{\partial y}\right) dA.\]

The left side is a measurement on the boundary and the right side one on the interior, and the theorem says the two agree when the boundary is traversed counterclockwise.

## Area from a boundary integral

Choosing \(P = 0\) and \(Q = x\) makes the integrand on the right equal 1, so the double integral is just the area of the region and the boundary integral computes it. For a rectangle from the origin to \((4, 3)\) that gives \(4 \times 3 = 12\).

The same idea computes the area of a region whose boundary is easy to parametrise and whose shape is not.

## Orientation

Reversing the direction of travel reverses the sign of the boundary integral, so the counterclockwise convention is part of the statement rather than a detail. For a region with holes the outer boundary is traversed counterclockwise and each inner one clockwise, which keeps the enclosed area positive.

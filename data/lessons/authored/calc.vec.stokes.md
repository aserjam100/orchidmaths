# Stokes' theorem

## Circulation and flux of curl

Stokes' theorem equates the circulation of a field around a closed curve with the flux of its curl through any surface the curve bounds:

\[\oint_C \mathbf{F} \cdot d\mathbf{r} = \iint_S (\nabla \times \mathbf{F}) \cdot d\mathbf{S}.\]

It generalises Green's theorem from the plane to a surface in space, and the plane version is the case where the surface is flat and lies in the \(xy\)-plane.

## A flat surface

For \((ay, bx, 0)\) the curl is \((0, 0, b - a)\), a constant pointing up. Over a rectangle of sides \(p\) and \(q\) in a horizontal plane the flux is \((b-a)pq\), and that is also the circulation around the boundary.

Because the answer does not depend on which surface spans the curve, a wobbly surface can be replaced by a flat one, which is the usual reason to reach for the theorem.

## Why the surface does not matter

Two surfaces with the same boundary enclose a volume, and the flux of a curl through the boundary of any closed volume is zero -- the divergence theorem applied to a curl, and the divergence of a curl is identically zero. The two fluxes therefore agree, which is why the statement can leave the surface unspecified.

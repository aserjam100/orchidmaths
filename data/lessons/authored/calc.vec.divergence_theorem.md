# The divergence theorem

## Flux out of a closed surface

The divergence theorem turns the total flux through a closed surface into a volume integral of the divergence:

\[\iint_S \mathbf{F} \cdot d\mathbf{S} = \iiint_V \nabla \cdot \mathbf{F}\,dV.\]

It is the three-dimensional counterpart of Green's theorem, and the two are the same statement read on a surface and in a plane.

## A box

For a linear field \((x, 2y, 3z)\) the divergence is \(1 + 2 + 3 = 6\), a constant. Over a box with sides 2, 3 and 4 the volume is 24, so the flux is \(6 \times 24 = 144\).

A constant divergence makes the volume integral a product, which is why the theorem is worth using on a linear field even when each face could be done separately.

## What the theorem says

The divergence at a point is the flux per unit volume out of a small box around it, so summing over the whole solid counts the interior contributions twice, once in and once out, and they cancel. Only the boundary survives, which is the theorem.

The same cancellation appears whenever a derivative is integrated: the interior disappears and the answer lives on the edge.

# Hydrostatic Force

**Hydrostatic force:** The force on a vertical plate is $F = \rho g \int y\, w(y)\, dy$, where $y$ is the depth below the surface, $w(y)$ the plate's width at that depth, and $\rho g$ the weight density of the fluid. Pressure grows linearly with depth, so the force is a first moment of the width.

## Pressure grows with depth

At depth $y$ the pressure is $\rho g y$. A thin horizontal strip of height $dy$ and width $w(y)$ has area $w(y)\,dy$, so the force on it is $\rho g\, y\, w(y)\, dy$. Adding the strips turns the force into the integral above.

## Worked: a rectangular plate

1. A plate 2 metres wide has its top at depth 0 and bottom at depth 3, with $\rho g = 10$.
2. The force is $10 \cdot 2 \int_0^3 y\,dy = 10 \cdot 2 \cdot 9/2 = 90$.
3. If a 2-metre plate sits between depths 1 and 3, the force is $10 \cdot 2\,(9 - 1)/2 = 80$.

## Why the integral appears

The pressure is not constant over the plate, so multiplying the total area by one pressure would be wrong. The integral weights each strip by its depth, which is why a deeper plate feels more force even when its area is unchanged. The same formula with a variable width handles triangles, semicircles and any other vertical shape.

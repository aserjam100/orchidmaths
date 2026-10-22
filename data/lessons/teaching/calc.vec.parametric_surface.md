# Parametric Surface Area

**Parametric surface:** a surface described by a map of two parameters, with area element the length of the cross product of the two tangent vectors.

## The two tangent vectors

A parametric surface is the image of a rectangle in the parameter plane under a map of two variables. Holding one parameter fixed gives a curve on the surface, and its velocity is one tangent vector; holding the other fixed gives the other tangent vector. The two tangents span the tangent plane at the point, and a tiny rectangle in the parameter plane maps to a tiny parallelogram on the surface whose sides are the two tangents scaled by the parameter steps. The local area is the length of the cross product of the tangents.

## The area element and integration

The area of the surface is the integral, over the parameter rectangle, of the length of the cross product of the two tangent vectors. When the tangent vectors are constant, as for a flat parallelogram, the cross product has constant length and the area is that length times the area of the parameter rectangle. The cross product is perpendicular to the surface, so its direction also gives the unit normal, which the flux integral reuses. Each degree of freedom in the parametrisation contributes one tangent vector and a factor in the area.

## Choosing a parametrisation

A graph given as a height over a plane region is the special case in which the parameter map is the identity on the plane and the two tangents come from the two partial derivatives, recovering the magnification formula for graphs. A surface of revolution or a sphere is parametrised by angles, and the tangent vectors are orthogonal, which makes the cross product length the product of their lengths. Choosing orthogonal parameters is the standard way to keep the area computation short.

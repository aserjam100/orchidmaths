# Scalar Surface Integrals

**Scalar surface integral:** the integral of a scalar function over a curved surface, with the surface area element, used for mass, average value and charge on a surface.

## The surface element

A scalar surface integral adds up the values of a function over a surface, weighting each contribution by the local surface area. For a surface given by a height over a plane region, the surface element is the area element of the plane multiplied by the magnification factor, which is the square root of one plus the squares of the two partial derivatives. For a surface given parametrically the surface element is the length of the cross product of the two tangent vectors. Both formulas encode how much a small patch of the surface stretches a small patch of the parameter domain.

## Working in spherical coordinates

A sphere or hemisphere is the standard example, and spherical coordinates make the surface element a product of the square of the radius, the sine of the polar angle and the angular differentials. An integral of the vertical coordinate over a hemisphere reduces to a single angular integral, because the radius is constant and the outer angular integral contributes a full factor of twice pi. The result is a power of the radius times pi, which matches the geometric intuition that the integral should scale with the cube of the size.

## Physical readings

If the function is a density per unit area, the scalar surface integral is the total mass; if it is a charge density, it is the total charge. Dividing by the total surface area gives the average value of the function on the surface, exactly as in the one-variable case. The scalar integral is the prerequisite for the flux integral, which adds a direction by dotting a vector field with the unit normal before integrating.

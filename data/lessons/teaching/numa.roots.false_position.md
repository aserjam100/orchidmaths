# The Method of False Position

**False position:** like the secant method, but the new point is chosen to keep the root bracketed, so convergence is guaranteed and often faster than bisection.

## Interpolating inside a bracket

False position starts with two points whose function values have opposite signs. It takes the intercept of the line joining them as a trial point, then keeps the half of the bracket that still contains the sign change. The interpolation uses the sizes of the function values, unlike bisection, which ignores them.

## Worked: why it can be slow

If the function is very curved, the interpolated point can land very close to one endpoint, so that endpoint barely moves for many iterations while the other creeps toward the root. The bracket stays valid, but one side may take a long time to close. The method is therefore dependable but can stall when the curve is far from a line.

## Compared with its cousins

Bisection converges linearly but ignores the function values; false position uses them and usually beats it. The secant method uses them too but drops the bracket, gaining speed and losing the guarantee. Newton is faster still but needs the derivative. Each method trades a guarantee for speed, and false position sits in the middle.

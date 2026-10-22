# The tangent plane and linear approximation

## The plane that matches the surface

The tangent plane to a surface at a point is the plane that matches the surface's value and both of its slopes there:

\[
z = f(a,b) + f_x(a,b)(x-a) + f_y(a,b)(y-b).
\]

It is the two-variable version of the tangent line, and its coefficients are the two partial derivatives because a plane has one slope in each direction.

## Linear approximation

Dropping the higher-order terms turns the tangent plane into an estimate: near the point, the change in the function is approximately the two partial derivatives times the two changes in the inputs.

That is how a small error in each input contributes to the error in the output, which is the same content as error propagation in one variable with a second input added.

## Where the approximation is exact

For a quadratic function the tangent plane's estimate is not an estimate at all: the second-order terms are exactly what the plane omits, and for a step of one unit those terms are constants that the formula happens to include.

Knowing when the approximation is good is the practical question, and the answer is that it is as good as the second derivatives are small over the step taken.

# The derivative of a parametrised curve

A parametrised curve describes a path by giving each coordinate as a function of a third variable, \(t\), which is why it can carry loops and vertical stretches that a graph cannot.

## Dividing the derivatives

For \(x = x(t)\) and \(y = y(t)\), the slope of the path is

\[\frac{dy}{dx} = \frac{dy/dt}{dx/dt},\]

which is the chain rule read backwards. The derivative in \(t\) is a velocity, and the slope is the ratio of the vertical speed to the horizontal speed.

## The slope at a point

For \(x = t^2\) and \(y = t^3\) we have \(dx/dt = 2t\) and \(dy/dt = 3t^2\), so \(dy/dx = 3t/2\). At \(t = 2\) the slope is 3, and at \(t = -2\) it is \(-3\): the two points have the same coordinates, and the path passes through them with opposite slopes.

## Curves a graph cannot carry

A circle parametrised as \(x = \cos t\), \(y = \sin t\) has \(dy/dx = -\cot t\), which is undefined where \(\sin t = 0\) even though the circle is smooth there: at those points the tangent is vertical, so a slope does not exist and the parametrisation is the honest description. That is the practical reason to parametrise: the path keeps its tangent where \(dy/dx\) does not.

# The Second Derivative of a Parametric Curve

**Parametric second derivative:** For a curve $x = x(t)$, $y = y(t)$, the slope is $dy/dx = y'/x'$ and the concavity is
$$ \frac{d^2y}{dx^2} = \frac{x'y'' - y'x''}{(x')^3}. $$
The denominator is $(x')^3$, not $x'$, because the slope is differentiated with respect to $t$ and then divided by $dx/dt$ once more.

## Differentiate the slope, not the coordinates

It is tempting to write $y''/x''$, which is wrong: the second derivative is the rate of change of the slope with respect to $x$. Starting from $dy/dx = y'/x'$, differentiate with respect to $t$ by the quotient rule to get $(x'y'' - y'x'')/(x')^2$, then divide by $x'$ to convert the $t$-derivative into an $x$-derivative.

## Worked: a parametric parabola

1. Take $x = t$, $y = t^2$, so $x' = 1$, $y' = 2t$, $x'' = 0$ and $y'' = 2$.
2. The formula gives $(1 \cdot 2 - 2t \cdot 0)/1^3 = 2$.
3. Since $y = x^2$ here, the result matches the ordinary second derivative, as it must.

At a point with $x' = 2$, $y' = 3$, $x'' = 2$ and $y'' = 6$, the value is $(12 - 6)/8 = 0.75$.

## Concavity from the second derivative

A positive second derivative means the curve is concave up with respect to $x$, a negative one concave down, provided $x'$ is nonzero. Where $x' = 0$ the curve has a vertical tangent and the formula does not apply; the concavity must be read from the geometry instead. The sign of the denominator matters too, because $(x')^3$ keeps the sign of $x'$.

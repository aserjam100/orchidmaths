# Transformations of Random Variables

**Transformation:** If $Y = g(X)$ for a strictly monotone differentiable $g$, then the density of $Y$ is $f_Y(y) = f_X(g^{-1}(y))\,|\tfrac{d}{dy}g^{-1}(y)|$.

## Changing variables

A transformation moves probability mass from the scale of $X$ to the scale of $Y$. The inverse function relocates the density, and the derivative of the inverse corrects for how much the transformation stretches or compresses the axis. The absolute value keeps the density positive regardless of whether $g$ increases or decreases, and the correction is the smooth change-of-variables factor from calculus.

## Worked: a linear map

If $X$ is uniform on the unit interval then $Y = aX + b$ is uniform on the interval from $b$ to $a + b$. The inverse is $y \mapsto (y - b)/a$, whose derivative is constant, so the density stays constant and the interval simply moves and stretches. The probability that $Y$ is at most a value is then the same linear proportion as before the transformation.

## Variance under scaling

For a linear transformation the effect on the first two moments is exact: $E[aX + b] = aE[X] + b$ and $\operatorname{Var}(aX + b) = a^2\operatorname{Var}(X)$. The additive constant shifts the location and drops out of the variance, while the multiplier scales the spread by its square. This is the linear transformation rule in its most used form.

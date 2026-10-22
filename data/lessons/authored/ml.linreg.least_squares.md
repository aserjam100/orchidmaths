# The least-squares line

## The line through the middle

Fitting a line to data means choosing the slope and intercept that make the vertical distances to the points as small as possible. Squaring those distances and adding them gives a single number to minimise, and the line that minimises it is the least-squares line.

The solution has a closed form, which is why a first course computes it by hand rather than by iterating.

## The formulas

The slope is the covariance of the two variables divided by the variance of the input:

\[
m = \frac{\sum (x_i - \bar{x})(y_i - \bar{y})}{\sum (x_i - \bar{x})^2},
\]

and the line then passes through the point of averages, which fixes the intercept. That the fit passes through the mean point is worth remembering as a check.

## Why squared

Squaring makes every distance positive, so errors in the two directions cannot cancel, and it makes the problem differentiable, so the minimum has a formula rather than needing a search. The cost is that a single outlier pulls the line much harder than it should.

Minimising the absolute distances instead is more robust and has no closed form, which is the trade written into almost every fitting method.

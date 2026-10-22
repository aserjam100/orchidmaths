# Simpson's Rule

**Simpson's rule:** fit a parabola through the two endpoints and the midpoint, integrate it exactly, and weight the endpoints once and the midpoint four times.

## Fitting a parabola

A straight line cannot capture curvature, so Simpson's rule fits a parabola through three values and integrates it exactly. The resulting weights are one for each endpoint and four for the midpoint, all over six, times the width. The rule is exact for any polynomial up to degree three, even though the parabola is only degree two, because the cubic error term cancels by symmetry.

## Worked: the square over an interval

For the square function on the interval from nought to one, the endpoint values are nought and one and the midpoint value is a quarter. The weighted sum is nought plus four times a quarter plus one, which is two, over six, giving one third. That is the exact integral, because the square is below degree three.

## Why it is the default

Simpson's rule is exact for cubics and has an error proportional to the fourth power of the width times the fourth derivative. That is two orders better than the trapezoid rule for the same number of samples, at almost no extra cost. It is the standard rule whenever the integrand is smooth and the sample points can be chosen.

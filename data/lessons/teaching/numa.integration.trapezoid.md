# The Trapezoid Rule

**Trapezoid rule:** estimate the integral over one interval by the area of the trapezoid under the straight line joining the two endpoint values.

## Replacing the curve by a chord

The integral is the area under the curve. The trapezoid rule replaces the curve on the interval by the chord joining its endpoints and takes the area of the resulting trapezoid, which is the width times the average of the two endpoint values. It is exact for straight lines and is the simplest of all quadrature rules.

## Worked: the square over an interval

For the square function on the interval from nought to one, the endpoint values are nought and one, so the estimate is one half of their sum, which is 0.5. Over the interval from one to two the endpoint values are one and four, so the estimate is three halves of their sum, which is 2.5. The exact integrals are one third and seven thirds, so the rule overestimates a convex function.

## When it is exact and when it is not

The trapezoid rule is exact whenever the function is a straight line, so its error comes from the curvature. For a convex function it overestimates, and for a concave function it underestimates, because the chord lies on one side of the curve. The error is proportional to the cube of the width times the second derivative, which is why the rule improves quickly as the interval shrinks.

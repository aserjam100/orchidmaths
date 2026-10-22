# The Midpoint Rule

**Midpoint rule:** estimate the integral over one interval by the area of the rectangle whose height is the function value at the midpoint.

## Replacing the curve by a level

The midpoint rule samples the function once, at the centre of the interval, and treats that value as the height of the whole region. The estimate is the width times the value at the midpoint. It ignores the endpoints entirely, which is why it behaves differently from the trapezoid rule on curved functions.

## Worked: the square over an interval

For the square function on the interval from nought to one, the midpoint is one half and its square is a quarter, so the estimate is 0.25, which underestimates the exact value of one third. Over the interval from nought to two the midpoint is one and its square is one, so the estimate is two. The midpoint rule underestimates a convex function, the opposite of the trapezoid rule.

## A surprising accuracy

Although the midpoint rule uses only one sample, its error is proportional to the square of the width, the same order as the trapezoid rule, and its constant is about half as large. This is why it is often the better single-interval rule for smooth functions. The two rules bracket the answer from opposite sides when the function is convex or concave.

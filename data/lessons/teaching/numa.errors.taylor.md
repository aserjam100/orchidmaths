# Taylor Truncation Error

**Truncation error:** the error a method makes by replacing an exact limiting process with a finite approximation, and it behaves like the next term left out of a Taylor expansion.

## Cutting a series short

Numerical methods approximate a limit, a derivative or an integral by finitely many operations, and the Taylor expansion shows the size of the term left behind. A method is of order n when its local error behaves like the step raised to the power n plus one. The truncation error is distinct from roundoff, which comes from finite precision rather than from the mathematics.

## Worked: halving the step

A second-order method has a global error proportional to the square of the step, so halving the step reduces the error by a factor of four. Doubling the accuracy therefore needs a quarter of the error, not half, and the ratio of successive errors is a quick diagnostic of the order. Fourth-order methods improve by a factor of sixteen per halving.

## Truncation versus roundoff

Making the step smaller reduces truncation error but increases the number of operations and hence the accumulated roundoff. The total error therefore has a minimum at some intermediate step, beyond which roundoff dominates. Choosing a step is a trade between the two, and neither can be driven to zero on its own.

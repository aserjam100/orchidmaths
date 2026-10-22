# Stepping through an interval

## How many steps fit

A numerical method steps across an interval in fixed increments, so the number of steps is the width of the interval divided by the step size. The count is arithmetic and needs no knowledge of the equation.

The step size is usually chosen to divide the interval, which makes the count whole and the last step land exactly on the far end.

## Why the step size divides the interval

A step size that does not divide the interval leaves a partial step, and the method then either overshoots or stops short. Choosing the step to suit the interval avoids the question entirely.

That is why numerical work starts by fixing the interval and then choosing a step size from a list of divisors.

## The cost of a smaller step

Halving the step size doubles the number of steps, and each step costs an evaluation of the equation. So the accuracy of the estimate is bought at a cost that grows in proportion to the number of steps, and the method's error shrinks in proportion to the step for Euler's method.

That trade is why a method with a smaller error per step is worth the extra arithmetic, and why the count matters before the estimate does.

# Changing the order of integration

## Two ways to describe a region

A region in the plane can be written as limits in two orders. A triangle below the line \(y = x\) is described from \(y = 0\) to \(y = x\) with \(x\) running over an interval, or from \(x = y\) to \(x\) up to the widest point with \(y\) running over its own interval.

The value of the integral is the same either way, but the work is not: one order often gives an inner integral that can be evaluated and the other does not.

## Writing the other limits

The way to swap is to draw the region and read off the new limits from the picture, not to manipulate the old ones algebraically. For a triangle the swapped limits turn a constant upper limit into a variable one and vice versa, which is the whole content of the change.

Getting the new limits right is the entire difficulty, and the check is that the two descriptions cover the same region: a point inside one is inside the other.

## Why the value does not change

Fubini's theorem says that for a continuous function on the region the two iterated integrals agree, so the order is a matter of convenience rather than correctness. The theorem needs the function to be reasonable on the region, which is why a discontinuity inside would have to be handled by splitting the region.

The practical use is not to change the answer but to find the order in which the inner integral is doable, and a region whose first order gives an impossible antiderivative is the standard reason to swap.

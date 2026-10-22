# The Limit of a Function

**Limit of a function:** the value a function approaches as its input approaches a point, defined so that the output can be held within any positive distance of the limit by keeping the input close enough to the point, but not equal to it.

## The epsilon-delta definition

The limit of a function at a point c is the number L such that for every positive epsilon there is a positive delta with the following property: whenever the input is within delta of c and is not equal to c, the output is within epsilon of L. The order of the quantifiers is the same as for sequences, with epsilon first and delta allowed to depend on it. The punctured condition that the input is not equal to c is what makes the limit ignore the value, or even the definition, of the function at the point itself.

## Worked: a linear function

For the linear function which sends x to m x plus b with m nonzero, the output difference from the value at two is m times the input difference from two. So the condition that the output be within epsilon of the value at two becomes the condition that the input be within epsilon over m of two. The largest delta that works is therefore epsilon divided by the slope. This is why, for straight lines, delta shrinks as epsilon shrinks and as the line gets steeper.

## Uniqueness and locality

A function has at most one limit at a point, because two different limits would need disjoint neighbourhoods of the output while the inputs approach the same point. The limit is local: only the behaviour of the function near the point matters, and changing the function far away, or at the point itself, does not change the limit. One-sided limits restrict the input to approach from the left or the right, and the two-sided limit exists exactly when both one-sided limits exist and agree.

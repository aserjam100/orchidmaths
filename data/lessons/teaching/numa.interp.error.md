# The Interpolation Error

**Interpolation error:** the difference between a function and its interpolating polynomial, proportional to the product of the distances to the nodes and to the next derivative of the function.

## The error formula

The error at a point is the product of the distances from that point to every node, times the value of the (n+1)-th derivative at some unknown point inside the interval, divided by a factorial. The node product vanishes at the nodes, so the interpolant matches the function exactly there, and it grows between the nodes, where the error lives.

## Worked: the node product

With equally spaced nodes at nought, one and two, the node product at a point between one and two is small, because the point is close to two nodes. Near the midpoint of a gap the product is largest. This is the shape of the error curve and the reason more nodes do not always help at the ends.

## Runge's phenomenon

Adding equally spaced nodes can make the interpolant oscillate wildly near the ends of the interval even when the function is smooth. The higher derivatives grow faster than the node product shrinks, so the maximum error diverges. Choosing the nodes differently, or giving up the single polynomial for splines, is the remedy.

# The norm of a vector

## The squared length

The length of a vector is the square root of the sum of its squared components, and its square is that sum:

\[
\|v\|^2 = v \cdot v.
\]

For the vector with components 3 and 4 the squared length is 25 and the length is 5. Working with the square avoids the root, and any comparison of lengths can be made with the squares because squaring is increasing for positive numbers.

## Why squared

The squared length is a dot product, so it is bilinear: it can be expanded, differentiated and combined with the rest of the algebra. The norm itself is not, which is why computations are usually carried out squared and the root is taken only at the end.

That is the same reason a distance formula uses squared differences before the square root, and why a nearest-neighbour search compares squares.

## In any dimension

The same sum of squared components defines the length in any number of dimensions, so the concept does not need a separate formula per dimension. That is what makes the norm the first example of an inner product, whose axioms are exactly the properties this sum has.

A norm of zero means every component is zero, and that only holds because the sum of squares cannot cancel.

# Chebyshev Nodes

**Chebyshev nodes:** a set of interpolation nodes clustered toward the ends of the interval that minimises the largest absolute value of the node product, controlling Runge's phenomenon.

## Clustering toward the ends

Equally spaced nodes are dense in the middle and sparse at the ends, which is exactly where the node product is largest and the interpolant oscillates. Chebyshev nodes are spaced by a cosine rule so that they crowd toward the endpoints and thin out in the middle. The crowding cancels the growth of the derivative term and keeps the error bounded.

## Worked: a bounded error

With Chebyshev nodes the node product has the smallest possible maximum absolute value among all node sets of the same size. The minimax property is the definition of the Chebyshev polynomial, and it makes the node product bounded instead of growing. The interpolant then converges for any function with a little smoothness.

## From theory to practice

Chebyshev nodes are used whenever interpolation or quadrature must be accurate on a fixed interval. The same clustering underlies Chebyshev quadrature and the fast methods for approximating smooth functions. They are the standard answer to the question of where to place sample points.

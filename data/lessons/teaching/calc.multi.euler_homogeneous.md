# Euler's Homogeneous Function Theorem

**Homogeneous function:** one whose value at a scaled input is the scale raised to a power times the value; Euler's theorem ties its partial derivatives to the function itself.

## Homogeneity and its degree

A function of several variables is homogeneous of degree k when scaling every input by a factor scales the value by that factor raised to k. Polynomials whose every term has the same total degree are homogeneous of that degree, and the reciprocal of a distance-like expression is homogeneous of a negative degree. Homogeneity captures the idea that the function has no preferred length scale, so its values depend on direction and overall size but not on absolute position.

## Euler's identity

Euler's theorem says that for a differentiable homogeneous function, the sum of each input times the partial derivative in that input equals the degree times the function. It is obtained by differentiating the scaling identity with respect to the scale and then setting the scale to one. Differentiating repeatedly gives further identities for the higher partial derivatives, and the theorem is the multidimensional form of the one-variable statement that a power function satisfies input times its derivative equals the exponent times the function.

## Using the identity

Euler's identity computes the weighted sum of the partial derivatives from the value of the function alone, without finding the derivatives first, which is its practical advantage. In economics a production function homogeneous of degree one has total output equal to the sum of the payments to each input at its marginal product, the statement that the distribution of output exhausts the product. The identity also gives quick consistency checks on computed gradients and gradients of power-law functions.

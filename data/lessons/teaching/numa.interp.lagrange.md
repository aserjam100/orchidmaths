# Lagrange Interpolation

**Lagrange interpolation:** the unique polynomial of degree at most n through n plus one points, written as a weighted sum of basis polynomials that are one at their own node and zero at the others.

## Building the basis

Each Lagrange basis polynomial is constructed to be one at its own node and zero at all the others, so it can be built from the products of the distances to the other nodes. Weighting each node value by its basis polynomial and adding gives a polynomial that passes through every point. The construction is unique for a given set of nodes.

## Worked: three points

Through the points at nought, one and two, the interpolant is the quadratic that matches the three values. Each basis polynomial is a product of two linear factors divided by the corresponding constant, and the interpolated value at any query point is the weighted sum. For equally spaced nodes the weights follow a simple pattern.

## Strengths and costs

Lagrange interpolation is elegant and needs no preliminary computation of differences, which makes it easy to use for a single evaluation. Its weakness is that adding a new node forces all the basis polynomials to be rebuilt from scratch. That cost is why the Newton form is preferred when nodes are added one at a time.

# Variance explained by a component

## Eigenvalues as variances

Principal-component analysis finds the directions in which the data varies most. The eigenvalue of a component is the variance of the data along it, so the components can be ranked by how much of the data's spread each one accounts for.

With eigenvalues of 12 and 4, the first component explains 12 over 16, which is 75 percent of the total.

## The share of the total

A component's share of the variance is its eigenvalue divided by the sum of them all. The shares add to one across the components, which is why they are a partition of the total spread rather than independent numbers.

That partition is what makes the analysis useful for compression: a few components with large shares summarise the data, and the rest can be dropped.

## Choosing how many components

The usual rule is to keep enough components to reach a chosen share of the variance, such as 90 percent. The rule is a convention rather than a theorem, and the right number depends on whether the components are being used for compression, for visualisation or as the input to another model.

Dropping components always loses something, and the share of the variance is how much: it is the size of what was thrown away rather than a measure of what is left.

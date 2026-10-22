# Correlation

**Correlation:** the covariance rescaled by the two standard deviations, $\rho = Cov(X, Y)/(\sigma_X \sigma_Y)$, which lies between minus one and one.

## Standardising the covariance

Covariance carries the units of the two variables, so its size depends on the scales and is hard to interpret. Dividing by the product of the standard deviations removes the units and confines the result to the interval from minus one to one. Correlation is the scale-free version of the same linear association.

## Worked: from covariance and scales

A pair has covariance six and standard deviations three and five. The correlation is $6/(3 \cdot 5) = 0.4$, a moderate positive linear association. Changing the units of either variable would change the covariance but leave this number untouched, which is the point of standardising.

## Interpreting the sign and size

A correlation of plus or minus one means the pair lies exactly on a line, and a correlation near zero means the linear association is weak. Crucially, a zero correlation does not mean independence, because a purely nonlinear relation can have zero correlation. Correlation measures linear dependence only, and that is both its strength and its limit.

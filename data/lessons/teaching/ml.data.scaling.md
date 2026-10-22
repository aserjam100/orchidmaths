# Feature Scaling

**Standardization:** Replace a feature $x$ by $z = (x - \text{mean})/\text{sd}$, using the mean and standard deviation of the training data. The result has mean 0 and standard deviation 1, so features measured in different units become comparable.

## Standardization and the z-score

The z-score says how many standard deviations a value sits from the mean. It is affine, so it preserves the shape of the distribution; it does not make skewed data normal. Min-max scaling to the interval from 0 to 1 is an alternative when bounds matter more than the spread.

## Worked: standardizing three values

Take mean 4 and standard deviation 2.

1. $x = 10$ gives $(10 - 4)/2 = 3$.
2. $x = 4$ gives $(4 - 4)/2 = 0$.
3. $x = 2$ gives $(2 - 4)/2 = -1$.

## Why unscaled features distort distance

A distance such as the one kNN and k-means use sums squared differences across features. A feature measured in thousands then dominates one measured in tenths, so the model effectively ignores the small-scale feature. Fit the mean and standard deviation on the training split only; using the whole dataset leaks test information into the transform.

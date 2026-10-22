# Least Squares

**Least squares:** the line or curve that minimises the sum of the squared vertical distances from the data, used when the points cannot all be fitted exactly.

## Fitting without matching

When there are more data points than degrees of freedom, no curve passes through all of them, so the goal changes from exact matching to the best compromise. Least squares measures badness by the sum of the squared vertical gaps and chooses the curve that makes it smallest. Squaring treats positive and negative gaps alike and makes the algebra tractable.

## Worked: the slope of a line

For the line y equals a plus b x the sum of squared gaps is a quadratic in the two unknowns, so it can be minimised directly. Subtracting the means of the data removes the intercept from the slope equation, leaving the centred sum of x times y divided by the centred sum of x squared. With points at values one, three and four over x values nought, one and two, the slope is three halves, which is 1.5.

## Why least squares is popular

The least-squares estimate is unique, easy to compute and has a clean statistical interpretation when the gaps are independent errors of equal variance. It is the maximum-likelihood fit under a Gaussian error model, which is why it is the default in science. Its weakness is sensitivity to outliers, because a single large gap is squared and can dominate the sum.

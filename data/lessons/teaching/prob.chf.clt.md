# The Central Limit Theorem by Characteristic Functions

**Central limit theorem by characteristic functions:** the transform of a standardised sum converges to the exponential of minus half the square, the transform of the standard normal.

## Expanding the log transform

Near the origin the logarithm of the transform of a single standardised summand is minus half the square of the parameter over the index, plus a remainder that vanishes faster. The expansion uses only the first two moments, which is why the limit depends on the summands only through their mean and variance. Independence is not even needed for the expansion step.

## Multiplying and taking the limit

Because the summands are independent, the transform of the sum is the product of the individual transforms, and taking logarithms turns that product into a sum. The sum of the quadratic terms accumulates to minus half the square, while the remainders vanish. Exponentiation returns the transform of the standard normal.

## Identifying the limit law

The limit transform is continuous at zero and is recognised as that of the standard normal. The continuity theorem then converts this pointwise convergence of transforms into weak convergence of the standardised sums. The whole proof is an analytic computation followed by a single probabilistic theorem.

# Covariance

**Covariance:** a signed measure of how two variables move together, $Cov(X, Y) = E[XY] - E[X]E[Y]$.

## The definition

Covariance is the average of the product of the two centred variables. Expanding the product gives the shortcut form, the mean of the product minus the product of the means, which is usually the faster computation. A positive value means the variables tend to be large together, a negative value that one tends to be large when the other is small.

## Worked: a joint table

A joint mass function on the four pairs from zero to one is proportional to the weights one, one, one and seven. Only the pair with both coordinates one contributes to the product, giving $E[XY] = 0.7$, while each mean is 0.8. The covariance is $0.7 - 0.64 = 0.06$, a small positive value from a strong diagonal pull.

## Properties

Covariance is symmetric and bilinear, so it distributes over sums and pulls out constants. The covariance of a variable with itself is its variance, and independent variables have zero covariance. These properties make covariance the natural inner product on centred variables, and they carry the whole algebra of linear combinations of variables.

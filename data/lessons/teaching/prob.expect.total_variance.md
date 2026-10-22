# The Law of Total Variance

**Law of total variance:** the total variance splits into the mean of the conditional variances plus the variance of the conditional means, $Var(X) = E[Var(X \mid Y)] + Var(E[X \mid Y])$.

## Variance within plus variance between

The first term is the variation left inside the slices, averaged over them, and the second is the variation of the slice centres. If the slices are internally spread out the first term is large; if their centres differ the second is. The two add because the cross term vanishes by the tower property.

## Worked: two equally likely groups

A variable $Y$ is equally likely to take either of its two values. Given $Y$, the conditional mean of $X$ is zero or two, and the conditional variance is one or three. The mean conditional variance is two, the variance of the means is $(0 - 2)^2/4 = 1$, and the total variance is three. Reading the two components separately shows where the spread comes from.

## Why the terms are the only ones

Expanding the variance as the second moment minus the square of the mean, and substituting the conditional versions, leaves exactly the two terms. The apparent cross term averages to zero because $E[X - E[X \mid Y] \mid Y]$ is zero inside each slice. Knowing the cancellation is why the identity holds without any extra assumption.

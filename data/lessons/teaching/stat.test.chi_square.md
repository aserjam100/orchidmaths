# The Chi-Square Goodness-of-Fit Test

**Goodness of fit:** Compare observed counts $O_i$ with the counts $E_i$ expected under a null model using $\chi^2 = \sum (O_i - E_i)^2/E_i$. Large values mean the observed counts stray far from the model, relative to the size of the expected counts.

## Comparing observed with expected counts

Each term squares a discrepancy and divides by the expected count, so a difference of 2 counts matters more when 10 were expected than when 1000 were. The statistic is nonnegative and is zero exactly when every observed count equals its expected count.

## Worked: three categories

1. Observed counts 12, 9, 9 against expected 10, 10, 10.
2. The terms are $4/10$, $1/10$ and $1/10$.
3. Summing gives $\chi^2 = 0.6$, a small value consistent with the null.

## Degrees of freedom

For $k$ categories the statistic has $k - 1$ degrees of freedom, because the counts must add to the fixed total. A larger statistic than the critical value for the chosen alpha rejects the null model. The test says the model fits poorly; it does not say which category is responsible, so inspect the largest terms.

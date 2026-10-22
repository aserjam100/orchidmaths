# The Central Limit Theorem

**Central limit theorem:** If $X_1, \dots, X_n$ are independent and identically distributed with mean $\mu$ and finite variance $\sigma^2$, the standardised sum $(S_n - n\mu)/(\sigma\sqrt{n})$ converges in distribution to the standard normal.

## Standardising the sum

The sum $S_n$ has mean $n\mu$ and variance $n\sigma^2$, so subtracting the mean and dividing by the standard deviation produces a variable with mean zero and variance one. The theorem says that this standardised variable forgets the shape of the original distribution and approaches the standard normal as $n$ grows. The sample mean is the same statement divided by $n$, with variance $\sigma^2/n$.

## Worked: standardising a value

Take one hundred independent variables with mean one and standard deviation two. The sum has mean one hundred and standard deviation twenty. The value one hundred and ten therefore standardises to $(110 - 100)/20 = 0.5$, a $z$-score of one half. Tables of the standard normal then approximate the probability that the sum is at most that value.

## What it does and does not say

The theorem is about the distribution of the standardised sum, not about the raw values, and it needs the variance to be finite. It says nothing about how fast the convergence is; that is the subject of the Berry-Esseen bound, which uses the third moment. Nor does it require the original variables to be normal, which is exactly why the normal distribution appears so widely as a limit.

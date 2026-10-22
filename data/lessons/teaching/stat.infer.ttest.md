# The One-Sample t Statistic

**t statistic:** $t = (\bar{x} - \mu_0)/(s/\sqrt{n})$ replaces the unknown population standard deviation with the sample standard deviation $s$. The reference distribution is Student's t with $n - 1$ degrees of freedom, which has heavier tails than the normal.

## Replacing sigma with s

Using $s$ adds variability, because $s$ is itself estimated from the same sample. The t distribution accounts for that extra uncertainty: it is wider than the normal, especially for small samples, so the same t value is less surprising than the same z value.

## Worked: computing t

1. A sample of 9 has mean 54, null mean 50 and sample standard deviation 6.
2. The standard error is $6/\sqrt{9} = 2$.
3. So $t = (54 - 50)/2 = 2$, to be compared with a t distribution on 8 degrees of freedom.

## Degrees of freedom and heavier tails

The degrees of freedom are $n - 1$, one less than the sample size because the deviations are measured from the estimated mean. As $n$ grows the t distribution approaches the normal, so the distinction matters most for small samples. A t test is the standard one-sample test whenever $\sigma$ is unknown.

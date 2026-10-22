# The One-Sample z Statistic

**z statistic:** $z = (\bar{x} - \mu_0)/(\sigma/\sqrt{n})$ measures how many standard errors the sample mean lies from the null mean $\mu_0$. It is used when the population standard deviation $\sigma$ is known and the sampling distribution is approximately normal.

## A signal measured in standard errors

The denominator is the standard error of the mean: the typical distance between a sample mean and the population mean for samples of size $n$. Dividing by it puts the gap on a universal scale. A z of 1 means the sample mean is one standard error above the null value; a z of 2.5 is much stronger evidence.

## Worked: computing z

1. A sample of 16 has mean 48, and the null mean is 50 with $\sigma = 4$.
2. The standard error is $4/\sqrt{16} = 1$.
3. So $z = (48 - 50)/1 = -2$: the sample mean is two standard errors below the null value.

## When the z statistic applies

The z version needs the population standard deviation, which is rarely known in practice. It also relies on the sampling distribution of the mean being normal, which the central limit theorem supplies for large samples even when the population is not normal. When $\sigma$ is estimated from the data, the t statistic replaces z.

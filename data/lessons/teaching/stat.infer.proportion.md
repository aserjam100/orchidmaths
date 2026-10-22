# Standard Error of a Proportion

**Standard error of a proportion:** For a sample proportion $\hat{p}$ from $n$ independent trials, $\mathrm{SE} = \sqrt{\hat{p}(1 - \hat{p})/n}$. It estimates the standard deviation of the sampling distribution of $\hat{p}$.

## The variability of a proportion

A proportion is a mean of zeros and ones, so its standard error follows the same pattern as the mean's: variability shrinks as the square root of the sample size. The factor $\hat{p}(1-\hat{p})$ is largest at 0.5, so a poll near a tied race is the most uncertain.

## Worked: a poll of 400

1. A poll of 400 gives $\hat{p} = 0.5$.
2. Then $\hat{p}(1-\hat{p})/n = 0.25/400 = 0.000625$.
3. The standard error is $\sqrt{0.000625} = 0.025$, or 2.5 percentage points.

## From standard error to margin

A 95 percent interval is roughly $\hat{p} \pm 1.96\,\mathrm{SE}$, so the poll above has a margin of about 4.9 percentage points. The interval is wider for smaller samples and for proportions near 0.5. Its accuracy also assumes independent, randomly selected respondents, which is often the harder assumption to defend.

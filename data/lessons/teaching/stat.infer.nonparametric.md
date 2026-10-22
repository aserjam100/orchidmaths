# Nonparametric Tests

**Nonparametric test:** A test that does not assume a particular distribution for the population, often working with signs or ranks instead of the raw values. The sign test and the Wilcoxon signed-rank test are the standard paired examples.

## When normality is not available

Small samples, ordinal data and heavy outliers all undermine tests that assume normality. A nonparametric test replaces the distributional assumption with a weaker one, usually only that the observations are independent and come from a continuous distribution under the null.

## Worked: the sign test

1. For paired data, record only whether each difference is positive or negative.
2. Under the null of no effect, each sign is equally likely, so the number of positive signs is binomial with $n$ trials and $p = 1/2$.
3. Compare the observed count with that binomial distribution; an extreme count is evidence against the null.

## Ranks and the cost of fewer assumptions

The Wilcoxon signed-rank test ranks the absolute differences before comparing the positive and negative sums, which recovers some information about magnitude. The price of dropping assumptions is lower power when the parametric assumptions actually hold. Choose the nonparametric route when the assumptions are doubtful, not merely because it sounds safer.

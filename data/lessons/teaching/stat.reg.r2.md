# The Coefficient of Determination

**R-squared:** $R^2 = 1 - \mathrm{SSE}/\mathrm{SST}$ is the fraction of the total variation in the response explained by the regression. It lies between 0 and 1, with 1 meaning the fitted values reproduce the data exactly.

## The share of variation explained

SST measures the total spread of the response around its mean; SSE measures what the regression leaves unexplained. Their ratio is the unexplained share, so one minus that ratio is the explained share. For a simple linear regression, $R^2$ is also the square of the correlation between the response and the predictor.

## Worked: R-squared from sums of squares

1. A regression has SST = 100 and SSE = 25.
2. The unexplained share is $25/100 = 0.25$.
3. So $R^2 = 1 - 0.25 = 0.75$: the line explains 75 percent of the variation.

## What R-squared misses

A high $R^2$ does not prove the model is correct, and a low one does not prove it is useless: the relationship may be nonlinear, the data may contain outliers, or the response may be genuinely noisy. Adding predictors can only raise $R^2$, which is why adjusted $R^2$ and residual plots are read alongside it. It also says nothing about causation.

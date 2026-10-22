# Mean Absolute Error

**Mean absolute error:** For residuals $r_i = y_i - \hat{y}_i$, the mean absolute error is $(1/n)\sum |r_i|$. It is the average size of the miss, in the units of the target, with no square to amplify large errors.

## The average absolute residual

Absolute value makes every residual positive, so errors in opposite directions cannot cancel. The mean keeps the result on the scale of the data: an MAE of 3 means the predictions are typically 3 units away, regardless of whether the misses are above or below.

## Worked: MAE of four residuals

1. Residuals 1, -2, 3 and 0.
2. Absolute values 1, 2, 3 and 0 sum to 6.
3. Divide by 4: MAE = 1.5.

## When MAE beats MSE

Squaring gives an outlier a weight proportional to its size squared, so one bad point can dominate the fit. Absolute loss grows linearly and is therefore robust to outliers, which is why the least-squares lesson points at it. The cost is that the absolute value is not differentiable at zero, so the minimiser is found by linear programming or subgradient methods rather than a closed form.

# Heun's Method

**Heun's method:** A second-order Runge-Kutta step. From $(t_0,y_0)$ with step $h$, form the Euler predictor $y^* = y_0 + h f(t_0,y_0)$, then correct with the average of the slopes at the two ends: $y_1 = y_0 + \frac{h}{2}(f(t_0,y_0) + f(t_0+h,y^*))$.

## Predict, then correct

Euler's method uses only the slope at the start, so it ignores how the slope changes across the step. Heun evaluates the slope again at the predicted endpoint and averages the two, which cancels the first-order error and leaves an error of order $h^3$ per step.

## Worked: y' = y

1. With $y(0)=1$ and $h=0.1$, the predictor is $1 + 0.1(1) = 1.1$.
2. The slopes are 1 at the start and 1.1 at the predictor.
3. The corrector is $1 + 0.05(2.1) = 1.105$, close to $e^{0.1} \approx 1.10517$.

## Error and stability

Heun's method is second-order: halving $h$ cuts the global error by about four. It is explicit and easy to implement, but like all explicit methods it has a stability limit on $h$ for stiff equations. RK4 averages four slopes and is fourth-order, at four evaluations per step instead of two.

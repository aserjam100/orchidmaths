# Applications of the Mean Value Theorem

**Mean value theorem:** a differentiable function on an interval reaches its average rate of change at some instant, so a global statement about change is witnessed by a local one.

## Average and instantaneous rates

The mean value theorem compares the average rate over an interval with the instantaneous rate at some point inside it. The average rate of a trip is the total distance divided by the total time, and the theorem says the speedometer read exactly that value at some moment, assuming the position is a differentiable function of time. This is the physical content of the theorem: you cannot average faster than you ever went, and you cannot average slower, so the average itself is attained. The same comparison applies to any quantity whose rate is defined.

## Why the equality must occur

If the instantaneous rate were always above the average, the total change would exceed the average times the length, which contradicts the definition of the average. The same argument rules out its being always below. So the rate crosses the average value, and by the intermediate value property of derivatives it equals it at some point. Rolle's theorem is the case where the average is zero and the endpoint values are equal, and the general theorem follows by subtracting the secant line, which is why the two are proved together.

## Consequences for estimates

The theorem bounds how much a function can change over an interval by the largest size of its derivative times the length of the interval. That bound is the basis of every error estimate that uses a derivative: the linear approximation, the trapezoid error, and the convergence of numerical methods. It also shows that a function whose derivative is zero on an interval is constant there, because any two values would have zero average rate and hence a zero derivative somewhere between them, which is the reasoning behind all integration constants.

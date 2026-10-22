# Finite Differences

**Finite difference:** An estimate of a derivative from values of the function at nearby points. The forward difference is $(f(x+h) - f(x))/h$, the backward difference is $(f(x) - f(x-h))/h$, and the central difference is $(f(x+h) - f(x-h))/(2h)$.

## Approximating a derivative with two evaluations

Each formula approximates the slope of a secant line. The forward and backward versions use one point on each side of the target, while the central version straddles it symmetrically. The step $h$ must be stated because the estimate depends on it, and a different $h$ gives a different secant.

## Worked: forward, backward and central

Take $f(x) = x^2$ at $x = 3$ with $h = 0.1$.

1. Forward: $(9.61 - 9)/0.1 = 6.1$.
2. Backward: $(9 - 8.41)/0.1 = 5.9$.
3. Central: $(9.61 - 8.41)/0.2 = 6$, which is the exact derivative $2x = 6$.

The two one-sided estimates bracket the true value and the central estimate averages them.

## Error and step size

The central difference has error of order $h^2$, while the one-sided differences have error of order $h$, so halving $h$ cuts the central error by about four and the one-sided error by about two. Too small an $h$ is also bad in floating point, because subtracting nearly equal numbers loses digits. The best step balances truncation error against roundoff rather than simply shrinking.

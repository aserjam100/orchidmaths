# Order of Convergence

**Order of convergence:** the exponent p in the relation between successive errors, $e_{k+1}$ like $e_k$ to the power p, which measures how fast an iteration improves.

## Defining the order

When the error of a method satisfies a relation in which each new error is proportional to a power of the previous one, that power is the order. Order one is linear convergence, each step cutting the error by a constant factor. Order two is quadratic, each step roughly doubling the number of correct digits, and higher orders improve still faster.

## Worked: reading the order from errors

If three successive errors are a tenth, a hundredth and a ten-thousandth, the logarithmic ratio of the last two divided by that of the first two gives two. The error is being squared at each step, so the method is quadratic. A run of errors that merely drops by a constant factor would give an order of one instead.

## Estimating the order in practice

The order is estimated from observed errors by the same logarithmic ratio, because the exact constant in the error relation is unknown. Superlinear methods have orders strictly between one and two, and the secant method's order is the golden ratio. The estimate is a diagnostic that tells you whether a method is behaving as its theory predicts.

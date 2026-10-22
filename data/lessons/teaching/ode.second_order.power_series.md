# Power-Series Solutions

**Power-series solution:** Assume $y = \sum a_n x^n$ (or a series about another point), substitute into the ODE, and match coefficients of each power. The result is a recurrence for the $a_n$, which determines the solution from its initial values.

## Coefficients from the recurrence

Differentiating the series shifts the index, so each term of the ODE becomes a series. Setting the coefficient of $x^n$ to zero for every $n$ gives the recurrence. For a first-order equation the recurrence is one step; for a second-order equation it links $a_{n+2}$ to earlier coefficients.

## Worked: exponential and cosine

1. For $y' = 2y$ with $y(0) = 1$, the solution is $e^{2x} = \sum (2x)^n/n!$, so $a_2 = 4/2 = 2$.
2. For $y'' + y = 0$ with $y(0)=1$ and $y'(0)=0$, the solution is $\cos x = 1 - x^2/2 + \cdots$, so $a_2 = -1/2$.
3. For $y'' - y = 0$ with the same data, the solution is $\cosh x = 1 + x^2/2 + \cdots$, so $a_2 = 1/2$.

## Radius and singular points

A power series solves the equation on an interval where the coefficient of $y''$ does not vanish and the other coefficients are analytic; a point where the leading coefficient vanishes is singular. At a regular singular point the Frobenius method allows a series with a fractional or shifted power. The radius of convergence is limited by the nearest singular point in the complex plane.

# Cauchy-Euler Equations

**Cauchy-Euler equation:** A linear equation of the form $x^2 y'' + a x y' + b y = 0$ with constant coefficients on the derivatives of descending powers of $x$. Substituting $y = x^m$ turns it into the algebraic equation $m(m-1) + a m + b = 0$.

## Powers of x as solutions

Each derivative of $x^m$ lowers the power by one and multiplies by the old exponent, so $x^2 y''$ and $x y'$ both come back as multiples of $x^m$. The powers cancel, leaving a quadratic in $m$.

## Worked: three characteristic cases

1. $x^2 y'' - 2x y' + 2y = 0$ gives $m^2 - 3m + 2 = 0$, roots 1 and 2, so $y = C_1 x + C_2 x^2$.
2. $x^2 y'' - 3x y' + 4y = 0$ gives $m^2 - 4m + 4 = 0$, a double root at 2, so $y = (C_1 + C_2 \ln x)x^2$.
3. $x^2 y'' + x y' + y = 0$ gives $m^2 + 1 = 0$, complex roots $\pm i$, so $y = C_1\cos(\ln x) + C_2\sin(\ln x)$.

## Why the substitution works

The equation is invariant under the scaling $x \mapsto cx$, which is why power functions are the natural solutions. The three root cases mirror the constant-coefficient equation, with $\ln x$ replacing $x$ in the repeated-root and oscillatory cases. The equation is singular at $x = 0$, so solutions live on an interval that avoids the origin.

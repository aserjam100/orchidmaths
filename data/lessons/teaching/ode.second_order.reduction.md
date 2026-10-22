# Reduction of Order

**Reduction of order:** Given one nonzero solution $y_1$ of a second-order linear homogeneous ODE, substituting $y = y_1 v$ produces a first-order equation for $v'$. Solving it and integrating gives a second, independent solution $y_2 = y_1 v$.

## One solution buys the second

The substitution works because the $y_1 v$ terms cancel: $y_1$ already solves the homogeneous equation, so what remains involves only $v'$ and $v''$. Setting $w = v'$ drops the order by one, and $w$ is first-order linear or separable.

## Worked: y'' - y = 0

1. Take $y_1 = e^x$ and set $y = e^x v$.
2. Then $y'' - y = e^x(v'' + 2v')$, so the equation reduces to $v'' + 2v' = 0$.
3. With $w = v'$, $w' + 2w = 0$, so $w = Ce^{-2x}$ and $v = C_1 + C_2 e^{-2x}$.

The second solution is $e^x e^{-2x} = e^{-x}$, as expected.

## The reduced equation for v

In general the coefficient of $v$ vanishes and the equation becomes $v'' + (2y_1'/y_1 + p)v' = 0$. The Wronskian identity turns the integrating factor into $1/y_1^2$, giving $v = \int e^{-\int p\,dx}/y_1^2\,dx$. The method needs $y_1$ nonzero on the interval, because it divides by $y_1$.

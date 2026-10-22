# Bernoulli Equations

**Bernoulli equation:** A first-order equation of the form $y' + P(x)y = Q(x)y^n$ with $n$ neither 0 nor 1. The power makes it nonlinear, but the substitution $v = y^{1-n}$ turns it into a first-order linear equation.

## A linear term plus a power

When $n = 0$ the equation is already linear, and when $n = 1$ the two $y$ terms combine into one, so the genuinely nonlinear cases are the other exponents. Dividing by $y^n$ and setting $v = y^{1-n}$ removes the power: the derivative $v'$ is a multiple of $y^{-n}y'$, which is exactly the combination the equation provides.

## Worked: y' + y = y^2

1. Here $n = 2$, so set $v = y^{1-2} = 1/y$.
2. Then $v' = -y'/y^2$, and dividing the equation by $y^2$ gives $-v' + v = 1$, that is $v' - v = -1$.
3. The linear equation has integrating factor $e^{-x}$, so $v = 1 + Ce^x$ and $y = 1/(1 + Ce^x)$.

## Choosing the substitution

The exponent $1-n$ is the only power that makes the transformed equation linear. For $n = 2$ it is $v = 1/y$; for $n = 3$ it is $v = y^{-2}$; for $n = 1/2$ it is $v = y^{1/2}$. After solving for $v$, substitute back and check that no division by zero was introduced.

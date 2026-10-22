# Continuous Compounding

**Continuous compounding:** At nominal annual rate $r$, a principal $P$ grows to $A = P e^{rt}$ after $t$ years. It is the limit of $P(1+r/n)^{nt}$ as the number of compounding periods $n$ grows without bound.

## The limit of more frequent compounding

Compounding more often raises the effective rate, but the increases shrink: monthly is slightly better than quarterly, daily slightly better than monthly, and the sequence converges. The limit is the exponential $e^{rt}$, where $e$ is the base that makes the growth rate exactly $r$ at every instant.

## Worked: doubling with continuous compounding

1. Start with $P = 1000$ and $r = 0.05$.
2. The amount is $A(t) = 1000 e^{0.05t}$.
3. At $t = 0$ the amount is 1000; at $t = 10$ it is $1000 e^{0.5} \approx 1648.72$.

## Why e appears

The limit $\lim_{n\to\infty}(1+r/n)^n = e^r$ is the definition of the exponential in disguise. For the same nominal rate, continuous compounding gives the largest amount, so annual compounding is a lower bound and continuous compounding is the upper bound. In practice the difference is small for modest rates, but it is the clean formula for growth at every instant.

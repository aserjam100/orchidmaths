# Markov's Inequality

**Markov's inequality:** If $X$ is nonnegative, then $P(X \ge a) \le E[X]/a$ for every positive $a$.

## A bound from the mean

On the event that $X$ is at least $a$, the variable itself is at least $a$, so its contribution to the expectation is at least $a$ times the probability of that event. Since the rest of the expectation is nonnegative, the whole expectation is at least that contribution. Dividing by $a$ gives the bound. Only nonnegativity and a finite mean are needed.

## Worked: an application

A nonnegative variable has mean ten. Markov's inequality bounds $P(X \ge 25)$ by $10/25 = 0.4$. The bound is valid for every distribution with that mean, which makes it a useful first estimate when the shape is unknown. It is the ancestor of the second-moment bounds such as Chebyshev's inequality.

## Why it is loose

Markov's inequality ignores everything except the mean, so it cannot be sharp for all distributions at once. For a variable that is often zero and occasionally large, the bound can be close to the truth; for a variable concentrated near its mean it is far too generous. Its value is not accuracy but universality, and it is the first step of many stronger inequalities.

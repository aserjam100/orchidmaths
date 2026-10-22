# Euclidean Domains

**Euclidean domain:** An integral domain $R$ is Euclidean if there is a function $N$ from its nonzero elements to the nonnegative integers such that for any $a$ and any nonzero $b$ there are $q$ and $r$ with $a = qb + r$, where $r = 0$ or $N(r) < N(b)$. The function $N$ is the Euclidean function.

## Division with a smaller remainder

The whole content is the inequality: the remainder is either zero or strictly smaller than the divisor as measured by $N$. In the integers $N$ is the absolute value, and in a polynomial ring $N$ is the degree. In both, repeatedly replacing $a$ by the remainder against $b$ must terminate, because $N$ takes values in the nonnegative integers and strictly decreases at each step.

## Worked: the integers and F[x]

1. In $Z$, divide 17 by 5: $17 = 3 \cdot 5 + 2$, and $|2| < |5|$.
2. In $F[x]$ for a field $F$, divide $x^3 + 1$ by $x + 1$: $x^3 + 1 = (x^2 - x + 1)(x + 1) + 0$, a zero remainder.
3. Repeating the step computes a greatest common divisor, which is the single generator of the ideal $(a, b)$.

## Euclidean implies PID implies UFD

Every ideal of a Euclidean domain is principal: choose an element of the ideal of least $N$, and division forces a zero remainder. A principal ideal domain has unique factorisation, because in a PID an irreducible element is prime. The reverse implications fail: $Z[x]$ is a UFD but not a PID, and $Z[\sqrt{-5}]$ is not even a UFD.

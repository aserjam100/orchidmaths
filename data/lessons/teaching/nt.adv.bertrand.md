# Bertrand's Postulate

**Bertrand's postulate:** For every integer $n > 1$ there is at least one prime $p$ with $n < p < 2n$. Equivalently, the prime counting function satisfies $\pi(2n) > \pi(n)$ for every $n > 1$.

## A prime in every interval (n, 2n)

The statement is a density guarantee: however far out you go, the interval from $n$ to twice $n$ is never empty of primes. It is weaker than the prime number theorem, which says the count grows like $n/\log n$, but it is elementary enough to prove without complex analysis.

## Worked: checking n = 100

1. Apply the postulate with $n = 100$: some prime lies strictly between 100 and 200.
2. Indeed 101 is prime, and there are many more in the interval.
3. For $n = 2$, the interval is $(2,4)$ and 3 is the guaranteed prime.

## From Bertrand to the prime number theorem

Chebyshev proved the postulate in 1850, and Erdős gave an elementary proof in 1932. The prime number theorem implies it for large $n$ because the count up to $2n$ eventually exceeds the count up to $n$ by a wide margin. Bertrand's postulate is the classical first step from Euclid's infinitude of primes toward a quantitative distribution.

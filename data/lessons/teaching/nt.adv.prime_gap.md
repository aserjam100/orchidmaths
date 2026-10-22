# Prime Gaps and the Next Prime

**Prime gap:** The difference between two consecutive primes. The gap after a prime $p$ is $q - p$, where $q$ is the next prime. Finding the next prime means testing $p+1, p+2, \dots$ until a prime appears.

## The distance between consecutive primes

The gap from 2 to 3 is 1, and every later gap is even because every prime beyond 2 is odd. The gap can be arbitrarily large: the $n-1$ consecutive numbers $n!+2, \dots, n!+n$ are all composite. Average gaps grow like $\log p$ by the prime number theorem, but individual gaps fluctuate.

## Worked: gaps around 23 and 89

1. After 23, the candidates 24 through 28 are composite and 29 is prime, so the gap is 6.
2. After 89, the next prime is 97, so the gap is 8.
3. The smallest prime above 100 is 101; above 200 it is 211.

## Large gaps and the prime number theorem

The prime number theorem says the number of primes up to $x$ is asymptotic to $x/\log x$, so the average gap near $x$ is about $\log x$. It does not control individual gaps, and the largest known gaps are found by searching constrained patterns. Bounded gaps between primes, on the other hand, are the subject of recent breakthroughs.

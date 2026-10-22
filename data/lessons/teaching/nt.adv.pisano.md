# Pisano Periods

**Pisano period:** The Fibonacci sequence reduced modulo $n$ is periodic, and because the pair $(F_k, F_{k+1})$ determines the rest, it is periodic from the start. The Pisano period $\pi(n)$ is the smallest $k > 0$ with $F_k \equiv 0$ and $F_{k+1} \equiv 1 \pmod n$.

## Fibonacci modulo n

The pair of consecutive values is the state. There are only $n^2$ possible pairs, so the sequence must repeat; the pair $(0,1)$ is the start, so the first return to it closes the period. The state view also explains why the period is finite and why it can be found by iteration.

## Worked: periods modulo 2, 5 and 6

1. Modulo 2 the pairs are $(0,1)$, $(1,1)$, $(1,0)$, $(0,1)$, so $\pi(2) = 3$.
2. Modulo 5 the pair returns after 20 steps, so $\pi(5) = 20$.
3. Modulo 6 the period is the least common multiple of the periods modulo 2 and 3, which is 24.

## Why the period exists

The Fibonacci recurrence is deterministic on the finite set of pairs modulo $n$, so every orbit is eventually periodic. Because the transition $(a,b) \mapsto (b,a+b)$ has determinant $-1$, which is a unit modulo $n$, it is invertible, so the orbit is purely periodic. The period divides $n^2$, and in practice it is much smaller.

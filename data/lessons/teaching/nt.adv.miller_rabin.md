# The Miller-Rabin Test

**Miller-Rabin test:** Write $n-1 = 2^s d$ with $d$ odd. For a base $a$, compute $a^d \bmod n$; if it is 1 or $-1$, the base passes; otherwise square repeatedly up to $s-1$ times, and the base passes if a $-1$ appears. If no $-1$ appears, $a$ is a witness that $n$ is composite.

## Strong probable primes

A number that passes for a base is a strong probable prime to that base. Every prime passes for every base coprime to it; a composite that passes is a strong pseudoprime. The test strengthens Fermat's congruence by examining the square roots of 1 that appear on the way to $a^{n-1}$.

## Worked: a base that exposes 9

1. Take $n = 9$, so $n-1 = 8 = 2^3 \cdot 1$, with $s = 3$ and $d = 1$.
2. For base $a = 2$, $2^1 = 2$, then $2^2 = 4$, then $2^4 = 7 \pmod 9$; no $-1$ appears.
3. So 2 is a witness and 9 is composite.

## Error control and witnesses

A single passing base is not proof. For a composite $n$, at most one quarter of the bases are liars, so $k$ independent random bases leave a false-positive probability at most $4^{-k}$. Deterministic variants use a fixed small set of bases that is sufficient up to a stated bound. Carmichael numbers defeat Fermat's test but not Miller-Rabin.

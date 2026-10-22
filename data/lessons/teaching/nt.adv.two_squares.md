# Sums of Two Squares

**Two-square theorem:** A positive integer $n$ is a sum of two squares exactly when every prime $q \equiv 3 \pmod 4$ occurs in its factorisation with an even exponent. The number of representations, counting signs and order, is $r_2(n) = 4(d_1(n) - d_3(n))$, where $d_1$ and $d_3$ count divisors congruent to 1 and 3 modulo 4.

## Which numbers are representable

The obstruction is the primes 3 mod 4. A single odd power of such a prime makes the number impossible; for example 3, 6 and 12 are not sums of two squares, while $9 = 0^2+3^2$ and $18 = 3^2+3^2$ are.

## Worked: counting representations of 25

1. Factor $25 = 5^2$; the only prime factor is 1 mod 4, so representations exist.
2. The representations are $0^2 + 5^2$ and $3^2 + 4^2$ up to signs and order.
3. Counting gives $(\pm5,0)$, $(0,\pm5)$, $(\pm3,\pm4)$ and $(\pm4,\pm3)$: twelve pairs.

## The r2 formula

The divisor formula gives $r_2(25) = 4(d_1(25)-d_3(25)) = 4(3-0) = 12$. It also explains the zeros: if a 3-mod-4 prime has odd exponent, the two divisor counts are equal and the difference vanishes. The formula counts all integer pairs, so it is larger than the number of essentially different representations.

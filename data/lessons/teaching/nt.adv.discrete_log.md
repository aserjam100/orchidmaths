# Discrete Logarithms

**Discrete logarithm:** In the multiplicative group modulo a prime $p$, fix a primitive root $g$. Every nonzero residue $h$ is $g^x$ for a unique $x$ in $1,\dots,p-1$, and that exponent is the discrete logarithm $\log_g h$.

## The inverse of exponentiation

Exponentiation is easy: repeated squaring computes $g^x \bmod p$ in time polynomial in $\log x$. The inverse problem, recovering $x$ from $g$ and $h$, is believed to be hard in general, and that asymmetry is the basis of Diffie-Hellman key exchange and the security of discrete-log cryptosystems.

## Worked: a logarithm modulo 11

1. Take $p = 11$ and the primitive root $g = 2$.
2. The powers are 2, 4, 8, 5, 10, 9, 7, 3, 6, 1 for exponents 1 through 10.
3. Since $2^6 = 9$, the discrete logarithm of 9 base 2 is 6.

## Hardness and cryptography

The naive attack lists all $p-1$ powers, which is exponential in the number of digits. Faster algorithms such as baby-step giant-step and index calculus improve the bound but remain impractical for large primes. Choosing a prime with a large prime factor in $p-1$ blocks the Pohlig-Hellman reduction to small subgroups.

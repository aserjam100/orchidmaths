# The Carmichael Function

**Carmichael function:** $\lambda(n)$ is the exponent of the multiplicative group modulo $n$: the smallest positive $k$ such that $a^k \equiv 1 \pmod n$ for every $a$ coprime to $n$. It divides Euler's totient $\varphi(n)$ and can be strictly smaller.

## The exponent of the unit group

For a prime power $p^e$, $\lambda(p^e) = \varphi(p^e) = p^{e-1}(p-1)$, except for powers of 2 with $e \ge 3$, where $\lambda(2^e) = 2^{e-2}$. For a product of coprime prime powers, $\lambda$ is the least common multiple of the values on the factors.

## Worked: prime powers and products

1. $\lambda(8) = 2^{3-2} = 2$, while $\varphi(8) = 4$.
2. $\lambda(9) = \varphi(9) = 6$.
3. $\lambda(10) = \operatorname{lcm}(\lambda(2),\lambda(5)) = \operatorname{lcm}(1,4) = 4$.
4. $\lambda(15) = \operatorname{lcm}(2,4) = 4$.

## Carmichael versus Euler

Euler's theorem says $a^{\varphi(n)} \equiv 1$, but the exponent $\varphi(n)$ need not be minimal. The group modulo 8 has exponent 2 although $\varphi(8) = 4$: every odd square is 1 mod 8. The Carmichael function is the sharp exponent, and it is the right modulus for RSA when the key is constructed from $\lambda(n)$ rather than $\varphi(n)$.

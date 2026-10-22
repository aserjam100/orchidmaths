# Wilson's Theorem

**Wilson's theorem:** For a prime $p$,

$$(p-1)! \equiv -1 \pmod{p}$$

Conversely, if $(n-1)! \equiv -1 \pmod{n}$ then $n$ is prime.

The theorem is an exact description of the factorial modulo a prime, and it is unusual in that it characterises primes: every other congruence in elementary number theory that holds for all primes also holds for some composites, while this one holds for no composite beyond 4. What it does not give is a practical test. The left side costs $p-2$ multiplications, so it is slower than trial division by a wide margin, and its value is theoretical rather than computational.

## Factorial Modulo a Prime

### The Pairing Argument
Modulo a prime $p$, every $a \in \{1,\dots,p-1\}$ has a unique inverse $a^{-1}$. Only 1 and $p-1$ are self-inverse. All other elements pair as $a\cdot a^{-1}\equiv1$, leaving $(p-1)!\equiv1\cdot(p-1)\equiv-1$.

The self-inverse elements are the solutions of $a^2\equiv1\pmod p$, which factors as $(a-1)(a+1)\equiv0\pmod p$. A prime has no zero divisors, so $a\equiv1$ or $a\equiv-1$: exactly two of them. Every other element of $\{2,\dots,p-2\}$ therefore sits in a pair $\{a,a^{-1}\}$ of distinct elements, and the product over each pair is 1. The whole factorial collapses to the unpaired two, and their product is $-1$.

The argument is worth reading as a counting statement rather than an algebraic one. It says that the product of the units modulo $p$ is $-1$, and it explains why the same computation over a composite modulus can give something else: there, an element can fail to have an inverse, and a self-inverse element need not satisfy $a\equiv\pm1$.

### Converse
If $n$ is composite and $n\neq4$, then $(n-1)!\equiv0\pmod{n}$. The case $n=4$ gives $3!=6\equiv2\pmod4\neq-1$.

The proof is a factoring argument. A composite $n$ has a divisor $a$ with $1<a<n$, and either $n=ab$ with $1<b<n$, in which case both $a$ and $b$ appear among $1,\dots,n-1$ and their product $n$ divides the factorial, or the only factorisation is $n=a^2$. In the square case $a$ and $2a$ are both at most $n-1$ as soon as $a\ge3$, so $a\cdot2a=2n$ divides the factorial; when $a=2$ the modulus is $n=4$ and the factorial is $3!=6$, which is 2 and not $-1$.

So the only composite modulus for which the congruence could hold is 4, and it does not: Wilson's theorem is an exact test in principle. What makes it useless in practice is that a single test costs $p$ multiplications where a hundred Miller-Rabin rounds cost a few thousand for a modulus of the same size.

## Example

Let $p=5$: $(5-1)!=4!=24$. $24\div5=4$ remainder 4, and $4\equiv-1\pmod5$. So $24\equiv-1\pmod5$.

Let $p=7$: $6!=720$. $720\div7=102$ remainder 6, and $6\equiv-1\pmod7$.

For a larger prime the pairing is easier to trust than the division. Take $p=11$. The pairs are $2\leftrightarrow6$, $3\leftrightarrow4$, $5\leftrightarrow9$, $7\leftrightarrow8$, since $2\cdot6=12\equiv1$, $3\cdot4=12\equiv1$, $5\cdot9=45\equiv1$ and $7\cdot8=56\equiv1$. The unpaired elements are 1 and 10, so

$$10! \equiv 1\cdot(2\cdot6)\cdot(3\cdot4)\cdot(5\cdot9)\cdot(7\cdot8)\cdot10 \equiv 1\cdot1\cdot1\cdot1\cdot1\cdot10 \equiv -1 \pmod{11}.$$

The direct computation agrees: $10!=3628800=11\cdot329890+10$.

A composite shows the failure. Modulo 6 the factorial is $5!=120\equiv0$, and modulo 10 it is $9!=362880\equiv0$: both are 0 rather than $-1$, because the factors of the modulus have already appeared in the product.

## Application

Wilson gives a (slow) primality test and is used to prove that $a^{p-1}\equiv1$ can be refined for factorial structure.

Two consequences are used elsewhere in this domain. The first is the congruence $(p-1)!\equiv-1\pmod p$ rewritten as $(p-2)!\equiv1\pmod p$, which turns a factorial into an inverse and appears when a binomial coefficient is reduced modulo a prime. The second is the generalisation: the pairing argument needs only that the elements being multiplied are exactly the units of the ring, so the product of the units modulo $n$ is $-1$ exactly when $n$ is 1, 2, 4, $p^k$ or $2p^k$ for an odd prime $p$, and is 1 for every other modulus — Gauss's theorem, of which Wilson's theorem is the case $n=p$.

The reason the list is short is the number of self-inverse units. Each pair $\{a,a^{-1}\}$ contributes 1, so the product is the product of the self-inverse units, and those form a subgroup in which every element squares to 1. A group of that shape has 1, 2 or more elements, and its product is the unique element of order 2 only when there is exactly one; when there are none the product is 1, and when there are several they cancel in pairs. The moduli in the list are exactly the ones whose unit group has a single element of order 2.

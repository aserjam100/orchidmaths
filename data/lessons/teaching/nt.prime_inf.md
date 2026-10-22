# Infinitude of Primes

**Theorem (Euclid):** There are infinitely many primes.

The proof is the oldest argument in the subject that is still used unchanged, and it is a proof by construction rather than by contradiction at heart: given any finite list of primes, it produces a number with a prime factor outside the list. Nothing about the size of the list is assumed, so no finite list can be complete. What the proof does *not* say is that the number it builds is prime; it says only that the number has some prime factor that is new.

## Euclid's Construction

### The Product Plus One
Given any finite set of primes $\{p_1,\dots,p_k\}$, form

$$N = p_1p_2\cdots p_k + 1$$

The construction has one job: make the division by each known prime leave remainder 1. If $p_i$ is in the list then $p_i$ divides the product $p_1\cdots p_k$, so modulo $p_i$ the number $N$ is 1. That is all the proof needs, and it is why the $+1$ cannot be replaced by anything else that is not itself divisible by some $p_i$.

### A New Prime Factor
$N$ is not divisible by any $p_i$ (it leaves remainder 1 upon division by each $p_i$). Hence either $N$ is prime, or $N$ has a prime divisor not in the set. In either case a new prime exists.

The two cases are exhaustive because every integer greater than 1 has a prime divisor, which is the fundamental theorem of arithmetic's existence half. If $N$ itself is prime, then $N$ is the new prime. If $N$ is composite, take any prime $q\mid N$: the remainder argument excludes every prime in the list, so $q$ is outside it. Note that the proof does not tell you which case holds, and it cannot, because both occur.

The numbers produced this way are called primorials plus one, $p_k\#+1$, and they grow fast: the product of the first $k$ primes is about $e^{k\log k}$, so the tenth such number already has eight digits.

## Example

Take primes 2,3,5: $N=2\cdot3\cdot5+1=31$. 31 is not divisible by 2,3, or 5; it is prime and is a new prime.

Take primes 2,3,5,7: $N=2\cdot3\cdot5\cdot7+1=211$. 211 is prime.

Take primes 2,3,5,7,11,13: $N=30031=59\times509$. Both 59 and 509 are new primes not in the original list.

The last case is the one worth keeping, because it shows the theorem's shape rather than its conclusion. The number built from the first six primes is composite, and neither of its factors is in the list — the argument never needed $N$ to be prime, only that its factors are new. Running the construction again on the six primes plus 59 and 509 would produce another number, and so on; each step adds at least one prime that was not there before.

## Why This Proves Infinitude

If there were only finitely many primes, Euclid's construction applied to the complete finite list would produce a prime outside the list — contradiction.

Written out, suppose the primes are exactly $p_1,\dots,p_k$, and let $N=p_1\cdots p_k+1$. Every prime divides the product, so none divides $N$, yet $N>1$ has a prime divisor. The divisor is a prime that is not one of the $p_i$, contradicting the assumption that the list was complete. The hypothesis is therefore false, which is the theorem.

Two neighbouring facts are worth separating from this one. The argument gives no bound on the smallest prime outside a given list, and in fact the first prime larger than $p_k$ can be much smaller than $N$. And the sum of the reciprocals of the primes diverges, which is a stronger statement than infinitude: it says the primes are dense enough that their reciprocals accumulate, not merely that there are unboundedly many of them. Euclid's proof does not reach that conclusion, and Euler's proof, which does, is a different argument about products over primes.

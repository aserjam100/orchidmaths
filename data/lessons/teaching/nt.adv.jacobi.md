# The Jacobi Symbol

**Jacobi symbol:** For an odd positive $n$ with prime factorisation $n = p_1^{e_1}\cdots p_k^{e_k}$, the Jacobi symbol is $(a/n) = (a/p_1)^{e_1}\cdots(a/p_k)^{e_k}$, the product of Legendre symbols. It is defined for composite odd $n$, which the Legendre symbol is not.

## Extending Legendre to composite moduli

The Jacobi symbol inherits multiplicativity in both arguments and the reciprocity law, so it can be computed by factoring $n$ or by the reciprocity algorithm without factoring. Its value is $\pm 1$ when $\gcd(a,n)=1$ and 0 otherwise.

## Worked: two factorisations

1. $(2/15) = (2/3)(2/5) = (-1)(-1) = 1$.
2. $(7/15) = (7/3)(7/5) = (1/3)(2/5) = 1 \cdot (-1) = -1$.
3. $(6/35) = (2/35)(3/35) = (-1)(1) = -1$.

## What Jacobi does and does not tell you

$(a/n) = -1$ proves that $a$ is not a square modulo $n$. But $(a/n) = 1$ does not prove that $a$ is a square when $n$ is composite: the symbol is a product of Legendre symbols, and an even number of failures can cancel. That gap is why the Jacobi symbol alone cannot decide quadratic residuosity for composite $n$.

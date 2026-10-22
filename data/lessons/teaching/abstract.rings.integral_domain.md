# Integral Domains

**Integral domain:** A commutative ring with identity is an integral domain if $1 \neq 0$ and it has no zero divisors: $ab = 0$ forces $a = 0$ or $b = 0$. The integers $Z$ and every field are the standard examples.

## A commutative ring with no zero divisors

The defining property is multiplicative, and it is what makes the ring behave like $Z$: products of nonzero elements stay nonzero and cancellation is available. The condition $1 \neq 0$ excludes the zero ring. The polynomial ring over an integral domain is again an integral domain, because the leading coefficient of a product is the product of the leading coefficients.

## Worked: Z_7 is a domain, M_2(R) is not

1. $Z_7$ is a field, and every field is an integral domain: a nonzero $a$ has an inverse, so $ab = 0$ gives $b = a^{-1} \cdot 0 = 0$.
2. In the ring of 2x2 real matrices take a matrix with a 1 in its top-left corner and a matrix with a 1 in its bottom-right corner.
3. Their product is the zero matrix, yet neither matrix is zero, so $M_2(R)$ has zero divisors and is not an integral domain.

## Finite domains are fields

If $D$ is a finite integral domain and $a \neq 0$, the map $x \mapsto ax$ is injective, since $ax = ay$ implies $a(x - y) = 0$ and hence $x = y$. An injective map from a finite set to itself is surjective, so $ax = 1$ has a solution and $a$ is a unit. Every nonzero element is invertible, so $D$ is a field.

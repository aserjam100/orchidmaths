# Vector Spaces and Their Axioms

**Vector space:** A set $V$ over a field $F$ with an addition and a scalar multiplication, satisfying closure, associativity, commutativity, an additive identity, additive inverses, and the compatibility and distributivity laws. The axioms are what every proof in linear algebra is allowed to use.

## One set, two operations, eight rules

Addition makes $V$ an abelian group: it is closed, associative and commutative, has a zero vector, and every vector has a negative. Scalar multiplication adds the field action: $1v = v$, $(ab)v = a(bv)$, and the two distributive laws $a(u+v) = au + av$ and $(a+b)v = av + bv$. Nothing else is assumed, which is why the same theory covers arrows, polynomials, matrices and functions.

## Worked: checking a candidate space

1. Take the set of polynomials of degree at most 2 with real coefficients.
2. Adding two such polynomials and multiplying by a real scalar stay within degree at most 2, so the set is closed.
3. The zero polynomial, negatives and the distributive laws are inherited from the usual algebra of polynomials, so the set is a vector space.

The set of vectors $(x,y)$ in the plane with $x \ge 0$ and $y \ge 0$ fails: it is not closed under multiplication by $-1$.

## Why the axioms matter

Every theorem later in the subject is proved from the axioms alone, so a result about $\mathbb{R}^n$ automatically holds for polynomial spaces and function spaces. When a candidate set fails one axiom, that single failure is enough to disqualify it, and the failure usually shows up as a missing zero, a missing negative, or a failed closure.

# Prime and Maximal Ideals

Throughout this lesson $R$ is a commutative ring with identity. That convention matters, because the equivalences below read the quotient $R/I$, and the commutative setting is what makes $R/I$ a ring whose arithmetic reflects divisibility in $R$.

## Quotients that are domains or fields

An ideal $P$ is prime when $R/P$ is an integral domain, and an ideal $M$ is maximal when $R/M$ is a field. Multiplication in the quotient is $(a + P)(b + P) = ab + P$, so $ab \in P$ exactly when the product of the two classes is zero. Hence $P$ is prime exactly when $ab \in P$ forces $a \in P$ or $b \in P$.

## Worked: prime and maximal ideals of Z

1. $Z/(n)$ is isomorphic to $Z_n$, which is an integral domain exactly when $n$ is prime.
2. So the prime ideals of $Z$ are $(0)$ together with $(p)$ for each prime $p$.
3. Every nonzero prime ideal $(p)$ is maximal, because $Z_p$ is a field; the ideal $(0)$ is prime because $Z$ is a domain but is not maximal because $Z$ is not a field.

## Prime does not always mean maximal

A maximal ideal is always prime, since a field is an integral domain. The converse fails. In $Z[x]$ the ideal $(0)$ is prime but not maximal, and $(x)$ is prime because $Z[x]/(x)$ is isomorphic to $Z$, which is a domain but not a field. The two notions are kept distinct precisely because containment can continue past a prime ideal.

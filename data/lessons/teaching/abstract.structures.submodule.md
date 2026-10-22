# Submodules and Quotient Modules

**Submodule:** Let $M$ be a module over a ring $R$. A subset $N \subseteq M$ is a submodule if it is an additive subgroup of $M$ and is closed under scalar multiplication: $n \in N$ and $r \in R$ imply $rn \in N$.

## Closed under addition and scaling

The two closure conditions mirror what a subgroup needs from a group, and add the ring action. A submodule contains 0, because an additive subgroup does, and it is closed under negation. When $R$ is a field, modules are vector spaces and submodules are exactly the subspaces; when $R$ is $Z$, modules are abelian groups and submodules are exactly the subgroups.

## Worked: submodules of Z_6

The submodules of the $Z$-module $Z_6$ are its subgroups.

1. $\{0\}$ and $Z_6$ itself.
2. The subgroup of order 2, namely $\{0, 3\}$.
3. The subgroup of order 3, namely $\{0, 2, 4\}$.

Each is closed under addition mod 6 and under multiplication by any integer, so each is a submodule.

## Quotients of modules

Given a submodule $N$, the quotient $M/N$ is the set of cosets $m + N$ with addition and scalar multiplication defined on representatives. Well-definedness uses that $N$ is a submodule, not merely a subset. The quotient is again an $R$-module, and the map $m \mapsto m + N$ is a module homomorphism whose kernel is exactly $N$.

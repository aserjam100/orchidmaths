# Simple Groups

**Simple group:** A nontrivial group $G$ is simple if its only normal subgroups are $\{e\}$ and $G$ itself. Simplicity is the group analogue of primality: the group cannot be split into a nontrivial normal subgroup and a quotient.

## No proper nontrivial normal subgroups

To prove a group is simple, show that any normal subgroup is trivial. To refute simplicity, one counterexample is enough: a normal subgroup strictly between $\{e\}$ and $G$. Abelian groups give the first family of examples, because every subgroup of an abelian group is normal.

## Worked: Z_7 is simple and Z_6 is not

1. The subgroups of $Z_7$ correspond to the divisors of 7: only 1 and 7. So $Z_7$ has no proper nontrivial subgroup, hence no proper nontrivial normal subgroup.
2. The divisors of 6 give the subgroup $\{0, 3\}$ of order 2 inside $Z_6$.
3. Because $Z_6$ is abelian, that subgroup is normal, so $Z_6$ is not simple.

In general a cyclic group is simple exactly when its order is prime.

## The smallest nonabelian simple group

The alternating group $A_5$, of order 60, is the smallest nonabelian simple group. Every smaller nonabelian group has a nontrivial proper normal subgroup. The simple groups therefore begin with the cyclic groups $Z_p$ of prime order and the alternating group $A_5$.

# Composition Series

**Composition series:** A composition series of a group $G$ is a chain $\{e\} = G_0 \triangleleft G_1 \triangleleft \cdots \triangleleft G_n = G$ in which each quotient $G_{i+1}/G_i$ is simple. Those quotients are the composition factors, and $n$ is the length.

## Refining a subgroup chain until the factors are simple

Start with a normal series. Whenever some $G_{i+1}/G_i$ is not simple it has a proper nontrivial normal subgroup, and the correspondence theorem lifts it to a subgroup strictly between $G_i$ and $G_{i+1}$. Insert that subgroup and repeat. For a finite group the process stops, so every finite group has a composition series.

## Worked: the composition factors of Z_12

Refine the factorisation $12 = 2 \cdot 2 \cdot 3$.

1. Build $\{0\} \triangleleft \langle 6 \rangle \triangleleft \langle 3 \rangle \triangleleft Z_{12}$, using the subgroups of orders 2, 4 and 12.
2. The quotients $\langle 6 \rangle/\{0\}$, $\langle 3 \rangle/\langle 6 \rangle$ and $Z_{12}/\langle 3 \rangle$ are $Z_2$, $Z_2$ and $Z_3$.
3. Each is simple, so the length is 3 and the factors are $Z_2, Z_2, Z_3$.

## Jordan-Holder and solvable groups

The Jordan-Holder theorem says any two composition series of the same finite group have the same multiset of composition factors, so the factors are invariants even though the series need not be unique. A group is solvable exactly when all its composition factors are cyclic of prime order, because a simple abelian group must be cyclic of prime order.

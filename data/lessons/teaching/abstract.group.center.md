# Center of a Group

**Center:** The center $Z(G)$ is the set of elements of $G$ that commute with every element: $Z(G) = \{g \in G : gx = xg \text{ for all } x \in G\}$.

## The set of elements that commute with everything

Being central is a condition on one element against the whole group: $g$ is central exactly when $gx = xg$ holds for every $x$. The identity is always central, and a group is abelian exactly when $Z(G) = G$. A nonabelian group with $Z(G) = \{e\}$ is called centerless, and that is the typical case.

## Worked: the center of D_4

The dihedral group $D_4$ has rotations $e, r, r^2, r^3$ and reflections $s, sr, sr^2, sr^3$.

1. The half-turn $r^2$ commutes with every rotation and with every reflection, so $r^2 \in Z(D_4)$.
2. A reflection $s$ does not commute with $r$, because $sr = r^{-1}s \neq rs$.
3. The rotations $r$ and $r^3$ also fail to commute with $s$.

So $Z(D_4) = \{e, r^2\}$ has order 2, even though $D_4$ has order 8.

## The center is normal

For any $g$ and any $z \in Z(G)$, $gzg^{-1} = zgg^{-1} = z$, so conjugation fixes every central element. Therefore $Z(G)$ is a normal subgroup. The quotient $G/Z(G)$ measures how far $G$ is from abelian: it is trivial exactly when $G$ is abelian, and it is the group the inner automorphisms are isomorphic to.

# Group Automorphisms

**Automorphism:** An automorphism of a group $G$ is an isomorphism from $G$ to itself. The automorphisms form a group $\operatorname{Aut}(G)$ under composition, with the identity map as unit.

## Structure-preserving bijections

An automorphism must be a bijection and satisfy $\varphi(xy) = \varphi(x)\varphi(y)$. It sends the identity to the identity and preserves the order of every element, so it permutes the elements of each order among themselves. That constraint is what keeps $\operatorname{Aut}(G)$ from being all of the symmetric group on the elements.

## Worked: Aut(Z_5) has four elements

A homomorphism from $Z_5$ is fixed by the image of the generator 1.

1. If $\varphi(1) = k$, then $\varphi(n) = nk \bmod 5$.
2. For $\varphi$ to be bijective, $k$ must generate $Z_5$, so $k \in \{1, 2, 3, 4\}$.
3. Multiplication by each of those four values is bijective.

So $\operatorname{Aut}(Z_5)$ has $4 = \varphi(5)$ elements. More generally $\operatorname{Aut}(Z_n)$ has $\varphi(n)$ elements, isomorphic to the unit group $(Z/nZ)^\times$.

## Inner automorphisms

Conjugation by a fixed element, $\iota_g(x) = gxg^{-1}$, is always an automorphism, called inner. The map $g \mapsto \iota_g$ is a homomorphism from $G$ to $\operatorname{Aut}(G)$ whose kernel is $Z(G)$, so the inner automorphisms form a subgroup isomorphic to $G/Z(G)$, and that subgroup is normal in $\operatorname{Aut}(G)$.

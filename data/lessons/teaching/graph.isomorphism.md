# Graph Isomorphism

**Isomorphism:** Two graphs $G$ and $H$ are isomorphic when a bijection $\varphi: V(G) \to V(H)$ preserves adjacency and non-adjacency: $uv$ is an edge of $G$ exactly when $\varphi(u)\varphi(v)$ is an edge of $H$. The map is an isomorphism, and isomorphism is an equivalence relation on graphs.

## Same structure, different labels

An isomorphism is a relabelling. The drawing, the vertex names and the order of the edge list are irrelevant; only the incidence pattern matters. Every graph is isomorphic to itself, the inverse of an isomorphism is an isomorphism, and a composition of isomorphisms is an isomorphism, which is why "isomorphic" is an equivalence relation.

## Worked: degree sequences

1. The path $P_4$ has degree sequence $(1,1,2,2)$.
2. The cycle $C_4$ has degree sequence $(2,2,2,2)$.
3. The sequences differ, so the graphs are not isomorphic.

The degree sequence is necessary but not sufficient: two graphs can share it and still differ in structure. The complement, the number of triangles and the spectrum are further invariants that can separate them.

## Invariants and non-invariants

A quantity preserved by every isomorphism is an invariant. Vertex count, edge count, degree multiset, number of connected components and diameter are invariants. The drawing, the labels and the edge ordering are not. To prove two graphs isomorphic, exhibit the bijection; to prove them non-isomorphic, find one invariant they disagree on.

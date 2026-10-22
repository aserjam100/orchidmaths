# Graph Homomorphisms

**Homomorphism:** A graph homomorphism $f: G \to H$ maps every vertex of $G$ to a vertex of $H$ so that every edge $uv$ of $G$ maps to an edge $f(u)f(v)$ of $H$. It need not be injective or surjective, and non-adjacent vertices may map to adjacent ones.

## Maps that preserve edges

The condition is one-directional: adjacency in $G$ forces adjacency in $H$, but adjacency in $H$ imposes nothing on the preimages. This is exactly what a coloring does: it collapses each independent color class to a single target vertex while keeping every edge visible in the target.

## Worked: colorings as homomorphisms

1. A proper $k$-coloring of $G$ assigns one of $k$ colors so that adjacent vertices differ.
2. Regard the colors as the vertices of $K_k$.
3. The coloring is a homomorphism $G \to K_k$, because every edge of $G$ joins two different colors, hence an edge of $K_k$.

A graph is $k$-colorable exactly when it has a homomorphism to $K_k$.

## Isomorphism as the bijective case

An isomorphism is a homomorphism that is bijective and whose inverse is also a homomorphism. Requiring only a bijective homomorphism is not enough in general, because it can identify a non-edge with an edge. The homomorphism viewpoint also gives the clique number a dual reading: a clique in $G$ is a homomorphism from $K_r$ into $G$.

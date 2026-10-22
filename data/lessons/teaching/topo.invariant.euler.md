# Euler Characteristic

**Euler characteristic:** For a finite cell complex, $\chi = V - E + F$, and more generally the alternating sum of the number of cells in each dimension. For a closed orientable surface, $\chi = 2 - 2g$, where $g$ is the genus.

## Vertices minus edges plus faces

The alternating sum is unchanged by subdividing a cell complex: adding a vertex to an edge adds one vertex and one edge, and adding an edge across a face adds one edge and one face. That invariance is why $\chi$ can be computed from any convenient triangulation.

## Worked: the tetrahedron and the torus

1. A tetrahedron has $V = 4$, $E = 6$ and $F = 4$, so $\chi = 4 - 6 + 4 = 2$.
2. A triangulated torus has $V = 9$, $E = 27$ and $F = 18$, so $\chi = 9 - 27 + 18 = 0$.
3. A triangulated projective plane has $V = 6$, $E = 15$ and $F = 10$, so $\chi = 1$.

## Why it is a topological invariant

The Euler characteristic is the alternating sum of the Betti numbers, $\chi = b_0 - b_1 + b_2 - \cdots$, and Betti numbers are unchanged by homeomorphism. It therefore detects some differences: a sphere has $\chi = 2$, a torus has $\chi = 0$, and no amount of bending turns one into the other. It is not a complete invariant: the torus and the Klein bottle both have $\chi = 0$ but are not homeomorphic.

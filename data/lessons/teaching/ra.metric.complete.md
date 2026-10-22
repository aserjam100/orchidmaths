# Completeness

**Complete metric space:** one in which every Cauchy sequence converges to a point of the space; completeness is what the rationals lack and the reals have.

## Completeness as an axiom

A metric space is complete when every Cauchy sequence in it converges to a point of the space. Completeness is not a topological property: two equivalent metrics can give the same open sets while one is complete and the other is not, since the metric itself, not just its topology, decides which sequences are Cauchy. On the line completeness is equivalent to the least upper bound property, and the real line with the usual metric is complete. The rationals with the usual metric are not, which is why they are enlarged to the reals.

## Completions and subspaces

Every metric space has a completion, the smallest complete space containing it, built from the Cauchy sequences modulo the relation of having distance tending to zero. The property of a subspace is not automatic: a closed subset of a complete space is complete, because the limit of a Cauchy sequence stays in the closed set, while an open subset such as the interval from zero to one in the line is not complete. Completeness is therefore inherited by closed subsets and by products, and it can always be obtained by completion.

## Uses

Completeness is what turns an approximation into existence: a sequence of better and better candidates converges precisely when the space is complete. The contraction mapping principle and the existence theorems for differential equations depend on it, as does the Baire category theorem. In each case a Cauchy sequence of approximate solutions is produced and completeness supplies the actual solution, which is the pattern behind numerical analysis as much as behind pure existence proofs.

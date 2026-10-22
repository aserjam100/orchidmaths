# Homeomorphisms

**Homeomorphism:** A homeomorphism is a bijection $f: X \to Y$ that is continuous and whose inverse is continuous. When one exists, $X$ and $Y$ are topologically the same: they have the same open sets up to renaming the points.

## Same shape, different metric

Distance is not part of the topology. Stretching or bending a space without tearing or gluing gives a homeomorphism, even though distances change. A square and a circle are homeomorphic, and so are the open interval $(0,1)$ and the whole real line.

## Worked: the open interval is the line

1. Define $f(x) = \tan(\pi(x - 1/2))$ on $(0, 1)$.
2. As $x$ runs from 0 to 1, the angle runs from $-\pi/2$ to $\pi/2$, and the tangent runs over all of $\mathbb{R}$.
3. The inverse is $\arctan$, which is continuous, so $f$ is a homeomorphism.

## Topological invariants

A property preserved by every homeomorphism is a topological invariant. Connectedness, compactness, Hausdorffness and the number of connected components are invariants; boundedness, diameter and the particular metric are not. Invariants are how one proves two spaces are not homeomorphic: find a property one has and the other lacks.

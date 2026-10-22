# Joint Densities

**Joint density:** a nonnegative function $f(x, y)$ whose integral over a region gives the probability that the pair lies in that region, and which integrates to one over the whole plane.

## Density over a region

When a pair has a joint density $f$, the probability that $(X, Y)$ falls in a region $A$ is the double integral of $f$ over $A$. The density must be nonnegative and integrate to one. A point has probability zero, so probabilities are always attached to areas, and the density is the local mass per unit area.

## Worked: the uniform rectangle

A pair uniform on the rectangle from zero to four in $x$ and zero to five in $y$ has density one twentieth throughout the rectangle. The probability that $X$ is at most two and $Y$ at most three is $(2/4)(3/5) = 0.3$, the area fraction of the sub-rectangle. For a uniform pair the answer is always a ratio of areas.

## Independence and the product

A uniform rectangle has independent coordinates, and more generally a joint density factors as the product of the two marginal densities exactly when the variables are independent. When the rectangle is tilted, or the density does not separate, the coordinates are dependent. The factorisation is the density version of the product rule for independent events.

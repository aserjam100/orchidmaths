# Projection onto a vector

## Projecting a vector

The projection of a vector onto a direction is the part of it that lies along that direction, and the rest is perpendicular to it. Splitting a vector into those two parts is one of the most used operations in the subject.

The projection is a multiple of the direction vector, and the multiplier is the scalar the projection is described by.

## The formula

For a vector \(v\) and a direction \(u\), the scalar is

\[
c = \frac{v \cdot u}{u \cdot u},
\]

and the projection is \(cu\). The numerator measures how much of \(v\) points along \(u\) and the denominator corrects for the length of \(u\), which is why the direction only has to be non-zero rather than unit.

## What the coefficient means

The coefficient is the projection's length measured in units of the direction vector, so it is not the length itself unless the direction is a unit vector. That distinction is the usual error in a computation: a coefficient of three says the projection is three copies of the direction.

Choosing a unit direction removes the denominator, which is why orthonormal bases make projection arithmetic trivial.

# Rotation-scaling matrices

## Rotation and scaling

A matrix of the form with equal diagonal entries and opposite off-diagonal entries is a rotation combined with a scaling: it turns every vector through the same angle and multiplies its length by the same factor.

That single shape covers both operations at once, which is why it is the standard example of a transformation that is not a pure rotation or a pure scaling.

## The determinant as a squared scale

Its determinant is the sum of the two squared entries, which is the square of the scaling factor. A determinant of the sum of squares can never be negative, and that is a property of rotations: they never reverse orientation, so the signed area they scale cannot change sign.

A pure rotation has a determinant of one, which is the case where the scaling factor is one.

## When the determinant is zero

Since the determinant is a sum of squares, it is zero only when both entries are zero, in which case the matrix sends everything to the origin. So a matrix of this shape is either invertible or the zero matrix, with nothing in between.

That is a useful thing to notice: for the shapes that represent a geometric operation, the algebra of the determinant says something about the geometry directly.

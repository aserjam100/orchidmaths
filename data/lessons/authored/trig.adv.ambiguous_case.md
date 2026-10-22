# The ambiguous case

## Why side-side-angle is not enough

Two sides and an angle determine a triangle when the angle is between the sides. When it is not, the
same three numbers can describe two different triangles, and that is not a failure of arithmetic:
the sine of an angle and the sine of its supplement are equal, so an angle recovered from a sine has
two candidates.

A side-side-angle triple is therefore genuinely open.

## The height decides

The pair of the non-included angle \(A\) and the side \(b\) opposite the other angle fixes a height.
Dropping a perpendicular from the far end of \(b\) gives \(h = b\sin A\), and the remaining side
\(a\), which is opposite \(A\), has to reach from the third vertex to that perpendicular's line.

Comparing \(a\) with \(h\) and with \(b\) is the whole test. A side shorter than the height never
reaches the opposite ray. A side between the height and \(b\) reaches it twice, once on each side of
the perpendicular, and those are the two triangles. A side at least as long as \(b\) reaches it
once, because the only candidate angle left is the acute one.

## Counting the triangles

A side exactly equal to the height is the boundary: it meets the ray once, at a right angle, so it
gives one triangle. That case is easy to miss, because it is the one value where the two
possibilities collapse rather than disappear.

The test is worth running before any angle is computed, since the answer to a side-side-angle
question is sometimes just this count.

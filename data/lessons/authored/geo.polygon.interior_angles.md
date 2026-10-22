# The interior-angle sum of a polygon

## Triangulating the polygon

Drawing all the diagonals from one vertex cuts a polygon with \(n\) sides into \(n-2\) triangles. Each triangle's angles sum to 180 degrees, and together they account for every interior angle of the polygon exactly once.

So the interior angles sum to

\[
180(n - 2)
\]

degrees. For a pentagon that is \(180 \times 3 = 540\).

## The formula

The formula works for any simple polygon, convex or not, because the triangulation only needs the polygon to have no crossing sides. A quadrilateral gives two triangles and a sum of 360, and a triangle gives one triangle and 180, recovering the facts already known.

The number of triangles is always two fewer than the number of sides, which is the whole content of the formula.

## Regular polygons

In a regular polygon every interior angle is equal, so each is the sum divided by the number of sides. A regular pentagon has interior angles of \(540/5 = 108\) degrees.

Combining the two facts, each interior angle of a regular \(n\)-gon is \(180(n-2)/n\) degrees, which is worth reading as "the sum, shared out equally" rather than memorising separately.

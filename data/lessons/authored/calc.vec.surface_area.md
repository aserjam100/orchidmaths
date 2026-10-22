# Surface area of a graph

## The area element

For a surface given as \(z = f(x,y)\) above a region in the plane, a small rectangle \(dx\,dy\) in the plane covers a tilted patch whose area is larger by a factor of

\[\sqrt{1 + f_x^2 + f_y^2}.\]

The factor is the length of the cross product of the two tangent vectors, which is why it depends on both partial derivatives: a surface that tilts in two directions stretches in both.

## A tilted plane

For \(z = 2x + 2y\) over a rectangle of sides \(p\) and \(q\), the factor is \(\sqrt{1 + 4 + 4} = 3\), a constant. The area is therefore \(3pq\), three times the area of the rectangle it sits above.

A flat surface with \(f_x = f_y = 0\) gives the factor 1, which recovers the ordinary area — the surface integral over the plane \(z = 0\) is just a double integral.

## When the factor varies

For a curved surface the factor varies from point to point and stays inside the integral. For \(z = x^2\) the factor is \(\sqrt{1 + 4x^2}\), which grows with \(x\), so equal rectangles in the plane cover larger patches as \(x\) grows.

That is the difference between the surface integral and the double integral: the surface integral counts the tilted area, and the double integral counts the shadow.

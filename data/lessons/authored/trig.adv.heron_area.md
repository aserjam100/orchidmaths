# Area from three sides

## An area without a height

The usual area formula, \(\tfrac{1}{2} \cdot \text{base} \cdot \text{height}\), needs a height, and
three sides do not supply one. A triangle is still completely determined by its three sides,
though, so its area must be recoverable from them alone.

The bridge between the two facts is the law of cosines. It turns the sides into an angle, and the
angle turns one of the sides into a height.

## The semiperimeter

Heron's formula goes straight to the answer through the semiperimeter, \(s = (a+b+c)/2\), the
perimeter halved. Then

\(A = \sqrt{s(s-a)(s-b)(s-c)}\).

The three sides enter symmetrically, which matches the geometry: no side is the base and no vertex
is special. For a 13-14-15 triangle the semiperimeter is 21, and the product inside
the root is \(21 \cdot 8 \cdot 7 \cdot 6 = 7056\), whose square root is 84.

## Where the formula comes from

Each factor \(s-a\) is the excess of the semiperimeter over one side, and the product of the four
factors measures how much room the triangle has. When one side grows to the length of the other two
added together, the triangle flattens, one factor goes to zero, and so does the area.

The formula is worth reading as a distance between a triangle and a degenerate one. A triangle that
is nowhere near flat has a large product; a sliver has a product close to zero.

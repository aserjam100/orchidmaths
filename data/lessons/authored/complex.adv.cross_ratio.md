# The cross-ratio

The **cross-ratio** of four points is \((z_1, z_2; z_3, z_4) = \frac{(z_3 - z_1)(z_4 - z_2)}{(z_3 - z_2)(z_4 - z_1)}\). It is the one arrangement of four differences that is unchanged by Möbius transformations.

## Definition and order

The order of the four points matters: swapping the first two points inverts the value, and swapping the pairs sends it to \(1 - \) the value. A computation has to use the order the question puts the points in.

The denominator is nonzero when the four points are distinct, which is the standing assumption.

## Why real values matter

The cross-ratio is real exactly when the four points lie on a common circle or line. That is the practical use: with four concyclic points the value is a real number, and a question can ask for it as an ordinary rational.

For four points on a line the same statement holds, because a line is a circle through infinity. Computing the cross-ratio of real points therefore always gives a real number.

## Computing one

For \(z_1 = 0\), \(z_2 = 2\), \(z_3 = 4\), \(z_4 = 8\): the numerator is \((4-0)(8-2) = 24\) and the denominator is \((4-2)(8-0) = 16\), so the value is \(24/16 = 3/2\).

The result is real, as the criterion promises. A complex value from four real inputs would mean an arithmetic error, which makes the computation self-checking.
